<template>
  <div>
    <div class="fl-left left-mg-xs">
        <img :src="userInfo.headerImg" class="avatar" v-if="userInfo.headerImg" />
        <i class="el-icon-plus avatar-uploader-icon" v-else></i>
      <!-- <el-avatar :size="120" :src="userInfo.headerImg" shape="square"></el-avatar> -->
    </div>
    <div class="fl-left left-mg-lg">
      <div>用户名：{{userInfo.userName}}</div>
      <div v-if="enPriseId != 0">企业名称：{{ enterpriseInfo.enterprise_name }}</div>
      <div>用户角色：{{userInfo.authority&&userInfo.authority.authorityName}}</div>
    </div>
  </div>
</template>
<script>
import { mapGetters, mapMutations } from "vuex";
import { findEnterpriseInfo } from "@/api/enterpriseinfo";
const path = process.env.VUE_APP_BASE_API;
export default {
  name: "Person",
  data() {
    return {
      path: path,
      enterpriseInfo:{}
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo", "token", "enPriseId"])
  },
  methods: {
    ...mapMutations("user", ["ResetUserInfo"]),
    handleAvatarSuccess(res) {
      this.ResetUserInfo({ headerImg: res.data.user.headerImg });
    }
  },
  async created() {
    if (this.enPriseId != 0) {
      const res = await findEnterpriseInfo({ ID: this.enPriseId });
      if (res.success) {
        this.enterpriseInfo = res.data.reinfo;
        if (this.enterpriseInfo.enterprise_logo != ""){
          this.userInfo.headerImg = this.enterpriseInfo.enterprise_logo;
        }
      }
    }
  }
};
</script>
<style lang="scss">
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>