<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="姓名">
          <el-input placeholder="姓名" clearable v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary" icon="el-icon-search">查询</el-button>
        </el-form-item>

        <el-form-item>
          <el-button @click="addSU" type="primary" icon="el-icon-plus">新增用户</el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-table :data="tableData" border stripe>
      <el-table-column label="ID" min-width="60" type="index" :index="indexMethod"></el-table-column>
      <el-table-column label="姓名" min-width="150" prop="name"></el-table-column>
      <el-table-column label="身份证号" min-width="100" prop="id_card"></el-table-column>
      <el-table-column label="手机号" min-width="100" prop="mobile"></el-table-column>

      <el-table-column label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editSU(scope.row)" icon="el-icon-edit" size="small" type="text">编辑</el-button>
          <el-button @click="deleteSU(scope.row)" icon="el-icon-delete" size="small" type="text">删除</el-button>
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
      :visible.sync="addSocialUserDialog"
      custom-class="user-dialog"
      @close="closeaddSocialUserDialog"
      :title="title"
      width="600px"
    >
      <el-form :rules="rules" ref="SocialUserForm" :model="socialuserInfo">
        <el-form-item label="姓名" label-width="80px" prop="name">
          <el-input v-model="socialuserInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="手机号" label-width="80px" prop="mobile">
          <el-input v-model="socialuserInfo.mobile"></el-input>
        </el-form-item>
        <el-form-item label="身份证号" label-width="80px" prop="id_card">
          <el-input v-model="socialuserInfo.id_card"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeaddSocialUserDialog">取 消</el-button>
        <el-button @click="enteraddSocialUserDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {
  getSocialUserList,
  createSocialUser,
  updateSocialUser,
  findSocialUser,
  deleteSocialUser
} from "@/api/socialuser";
import infoList from "@/components/mixins/infoList";
export default {
  name: "SocialUser",
  mixins: [infoList],
  data() {
    return {
      listApi: getSocialUserList,
      listKey: "socialUserList",
      path: path,
      addSocialUserDialog: false,
      isEdit: false,
      title: "",
      socialuserInfo: {
        name: "",
        mobile: "",
        id_card: ""
      },
      rules: {
        name: [
          {
            required: true,
            message: "请输入姓名",
            trigger: "blur"
          }
        ],
        mobile: [
          {
            required: true,
            message: "请输入手机号",
            trigger: "blur"
          }
        ],
        id_card: [
          { required: true, message: "请输入身份证号码", trigger: "blur" },
          {
            pattern: /(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$)/,
            message: "身份证号码格式不正确",
            trigger: "blur"
          }
        ]
      }
    };
  },
  methods: {
    indexMethod(index) {
      return index + 1 + (this.page - 1) * this.pageSize;
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    async enteraddSocialUserDialog() {
      this.$refs.SocialUserForm.validate(async valid => {
        if (valid) {
          let res;
          if (this.isEdit) {
            res = await updateSocialUser(this.socialuserInfo);
          } else {
            res = await createSocialUser(this.socialuserInfo);
          }
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          } else {
            this.$message({ type: "error", message: "添加失败!" });
          }
          await this.getTableData();
          this.closeaddSocialUserDialog();
        }
      });
    },
    closeaddSocialUserDialog() {
      this.socialuserInfo = {
        name: "",
        mobile: "",
        id_card: ""
      };
      this.$refs.SocialUserForm.resetFields();
      this.addSocialUserDialog = false;
    },
    //新增企业
    addSU() {
      this.title = "新增用户";
      this.isEdit = false;
      this.addSocialUserDialog = true;
    },
    //编辑企业
    async editSU(row) {
      this.title = "编辑用户";
      const res = await findSocialUser(row);
      this.socialuserInfo = res.data.reet;
      this.isEdit = true;
      this.addSocialUserDialog = true;
    },
    //删除企业
    deleteSU(row) {
      this.$confirm("此操作将永久删除该用户信息, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteSocialUser(row);
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
</style>