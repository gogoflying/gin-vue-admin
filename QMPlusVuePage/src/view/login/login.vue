<template>
  <el-container class="login-register-box">
    <el-dialog title="找回密码" :visible.sync="dialogFormVisible">
      <el-form :model="forgetForm">
        <el-form-item>
          <el-input
            style="width:80%;"
            placeholder="请输入用户名"
            v-model="forgetForm.userName"
            autocomplete="off"
          >
            <i class="el-input__icon el-icon-user" slot="suffix"></i>
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-select style="width:80%;" placeholder="请选择用户角色" v-model="forgetForm.roleId">
            <el-option :key="role.name" :label="role.name" :value="role.id" v-for="role in roles"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item style="position:relative;width:80%;">
          <el-row :gutter="10">
            <el-col :span="20">
              <el-input
                v-model="forgetForm.captcha"
                name="logVerify"
                placeholder="请输入验证码"
                maxlength="10"
              />
              <img :src="path + picPath" alt="请输入验证码" @click="loginVefify1()" class="vPic1" />
            </el-col>
            <el-col :span="4">
              <el-button type="primary" size="small" :disabled="btnEnable" @click="sendEmail()" round>{{btnName}}</el-button>
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item>
          <el-input
            style="width:80%;"
            placeholder="请输入邮箱验证码"
            v-model="forgetForm.code"
            autocomplete="off"
          >
            <i class="el-input__icon el-icon-orange" slot="suffix"></i>
          </el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="onOK()">确 定</el-button>
      </div>
    </el-dialog>

    <vue-particle-line></vue-particle-line>
    <el-main class="login-box">
      <h1 class="title-1">
        <img class="logo" :src="require('@/assets/logo.jpg')" alt srcset />
      </h1>
      <div class="main">
        <el-form :model="loginForm" :rules="rules" ref="loginForm">
          <el-form-item prop="username">
            <el-input placeholder="请输入用户名" v-model="loginForm.username">
              <i class="el-input__icon el-icon-user" slot="suffix"></i>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              :type="lock==='lock'?'password':'text'"
              placeholder="请输入密码"
              v-model="loginForm.password"
            >
              <i :class="'el-input__icon el-icon-' + lock" @click="changeLock" slot="suffix"></i>
            </el-input>
          </el-form-item>
          <el-form-item style="position:relative">
            <el-input
              v-model="loginForm.captcha"
              name="logVerify"
              placeholder="请输入验证码"
              maxlength="10"
              @keyup.enter.native="submitForm('ruleForm2')"
            />
            <img :src="path + picPath" alt="请输入验证码" @click="loginVefify()" class="vPic" />
          </el-form-item>
          <el-form-item>
            <el-row :gutter="20">
              <el-col :span="8">
                <el-checkbox v-model="checked" style="width:100%">记住密码</el-checkbox>
              </el-col>
              <el-col :span="16">
                <el-link class="fl-right" @click="forgetPass()" type="primary">忘记密码</el-link>
              </el-col>
            </el-row>
          </el-form-item>
          <el-form-item>
            <el-button @click="submitForm" style="width:100%">登 录</el-button>
          </el-form-item>
          <el-form-item>
            <el-button @click="register" style="width:100%">注 册</el-button>
          </el-form-item>
        </el-form>
      </div>
      <!-- <h3 class="title-3 fl-right">测试用户:admin 密码:123456</h3> -->
    </el-main>
  </el-container>
</template>

