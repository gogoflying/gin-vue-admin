<template>
  <el-container class="login-register-box">
    <vue-particle-line></vue-particle-line>
    <el-main class="login-box">
      <h1 class="title-1">
        <img class="logo" :src="require('@/assets/logo.jpg')" />
      </h1>
      <div class="main">
        <el-form :model="registerForm" :rules="rules" ref="registerForm">
          <el-form-item label="用户名" label-width="80px" prop="username">
            <el-input placeholder="请输入用户名" v-model="registerForm.username"></el-input>
          </el-form-item>
          <el-form-item label="密码" label-width="80px" prop="password">
            <el-input
              :type="lock==='lock'?'password':'text'"
              placeholder="请输入密码"
              v-model="registerForm.password"
            >
              <i :class="'el-input__icon el-icon-' + lock" @click="changeLock" slot="suffix"></i>
            </el-input>
          </el-form-item>
          <el-form-item label="确认密码" label-width="80px" prop="rePassword">
            <el-input
              :type="lock==='lock'?'password':'text'"
              placeholder="请再一次输入密码"
              v-model="registerForm.rePassword"
            >
              <i :class="'el-input__icon el-icon-' + lock" @click="changeLock" slot="suffix"></i>
            </el-input>
          </el-form-item>
          <br />
          <br />
          <el-form-item label="企业名称" label-width="80px" prop="enterprise_name">
            <el-input placeholder="请输入与营业执照一致的企业名称" v-model="registerForm.enterprise_name"></el-input>
          </el-form-item>
          <el-form-item label="用户角色" label-width="80px" prop="authorityId">
            <el-select placeholder="请选择用户角色" v-model="registerForm.authorityId" style="width:100%">
              <el-option :key="role.name" :label="role.name" :value="role.id" v-for="role in roles"></el-option>
            </el-select>
          </el-form-item>
          <!-- <el-form-item  label="联系人" label-width="80px" prop="telname">
          <el-input placeholder="请输入联系人全名" v-model="registerForm.telname"></el-input>
        </el-form-item>
        <el-form-item label="电话" label-width="80px" prop="telphone">
          <el-input placeholder="请输入电话" v-model="registerForm.telphone"></el-input>
        </el-form-item>
        <el-form-item label="企业邮箱" label-width="80px" prop="companyemail">
          <el-input placeholder="请输入企业邮箱" v-model="registerForm.companyemail"></el-input>
          </el-form-item>-->
          <el-form-item>
            <el-button @click="submitForm" style="width:100%">立即注册</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-main>
  </el-container>
</template>

<script>
import { mapActions } from "vuex";
import { register } from "@/api/user";
export default {
  name: "Register",
  data() {
    const ratioPassword = (rule, value, callback) => {
      if (value != this.registerForm.password) {
        return callback(new Error("两次密码不同"));
      } else {
        callback();
      }
    };
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
    const checkRole = (rule, value, callback) => {
      if (value == "" || value == null) {
        return callback(new Error("请选择角色"));
      } else {
        callback();
      }
    };

    return {
      lock: "lock",
      registerForm: {
        username: "",
        password: "",
        rePassword: "",
        enterprise_name: "",
        authorityId: ""
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
        password: [{ validator: checkPassword, trigger: "blur" }],
        rePassword: [{ validator: ratioPassword, trigger: "blur" }],
        authorityId: [{ validator: checkRole, trigger: "blur" }]
      }
    };
  },
  methods: {
    ...mapActions("user", ["LoginIn"]),
    async submitForm() {
      this.$refs.registerForm.validate(async v => {
        if (v) {
          const res = await register(this.registerForm);
          if (res.success) {
            this.$message({
              type: "success",
              message: "注册成功",
              showClose: true
            });
            this.$router.push({ name: "login" });
          }
        } else {
          this.$message({
            type: "error",
            message: "请正确填写注册信息",
            showClose: true
          });
          return false;
        }
      });
    },
    changeLock() {
      this.lock === "lock" ? (this.lock = "unlock") : (this.lock = "lock");
    }
  }
};
</script>

<style scoped lang="scss">
.login-register-box {
  background: #fff;
  height: 100vh;
  .login-box {
    width: 40vw;
    position: absolute;
    left: 50%;
    margin-left: -22vw;
    top: 5vh;
    .logo {
      height: 14vh;
      width: 49vh;
    }
    .main {
      min-width: 260px;
      width: 358px;
      margin: 0 auto;
    }
  }
}
</style>