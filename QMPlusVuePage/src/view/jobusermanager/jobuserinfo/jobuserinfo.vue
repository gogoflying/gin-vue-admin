<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="手机号">
          <el-input placeholder="手机号" v-model="searchInfo.mobile"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="用户标识" min-width="150" prop="openid"></el-table-column>
      <el-table-column label="手机号" min-width="150" prop="mobile"></el-table-column>
      <el-table-column label="是否存在简历" min-width="150" prop="is_resume">
        <template slot-scope="scope">
          <el-radio-group @change="changestatus(scope.row)" v-model="scope.row.is_resume">
            <el-radio :label="0">否</el-radio>
            <el-radio :label="1">是</el-radio>
          </el-radio-group>
        </template>
      </el-table-column>
      <el-table-column label="用户状态" min-width="250" prop="status">
        <template slot-scope="scope">
          <el-switch
            @change="changestatus(scope.row)"
            active-text="开启"
            inactive-text="禁用"
            :active-value="1"
            :inactive-value="0"
            v-model.number="scope.row.status"
          ></el-switch>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <!-- <el-button @click="viewResume(scope.row)" size="small" type="text">查看</el-button> -->
          <router-link :to="{name:'editresume', query: { id: scope.row.ID }}">
            <el-button type="primary" size="small" icon="el-icon-edit">编辑</el-button>
          </router-link>
          <router-link :to="{name:'userpreview', query: { id: scope.row.ID }}">
            <el-button type="primary" size="small" icon="el-icon-view">查看</el-button>
          </router-link>
          <router-link :to="{name:'resumelist', query: { id: scope.row.ID }}">
            <el-button type="primary" size="small" icon="el-icon-view">投递记录</el-button>
          </router-link>
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
import { getUsersList, updateUsers } from "@/api/jobuser";
import infoList from "@/components/mixins/infoList";
export default {
  name: "Jobuser",
  mixins: [infoList],
  data() {
    return {
      listApi: getUsersList,
      listKey: "userList",
      path: path
    };
  },
  methods: {
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    async changestatus(row) {
      const res = await updateUsers(row);
      if (res.success) {
        this.$message({ type: "success", message: "状态设置成功" });
      }
    }
  }
};
</script>