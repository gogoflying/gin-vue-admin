<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addSalaryTemplates" type="primary" icon="el-icon-plus">新增模板</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="ID" min-width="60" type="index" :index="indexMethod"></el-table-column>
      <el-table-column label="所属企业" min-width="120" prop="enterprise" v-if="enPriseId == 0"></el-table-column>
      <el-table-column label="模板名称" min-width="60" prop="gwtemp_name"></el-table-column>
      <el-table-column label="薪资构成" align="center">
        <el-table-column label="基本工资" min-width="100" prop="base"></el-table-column>
        <el-table-column label="岗位工资" min-width="100" prop="gangwei"></el-table-column>
        <el-table-column label="奖金基数" min-width="80" prop="jjjs"></el-table-column>
        <el-table-column label="工作天数" min-width="120" prop="gzts"></el-table-column>
      </el-table-column>
      <el-table-column label="补贴明细" align="center">
        <el-table-column label="通讯补贴" min-width="80" prop="txbt"></el-table-column>
        <el-table-column label="餐食补贴" min-width="80" prop="csbt"></el-table-column>
        <el-table-column label="交通补贴" min-width="80" prop="jtbt"></el-table-column>
        <el-table-column label="其他补贴" min-width="80" prop="qtbt"></el-table-column>
        <el-table-column label="补贴合计" min-width="80" prop="bthj"></el-table-column>
      </el-table-column>
      <el-table-column label="五险一金" align="center">
        <el-table-column label="代扣五险" min-width="80" prop="dkwx"></el-table-column>
        <el-table-column label="代扣住房公积金" min-width="80" prop="gjj"></el-table-column>
        <el-table-column label="代扣个人所得税" min-width="80" prop="dkgs"></el-table-column>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="180">
        <template slot-scope="scope">
          <el-button
            @click="editSalaryTemplates(scope.row)"
            icon="el-icon-edit"
            size="small"
            type="text"
          >编辑</el-button>
          <el-button
            @click="deleteSalaryTemplates(scope.row)"
            icon="el-icon-delete"
            size="small"
            type="text"
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

    <el-dialog
      :visible.sync="addSalaryTemplatesDialog"
      custom-class="user-dialog"
      @close="closeAddSalaryTemplatesDialog"
      title="新增模板"
      width="800px"
    >
      <el-form :rules="rules" ref="SalaryTemplatesForm" :model="SalaryTemplatesinfo">
        <section>
          <h3>基本信息</h3>
          <el-row>
            <el-col :span="12">
              <el-form-item
                v-show="enPriseId == 0"
                label="所属企业"
                label-width="80px"
                prop="enterprise_id"
              >
                <el-select
                  @change="selectEp"
                  v-model="SalaryTemplatesinfo.enterprise_id"
                  filterable
                  remote
                  reserve-keyword
                  placeholder="请选择所属公司"
                  :remote-method="remoteMethod"
                  :loading="loading"
                >
                  <el-option
                    v-for="ep in enterpriseInfo"
                    :key="ep.enterprise_name"
                    :value="ep.ID"
                    :label="ep.enterprise_name"
                  ></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="模板名称" label-width="80px" prop="gwtemp_name">
                <el-input v-model="SalaryTemplatesinfo.gwtemp_name"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>薪资构成</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="基本工资" label-width="80px" prop="base">
                <el-input v-model="SalaryTemplatesinfo.base"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="岗位工资" label-width="80px" prop="gangwei">
                <el-input v-model="SalaryTemplatesinfo.gangwei"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="奖金基数" label-width="80px" prop="jjjs">
                <el-input v-model="SalaryTemplatesinfo.jjjs"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="工作天数" label-width="80px" prop="gzts">
                <el-input v-model="SalaryTemplatesinfo.gzts"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>补贴明细</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="通讯补贴" label-width="80px" prop="txbt">
                <el-input v-model="SalaryTemplatesinfo.txbt"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="餐食补贴" label-width="80px" prop="csbt">
                <el-input v-model="SalaryTemplatesinfo.csbt"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="交通补贴" label-width="80px" prop="jtbt">
                <el-input v-model="SalaryTemplatesinfo.jtbt"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="其他补贴" label-width="80px" prop="qtbt">
                <el-input v-model="SalaryTemplatesinfo.qtbt"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row style="margin-top:-20px">
            <el-col :span="12">
              <el-form-item label="补贴合计" label-width="80px">
                <el-input style="200px" :value="bthj"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>五险一金</h3>
          <el-row>
            <el-col :span="8">
              <el-form-item label="代扣五险" label-width="80px" prop="dkwx">
                <el-input v-model="SalaryTemplatesinfo.dkwx"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="代扣住房公积金" label-width="120px" prop="gjj">
                <el-input v-model="SalaryTemplatesinfo.gjj"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="代扣个人所得税" label-width="120px" prop="dkgs">
                <el-input v-model="SalaryTemplatesinfo.dkgs"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddSalaryTemplatesDialog">取 消</el-button>
        <el-button @click="enterAddSalaryTemplatesDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {
  getSalaryTemplatesList,
  createSalaryTemplates,
  updateSalaryTemplates,
  findSalaryTemplates,
  deleteSalaryTemplates
} from "@/api/salarytemplate";
import { getEnterpriseListBySearch } from "@/api/enterpriseinfo";
import { mapGetters } from "vuex";
import infoList from "@/components/mixins/infoList";
export default {
  name: "SalaryTemplatess",
  mixins: [infoList],
  data() {
    const isFloatVlidator = (rule, value, callback) => {
      var pattern = /^\d+.?\d{0,2}$/;
      if (value.length <= 0) {
        return callback(new Error("不能为空"));
      }
      if (!pattern.test(value)) {
        return callback(new Error("请输入非负数并且最多两位小数"));
      } else {
        return callback();
      }
    };
    return {
      listApi: getSalaryTemplatesList,
      listKey: "stList",
      path: path,
      addSalaryTemplatesDialog: false,
      isEdit: false,
      loading: false,
      title: "",
      enterpriseInfo: [],
      SalaryTemplatesinfo: {
        enterprise: "",
        enterprise_id: "",
        gwtemp_name: "",
        base: "",
        gangwei: "",
        jjjs: "",
        gzts: "",
        txbt: "",
        csbt: "",
        jtbt: "",
        qtbt: "",
        bthj: "",
        dkwx: "",
        gjj: "",
        dkgs: ""
      },
      rules: {
        gwtemp_name: [
          { required: true, message: "请输入模板名称", trigger: "blur" }
        ],
        base: [{ validator: isFloatVlidator, trigger: "blur" }],
        gangwei: [{ validator: isFloatVlidator, trigger: "blur" }],
        jjjs: [{ validator: isFloatVlidator, trigger: "blur" }],
        gzts: [{ validator: isFloatVlidator, trigger: "blur" }],
        txbt: [{ validator: isFloatVlidator, trigger: "blur" }],
        csbt: [{ validator: isFloatVlidator, trigger: "blur" }],
        jtbt: [{ validator: isFloatVlidator, trigger: "blur" }],
        qtbt: [{ validator: isFloatVlidator, trigger: "blur" }],
        //bthj: [{ validator: isFloatVlidator, trigger: "blur" }],
        dkwx: [{ validator: isFloatVlidator, trigger: "blur" }],
        gjj: [{ validator: isFloatVlidator, trigger: "blur" }],
        dkgs: [{ validator: isFloatVlidator, trigger: "blur" }]
      }
    };
  },
  computed: {
    ...mapGetters("user", ["enPriseId"]),
    bthj: {
      get() {
        let sum = 0;
        if (!isNaN(parseFloat(this.SalaryTemplatesinfo.txbt))) {
          sum += parseFloat(this.SalaryTemplatesinfo.txbt);
        }
        if (!isNaN(parseFloat(this.SalaryTemplatesinfo.csbt))) {
          sum += parseFloat(this.SalaryTemplatesinfo.csbt);
        }
        if (!isNaN(parseFloat(this.SalaryTemplatesinfo.jtbt))) {
          sum += parseFloat(this.SalaryTemplatesinfo.jtbt);
        }
        if (!isNaN(parseFloat(this.SalaryTemplatesinfo.qtbt))) {
          sum += parseFloat(this.SalaryTemplatesinfo.qtbt);
        }
        //return sum.toFixed(2);
        return Math.round(sum * 100) / 100;
      }
    }
  },
  methods: {
    indexMethod(index) {
      return index + 1 + (this.page - 1) * this.pageSize;
    },
    remoteMethod(query) {
      if (query !== "") {
        this.loading = true;
        setTimeout(() => {
          this.loading = false;
          const params = {
            query: query
          };
          getEnterpriseListBySearch(params).then(res => {
            if (res.success) {
              this.enterpriseInfo = res.data.result;
            } else {
              this.enterpriseInfo = [];
            }
          });
        }, 200);
      } else {
        this.enterpriseInfo = [];
      }
    },
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    selectEp(val) {
      var selectedItem = {};
      selectedItem = this.enterpriseInfo.find(item => {
        return item.ID === val;
      });
      this.SalaryTemplatesinfo.enterprise = selectedItem.enterprise_name;
    },
    downSalaryTemplatestemplate() {
      const link = document.createElement("a");
      const url = `${path}/un/template/SalaryTemplatess.xlsx`;
      link.setAttribute("href", url);
      link.setAttribute("download", "薪资模板");
      link.click();
    },
    async enterAddSalaryTemplatesDialog() {
      this.SalaryTemplatesinfo.bthj = this.bthj.toString();
      this.$refs.SalaryTemplatesForm.validate(async valid => {
        if (valid) {
          let res;
          if (this.isEdit) {
            res = await updateSalaryTemplates(this.SalaryTemplatesinfo);
          } else {
            res = await createSalaryTemplates(this.SalaryTemplatesinfo);
          }
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          } else {
            this.$message({ type: "error", message: "添加失败!" });
          }
          await this.getTableData();
          this.closeAddSalaryTemplatesDialog();
        }
      });
    },
    closeAddSalaryTemplatesDialog() {
      this.SalaryTemplatesinfo = {
        enterprise: "",
        enterprise_id: "",
        gwtemp_name: "",
        base: "",
        gangwei: "",
        jjjs: "",
        gzts: "",
        txbt: "",
        csbt: "",
        jtbt: "",
        qtbt: "",
        bthj: "",
        dkwx: "",
        gjj: "",
        dkgs: ""
      };
      this.$refs.SalaryTemplatesForm.resetFields();
      this.addSalaryTemplatesDialog = false;
    },
    //新增企业
    addSalaryTemplates() {
      this.title = "新增模板";
      this.isEdit = false;
      this.enterpriseInfo = [];
      this.addSalaryTemplatesDialog = true;
    },
    //编辑企业
    async editSalaryTemplates(row) {
      this.title = "编辑模板";
      this.enterpriseInfo = [];
      const res = await findSalaryTemplates(row);
      this.SalaryTemplatesinfo = res.data.rest;
      if (this.SalaryTemplatesinfo != null) {
        this.enterpriseInfo.push({
          ID: this.SalaryTemplatesinfo.enterprise_id,
          enterprise_name: this.SalaryTemplatesinfo.enterprise
        });
      }
      this.isEdit = true;
      this.addSalaryTemplatesDialog = true;
    },
    //删除企业
    deleteSalaryTemplates(row) {
      this.$confirm("此操作将永久删除该薪资信息, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteSalaryTemplates(row);
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