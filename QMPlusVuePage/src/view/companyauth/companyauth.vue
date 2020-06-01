<template>
  <div>
    <el-form
      :rules="rules"
      ref="authEnform"
      :model="authform"
      label-width="80px"
      :disabled="authform.status == 1"
    >
      <el-form-item
        label="提示信息"
        prop="errors"
        v-show="authform.errors != '' && authform.status == 0"
      >
        <span style="color:red">{{authform.errors}}</span>
      </el-form-item>
      <el-form-item label="企业名称" prop="enterprise_name">
        <el-input v-model="authform.enterprise_name" style="width:50%;"></el-input>
      </el-form-item>
      <el-form-item label="企业规模" prop="enterprise_scale">
        <el-input
          type="enterprise_scale"
          v-model.number="authform.enterprise_scale"
          style="width:50%;"
        ></el-input>
      </el-form-item>
      <el-form-item label="企业性质" prop="enterprise_type">
        <el-select @change="selectEnpType" placeholder="请选择企业性质" v-model="authform.enterprise_type">
          <el-option
            :key="enptype.name"
            :label="enptype.name"
            :value="enptype.ID"
            v-for="enptype in enterprisetype"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="行业类别" prop="industry_type">
        <el-select @change="selectEnpIndust" placeholder="请选择行业类别" v-model="authform.industry_type">
          <el-option
            :key="industry.name"
            :label="industry.name"
            :value="industry.ID"
            v-for="industry in industrytype"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="企业描述" prop="enterprise_desc">
        <quill-editor
          class="ql-editor"
          ref="myQuillEditor"
          v-model="authform.enterprise_desc"
          :disabled="authform.status == 1"
        ></quill-editor>
      </el-form-item>
      <el-form-item label="所在城市" prop="city_id">
        <el-select @change="selectCityName" placeholder="请选择注册城市" v-model="authform.city_id">
          <el-option :key="city.name" :label="city.name" :value="city.ID" v-for="city in cityinfo"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="企业logo" prop="enterprise_logo">
        <el-upload
          :headers="{'x-token':token}"
          :on-success="handleAvatarLogoSuccess"
          :show-file-list="false"
          :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
          class="avatar-uploader"
          name="file"
          :before-upload="beforeAvatarUploadLogo"
        >
          <img :src="authform.enterprise_logo" class="avatar" v-if="authform.enterprise_logo" />
          <i class="el-icon-plus avatar-uploader-icon" v-else></i>
        </el-upload>
      </el-form-item>
      <el-form-item label="企业主图" prop="enterprise_img">
        <el-upload
          :headers="{'x-token':token}"
          :on-success="handleAvatarImgSuccess"
          :show-file-list="false"
          :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
          class="avatar-uploader"
          name="file"
          :before-upload="beforeAvatarUploadImg"
        >
          <img :src="authform.enterprise_img" class="avatar" v-if="authform.enterprise_img" />
          <i class="el-icon-plus avatar-uploader-icon" v-else></i>
        </el-upload>
      </el-form-item>
      <el-form-item label="企业资质" prop="enterprise_qfc">
        <el-upload
          :headers="{'x-token':token}"
          :on-success="handleAvatarQfcSuccess"
          :show-file-list="false"
          :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
          class="avatar-uploader"
          name="file"
          :before-upload="beforeAvatarUploadQfc"
        >
          <img :src="authform.enterprise_qfc" class="avatar" v-if="authform.enterprise_qfc" />
          <i class="el-icon-plus avatar-uploader-icon" v-else></i>
        </el-upload>
      </el-form-item>
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
  findEnterpriseInfo,
  getEnterpriseOptions
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
        enterprise_type_info: "",
        industry_type: null,
        enterprise_indust_info: "",
        enterprise_desc: "",
        city_id: null,
        enterprise_logo: "",
        enterprise_img: "",
        enterprise_qfc: "",
        errors: ""
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
          { required: true, message: "请输入企业规模", trigger: "blur" },
          { type: "number", message: "规模必须为数字", trigger: "blur" },
          {
            type: "number",
            min: 0,
            message: "规模必须为大于零",
            trigger: "blur"
          }
        ],
        enterprise_type: [
          { required: true, message: "请输入企业类型", trigger: "blur" }
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
      enterprisetype: [],
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
      path: path
    };
  },
  computed: {
    ...mapGetters("user", ["enPriseId", "token"])
  },
  methods: {
    selectCityName(val) {
      var selectedItem = {};
      selectedItem = this.cityinfo.find(item => {
        return item.ID === val;
      });
      this.authform.enterprise_address = selectedItem.name;
    },
    selectEnpType(val) {
      var selectedItem = {};
      selectedItem = this.enterprisetype.find(item => {
        return item.ID === val;
      });
      this.authform.enterprise_type_info = selectedItem.name;
    },
    selectEnpIndust(val) {
      var selectedItem = {};
      selectedItem = this.industrytype.find(item => {
        return item.ID === val;
      });
      this.authform.enterprise_indust_info = selectedItem.name;
    },
    filterMethod(query, item) {
      return item.pinyin.indexOf(query) > -1;
    },
    async onSubmit() {
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
        }
      });
    },
    handleAvatarLogoSuccess(res) {
      this.authform.enterprise_logo = res.data.file.url;
    },
    handleAvatarImgSuccess(res) {
      this.authform.enterprise_img = res.data.file.url;
    },
    handleAvatarQfcSuccess(res) {
      this.authform.enterprise_qfc = res.data.file.url;
    },
    beforeAvatarUploadLogo(file) {
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
    },
    beforeAvatarUploadImg(file) {
      var testmsg = file.name.substring(file.name.lastIndexOf(".") + 1);
      const extension = testmsg === "jpg";
      const extension2 = testmsg === "png";
      const isLt500KB = file.size / 1024 < 500;
      if (!extension && !extension2) {
        this.$message({
          message: "上传文件只能是 jpg、png格式!",
          type: "warning"
        });
        return false;
      }
      if (!isLt500KB) {
        this.$message({
          message: "上传文件大小不能超过 500KB!",
          type: "warning"
        });
        return false;
      }
      return (extension || extension2) && isLt500KB;
    },
    beforeAvatarUploadQfc(file) {
      var testmsg = file.name.substring(file.name.lastIndexOf(".") + 1);
      const extension = testmsg === "jpg";
      const extension2 = testmsg === "png";
      const isLt2MB = file.size / 1024 / 1024 < 2;
      if (!extension && !extension2) {
        this.$message({
          message: "上传文件只能是 jpg、png格式!",
          type: "warning"
        });
        return false;
      }
      if (!isLt2MB) {
        this.$message({
          message: "上传文件大小不能超过 2MB!",
          type: "warning"
        });
        return false;
      }
      return (extension || extension2) && isLt2MB;
    }
  },
  created() {
    getEnterpriseOptions().then(res => {
      if (res.success) {
        this.enterprisetype = res.data.enterprisetype;
        this.industrytype = res.data.industrytype;
        this.cityinfo = res.data.cityinfo;
      } else {
        this.enterprisetype = [];
        this.industrytype = [];
        this.cityinfo = [];
      }
    });
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
.ql-container {
  height: 200px;
}
.ql-editor strong {
  font-style: normal;
  font-weight: bold;
}
.ql-editor em {
  font-style: italic;
}
</style>