package middleware

import (
	//"errors"
	"gin-vue-admin/config"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/sysModel"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth_wx() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localSstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		ModelToken := sysModel.JwtBlacklist{
			Jwt: token,
		}
		if token == "" {
			servers.ReportFormat(c, false, "未登录或非法访问", gin.H{
				"reload": true,
			})
			c.Abort()
			return
		}
		if ModelToken.IsBlacklist(token) {
			servers.ReportFormat(c, false, "您的帐户异地登陆或令牌失效", gin.H{
				"reload": true,
			})
			c.Abort()
			return
		}
		j := NewJWT_wx()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken_wx(token)
		if err != nil {
			if err == TokenExpired {
				servers.ReportFormat(c, false, "授权已过期", gin.H{
					"reload": true,
				})
				c.Abort()
				return
			}
			servers.ReportFormat(c, false, err.Error(), gin.H{
				"reload": true,
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

type JWT_wx struct {
	SigningKey []byte
}

/*var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)*/

type CustomClaims_wx struct {
	Open_Id     string
	Session_key string
	jwt.StandardClaims
}

func NewJWT_wx() *JWT_wx {
	return &JWT_wx{
		[]byte(config.GinVueAdminconfig.JWT.SigningKey),
	}
}

//创建一个token
func (j *JWT_wx) CreateToken_wx(claims CustomClaims_wx) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//解析 token
func (j *JWT_wx) ParseToken_wx(tokenString string) (*CustomClaims_wx, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims_wx{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims_wx); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

// 更新token
func (j *JWT_wx) RefreshToken_wx(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims_wx{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims_wx); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken_wx(*claims)
	}
	return "", TokenInvalid
}
