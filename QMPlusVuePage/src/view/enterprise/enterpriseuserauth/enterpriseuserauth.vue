<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="企业名称">
          <el-input placeholder="请输入企业名称" v-model="searchInfo.enterprise_name"></el-input>
        </el-form-item>
        <el-form-item label="用户名">
          <el-input placeholder="请输入用户名" v-model="searchInfo.userName"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <!-- <el-form-item>
          <el-button @click="addjob" type="primary">新增职位</el-button>
        </el-form-item>-->
      </el-form>
    </div>
    <el-table :data="tableData" border stripe>
      <!-- <el-table-column label="id" min-width="60" prop="ID"></el-table-column> -->
      <el-table-column label="企业名称" min-width="150" prop="enterprise_name"></el-table-column>
      <el-table-column label="用户名" min-width="80" prop="userName"></el-table-column>
      <el-table-column label="角色" min-width="80" prop="authorityName"></el-table-column>
      <el-table-column label="审核" min-width="120"  prop="status">
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
      <el-table-column fixed="right" label="操作" width="80" v-if="false">
        <template slot-scope="scope">
          <el-button @click="deleteUserAuth(scope.row.ID)" size="small" type="text">删除</el-button>
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
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {
  getUserAuthList,
  updateUserAuth,
  deleteUserAuth
} from "@/api/userauth";
import infoList from "@/components/mixins/infoList";
export default {
  name: "Enterpriseuserauth",
  mixins: [infoList],
  data() {
    return {
      listApi: getUserAuthList,
      listKey: "userAuthsList",
      path: path
    };
  },
  methods: {
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    async changestatus(row) {
      const res = await updateUserAuth(row);
      if (res.success) {
        this.$message({ type: "success", message: "状态设置成功" });
      }
    },
    // 删除用户审核
    deleteUserAuth(ID) {
      this.$confirm("此操作将永久删除改用户审核, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteUserAuth({ ID });
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