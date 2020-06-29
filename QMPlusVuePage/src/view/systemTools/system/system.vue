<template>
  <div>
    <el-form :model="config" label-width="100px" ref="form" class="system">
      <h2>系统配置</h2>
      <el-form-item label="多点登录拦截">
        <el-checkbox v-model="config.system.useMultipoint">开启</el-checkbox>
      </el-form-item>
      <el-form-item label="环境值">
        <el-input v-model="config.system.env"></el-input>
      </el-form-item>
      <el-form-item label="端口值">
        <el-input v-model="config.system.addr"></el-input>
      </el-form-item>
      <h2>jwt签名</h2>
      <el-form-item label="jwt签名">
        <el-input v-model="config.jwt.signingKey"></el-input>
      </el-form-item>
      <h2>casbin配置</h2>
      <el-form-item label="模型地址">
        <el-input v-model="config.casbinConfig.modelPath"></el-input>
      </el-form-item>
      <h2>mysql admin数据库配置</h2>
      <el-form-item label="用户名">
        <el-input v-model="config.mysqlAdmin.username"></el-input>
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="config.mysqlAdmin.password"></el-input>
      </el-form-item>
      <el-form-item label="addr">
        <el-input v-model="config.mysqlAdmin.path"></el-input>
      </el-form-item>
      <el-form-item label="dbname">
        <el-input v-model="config.mysqlAdmin.dbname"></el-input>
      </el-form-item>
      <el-form-item label="maxIdleConns">
        <el-input v-model.number="config.mysqlAdmin.maxIdleConns"></el-input>
      </el-form-item>
      <el-form-item label="maxOpenConns">
        <el-input v-model.number="config.mysqlAdmin.maxOpenConns"></el-input>
      </el-form-item>
      <el-form-item label="logMode">
        <el-checkbox v-model="config.mysqlAdmin.logMode"></el-checkbox>
      </el-form-item>
      <h2>Redis admin数据库配置</h2>
      <el-form-item label="addr">
        <el-input v-model="config.redisAdmin.addr"></el-input>
      </el-form-item>
      <el-form-item label="password">
        <el-input v-model="config.redisAdmin.password"></el-input>
      </el-form-item>
      <el-form-item label="db">
        <el-input v-model="config.redisAdmin.db"></el-input>
      </el-form-item>
      <h2>七牛密钥配置</h2>
      <el-form-item label="accessKey">
        <el-input v-model="config.qiniu.accessKey"></el-input>
      </el-form-item>
      <el-form-item label="secretKey">
        <el-input v-model="config.qiniu.secretKey"></el-input>
      </el-form-item>
      <h2>验证码配置</h2>
      <el-form-item label="keyLong">
        <el-input v-model.number="config.captcha.keyLong"></el-input>
      </el-form-item>
      <el-form-item label="imgWidth">
        <el-input v-model.number="config.captcha.imgWidth"></el-input>
      </el-form-item>
      <el-form-item label="imgHeight">
        <el-input v-model.number="config.captcha.imgHeight"></el-input>
      </el-form-item>
      <h2>日志配置</h2>
      <el-form-item label="prefix">
        <el-input v-model.number="config.log.prefix"></el-input>
      </el-form-item>
      <el-form-item label="logFile">
        <el-checkbox v-model="config.log.logFile"></el-checkbox>
      </el-form-item>
      <h2>阿里云配置</h2>
      <el-form-item label="endPoint">
        <el-input v-model="config.aliyunoss.endPoint"></el-input>
      </el-form-item>
      <el-form-item label="accessKeyId">
        <el-input v-model="config.aliyunoss.accessKeyId"></el-input>
      </el-form-item>
      <el-form-item label="accessKeySecret">
        <el-input v-model="config.aliyunoss.accessKeySecret"></el-input>
      </el-form-item>
      <h2>小程序过滤配置</h2>
      <el-form-item label="appid">
        <el-input v-model="config.weiXin.resumeapp.appid"></el-input>
      </el-form-item>
      <el-form-item label="appsecret">
        <el-input v-model="config.weiXin.resumeapp.appsecret"></el-input>
      </el-form-item>
      <el-form-item label="appid">
        <el-input v-model="config.weiXin.salaryapp.appid"></el-input>
      </el-form-item>
      <el-form-item label="appsecret">
        <el-input v-model="config.weiXin.salaryapp.appsecret"></el-input>
      </el-form-item>
      <h2>短信配置</h2>
      <el-form-item label="user">
        <el-input v-model="config.sendSms.user"></el-input>
      </el-form-item>
      <el-form-item label="passwd">
        <el-input v-model="config.sendSms.passwd"></el-input>
      </el-form-item>
      <el-form-item label="smsKind">
        <el-input v-model="config.sendSms.smsKind"></el-input>
      </el-form-item>
      <h2>邮箱配置</h2>
      <el-form-item label="serverHost">
        <el-input v-model="config.sendEmail.serverHost"></el-input>
      </el-form-item>
      <el-form-item label="serverPort">
        <el-input v-model.number="config.sendEmail.serverPort"></el-input>
      </el-form-item>
      <el-form-item label="fromEmail">
        <el-input v-model="config.sendEmail.fromEmail"></el-input>
      </el-form-item>
      <el-form-item label="fromPasswd">
        <el-input v-model="config.sendEmail.fromPasswd"></el-input>
      </el-form-item>
      <h2>微信支付</h2>
      <el-form-item label="appid">
        <el-input v-model="config.weixinPayinfo.socialsecurityapp.appid"></el-input>
      </el-form-item>
      <el-form-item label="appsecret">
        <el-input v-model.number="config.weixinPayinfo.socialsecurityapp.appsecret"></el-input>
      </el-form-item>
      <el-form-item label="mchkey">
        <el-input v-model="config.weixinPayinfo.mchkey"></el-input>
      </el-form-item>
      <el-form-item label="mchid">
        <el-input v-model="config.weixinPayinfo.mchid"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="update" type="primary">立即更新</el-button>
        <el-button @click="reload" type="primary">重启服务（开发中）</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getSystemConfig, setSystemConfig } from "@/api/system";
export default {
  name: "Config",
  data() {
    return {
      config: {
        system: {},
        jwt: {},
        casbinConfig: {},
        mysqlAdmin: {},
        redisAdmin: {},
        qiniu: {},
        captcha: {},
        log: {},
        aliyunoss: {},
        weiXin: {
          resumeapp: {},
          salaryapp: {}
        },
        sendSms: {},
        sendEmail: {},
        weixinPayinfo: {
          socialsecurityapp: {}
        }
      }
    };
  },
  async created() {
    await this.initForm();
  },
  methods: {
    async initForm() {
      const res = await getSystemConfig();
      if (res.success) {
        this.config = res.data.config;
      }
    },
    reload() {},
    async update() {
      const res = await setSystemConfig({ config: this.config });
      if (res.success) {
        this.$message({
          type: "success",
          message: "配置文件设置成功"
        });
        await this.initForm();
      }
    }
  }
};
</script>
<style lang="scss">
h2 {
  padding: 10px;
  margin: 10px 0;
  font-size: 16px;
  box-shadow: -4px 1px 3px 0px #e7e8e8;
}
</style>