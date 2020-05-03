package middleware

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/userJobs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//拦截器  验证用户是否审核，审核通过，列出企业信息
func EnterpriseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*CustomClaims)
		if waitUse.EnterPriseId != 0 {
			ep := &userJobs.EnterpriseInfo{
				Model: gorm.Model{ID: waitUse.EnterPriseId},
			}
			err, enpInfo := ep.FindById()
			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("企业信息不存在"), gin.H{})
				c.Abort()
				return
			}
			c.Set("enpInfo", &enpInfo)
			//获取登录用户Id
			ua := &userJobs.UserAuth{
				UserId: waitUse.ID,
			}
			err = ua.FindByUserId()
			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("请点击头像上传企业认证"), gin.H{})
				c.Abort()
				return
			}
			if ua.Status != 1 {
				servers.ReportFormat(c, false, fmt.Sprintf("管理员审核中"), gin.H{})
				c.Abort()
				return
			}

		}
		c.Next()
	}
}
