<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="新闻标题">
          <el-input placeholder="新闻标题" v-model="searchInfo.title"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <router-link :to="{name:'newdetail'}">
            <el-button type="primary" size="small" icon="el-icon-edit">创建文章</el-button>
          </router-link>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="60" prop="ID" fixed></el-table-column>
      <el-table-column label="新闻标题" min-width="100" prop="title" fixed></el-table-column>
      <el-table-column label="新闻副标题" min-width="100" prop="sub_title" v-if="false"></el-table-column>
      <!-- <el-table-column label="排序" min-width="100" prop="order"></el-table-column> -->
      <el-table-column label="新闻类型" min-width="150" prop="news_type" :formatter="newTypeFormat"></el-table-column>
      <el-table-column label="状态" min-width="150" prop="status">
        <template slot-scope="scope">
          <el-tag>{{ scope.row.status | statusFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="浏览量" min-width="150" prop="count"></el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <router-link
            :to="{name:'newdetail', params:{row:scope.row}, query: { id: scope.row.ID }}"
          >
            <el-button type="primary" size="small" icon="el-icon-edit">编辑</el-button>
          </router-link>
          <el-button
            @click="deletenew(scope.row)"
            size="small"
            icon="el-icon-delete"
            type="primary"
          >删除</el-button>
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
import { deleteUserNews, getUserNewsList } from "@/api/newmanager";
import infoList from "@/components/mixins/infoList";
export default {
  name: "Jobmanagerinfo",
  mixins: [infoList],
  filters: {
    statusFilter(status) {
      if (status == 2) {
        return "已发布";
      } else {
        return "草稿";
      }
    }
  },
  data() {
    return {
      listApi: getUserNewsList,
      listKey: "userSalaryList",
      path: path
    };
  },
  methods: {
    newTypeFormat(row){
      if (row.news_type == 1) {
        return "社保";
      } else {
        return "个税";
      }
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    // 删除职位
    deletenew(row) {
      this.$confirm("此操作将永久删除该文章, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteUserNews(row);
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
  },
  async created() {
    // const res = await getCitynameList({ page: 1, pageSize: 999 });
    // this.cityinfo = res.data.userCityList;
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
</style>