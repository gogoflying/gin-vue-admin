<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addCompany" type="primary">新增企业</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column type="expand">
        <template scope="scope">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="企业logo:">
              <span>
                <img :src="scope.row.enterprise_logo" />
              </span>
            </el-form-item>
            <el-form-item label="企业资质:">
              <span>
                <img :src="scope.row.enterprise_img" />
              </span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column label="id" min-width="60" prop="ID"></el-table-column>
      <el-table-column label="企业名称" min-width="150" prop="enterprise_name"></el-table-column>
      <el-table-column label="注册地址" min-width="150" prop="enterprise_address"></el-table-column>
      <el-table-column label="企业规模" min-width="100" prop="enterprise_scale"></el-table-column>
      <el-table-column label="企业类型" min-width="100" prop="enterprise_type"></el-table-column>
      <el-table-column label="企业热度" min-width="100" prop="enterprise_hot"></el-table-column>
      <el-table-column label="企业性质" min-width="100" prop="industry_type"></el-table-column>
      <el-table-column label="企业描述" min-width="150" prop="enterprise_desc"></el-table-column>
      <el-table-column label="所在城市" min-width="100" prop="city_id"></el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editEnterprise(scope.row)" size="small" type="text">编辑</el-button>
          <el-button @click="deleteEnterprise(scope.row)" size="small" type="text">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :visible.sync="addCompanyDialog" custom-class="user-dialog" title="新增企业">
      <el-form :rules="rules" ref="enterpriseForm" :model="enterpriseInfo">
        <el-form-item label="企业名称" label-width="80px" prop="enterprise_name">
          <el-input v-model="enterpriseInfo.enterprise_name"></el-input>
        </el-form-item>
        <el-form-item label="注册地址" label-width="80px" prop="enterprise_address">
          <el-input v-model="enterpriseInfo.enterprise_address"></el-input>
        </el-form-item>
        <el-form-item label="企业规模" label-width="80px" prop="enterprise_scale">
          <el-input v-model.number="enterpriseInfo.enterprise_scale"></el-input>
        </el-form-item>
        <el-form-item label="企业类型" label-width="80px" prop="enterprise_type">
          <el-input v-model.number="enterpriseInfo.enterprise_type"></el-input>
        </el-form-item>
        <el-form-item label="企业热度" label-width="80px" prop="enterprise_hot">
          <el-input v-model.number="enterpriseInfo.enterprise_hot"></el-input>
        </el-form-item>
        <el-form-item label="企业性质" label-width="80px" prop="industry_type">
          <el-select placeholder="请选择企业性质" v-model="enterpriseInfo.industry_type">
            <el-option
              :key="industry.name"
              :label="industry.name"
              :value="industry.id"
              v-for="industry in industrytype"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="企业描述" label-width="80px" prop="enterprise_desc">
          <el-input v-model="enterpriseInfo.enterprise_desc"></el-input>
        </el-form-item>
        <el-form-item label="所在城市" label-width="80px" prop="city_id">
          <el-select placeholder="请选择注册城市" v-model="enterpriseInfo.city_id">
            <el-option
              :key="city.name"
              :label="city.name"
              :value="city.id"
              v-for="city in cityinfo"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="企业logo" label-width="80px" prop="enterprise_logo">
          <el-upload
            :headers="{'x-token':token}"
            :on-success="handleAvatarLogoSuccess"
            :show-file-list="false"
            :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
            class="avatar-uploader"
            name="file"
            :beforeUpload="beforeAvatarUpload"
          >
            <img
              :src="enterpriseInfo.enterprise_logo"
              class="avatar"
              v-if="enterpriseInfo.enterprise_logo"
            />
            <i class="el-icon-plus avatar-uploader-icon" v-else></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="企业资质" label-width="80px" prop="enterprise_img">
          <el-upload
            :headers="{'x-token':token}"
            :on-success="handleAvatarImgSuccess"
            :show-file-list="false"
            :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
            class="avatar-uploader"
            name="file"
            :beforeUpload="beforeAvatarUpload"
          >
            <img
              :src="enterpriseInfo.enterprise_img"
              class="avatar"
              v-if="enterpriseInfo.enterprise_img"
            />
            <i class="el-icon-plus avatar-uploader-icon" v-else></i>
          </el-upload>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddCompanyDialog">取 消</el-button>
        <el-button @click="enterAddCompanyDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {
  getEnterpriseInfoList,
  createEnterpriseInfo
} from "@/api/enterpriseinfo";
import infoList from "@/components/mixins/infoList";
import { mapGetters } from "vuex";
export default {
  name: "Enterpriseinfo",
  mixins: [infoList],
  data() {
    return {
      listApi: getEnterpriseInfoList,
      listKey: "result",
      path: path,
      authOptions: [],
      addCompanyDialog: false,
      enterpriseInfo: {
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
      rules: {
        enterprise_name: [
          {
            required: true,
            message: "请输入与营业执照一致的企业名称",
            trigger: "blur"
          }
        ],
        enterprise_address: [
          { required: true, message: "请输入联系人全名", trigger: "blur" }
        ],
        enterprise_type: [
          { required: true, message: "请输入电话", trigger: "blur" }
        ],
        enterprise_desc: [
          { required: true, message: "请输入企业邮箱", trigger: "blur" }
        ]
      }
    };
  },
  computed: {
    ...mapGetters("user", ["token"])
  },
  methods: {
    async enterAddCompanyDialog() {
      this.$refs.enterpriseForm.validate(async valid => {
        if (valid) {
          const res = await createEnterpriseInfo(this.enterpriseInfo);
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          }
          await this.getTableData();
          this.closeAddCompanyDialog();
        }
      });
    },
    closeAddCompanyDialog() {
      this.enterpriseInfo = {
        enterprise_name: "",
        enterprise_address: "",
        enterprise_scale: "",
        enterprise_type: "",
        enterprise_hot: "",
        industry_type: "",
        enterprise_desc: "",
        city_id: "",
        enterprise_logo: "",
        enterprise_img: ""
      };
      this.addCompanyDialog = false;
    },
    addCompany() {
      this.addCompanyDialog = true;
    },
    handleAvatarLogoSuccess(res) {
      this.enterpriseInfo.enterprise_logo = res.data.file.url;
    },
    handleAvatarImgSuccess(res) {
      this.enterpriseInfo.enterprise_img = res.data.file.url;
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
      return (extension || extension2) && isLt50KB
    }
  }
};
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}

.user-dialog {
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
}

.demo-table-expand {
  font-size: 0;
}
.demo-table-expand label {
  width: 150px;
  color: #4e5868;
}
.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 100%;
}
</style>