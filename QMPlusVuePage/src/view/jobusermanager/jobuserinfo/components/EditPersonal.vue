<!-- 个人信息 -->
<template>
  <div>
    <h2>
      基本信息
      <el-button
        @click="addPerson"
        size="small"
        type="text"
        style="margin-left:50%;font-size: 15px;"
      >保存</el-button>
    </h2>
    <el-form :rules="rules" ref="baseForm" :model="user_base_info" label-width="80px">
      <el-form-item label="姓名" prop="userName">
        <el-input v-model="user_base_info.userName" placeholder="请输入姓名"></el-input>
      </el-form-item>
      <el-form-item label="性别" prop="genderindex">
        <el-radio-group v-model="user_base_info.genderindex">
          <el-radio :label="0">男</el-radio>
          <el-radio :label="1">女</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="最高学历" prop="edulevelindex">
        <el-select
          @change="selectJobedu"
          placeholder="请输入学历"
          v-model="user_base_info.edulevelindex"
        >
          <el-option :key="item.name" :label="item.name" :value="item.ID" v-for="item in option.el"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="工作年限" prop="worksYearindex">
        <el-select
          @change="selectWorkYear"
          placeholder="请输入工作年限"
          v-model="user_base_info.worksYearindex"
        >
          <el-option
            :key="item.name"
            :label="item.name"
            :value="item.ID"
            v-for="item in option.jwe"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="所在城市" prop="cityindex">
        <el-select @change="selectCityName" placeholder="所在城市" v-model="user_base_info.cityindex">
          <el-option
            :key="item.name"
            :label="item.name"
            :value="item.ID"
            v-for="item in option.cityinfo"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="联系电话" prop="contact">
        <el-input v-model="user_base_info.contact" placeholder="请输入联系电话"></el-input>
      </el-form-item>
      <el-form-item label="联系邮箱" prop="email">
        <el-input v-model="user_base_info.email" placeholder="请输入联系邮箱"></el-input>
      </el-form-item>
      <el-form-item label="身份证号" prop="idcard">
        <el-input v-model="user_base_info.idcard" placeholder="请输入身份证号码"></el-input>
      </el-form-item>
      <el-form-item label="生日" prop="birthday">
        <el-date-picker
          placeholder="选择生日"
          type="date"
          :editable="false"
          format="yyyy-MM"
          value-format="yyyy-MM"
          style="width:80%;"
          v-model="user_base_info.birthday"
        ></el-date-picker>
      </el-form-item>
      <el-form-item label="头像" prop="avatarUrl">
        <el-upload
          :headers="{'x-token':token}"
          :on-success="handleAvatarSuccess"
          :show-file-list="false"
          :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
          class="avatar-uploader"
          name="file"
          :before-upload="beforeAvatarUpload"
        >
          <img :src="user_base_info.avatarUrl" class="avatar" v-if="user_base_info.avatarUrl" />
          <i class="el-icon-plus avatar-uploader-icon" v-else></i>
        </el-upload>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
const path = process.env.VUE_APP_BASE_API;
import { createUserBaseInfo, updateUserBaseInfo } from "@/api/jobuser";
import { mapGetters } from "vuex";
export default {
  props: ["user_base_info", "option", "openid"],
  data() {
    return {
      path: path,
      rules: {
        userName: [
          {
            required: true,
            message: "请输入姓名",
            trigger: "blur"
          }
        ],
        genderindex: [
          { required: true, message: "请选择年龄", trigger: "blur" }
        ],
        edulevelindex: [
          { required: true, message: "请选择学历", trigger: "blur" }
        ],
        worksYearindex: [
          { required: true, message: "请选择工作年限", trigger: "blur" }
        ],
        cityindex: [{ required: true, message: "请选择城市", trigger: "blur" }],
        contact: [
          { required: true, message: "请输入联系方式", trigger: "blur" }
        ],
        email: [
          { required: true, message: "请输入邮箱", trigger: "blur" },
          {
            type: "email",
            message: "邮箱格式不正确",
            trigger: "blur"
          }
        ],
        idcard: [
          {
            pattern: /(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$)/,
            message: "身份证号码格式不正确",
            trigger: "blur"
          }
        ],
        birthday: [{ required: true, message: "请输入生日", trigger: "blur" }],
        //avatarUrl: [{ required: true, message: "请上传头像", trigger: "blur" }]
      }
    };
  },
  computed: {
    ...mapGetters("user", ["token"])
  },
  methods: {
    selectCityName(val) {
      var selectedItem = {};
      selectedItem = this.option.cityinfo.find(item => {
        return item.ID === val;
      });
      this.user_base_info.city_name = selectedItem;
    },
    selectWorkYear(val) {
      var selectedItem = {};
      selectedItem = this.option.jwe.find(item => {
        return item.ID === val;
      });
      this.user_base_info.job_work_expire = selectedItem;
    },
    selectJobedu(val) {
      var selectedItem = {};
      selectedItem = this.option.el.find(item => {
        return item.ID === val;
      });
      this.user_base_info.edu_level = selectedItem;
    },
    async addPerson() {
      this.$refs.baseForm.validate(async valid => {
        if (valid) {
          let res;
          if (this.user_base_info.ID == null) {
            this.user_base_info.openid = this.openid;
            res = await createUserBaseInfo(this.user_base_info);
          } else {
            res = await updateUserBaseInfo(this.user_base_info);
          }
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          } else {
            this.$message({ type: "error", message: "添加失败!" });
          }
          this.$emit("freshResume");
        }
      });
    },
    handleAvatarSuccess(res) {
      this.user_base_info.avatarUrl = res.data.file.url;
    },
    beforeAvatarUpload(file) {
      var testmsg = file.name.substring(file.name.lastIndexOf(".") + 1);
      const extension = testmsg === "jpg";
      const extension2 = testmsg === "png";
      const isLt50KB = file.size / 1024 < 50;
      if (!extension && !extension2) {
        this.$message({
          message: "上传文件只能是 jpg、png格式!",
          type: "warning"
        });
        return false;
      }
      if (!isLt50KB) {
        this.$message({
          message: "上传文件大小不能超过 50KB!",
          type: "warning"
        });
        return false;
      }
      return (extension || extension2) && isLt50KB;
    }
  }
};
</script>

<style lang="scss">
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  border: 1px dashed #d9d9d9 !important;
  border-radius: 6px;
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
// .el-form-item {
//   margin-bottom: 3px;
//   margin-top:10px;
// }
</style>



