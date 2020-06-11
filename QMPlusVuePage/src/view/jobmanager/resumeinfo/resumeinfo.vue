<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item v-show="enPriseId == 0" label="企业名称">
          <el-input placeholder="企业名称" v-model="searchInfo.enterprise_name"></el-input>
        </el-form-item>
        <el-form-item label="职位名称">
          <el-input placeholder="职位名称" v-model="searchInfo.p_name"></el-input>
        </el-form-item>
        <el-form-item label="投递人">
          <el-input placeholder="投递人" v-model="searchInfo.userName"></el-input>
        </el-form-item>
        <el-form-item label="手机号码">
          <el-input placeholder="手机号" v-model="searchInfo.mobile"></el-input>
        </el-form-item>

        <el-form-item label="简历状态">
          <el-select
            clearable
            @clear="clearStatus"
            placeholder="请选择"
            v-model.number="searchInfo.resume_status_search"
          >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.id"
              v-for="item in resume_status_infos"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-dialog title="备注信息" :visible.sync="dialogFormVisible" append-to-body>
      <el-input type="textarea" :rows="3" placeholder="请输入备注" v-model="p_memo"></el-input>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="dialogOK()">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog title="面试提示" :visible.sync="dialogRemarkVisible" append-to-body>
      <el-form :model="resume_info" label-width="80px">
        <el-form-item label="联系人" prop="p_contact">
          <el-input placeholder="请输入联系人" v-model="resume_info.p_contact"></el-input>
        </el-form-item>
        <el-form-item label="面试时间" prop="p_interview_time">
          <el-date-picker
            placeholder="请输入面试时间"
            type="datetime"
            :editable="false"
            format="yyyy-MM-dd HH:mm:ss"
            value-format="yyyy-MM-dd HH:mm:ss"
            v-model="resume_info.p_interview_time"
          ></el-date-picker>
        </el-form-item>
        <el-form-item label="面试地点" prop="p_interview_loc">
          <el-input placeholder="请输入面试地点" v-model="resume_info.p_interview_loc"></el-input>
        </el-form-item>
        <el-form-item label="备注" prop="p_remark">
          <el-input type="textarea" :rows="3" placeholder="请输入备注信息" v-model="resume_info.p_remark"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogRemarkVisible = false">取 消</el-button>
        <el-button type="primary" @click="dialogRemarkOK()">确 定</el-button>
      </div>
    </el-dialog>

    <el-table :data="tableData" border stripe>
      <!-- <el-table-column type="selection" min-width="55"></el-table-column> -->
      <el-table-column label="ID" min-width="60" type="index" :index="indexMethod"></el-table-column>
      <el-table-column label="职位名称" min-width="150" prop="job_info.p_name"></el-table-column>
      <el-table-column label="企业名称" min-width="150" prop="enterprise_name" v-if="enPriseId == 0"></el-table-column>
      <el-table-column label="投递人" min-width="100" prop="user_info.userName"></el-table-column>
      <el-table-column label="手机号" min-width="150" prop="user_info.contact"></el-table-column>
      <el-table-column
        label="工作年限"
        min-width="150"
        prop="user_info.worksYearindex"
        :formatter="wyFormat"
      ></el-table-column>
      <el-table-column
        label="学历"
        min-width="150"
        prop="user_info.edulevelindex"
        :formatter="elFormat"
      ></el-table-column>
      <el-table-column label="年龄" min-width="150" prop="user_info.birthday" :formatter="ageFormat"></el-table-column>
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
            type="text"
            size="small"
            icon="el-icon-view"
          >查看</el-button>
          <el-button
            @click="memo(scope.row)"
            size="small"
            type="text"
            icon="el-icon-collection-tag"
          >备注</el-button>
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
  getResumeStatusList,
  updateResumeStatus,
  updateResumeMemo
} from "@/api/resumestatus";
import { getUserOptions } from "@/api/jobuser";
import { mapGetters } from "vuex";
import infoList from "@/components/mixins/infoList";
export default {
  name: "Api",
  mixins: [infoList],
  data() {
    return {
      listApi: getResumeStatusList,
      listKey: "resumeStatusList",
      path: path,
      p_memo: "",
      Id: 0,
      option: {},
      dialogFormVisible: false,
      dialogRemarkVisible: false,
      resume_info: {},
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
        },
        {
          id: 5,
          name: "待入职"
        },
        {
          id: 6,
          name: "已入职"
        },
        {
          id: 7,
          name: "已拒绝"
        }
      ]
    };
  },
  computed: {
    ...mapGetters("user", ["enPriseId"])
  },
  methods: {
    indexMethod(index) {
      return index + 1 + (this.page - 1) * this.pageSize;
    },
    wyFormat(row) {
      if (row.user_info.worksYearindex == null || this.option.jwe == null) {
        return "";
      } else {
        var selectedItem = null;
        selectedItem = this.option.jwe.find(item => {
          return item.ID === row.user_info.worksYearindex;
        });
        return selectedItem == null ? "" : selectedItem.name;
      }
    },
    elFormat(row) {
      if (row.user_info.edulevelindex == null || this.option.el == null) {
        return "";
      } else {
        var selectedItem = null;
        selectedItem = this.option.el.find(item => {
          return item.ID === row.user_info.edulevelindex;
        });
        return selectedItem === null ? "" : selectedItem.name;
      }
    },
    ageFormat(row) {
      if (row.user_info.birthday != "") {
        var date = new Date(row.user_info.birthday);
        var d = new Date();
        return d.getFullYear() - date.getFullYear();
      } else {
        return "";
      }
    },
    async editpreview(row, column) {
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
    async memo(row) {
      this.Id = row.ID;
      this.p_memo = row.p_memo;
      this.dialogFormVisible = !this.dialogFormVisible;
    },
    clearStatus() {
      this.searchInfo.resume_status_search = null;
    },
    async dialogOK() {
      const res = await updateResumeMemo({ ID: this.Id, p_memo: this.p_memo });
      if (res.success) {
        this.$message({ type: "success", message: "更新备注成功" });
      }
      this.getTableData();
      this.dialogFormVisible = false;
    },
    async dialogRemarkOK() {
      if (
        this.resume_info.p_contact == "" ||
        this.resume_info.p_interview_time == "" ||
        this.resume_info.p_interview_loc == ""
      ) {
        this.$message({ type: "error", message: "请填写必要信息" });
        return;
      }
      const res = await updateResumeStatus(this.resume_info);
      if (res.success) {
        this.$message({ type: "success", message: "更新提示信息成功" });
      }
      this.getTableData();
      this.dialogRemarkVisible = false;
    },
    async changeResumeStatus(row) {
      row.p_badge = 1;
      if (row.resume_status != 3) {
        const res = await updateResumeStatus(row);
        if (res.success) {
          this.$message({ type: "success", message: "状态设置成功" });
        }
      }
      if (row.resume_status == 3) {
        if (row.p_contact == "") {
          row.p_contact = row.job_info.p_contact;
        }
        if (row.p_interview_loc == "") {
          row.p_interview_loc = row.job_info.p_interview;
        }
        this.resume_info = row;
        this.dialogRemarkVisible = !this.dialogRemarkVisible;
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