<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addEnterprise" type="primary" icon="el-icon-plus">新增企业</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column type="expand">
        <template slot-scope="scope">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="企业logo:">
              <span v-if="scope.row.enterprise_logo">
                <img :src="scope.row.enterprise_logo" height="500" width="500" />
              </span>
              <span v-else style="color:red">企业logo未上传</span>
            </el-form-item>
            <el-form-item label="企业主图:">
              <span v-if="scope.row.enterprise_img">
                <img :src="scope.row.enterprise_img" height="500" width="500" />
              </span>
              <span v-else style="color:red">企业主图未上传</span>
            </el-form-item>
            <el-form-item label="企业资质:">
              <span v-if="scope.row.enterprise_qfc">
                <img :src="scope.row.enterprise_qfc" height="500" width="500" />
              </span>
              <span v-else style="color:red">企业资质未上传</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <!-- <el-table-column label="id" min-width="60" prop="ID"></el-table-column> -->
      <el-table-column label="企业名称" min-width="150" prop="enterprise_name"></el-table-column>
      <el-table-column label="企业性质" min-width="100" prop="enterprise_type_info"></el-table-column>
      <el-table-column label="行业类别" min-width="100" prop="enterprise_indust_info"></el-table-column>
      <el-table-column label="所在城市" min-width="100" prop="enterprise_address"></el-table-column>
      <el-table-column label="企业规模" min-width="100" prop="enterprise_scale"></el-table-column>
      <el-table-column label="职位数量" min-width="80" prop="job_count"></el-table-column>
      <el-table-column label="企业审核" min-width="200" prop="status">
        <template slot-scope="scope">
          <el-switch
            @change="changestatus(scope.row)"
            active-text="已审核"
            inactive-text="未审核"
            :active-value="1"
            :inactive-value="0"
            v-model.number="scope.row.status"
          ></el-switch>
        </template>
      </el-table-column>
      <!-- <el-table-column label="企业热度" min-width="100" prop="enterprise_hot"></el-table-column> -->
      <!-- <el-table-column label="企业描述" min-width="150" prop="enterprise_desc" :show-overflow-tooltip="true"></el-table-column> -->

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

    <el-dialog
      :visible.sync="addCompanyDialog"
      custom-class="user-dialog"
      @close="closeAddCompanyDialog"
      title="新增企业"
    >
      <el-form :rules="rules" ref="enterpriseForm" :model="enterpriseInfo">
        <el-form-item label="企业名称" label-width="80px" prop="enterprise_name">
          <el-input v-model="enterpriseInfo.enterprise_name"></el-input>
        </el-form-item>
        <el-form-item label="企业规模" label-width="80px" prop="enterprise_scale">
          <el-input v-model.number="enterpriseInfo.enterprise_scale"></el-input>
        </el-form-item>
        <el-form-item label="企业性质" label-width="80px" prop="enterprise_type">
          <el-select
            @change="selectEnpType"
            placeholder="请选择企业性质"
            v-model="enterpriseInfo.enterprise_type"
          >
            <el-option
              :key="enptype.name"
              :label="enptype.name"
              :value="enptype.ID"
              v-for="enptype in enterprisetype"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="行业类别" label-width="80px" prop="industry_type">
          <el-select
            @change="selectEnpIndust"
            placeholder="请选择行业类别"
            v-model="enterpriseInfo.industry_type"
          >
            <el-option
              :key="industry.name"
              :label="industry.name"
              :value="industry.ID"
              v-for="industry in industrytype"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="企业热度" label-width="80px" prop="enterprise_hot">
          <el-input v-model.number="enterpriseInfo.enterprise_hot"></el-input>
        </el-form-item>
        <el-form-item label="企业描述" label-width="80px" prop="enterprise_desc">
          <quill-editor
            class="ql-editor"
            ref="myQuillEditor"
            v-model="enterpriseInfo.enterprise_desc"
          ></quill-editor>
        </el-form-item>
        <el-form-item label="所在城市" label-width="80px" prop="city_id">
          <el-select
            @change="selectCityName"
            placeholder="请选择注册城市"
            v-model="enterpriseInfo.city_id"
          >
            <el-option
              :key="city.name"
              :label="city.name"
              :value="city.ID"
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
            :before-upload="beforeAvatarUpload"
          >
            <img
              :src="enterpriseInfo.enterprise_logo"
              class="avatar"
              v-if="enterpriseInfo.enterprise_logo"
            />
            <i class="el-icon-plus avatar-uploader-icon" v-else></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="企业主图" label-width="80px" prop="enterprise_img">
          <el-upload
            :headers="{'x-token':token}"
            :on-success="handleAvatarImgSuccess"
            :show-file-list="false"
            :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
            class="avatar-uploader"
            name="file"
            :before-upload="beforeAvatarUpload"
          >
            <img
              :src="enterpriseInfo.enterprise_img"
              class="avatar"
              v-if="enterpriseInfo.enterprise_img"
            />
            <i class="el-icon-plus avatar-uploader-icon" v-else></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="企业资质" label-width="80px" prop="enterprise_qfc">
          <el-upload
            :headers="{'x-token':token}"
            :on-success="handleAvatarQfcSuccess"
            :show-file-list="false"
            :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
            class="avatar-uploader"
            name="file"
            :before-upload="beforeAvatarUpload"
          >
            <img
              :src="enterpriseInfo.enterprise_qfc"
              class="avatar"
              v-if="enterpriseInfo.enterprise_qfc"
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
  createEnterpriseInfo,
  updateEnterpriseInfo,
  findEnterpriseInfo,
  deleteEnterpriseInfo,
  getEnterpriseOptions
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
      addCompanyDialog: false,
      isEdit: false,
      title: "",
      enterpriseInfo: {
        enterprise_name: "",
        enterprise_address: "",
        enterprise_scale: null,
        enterprise_type: null,
        enterprise_type_info: "",
        enterprise_hot: null,
        industry_type: null,
        enterprise_indust_info: "",
        enterprise_desc: "",
        city_id: null,
        enterprise_logo: "",
        enterprise_img: "",
        enterprise_qfc: ""
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
          { required: true, message: "请输入企业性质", trigger: "blur" }
        ],
        industry_type: [
          { required: true, message: "请输入行业类别", trigger: "blur" }
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
          { required: true, message: "请上传企业主图", trigger: "blur" }
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
      ]
    };
  },
  computed: {
    ...mapGetters("user", ["token"]),
    editor() {
      return this.$refs.myQuillEditor.quill;
    }
  },
  methods: {
    async changestatus(row) {
      const res = await updateEnterpriseInfo(row);
      if (res.success) {
        this.$message({ type: "success", message: "状态设置成功" });
      }
    },
    selectCityName(val) {
      var selectedItem = {};
      selectedItem = this.cityinfo.find(item => {
        return item.ID === val;
      });
      this.enterpriseInfo.enterprise_address = selectedItem.name;
    },
    selectEnpType(val) {
      var selectedItem = {};
      selectedItem = this.enterprisetype.find(item => {
        return item.ID === val;
      });
      this.enterpriseInfo.enterprise_type_info = selectedItem.name;
    },
    selectEnpIndust(val) {
      var selectedItem = {};
      selectedItem = this.industrytype.find(item => {
        return item.ID === val;
      });
      this.enterpriseInfo.enterprise_indust_info = selectedItem.name;
    },
    onEditorBlur() {}, // 失去焦点事件
    onEditorFocus() {}, // 获得焦点事件
    onEditorChange() {}, // 内容改变事件
    async enterAddCompanyDialog() {
      this.$refs.enterpriseForm.validate(async valid => {
        if (valid) {
          let res;
          if (this.isEdit) {
            res = await updateEnterpriseInfo(this.enterpriseInfo);
          } else {
            res = await createEnterpriseInfo(this.enterpriseInfo);
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
        enterprise_img: "",
        enterprise_qfc: ""
      };
      this.addCompanyDialog = false;
    },
    //新增企业
    addEnterprise() {
      this.title = "新增企业";
      this.isEdit = false;
      this.addCompanyDialog = true;
    },
    //编辑企业
    async editEnterprise(row) {
      this.title = "编辑企业";
      const res = await findEnterpriseInfo(row);
      this.enterpriseInfo = res.data.reinfo;
      this.isEdit = true;
      this.addCompanyDialog = true;
    },
    //删除企业
    deleteEnterprise(row) {
      this.$confirm("此操作将永久删除该企业信息, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteEnterpriseInfo(row);
          if (res.success) {
            this.$message({
              type: "success",
              message: "删除成功!"
            });
            this.getTableData();
          }
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    },
    handleAvatarLogoSuccess(res) {
      this.enterpriseInfo.enterprise_logo = res.data.file.url;
    },
    handleAvatarImgSuccess(res) {
      this.enterpriseInfo.enterprise_img = res.data.file.url;
    },
    handleAvatarQfcSuccess(res) {
      this.enterpriseInfo.enterprise_qfc = res.data.file.url;
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
  },
  async created() {
    const res = await getEnterpriseOptions();
    this.enterprisetype = res.data.enterprisetype;
    this.industrytype = res.data.industrytype;
    this.cityinfo = res.data.cityinfo;
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
<style lang="scss">
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