<script>
import { mapActions } from "vuex";
import { captcha, forgetVerify, forgetSendEmail } from "@/api/user";
const path = process.env.VUE_APP_BASE_API;
export default {
  name: "Login",
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5 || value.length > 12) {
        return callback(new Error("请输入正确的用户名"));
      } else {
        callback();
      }
    };
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6 || value.length > 12) {
        return callback(new Error("请输入正确的密码"));
      } else {
        callback();
      }
    };

    return {
      checked: false,
      dialogFormVisible: false,
      lock: "lock",
      loginForm: {
        username: "",
        password: "",
        captcha: "",
        captchaId: ""
      },
      forgetForm: {
        userName: "",
        code: "",
        captcha: "",
        captchaId: "",
        roleId:""
      },
      roles: [
        {
          id: "9528",
          name: "招聘管理员"
        },
        {
          id: "9529",
          name: "薪酬管理员"
        }
      ],
      rules: {
        username: [{ validator: checkUsername, trigger: "blur" }],
        password: [{ validator: checkPassword, trigger: "blur" }]
      },
      path: path,
      logVerify: "",
      picPath: "",
      isOk: false,
      btnName:"发送验证码",
      btnEnable:false,
      count:90,
      timer:null
    };
  },
  created() {
    this.loginVefify();
  },
  methods: {
    ...mapActions("user", ["LoginIn"]),
    async login() {
      await this.LoginIn(this.loginForm);
    },
    async sendEmail() {
      if(this.forgetForm.userName == "" || this.forgetForm.roleId == "" || this.forgetForm.captcha == ""){
        this.$message({ type: "error", message: "不能为空" });
        return;
      }
      var res = await forgetSendEmail(this.forgetForm);
      if (res.success) {
        this.isOk = true;
        this.timer = setInterval(this.callbackTimer, 1000);
        this.btnEnable = true;
         this.$message({ type: "success", message: res.msg });
      } else {
        this.$message({ type: "error", message: res.msg });
      }
      this.loginVefify1() 
    },

    callbackTimer(){
      this.count --
      if(this.count == 0){
        this.btnEnable = false;
        this.btnName = "发送验证码";
        this.count = 90;
        clearInterval(this.timer);
      }else{
        this.btnName = "发送验证码（ " + this.count + "）";
      }
    },

    async onOK() {
      if (this.isOk) {
        if(this.forgetForm.userName == "" || this.forgetForm.code == ""){
          this.$message({ type: "error", message: "不能为空" });
          return;
        }
        var res = await forgetVerify(this.forgetForm);
        if (res.success) {
          this.$message({ type: "success", message: res.msg });
          this.dialogFormVisible = !this.dialogFormVisible;
        } else {
          this.$message({ type: "error", message: res.msg });
        }
      } else {
        this.$message({ type: "error", message: "请先发送邮箱验证码" });
      }
    },
    async forgetPass() {
      this.dialogFormVisible = !this.dialogFormVisible;
    },
    async submitForm() {
      this.$refs.loginForm.validate(async v => {
        if (v) {
          this.cookieValid();
          this.login();
          this.loginVefify();
        } else {
          this.$message({
            type: "error",
            message: "请正确填写登录信息",
            showClose: true
          });
          this.loginVefify();
          return false;
        }
      });
    },
    register() {
      this.$router.push({ name: "register" });
    },
    changeLock() {
      this.lock === "lock" ? (this.lock = "unlock") : (this.lock = "lock");
    },
    loginVefify() {
      captcha({}).then(ele => {
        this.picPath = ele.data.picPath;
        this.loginForm.captchaId = ele.data.captchaId;
      });
    },
    loginVefify1() {
      captcha({}).then(ele => {
        this.picPath = ele.data.picPath;
        this.forgetForm.captchaId = ele.data.captchaId;
      });
    },
    //校验cookie
    cookieValid() {
      const self = this;
      if (self.checked == true) {
        self.setCookie(self.loginForm.username, self.loginForm.password, 7);
      } else {
        self.clearCookie();
      }
    },
    //设置cookie
    setCookie(c_name, c_pwd, exdays) {
      var exdate = new Date(); //获取时间
      exdate.setTime(exdate.getTime() + 24 * 60 * 60 * 1000 * exdays); //保存的天数
      //字符串拼接cookie
      window.document.cookie =
        "userName" + "=" + c_name + ";path=/;expires=" + exdate.toGMTString();
      window.document.cookie =
        "userPwd" + "=" + c_pwd + ";path=/;expires=" + exdate.toGMTString();
    },
    //读取cookie
    getCookie: function() {
      if (document.cookie.length > 0) {
        var arr = document.cookie.split("; "); //这里显示的格式需要切割一下自己可输出看下
        for (var i = 0; i < arr.length; i++) {
          var arr2 = arr[i].split("="); //再次切割
          //判断查找相对应的值
          if (arr2[0] == "userName") {
            this.loginForm.username = arr2[1]; //保存到保存数据的地方
          } else if (arr2[0] == "userPwd") {
            this.loginForm.password = arr2[1];
          }
        }
      }
    },
    //清除cookie
    clearCookie: function() {
      this.setCookie("", "", -1); //修改2值都为空，天数为负1天就好了
    }
  },
  mounted() {
    this.getCookie();
  }
};
</script>

<style scoped lang="scss">
.login-register-box {
  background: #fff;
  height: 100vh;
  .login-box {
    width: 80vw;
    position: absolute;
    left: 50%;
    margin-left: -40vw;
    top: 5vh;
    .logo {
      height: 14vh;
      width: 49vh;
    }
    .main {
      min-width: 200px;
      width: 260px;
      margin: 0 auto;
    }
  }
  .vPic {
    position: absolute;
    right: 10px;
    bottom: 0px; // 适配ie
  }
  .vPic1 {
    position: absolute;
    right: 90px;
    bottom: 10px; // 适配ie
  }
}
</style>
