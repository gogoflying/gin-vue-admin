<template>
  <div>
    <el-form :rules="rules" :model="authform" label-width="80px" ref="authEnform">
      <el-row>
        <el-col :span="3">
          <label for>企业名称</label>
        </el-col>
        <el-col :span="10">
          <el-input v-model="authform.enterprise_name"></el-input>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="3">
          <label for>注册地址</label>
        </el-col>
        <el-col :span="10">
          <el-input v-model="authform.enterprise_address"></el-input>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="3">
          <label for>企业规模</label>
        </el-col>
        <el-col :span="10">
          <el-input v-model.number="authform.enterprise_scale"></el-input>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="3">
          <label for>企业类型</label>
        </el-col>
        <el-col :span="10">
          <el-input v-model.number="authform.enterprise_type"></el-input>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="3">
          <label for>企业热度</label>
        </el-col>
        <el-col :span="10">
          <el-input v-model.number="authform.enterprise_hot"></el-input>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="3">
          <label for>企业性质</label>
        </el-col>
        <el-col :span="10">
          <el-select placeholder="请选择企业性质" v-model="authform.industry_type">
            <el-option
              :key="industry.name"
              :label="industry.name"
              :value="industry.id"
              v-for="industry in industrytype"
            ></el-option>
          </el-select>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="3">
          <label for>企业描述</label>
        </el-col>
        <el-col :span="10">
          <el-input v-model="authform.enterprise_desc"></el-input>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="3">
          <label for>所在城市</label>
        </el-col>
        <el-col :span="10">
          <el-select placeholder="请选择注册城市" v-model="authform.city_id">
            <el-option
              :key="city.name"
              :label="city.name"
              :value="city.id"
              v-for="city in cityinfo"
            ></el-option>
          </el-select>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="3">
          <label for>企业logo</label>
        </el-col>
        <el-col :span="10">
          <el-upload
            :headers="{'x-token':token}"
            :on-success="handleAvatarLogoSuccess"
            :show-file-list="false"
            :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
            class="avatar-uploader"
            name="file"
            :beforeUpload="beforeAvatarUpload"
          >
            <img :src="authform.enterprise_logo" class="avatar" v-if="authform.enterprise_logo" />
            <i class="el-icon-plus avatar-uploader-icon" v-else></i>
          </el-upload>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="3">
          <label for>企业资质</label>
        </el-col>
        <el-col :span="10">
          <el-upload
            :headers="{'x-token':token}"
            :on-success="handleAvatarImgSuccess"
            :show-file-list="false"
            :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
            class="avatar-uploader"
            name="file"
            :beforeUpload="beforeAvatarUpload"
          >
            <img :src="authform.enterprise_img" class="avatar" v-if="authform.enterprise_img" />
            <i class="el-icon-plus avatar-uploader-icon" v-else></i>
          </el-upload>
        </el-col>
      </el-row>
      <el-row type="flex" justify="center">
        <el-col :span="16">
          <el-button @click="onSubmit" type="primary">立即认证</el-button>
          <el-button type="primary">取消</el-button>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>
<script>
import { mapGetters } from "vuex";
const path = process.env.VUE_APP_BASE_API;
import {
  createEnterpriseInfo,
  updateEnterpriseInfo,
  findEnterpriseInfo
} from "@/api/enterpriseinfo";
export default {
  name: "companyauth",
  data() {
    return {
      isEdit: false,
      authform: {
        enterprise_name: "",
        enterprise_address: "",
        enterprise_scale: null,
        enterprise_type: null,
        enterprise_hot: null,
        industry_type: null,
        enterprise_desc: "",
        city_id: null,
        enterprise_logo: "",
        enterprise_img: ""
      },
      rules: {
        enterprise_name: [
          {
            required: true,
            message: "请输入与营业执照一致的企业名称",
            trigger: "blur"
          }
        ],
        enterprise_address: [
          { required: true, message: "请输入注册地址", trigger: "blur" }
        ],
        enterprise_scale: [
          { required: true, message: "请输入企业规模", trigger: "blur" }
        ],
        enterprise_type: [
          { required: true, message: "请输入企业类型", trigger: "blur" }
        ],
        enterprise_hot: [
          { required: true, message: "请输入企业热度", trigger: "blur" }
        ],
        industry_type: [
          { required: true, message: "请输入企业性质", trigger: "blur" }
        ],
        enterprise_desc: [
          { required: true, message: "请输入企业描述", trigger: "blur" }
        ],
        city_id: [
          { required: true, message: "请输入所在城市", trigger: "blur" }
        ],
        enterprise_logo: [
          { required: true, message: "请上传企业logo", trigger: "blur" }
        ],
        enterprise_img: [
          { required: true, message: "请上传企业资质", trigger: "blur" }
        ]
      },
      cityinfo: [
        {
          id: 1,
          name: "北京"
        },
        {
          id: 2,
          name: "上海"
        },
        {
          id: 3,
          name: "广州"
        }
      ],
      industrytype: [
        {
          id: 1,
          name: "国企"
        },
        {
          id: 2,
          name: "私企"
        },
        {
          id: 3,
          name: "个体"
        }
      ],
      value: [],
      path: path
    };
  },
  computed: {
    ...mapGetters("user", ["enPriseId", "token"])
  },
  methods: {
    filterMethod(query, item) {
      return item.pinyin.indexOf(query) > -1;
    },
    onSubmit() {
      this.$refs.authEnform.validate(async valid => {
        if (valid) {
          let res;
          if (this.isEdit) {
            res = await updateEnterpriseInfo(this.authform);
          } else {
            res = await createEnterpriseInfo(this.authform);
          }
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          } else {
            this.$message({ type: "error", message: "添加失败!" });
          }
          await this.getTableData();
          this.closeAddCompanyDialog();
        }
      });
    },
    handleAvatarLogoSuccess(res) {
      this.authform.enterprise_logo = res.data.file.url;
    },
    handleAvatarImgSuccess(res) {
      this.authform.enterprise_img = res.data.file.url;
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
      }
      if (!isLt50KB) {
        this.$message({
          message: "上传文件大小不能超过 50KB!",
          type: "warning"
        });
      }
      return (extension || extension2) && isLt50KB;
    }
  },
  created() {
    let row = {
      ID: this.enPriseId
    };
    findEnterpriseInfo(row).then(res => {
      if (res.success) {
        this.isEdit = true;
        this.authform = res.data.reinfo;
      } else {
        this.isEdit = false;
        this.authform = {};
      }
    });
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