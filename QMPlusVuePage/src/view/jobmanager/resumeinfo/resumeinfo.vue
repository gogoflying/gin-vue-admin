<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="职位名称">
          <el-input placeholder="职位名称" v-model="searchInfo.p_name"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe @cell-click="editpreview">
      <el-table-column type="selection" min-width="55"></el-table-column>
      <el-table-column label="职位名称" min-width="150" prop="job_info.p_name"></el-table-column>
      <el-table-column label="手机号" min-width="150" prop="user_info.mobile"></el-table-column>
      <el-table-column label="简历状态" min-width="150">
        <template slot-scope="scope">
          <el-select
            @change="changeResumeStatus(scope.row)"
            placeholder="请选择"
            v-model="scope.row.resume_status"
          >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.id"
              v-for="item in resume_status_infos"
            ></el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button
            @click="editpreview(scope.row,scope.column)"
            type="primary"
            size="small"
            icon="el-icon-view"
          >查看</el-button>
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
import { getResumeStatusList, updateResumeStatus } from "@/api/resumestatus";
import infoList from "@/components/mixins/infoList";
export default {
  name: "Api",
  mixins: [infoList],
  data() {
    return {
      listApi: getResumeStatusList,
      listKey: "resumeStatusList",
      path: path,
      resume_status_infos: [
        {
          id: 0,
          name: "未读"
        },
        {
          id: 1,
          name: "已读"
        },
        {
          id: 2,
          name: "有意向"
        },
        {
          id: 3,
          name: "约面试"
        },
        {
          id: 4,
          name: "不合适"
        }
      ]
    };
  },
  methods: {
    async editpreview(row,column) {
      //打开简历详情
      if (column.type !== "selection") {
        const { href } = this.$router.resolve({
          name: "resumepreview",
          query: { id: row.ID }
        });
        window.open(href, "_blank");
        if (row.resume_status == 0) {
          row.resume_status = 1;
          await updateResumeStatus(row);
        }
      }
    },
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    async changeResumeStatus(row) {
      const res = await updateResumeStatus(row);
      if (res.success) {
        this.$message({ type: "success", message: "状态设置成功" });
      }
    }
  }
};
</script>