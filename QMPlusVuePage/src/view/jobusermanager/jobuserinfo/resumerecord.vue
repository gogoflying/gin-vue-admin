<template>
  <div>
    <div class="search-panel clear">
      <el-form :model="searchInfo">
        <div class="search-condition left">
          <el-row :gutter="10">
            <el-col :span="12" :xs="24">
              <el-form-item label="时间" label-width="80px">
                <el-date-picker
                  v-model="searchInfo.dateRange"
                  type="daterange"
                  align="right"
                  unlink-panels
                  range-separator="至"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
                ></el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="6" :xs="24" :offset="1">
              <el-button @click="onSearch" type="primary">搜索</el-button>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="24" :xs="12" :offset="1">
              <el-radio-group @change="onSearch" v-model="searchInfo.resume_status">
                <el-radio :label="-1">全部</el-radio>
                <el-radio :label="0">未读</el-radio>
                <el-radio :label="1">已读</el-radio>
                <el-radio :label="2">有意向</el-radio>
                <el-radio :label="3">约面试</el-radio>
                <el-radio :label="4">不合适</el-radio>
                <el-radio :label="5">待入职</el-radio>
                <el-radio :label="6">已入职</el-radio>
                <el-radio :label="7">已拒绝</el-radio>
              </el-radio-group>
            </el-col>
          </el-row>
        </div>
      </el-form>
      <div></div>
    </div>

    <el-row class="grid-warp">
      <el-col :span="24">
        <div class="content">
          <el-table :data="tableData" style="width: 100%" max-height="370">
            <el-table-column fixed label="ID" min-width="60" type="index" :index="indexMethod"></el-table-column>
            <el-table-column fixed prop="job_info.enterprise_name" label="公司名字" min-width="100"></el-table-column>
            <el-table-column fixed prop="job_info.p_name" label="职位" align="center" min-width="100"></el-table-column>
            <el-table-column
              prop="CreatedAt"
              label="时间"
              class="labelStyle"
              align="center"
              min-width="150"
              :formatter="dateFormat"
            ></el-table-column>
            <el-table-column
              prop="resume_status"
              label="状态"
              class="labelStyle"
              align="center"
              min-width="150"
            >
              <template slot-scope="scope">
                <span v-if="scope.row.resume_status == 0">未读</span>
                <span v-if="scope.row.resume_status == 1">已读</span>
                <span v-if="scope.row.resume_status == 2">有意向</span>
                <span v-if="scope.row.resume_status == 3">约面试</span>
                <span v-if="scope.row.resume_status == 4">不合适</span>
                <span v-if="scope.row.resume_status == 5">待入职</span>
                <span v-if="scope.row.resume_status == 6">已入职</span>
                <span v-if="scope.row.resume_status == 7">已拒绝</span>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-col>
    </el-row>
    <div class="bottom-handle">
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
  </div>
</template>
<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import { findJobUserById, getUserResumeList } from "@/api/jobuser";
import infoList from "@/components/mixins/infoList";
import { formatTimeToStr } from "@/utils/data";
export default {
  name: "Resumelist",
  mixins: [infoList],
  data() {
    return {
      listApi: getUserResumeList,
      listKey: "resumeList",
      path: path
    };
  },
  mounted() {
    if (this.$route.query && this.$route.query.id) {
      const id = this.$route.query && this.$route.query.id;
      const row = {
        ID: Number(id)
      };
      this.searchInfo.resume_status = -1;
      findJobUserById(row).then(res => {
        if (res.success) {
          this.searchInfo.openid = res.data.reusers.openid;
          this.onSearch();
        }
      });
    }
  },
  methods: {
    indexMethod(index) {
      return index + 1 + (this.page - 1) * this.pageSize;
    },
    dateFormat(row) {
      if (row.CreatedAt != null && row.CreatedAt != 0) {
        var date = new Date(row.CreatedAt);
        return formatTimeToStr(date, "yyyy-MM-dd");
      } else {
        return "";
      }
    },
    //条件搜索前端看此方法
    onSearch() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    }
  }
};
</script>

<style lang="scss" scoped>
.search-panel {
  width: 900px;
  margin-bottom: 20px;
  .search-condition {
    padding-right: 20px;
  }
}
.bottom-handle {
  margin-top: 10px;
}
</style>