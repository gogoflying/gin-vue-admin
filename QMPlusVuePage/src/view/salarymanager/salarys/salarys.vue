<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addSalaryDetail" type="primary">新增薪资</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="60" prop="ID"></el-table-column>
      <el-table-column label="openid" min-width="150" prop="openid"></el-table-column>
      <el-table-column label="姓名" min-width="150" prop="name"></el-table-column>
      <el-table-column label="手机" min-width="100" prop="mobile"></el-table-column>
      <el-table-column label="身份证号" min-width="100" prop="card"></el-table-column>
      <el-table-column label="邮箱" min-width="100" prop="email"></el-table-column>
      <el-table-column label="入职企业" min-width="100" prop="enterprise"></el-table-column>
      <el-table-column label="入职日期" min-width="100" prop="enter_time"></el-table-column>
      <el-table-column label="离职日期" min-width="150" prop="leave_time"></el-table-column>
      <el-table-column label="入职进度" min-width="100" prop="enter_step"></el-table-column>
      <el-table-column label="离职进度" min-width="100" prop="leave_step"></el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editSalaryDetail(scope.row)" size="small" type="text">编辑</el-button>
          <el-button @click="deleteSalaryDetail(scope.row)" size="small" type="text">删除</el-button>
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

    <el-dialog :visible.sync="addSalaryDetailDialog" custom-class="user-dialog" title="新增薪资">
      <el-form :rules="rules" ref="salarydetailForm" :model="salarydetailinfo">
        <el-form-item label="姓名" label-width="80px" prop="name">
          <el-input v-model="salarydetailinfo.name"></el-input>
        </el-form-item>
        <el-form-item label="手机" label-width="80px" prop="mobile">
          <el-input v-model="salarydetailinfo.mobile"></el-input>
        </el-form-item>
        <el-form-item label="身份证号" label-width="80px" prop="card">
          <el-input v-model.number="salarydetailinfo.card"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" label-width="80px" prop="email">
          <el-input v-model="salarydetailinfo.email"></el-input>
        </el-form-item>
        <el-form-item label="入职企业" label-width="80px" prop="email">
          <el-select placeholder="请选择企业" v-model="salarydetailinfo.enterprise_id">
            <el-option
              :key="industry.enterprise_name"
              :label="industry.enterprise_name"
              :value="industry.id"
              v-for="industry in industrytype"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="入职日期" label-width="80px" prop="enter_time">
          <el-input v-model="salarydetailinfo.enter_time"></el-input>
        </el-form-item>
        <el-form-item label="入职日期" label-width="80px" prop="leave_time">
          <el-input v-model="salarydetailinfo.leave_time"></el-input>
        </el-form-item>
        <el-form-item label="入职进度" label-width="80px" prop="enter_step">
          <el-select placeholder="请选择入职进度" v-model="salarydetailinfo.enter_step">
            <el-option :key="es.name" :label="es.name" :value="es.id" v-for="es in enterSteps"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="离职进度" label-width="80px" prop="leave_step">
          <el-select placeholder="请选择离职进度" v-model="salarydetailinfo.leave_step">
            <el-option :key="ls.name" :label="ls.name" :value="ls.id" v-for="ls in leaveSteps"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddSalaryDetailDialog">取 消</el-button>
        <el-button @click="enterAddSalaryDetailDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {
  getSalarysList,
  createSalarys,
  updateSalarys,
  findSalarys,
  deleteSalarys
} from "@/api/salarydetail";
import infoList from "@/components/mixins/infoList";
export default {
  name: "Salarys",
  mixins: [infoList],
  data() {
    return {
      listApi: getSalarysList,
      listKey: "userSalaryList",
      path: path,
      addSalaryDetailDialog: false,
      isEdit: false,
      title: "",
      salarydetailinfo: {
        openid: "",
        year: "",
        month: "",
        base: "",
        gangwei: "",
        xzhj: "",
        yjtc: "",
        jjjs: "",
        fdxs: "",
        ydjj: "",
        gzts: "",
        jbf: "",
        txbt: "",
        csbt: "",
        jtbt: "",
        qtbt: "",
        bthj: "",
        qtjq: "",
        njts: "",
        cdkk: "",
        bjts: "",
        bjkk: "",
        sjts: "",
        sjkk: "",
        kkhj: "",
        yftz: "",
        byyf: "",
        dkwx: "",
        gjj: "",
        dkgs: "",
        sfgz: ""
      },
      rules: {
        name: [{ required: true, message: "请输入用户名", trigger: "blur" }],
        mobile: [
          { required: true, message: "请输入电话号码", trigger: "blur" }
        ],
        card: [
          { required: true, message: "请输入身份证号码", trigger: "blur" },
          {
            pattern: /(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$)/,
            message: "身份证号码格式不正确",
            trigger: "blur"
          }
        ],
        email: [
          { required: true, message: "请输入邮箱地址", trigger: "blur" },
          {
            type: "email",
            message: "邮箱格式不正确",
            trigger: "blur"
          }
        ],
        enterprise_id: [
          { required: true, message: "请输入企业", trigger: "blur" }
        ],
        enter_time: [
          { required: true, message: "请输入入职日期", trigger: "blur" }
        ]
      },
      enterSteps: [
        {
          id: 0,
          name: "待用户确认信息"
        },
        {
          id: 1,
          name: "已确认信息"
        },

        {
          id: 2,
          name: "待用户网签合同"
        },
        {
          id: 3,
          name: "合同签完"
        }
      ],
      leaveSteps: [
        {
          id: 0,
          name: "待用户审核"
        },
        {
          id: 1,
          name: "用户确认信息"
        },
        {
          id: 2,
          name: "离职"
        }
      ]
    };
  },
  methods: {
    async enterAddSalaryDetailDialog() {
      this.$refs.salarydetailForm.validate(async valid => {
        if (valid) {
          let res;
          if (this.isEdit) {
            res = await updateSalarys(this.salarydetailinfo);
          } else {
            res = await createSalarys(this.salarydetailinfo);
          }
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          } else {
            this.$message({ type: "error", message: "添加失败!" });
          }
          await this.getTableData();
          this.closeAddSalaryDetailDialog();
        }
      });
    },
    closeAddSalaryDetailDialog() {
      this.salarydetailinfo = {
        openid: "",
        year: "",
        month: "",
        base: "",
        gangwei: "",
        xzhj: "",
        yjtc: "",
        jjjs: "",
        fdxs: "",
        ydjj: "",
        gzts: "",
        jbf: "",
        txbt: "",
        csbt: "",
        jtbt: "",
        qtbt: "",
        bthj: "",
        qtjq: "",
        njts: "",
        cdkk: "",
        bjts: "",
        bjkk: "",
        sjts: "",
        sjkk: "",
        kkhj: "",
        yftz: "",
        byyf: "",
        dkwx: "",
        gjj: "",
        dkgs: "",
        sfgz: ""
      };
      this.addSalaryDetailDialog = false;
    },
    //新增企业
    addSalaryDetail() {
      this.title = "新增薪资";
      this.isEdit = false;
      this.addSalaryDetailDialog = true;
    },
    //编辑企业
    async editSalaryDetail(row) {
      this.title = "编辑薪资";
      const res = await findSalarys(row);
      this.salarydetailinfo = res.data.reinfo;
      this.isEdit = true;
      this.addSalaryDetailDialog = true;
    },
    //删除企业
    deleteSalaryDetail(row) {
      this.$confirm("此操作将永久删除该薪资信息, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteSalarys(row);
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