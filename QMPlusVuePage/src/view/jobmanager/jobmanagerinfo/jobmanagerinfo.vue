<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="工作城市">
          <el-select
            clearable
            @clear="clearOptionCity"
            placeholder="请选择工作城市"
            v-model.number="searchInfo.p_city_id"
          >
            <el-option
              :key="city.name"
              :label="city.name"
              :value="city.ID"
              v-for="city in cityinfo"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-show="enPriseId == 0" label="企业名称">
          <el-input placeholder="企业名称" clearable v-model="searchInfo.enterprise_name"></el-input>
        </el-form-item>
        <el-form-item label="职位名称" label-width="85px">
          <el-input placeholder="职位名称" clearable v-model="searchInfo.p_name"></el-input>
        </el-form-item>
        <el-form-item label="工作年限" prop="p_edujy_id">
          <el-select
            placeholder="请选择工作年限"
            v-model="searchInfo.p_edujy_id"
            clearable
            @clear="clearOptionJobyear"
          >
            <el-option
              :key="jobyear.name"
              :label="jobyear.name"
              :value="jobyear.ID"
              v-for="jobyear in jobyears"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="最低学历" prop="p_education_id">
          <el-select
            placeholder="请选择最低学历"
            v-model="searchInfo.p_education_id"
            clearable
            @clear="clearOptionJobedu"
          >
            <el-option
              :key="jobedu.name"
              :label="jobedu.name"
              :value="jobedu.ID"
              v-for="jobedu in jobedus"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工作类型" prop="p_type_id">
          <el-select
            placeholder="请选择工作类型"
            v-model.number="searchInfo.p_type_id"
            clearable
            @clear="clearOptionPType"
          >
            <el-option
              :key="jobtype.name"
              :label="jobtype.name"
              :value="jobtype.ID"
              v-for="jobtype in jobtypes"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="职位状态" prop="p_status">
          <el-select
            placeholder="请选择职位状态"
            v-model.number="searchInfo.p_status"
            clearable
            @clear="clearOptionJobstatus"
          >
            <el-option
              :key="jobstatus.name"
              :label="jobstatus.name"
              :value="jobstatus.ID"
              v-for="jobstatus in jobstatuss"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="是否置顶" prop="p_top">
          <el-select
            placeholder="请选择是否置顶"
            v-model.number="searchInfo.p_top"
            clearable
            @clear="clearOptionPTop"
          >
            <el-option
              :key="jobtop.name"
              :label="jobtop.name"
              :value="jobtop.ID"
              v-for="jobtop in jobtops"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button @click="onSubmit" type="primary" icon="el-icon-search">查询</el-button>
        </el-form-item>
        <el-form-item>
          <router-link :to="{name:'newjobinfo'}">
            <el-button type="primary" size="small" icon="el-icon-plus">新增职位</el-button>
          </router-link>
        </el-form-item>
      </el-form>
    </div>

    <el-table
      :data="tableData"
      border
      stripe
      @selection-change="handleSelectionChange"
      ref="multipleTable"
      style="width: 100%"
      tooltip-effect="dark"
    >
      <!-- <el-table-column type="selection" min-width="55"></el-table-column> -->
      <el-table-column label="id" min-width="60" prop="ID" ></el-table-column>
      <el-table-column label="职位名称" min-width="100" prop="p_name"></el-table-column>
      <el-table-column label="企业名称" min-width="150" prop="enterprise_name" v-if="enPriseId == 0"></el-table-column>
      <!-- <el-table-column label="薪资上限" min-width="100" prop="p_salary_high"></el-table-column>
      <el-table-column label="薪资下限" min-width="100" prop="p_salary_low"></el-table-column>
      <el-table-column label="工作地点纬度" min-width="150" prop="p_latitude"></el-table-column>
      <el-table-column label="工作地点经度" min-width="150" prop="p_longitude"></el-table-column>-->
      <el-table-column label="招聘人数" min-width="80" prop="p_count"></el-table-column>
      <el-table-column label="工作类型" min-width="100" prop="p_type"></el-table-column>
      <el-table-column label="工作城市" min-width="100" prop="p_city"></el-table-column>
      <!-- <el-table-column label="工作地点" min-width="100" prop="p_address"></el-table-column> 
      <el-table-column label="工作年限" min-width="100" prop="p_edujy"></el-table-column>
      <el-table-column label="最低学历" min-width="150" prop="p_education"></el-table-column>-->
      <el-table-column label="状态" min-width="100" prop="p_status" :formatter="StatusFormat"></el-table-column>
      <el-table-column label="浏览数" min-width="100" prop="p_views"></el-table-column>
      <el-table-column label="投递数" min-width="100" prop="p_count_todi"></el-table-column>
      <el-table-column label="失效时间" min-width="100" prop="p_outdate" :formatter="dateFormat"></el-table-column>
      <!-- <el-table-column label="工作描述" min-width="150" prop="p_des" :show-overflow-tooltip="true"></el-table-column> -->
      <el-table-column label="操作" width="300">
        <template slot-scope="scope">
          <router-link :to="{name:'newjobinfo', query: { id: scope.row.ID }}">
            <el-button type="text" size="small" icon="el-icon-edit">编辑</el-button>
          </router-link>
          <el-button
            @click="changeTop(scope.row)"
            size="small"
            type="text"
            :icon="scope.row.p_top == 0 ? 'el-icon-top':'el-icon-bottom'"
          >{{scope.row.p_top == 0 ? '置顶':'取消置顶'}}</el-button>
          <el-button @click="flush(scope.row)" size="small" type="text" icon="el-icon-refresh">刷新</el-button>
          <el-button @click="deletejob(scope.row)" size="small" type="text" icon="el-icon-delete">删除</el-button>
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
  getJoblistListBackend,
  deleteJoblist,
  updateJoblist,
  getjoblistOptions
} from "@/api/jobmanagerinfo";
import { formatTimeToStr } from "@/utils/data";
import { mapGetters } from "vuex";
//import { getCitynameList } from "@/api/cityname";
//import { getEnterpriseAllInfo } from "@/api/enterpriseinfo";
import infoList from "@/components/mixins/infoList";
export default {
  name: "Jobmanagerinfo",
  mixins: [infoList],
  data() {
    return {
      listApi: getJoblistListBackend,
      listKey: "RspJoblistList",
      path: path,
      multipleSelection: [],
      enterpriseInfo: [],
      jobyears: [],
      jobsalary: [],
      jobtypes: [],
      jobedus: [],
      cityinfo: [],
      jobtops: [
        {
          ID: 0,
          name: "否"
        },
        {
          ID: 1,
          name: "是"
        }
      ],
      jobstatuss: [
        {
          ID: 0,
          name: "普通"
        },
        {
          ID: 1,
          name: "急招"
        },
        {
          ID: 2,
          name: "热门"
        },
        {
          ID: 3,
          name: "下线"
        }
      ]
      // searchInfo: {
      //   p_name: "",
      //   p_salary_id: "",
      //   p_salary: "",
      //   p_city_id: "",
      //   p_edujy: "",
      //   p_edujy_id: "",
      //   p_education: "",
      //   p_education_id: "",
      //   p_type: "",
      //   p_type_id: "",
      //   p_desc: "",
      //   p_status: 0,
      //   p_count: "",
      //   p_outdate: "",
      //   enterprise_id: "",
      //   enterprise_name: "",
      //   p_top:0,
      // },
    };
  },
  computed: {
    ...mapGetters("user", ["enPriseId"])
  },
  methods: {
    clearOptionPTop() {
      this.searchInfo.p_top = null;
    },
    clearOptionJobstatus() {
      this.searchInfo.p_status = null;
    },
    clearOptionCity() {
      this.searchInfo.p_city_id = null;
    },
    clearOptionJobyear() {
      this.searchInfo.p_edujy_id = null;
    },
    clearOptionJobedu() {
      this.searchInfo.p_education_id = null;
    },
    clearOptionPType() {
      this.searchInfo.p_type_id = null;
    },

    StatusFormat(row) {
      if (row.p_status == 1) {
        return "急招";
      } else if (row.p_status == 2) {
        return "热门";
      } else if (row.p_status == 3) {
        return "下线";
      } else {
        return "普通";
      }
    },
    dateFormat(row) {
      if (row.p_outdate != null && row.p_outdate != 0) {
        var date = new Date(row.p_outdate);
        return formatTimeToStr(date, "yyyy-MM-dd");
      } else {
        return "永久";
      }
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    // 删除职位
    deletejob(row) {
      this.$confirm("此操作将永久删除该职位, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteJoblist(row);
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
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    async flush(row) {
      const res = await updateJoblist(row);
      if (res.success) {
        this.$message({ type: "success", message: "更新成功" });
      }
      this.getTableData();
    },
    async changeEp(row) {
      var selectedItem = {};
      selectedItem = this.enterpriseInfo.find(item => {
        return item.ID === row.enterprise_id;
      });
      row.enterprise_name = selectedItem.enterprise_name;
      row.enterprise_logo = selectedItem.enterprise_logo;
      const res = await updateJoblist(row);
      if (res.success) {
        this.$message({ type: "success", message: "指派简历成功" });
      }
    },
    async changeTop(row) {
      if (row.p_top == 0) {
        row.p_top = 1;
      } else {
        row.p_top = 0;
      }
      const res = await updateJoblist(row);
      if (res.success) {
        this.$message({ type: "success", message: "切换成功" });
      }
      this.getTableData();
    }
  },
  async created() {
    // const res = await getCitynameList({ page: 1, pageSize: 999 });
    // this.cityinfo = res.data.userCityList;
    // getEnterpriseAllInfo().then(res => {
    //   if (res.success) {
    //     this.enterpriseInfo = res.data.result;
    //   } else {
    //     this.enterpriseInfo = [];
    //   }
    // });
    getjoblistOptions().then(res => {
      if (res.success) {
        this.jobyears = res.data.jbwe;
        this.jobsalary = res.data.js;
        this.jobtypes = res.data.jwt;
        this.jobedus = res.data.el;
        this.cityinfo = res.data.cityinfo;
      } else {
        this.jobyears = [];
        this.jobsalary = [];
        this.jobtypes = [];
        this.jobedus = [];
        this.cityinfo = [];
      }
    });
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
</style>