<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" label-width="80px" class="demo-form-inline">
        <el-form-item label="姓名">
          <el-input placeholder="姓名" clearable v-model="searchInfo.userName" size="small"></el-input>
        </el-form-item>
        <el-form-item label="手机号" label-width="80px">
          <el-input placeholder="手机号" clearable v-model="searchInfo.contact" size="small"></el-input>
        </el-form-item>
        <el-form-item label="期望职位" label-width="80px">
          <el-input placeholder="期望职位" clearable v-model="searchInfo.dreamposi" size="small"></el-input>
        </el-form-item>
        <el-form-item label="性别">
          <!-- <el-radio-group v-model="searchInfo.genderindex">
                <el-radio :label="0">男</el-radio>
                <el-radio :label="1">女</el-radio>
          </el-radio-group>-->
          <el-select
            placeholder="请选择性别"
            clearable
            @clear="clearOptionGender"
            v-model="searchInfo.genderindex"
          >
            <el-option label="男" :value="0"></el-option>
            <el-option label="女" :value="1"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工作年限" style="workTime">
          <el-select
            placeholder="请选择工作年限"
            clearable
            @clear="clearOptionWY"
            v-model="searchInfo.worksYearindex"
          >
            <!-- <el-option></el-option> -->
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.ID"
              v-for="item in option.jwe"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="学历">
          <el-select
            placeholder="请选择学历"
            clearable
            @clear="clearOptionEdu"
            v-model="searchInfo.edulevelindex"
          >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.ID"
              v-for="item in option.el"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工作类型">
          <el-select
            placeholder="请选择工作类型"
            clearable
            @clear="clearOptionWT"
            v-model="searchInfo.workTypeindex"
          >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.ID"
              v-for="item in option.wt"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="期望薪资">
          <el-select
            placeholder="请选择期望薪资"
            clearable
            @clear="clearOptionSalary"
            v-model="searchInfo.salaryindex"
          >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.ID"
              v-for="item in option.js"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="期望城市">
          <el-select
            placeholder="请选择期望城市"
            clearable
            @clear="clearOptionCity"
            v-model="searchInfo.cityindex"
          >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.ID"
              v-for="item in option.cityinfo"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="到岗时间">
          <el-select
            placeholder="请选择到岗时间"
            clearable
            @clear="clearOptionDT"
            v-model="searchInfo.dutyTimeindex"
          >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.ID"
              v-for="item in option.jdt"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="照片">
          <el-select
            placeholder="请选择"
            clearable
            @clear="clearOptionPhoto"
            v-model="searchInfo.hasAvatarurl"
          >
            <el-option label="无" :value="0"></el-option>
            <el-option label="有" :value="1"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary" icon="el-icon-search">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe @cell-click="gotopreview">
      <!-- <el-table-column label="用户标识" min-width="150" prop="openid"></el-table-column> -->
      <el-table-column label="ID" min-width="60" type="index"></el-table-column>
      <el-table-column label="姓名" min-width="100" prop="userName">
        <template slot-scope="scope">
          <el-link type="primary">{{scope.row.userName}}</el-link>
        </template>
      </el-table-column>
      <el-table-column label="手机号" min-width="100" prop="contact"></el-table-column>
      <el-table-column label="期望职位" min-width="100" prop="dreamposi"></el-table-column>
      <el-table-column label="性别" min-width="100" prop="genderindex" :formatter="genderFormat"></el-table-column>
      <el-table-column label="工作年限" min-width="100" prop="worksYearindex" :formatter="wyFormat"></el-table-column>
      <el-table-column label="学历" min-width="100" prop="edulevelindex" :formatter="elFormat"></el-table-column>
      <el-table-column label="工作类型" min-width="100" prop="workTypeindex" :formatter="wtFormat"></el-table-column>
      <el-table-column label="期望薪资" min-width="100" prop="salaryindex" :formatter="salaryFormat"></el-table-column>
      <el-table-column label="期望城市" min-width="100" prop="cityindex" :formatter="cityFormat"></el-table-column>
      <el-table-column label="到岗时间" min-width="100" prop="dutyTimeindex" :formatter="dtFormat"></el-table-column>
      <el-table-column label="有无简历" min-width="130" prop="is_resume">
        <template slot-scope="scope">
          <el-radio-group @change="changestatus(scope.row)" v-model="scope.row.is_resume">
            <el-radio :label="0">无</el-radio>
            <el-radio :label="1">有</el-radio>
          </el-radio-group>
        </template>
      </el-table-column>
      <el-table-column label="用户状态" min-width="150" prop="status">
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
      <el-table-column label="操作" width="250">
        <template slot-scope="scope">
          <!-- <el-button @click="viewResume(scope.row)" size="small" type="text">查看</el-button> -->
          <router-link :to="{name:'editresume', query: { id: scope.row.ID }}">
            <el-button
              type="text"
              size="small"
              :icon="scope.row.is_resume == 0 ? 'el-icon-plus':'el-icon-edit'"
            >{{scope.row.is_resume == 0 ? '新建':'编辑'}}</el-button>
          </router-link>
          <el-button @click="editpreview(scope.row)" type="text" size="small" icon="el-icon-view">查看</el-button>
          <router-link :to="{name:'resumelist', query: { id: scope.row.ID }}">
            <el-button type="text" size="small" icon="el-icon-more">投递记录</el-button>
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
import { getUsersList, updateUsers, getUserOptions } from "@/api/jobuser";
import { mapGetters } from "vuex";
import infoList from "@/components/mixins/infoList";
export default {
  name: "Jobuser",
  mixins: [infoList],
  data() {
    return {
      listApi: getUsersList,
      listKey: "userList",
      path: path,
      option: {}
    };
  },
  computed: {
    ...mapGetters("user", ["enPriseId"])
  },
  methods: {
    genderFormat(row) {
      return row.genderindex === 0 ? "男" : row.genderindex === 1 ? "女" : "";
    },
    wyFormat(row) {
      if (row.worksYearindex == null || this.option.jwe == null) {
        return "";
      } else {
        var selectedItem = null;
        selectedItem = this.option.jwe.find(item => {
          return item.ID === row.worksYearindex;
        });
        return selectedItem == null ? "" : selectedItem.name;
      }
    },
    elFormat(row) {
      if (row.edulevelindex == null || this.option.el == null) {
        return "";
      } else {
        var selectedItem = null;
        selectedItem = this.option.el.find(item => {
          return item.ID === row.edulevelindex;
        });
        return selectedItem === null ? "" : selectedItem.name;
      }
    },
    wtFormat(row) {
      if (row.workTypeindex == null || this.option.wt == null) {
        return "";
      } else {
        var selectedItem = null;
        selectedItem = this.option.wt.find(item => {
          return item.ID === row.workTypeindex;
        });
        return selectedItem === null ? "" : selectedItem.name;
      }
    },
    salaryFormat(row) {
      if (row.salaryindex == null || this.option.js == null) {
        return "";
      } else {
        var selectedItem = null;
        selectedItem = this.option.js.find(item => {
          return item.ID === row.salaryindex;
        });
        return selectedItem === null ? "" : selectedItem.name;
      }
    },
    cityFormat(row) {
      if (row.cityindex == null || this.option.cityinfo == null) {
        return "";
      } else {
        var selectedItem = null;
        selectedItem = this.option.cityinfo.find(item => {
          return item.ID === row.cityindex;
        });
        return selectedItem === null ? "" : selectedItem.name;
      }
    },
    dtFormat(row) {
      if (row.dutyTimeindex == null || this.option.jdt == null) {
        return "";
      } else {
        var selectedItem = null;
        selectedItem = this.option.jdt.find(item => {
          return item.ID === row.dutyTimeindex;
        });
        return selectedItem === null ? "" : selectedItem.name;
      }
    },
    clearOptionGender() {
      this.searchInfo.genderindex = null;
    },
    clearOptionEdu() {
      this.searchInfo.edulevelindex = null;
    },
    clearOptionWY() {
      this.searchInfo.worksYearindex = null;
    },
    clearOptionWT() {
      this.searchInfo.workTypeindex = null;
    },
    clearOptionCity() {
      this.searchInfo.cityindex = null;
    },
    clearOptionSalary() {
      this.searchInfo.salaryindex = null;
    },
    clearOptionDT() {
      this.searchInfo.dutyTimeindex = null;
    },
    clearOptionPhoto() {
      this.searchInfo.hasAvatarurl = null;
    },
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    async gotopreview(row, column) {
      //当某行被点击时，去到某行
      if (column.label == "姓名") {
        const { href } = this.$router.resolve({
          name: "userpreview",
          query: { id: row.ID }
        });
        window.open(href, "_blank");
      }
    },
    async editpreview(row) {
      const { href } = this.$router.resolve({
        name: "userpreview",
        query: { id: row.ID }
      });
      window.open(href, "_blank");
    },
    async changestatus(row) {
      const res = await updateUsers(row);
      if (res.success) {
        this.$message({ type: "success", message: "状态设置成功" });
      }
    }
  },
  async created() {
    const res = await getUserOptions();
    if (res.success) {
      this.option = res.data;
    }
  }
};
</script>
<style lang="scss" scoped>
.workTime {
  .el-input {
    width: 40%;
  }
}
</style>>