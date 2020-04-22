<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="工作名称">
          <el-input placeholder="工作名称" v-model="searchInfo.p_name"></el-input>
        </el-form-item>
        <el-form-item label="工作城市">
          <el-input placeholder="工作城市" v-model="searchInfo.p_city"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="addjob" type="primary">新增职位</el-button>
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
      <el-table-column label="id" min-width="60" prop="ID" fixed></el-table-column>
      <el-table-column label="工作名称" min-width="100" prop="p_name"></el-table-column>
      <el-table-column label="薪资上限" min-width="100" prop="p_salary_high"></el-table-column>
      <el-table-column label="薪资下限" min-width="100" prop="p_salary_low"></el-table-column>
      <el-table-column label="工作地点纬度" min-width="150" prop="p_latitude"></el-table-column>
      <el-table-column label="工作地点经度" min-width="150" prop="p_longitude"></el-table-column>
      <el-table-column label="工作地点" min-width="100" prop="p_address"></el-table-column>
      <el-table-column label="工作城市" min-width="150" prop="p_city"></el-table-column>
      <el-table-column label="工作年限" min-width="100" prop="p_edujy"></el-table-column>
      <el-table-column label="最低学历" min-width="150" prop="p_education"></el-table-column>
      <el-table-column label="工作类型" min-width="100" prop="p_type"></el-table-column>
      <el-table-column label="工作描述" min-width="150" prop="p_des"></el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editjob(scope.row)" size="small" type="text">编辑</el-button>
          <el-button @click="deletejob(scope.row)" size="small" type="text">删除</el-button>
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

    <el-dialog :visible.sync="addjobDialog" custom-class="user-dialog" :title="title">
      <el-form :rules="rules" ref="jobForm" :model="jobmanagerinfo">
        <el-form-item label="工作名称" label-width="100px" prop="p_name">
          <el-input v-model="jobmanagerinfo.p_name" placeholder="请输入工作名称"></el-input>
        </el-form-item>
        <el-form-item label="薪资上限" label-width="100px" prop="p_salary_high">
          <el-input v-model.number="jobmanagerinfo.p_salary_high" placeholder="请输入薪资上限"></el-input>
        </el-form-item>
        <el-form-item label="薪资下限" label-width="100px" prop="p_salary_low">
          <el-input v-model.number="jobmanagerinfo.p_salary_low" placeholder="请输入薪资下限"></el-input>
        </el-form-item>
        <el-form-item label="工作地点纬度" label-width="100px" prop="p_latitude">
          <el-input v-model.number="jobmanagerinfo.p_latitude" placeholder="请输入工作地点纬度"></el-input>
        </el-form-item>
        <el-form-item label="工作地点经度" label-width="100px" prop="p_longitude">
          <el-input v-model.number="jobmanagerinfo.p_longitude" placeholder="请输入工作地点经度"></el-input>
        </el-form-item>
        <el-form-item label="工作地点" label-width="100px" prop="p_address">
          <el-input v-model="jobmanagerinfo.p_address" placeholder="请输入工作地点"></el-input>
        </el-form-item>
        <el-form-item label="工作城市" label-width="100px" prop="p_city_id" @change="selectCityName">
          <el-select placeholder="请选择工作城市" v-model.number="jobmanagerinfo.p_city_id">
            <el-option
              :key="city.name"
              :label="city.name"
              :value="city.id"
              v-for="city in cityinfo"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工作年限" label-width="100px" prop="p_edujy_id" @change="selectJobyear">
          <el-select placeholder="请选择工作年限" v-model.number="jobmanagerinfo.p_edujy_id">
            <el-option
              :key="jobyear.name"
              :label="jobyear.name"
              :value="jobyear.id"
              v-for="jobyear in jobyears"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="最低学历" label-width="100px" prop="p_education_id" @change="selectJobedu">
          <el-select placeholder="请选择最低学历" v-model.number="jobmanagerinfo.p_education_id">
            <el-option
              :key="jobedu.name"
              :label="jobedu.name"
              :value="jobedu.id"
              v-for="jobedu in jobedus"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工作类型" label-width="100px" prop="p_type_id" @change="selectJobtype">
          <el-select placeholder="请选择工作类型" v-model.number="jobmanagerinfo.p_type_id">
            <el-option
              :key="jobtype.name"
              :label="jobtype.name"
              :value="jobtype.id"
              v-for="jobtype in jobtypes"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工作描述" label-width="100px" prop="p_desc">
          <el-input
            type="textarea"
            :autosize="{ minRows: 4}"
            v-model="jobmanagerinfo.p_desc"
            placeholder="请输入工作描述"
          ></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddjobDialog">取 消</el-button>
        <el-button @click="enterAddjobDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {
  getJoblistListBackend,
  createJoblist,
  findJoblist,
  deleteJoblist
} from "@/api/jobmanagerinfo";
//import { getCitynameList } from "@/api/cityname";
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
      addjobDialog: false,
      title: "",
      jobmanagerinfo: {
        p_name: "",
        p_salary_high: "",
        p_salary_low: "",
        p_latitude: "",
        p_longitude: "",
        p_address: "",
        p_city: "",
        p_city_id: "",
        p_edujy: "",
        p_edujy_id: "",
        p_education: "",
        p_education_id: "",
        p_type: "",
        p_type_id: "",
        p_des: ""
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
      jobyears: [
        {
          id: 0,
          name: "1年"
        },
        {
          id: 1,
          name: "3年"
        },
        {
          id: 2,
          name: "5年"
        }
      ],
      jobtypes: [
        {
          id: 0,
          name: "全职"
        },
        {
          id: 1,
          name: "兼职"
        },
        {
          id: 2,
          name: "实习"
        },
        {
          id: 3,
          name: "志愿者"
        }
      ],
      jobedus: [
        {
          id: -1,
          name: "其他"
        },
        {
          id: 0,
          name: "不限"
        },
        {
          id: 1,
          name: "初中"
        },
        {
          id: 2,
          name: "高中"
        },
        {
          id: 3,
          name: "中专"
        },
        {
          id: 4,
          name: "大专"
        },
        {
          id: 5,
          name: "本科"
        },
        {
          id: 6,
          name: "硕士"
        },
        {
          id: 7,
          name: "博士"
        }
      ],
      rules: {
        enterprise_name: [
          {
            required: true,
            message: "请输入与营业执照一致的企业名称",
            trigger: "blur"
          }
        ],
        enterprise_address: [
          { required: true, message: "请输入联系人全名", trigger: "blur" }
        ],
        enterprise_type: [
          { required: true, message: "请输入电话", trigger: "blur" }
        ],
        enterprise_desc: [
          { required: true, message: "请输入企业邮箱", trigger: "blur" }
        ]
      }
    };
  },
  methods: {
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    async enterAddjobDialog() {
      this.$refs.jobForm.validate(async valid => {
        if (valid) {
          const res = await createJoblist(this.jobmanagerinfo);
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          }
          await this.getTableData();
          this.closeAddjobDialog();
        }
      });
    },
    closeAddjobDialog() {
      this.jobmanagerinfo = {
        p_name: "",
        p_salary_high: "",
        p_salary_low: "",
        p_latitude: "",
        p_longitude: "",
        p_address: "",
        p_city: "",
        p_city_id: "",
        p_edujy: "",
        p_edujy_id: "",
        p_education: "",
        p_education_id: "",
        p_type: "",
        p_type_id: "",
        p_des: ""
      };
      this.addjobDialog = false;
    },
    //新增职位
    addjob() {
      this.title = "新增职位";
      this.addjobDialog = true;
    },
    //编辑职位
    async editjob(row) {
      this.title = "编辑职位";
      const res = await findJoblist(row);
      this.jobmanagerinfo = res.data.rejl;
      this.addjobDialog = true;
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
    selectCityName(val) {
      var selectedItem = {};
      selectedItem = this.cityinfo.find(item => {
        return item.id === val;
      });
      this.jobmanagerinfo.p_city = selectedItem.name;
    },
    selectJobyear(val) {
      var selectedItem = {};
      selectedItem = this.jobyears.find(item => {
        return item.id === val;
      });
      this.jobmanagerinfo.p_edujy = selectedItem.name;
    },
    selectJobtype(val) {
      var selectedItem = {};
      selectedItem = this.jobtypes.find(item => {
        return item.id === val;
      });
      this.jobmanagerinfo.p_type = selectedItem.name;
    },
    selectJobedu(val) {
      var selectedItem = {};
      selectedItem = this.jobedus.find(item => {
        return item.id === val;
      });
      this.jobmanagerinfo.p_education = selectedItem.name;
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
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