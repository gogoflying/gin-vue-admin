<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addSalaryTemplates" type="primary" icon="el-icon-plus">新增模板</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="60" prop="ID" v-if="enPriseId == 0"></el-table-column>
      <el-table-column label="所属企业" min-width="120" prop="enterprise" v-if="enPriseId == 0"></el-table-column>
      <el-table-column label="模板名称" min-width="60" prop="gwtemp_name"></el-table-column>
      <el-table-column label="薪资构成" align="center">
        <el-table-column label="基本工资" min-width="100" prop="base"></el-table-column>
        <el-table-column label="岗位工资" min-width="100" prop="gangwei"></el-table-column>
        <el-table-column label="业绩提成" min-width="80" prop="yjtc"></el-table-column>
        <el-table-column label="奖金基数" min-width="80" prop="jjjs"></el-table-column>
        <el-table-column label="浮动系数" min-width="80" prop="fdxs"></el-table-column>
        <el-table-column label="月度奖金" min-width="80" prop="ydjj"></el-table-column>
        <el-table-column label="工作天数" min-width="120" prop="gzts"></el-table-column>
        <el-table-column label="加班费" min-width="80" prop="jbf"></el-table-column>
        <el-table-column label="薪资合计" min-width="80" prop="xzhj"></el-table-column>
      </el-table-column>
      <el-table-column label="补贴明细" align="center">
        <el-table-column label="通讯补贴" min-width="80" prop="txbt"></el-table-column>
        <el-table-column label="餐食补贴" min-width="80" prop="csbt"></el-table-column>
        <el-table-column label="交通补贴" min-width="80" prop="jtbt"></el-table-column>
        <el-table-column label="其他补贴" min-width="80" prop="qtbt"></el-table-column>
        <el-table-column label="补贴合计" min-width="80" prop="bthj"></el-table-column>
      </el-table-column>
      <!-- <el-table-column label="假期明细" align="center">
        <el-table-column label="年假天数" min-width="80" prop="njts"></el-table-column>
        <el-table-column label="病假天数" min-width="80" prop="bjts"></el-table-column>
        <el-table-column label="事假天数" min-width="80" prop="sjts"></el-table-column>
        <el-table-column label="其他假期" min-width="80" prop="qtjq"></el-table-column>
      </el-table-column>
      <el-table-column label="扣款明细" align="center">
        <el-table-column label="迟到扣款" min-width="80" prop="cdkk"></el-table-column>
        <el-table-column label="病假扣款" min-width="80" prop="bjkk"></el-table-column>
        <el-table-column label="事假扣款" min-width="80" prop="sjkk"></el-table-column>
        <el-table-column label="扣款合计" min-width="80" prop="kkhj"></el-table-column>
      </el-table-column> -->
      <el-table-column label="五险一金" align="center">
        <el-table-column label="代扣五险" min-width="80" prop="dkwx"></el-table-column>
        <el-table-column label="代扣住房公积金" min-width="80" prop="gjj"></el-table-column>
        <el-table-column label="代扣个人所得税" min-width="80" prop="dkgs"></el-table-column>
      </el-table-column>
      <!-- <el-table-column label="应发调整" align="center">
        <el-table-column label="应发工资" min-width="80" prop="byyf"></el-table-column>
        <el-table-column label="应发调整" min-width="80" prop="yftz"></el-table-column>
        <el-table-column label="实发工资" min-width="80" prop="sfgz"></el-table-column>
      </el-table-column> -->
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
              <el-form-item label="业绩提成" label-width="80px" prop="yjtc">
                <el-input v-model="SalaryTemplatesinfo.yjtc"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="奖金基数" label-width="80px" prop="jjjs">
                <el-input v-model="SalaryTemplatesinfo.jjjs"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row style="margin-top:-20px">
            <el-col :span="6">
              <el-form-item label="浮动系数" label-width="80px" prop="fdxs">
                <el-input v-model="SalaryTemplatesinfo.fdxs"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="月度奖金" label-width="80px" prop="ydjj">
                <el-input v-model="SalaryTemplatesinfo.ydjj"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="工作天数" label-width="80px" prop="gzts">
                <el-input v-model="SalaryTemplatesinfo.gzts"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="加班费" label-width="80px" prop="jbf">
                <el-input v-model="SalaryTemplatesinfo.jbf"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row style="margin-top:-20px">
            <el-col :span="12">
              <el-form-item label="薪资合计" label-width="80px" prop="xzhj">
                <el-input style="200px" v-model="SalaryTemplatesinfo.xzhj"></el-input>
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
              <el-form-item label="补贴合计" label-width="80px" prop="bthj">
                <el-input style="200px" v-model="SalaryTemplatesinfo.bthj"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <!-- <section>
          <h3>假期明细</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="年假天数" label-width="80px" prop="njts">
                <el-input v-model="SalaryTemplatesinfo.njts"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="病假天数" label-width="80px" prop="bjts">
                <el-input v-model="SalaryTemplatesinfo.bjts"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="事假天数" label-width="80px" prop="sjts">
                <el-input v-model="SalaryTemplatesinfo.sjts"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="其他假期" label-width="80px" prop="qtjq">
                <el-input v-model="SalaryTemplatesinfo.qtjq"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>扣款明细</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="迟到扣款" label-width="80px" prop="cdkk">
                <el-input v-model="SalaryTemplatesinfo.cdkk"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="病假扣款" label-width="80px" prop="bjkk">
                <el-input v-model="SalaryTemplatesinfo.bjkk"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="事假扣款" label-width="80px" prop="sjkk">
                <el-input v-model="SalaryTemplatesinfo.sjkk"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="扣款合计" label-width="80px" prop="kkhj">
                <el-input v-model="SalaryTemplatesinfo.kkhj"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>-->
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
        <!-- <section>
          <h3>应发调整</h3>
          <el-row>
            <el-col :span="8">
              <el-form-item label="应发工资" label-width="80px" prop="byyf">
                <el-input v-model="SalaryTemplatesinfo.byyf"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="应发调整" label-width="80px" prop="yftz">
                <el-input v-model="SalaryTemplatesinfo.yftz"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="实发工资" label-width="80px" prop="sfgz">
                <el-input v-model="SalaryTemplatesinfo.sfgz"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>-->
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
        gwtemp_name: "",
        base: "",
        gangwei: "",
        enterprise: "",
        enterprise_id: "",
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
        dkwx: "",
        gjj: "",
        dkgs: ""
        // qtjq: "",
        // njts: "",
        // cdkk: "",
        // bjts: "",
        // bjkk: "",
        // sjts: "",
        // sjkk: "",
        // kkhj: "",
        // yftz: "",
        // byyf: "",
        // sfgz: ""
      },
      rules: {}
    };
  },
  computed: {
    ...mapGetters("user", ["enPriseId"])
  },
  methods: {
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
        gwtemp_name: "",
        base: "",
        gangwei: "",
        enterprise: "",
        enterprise_id: "",
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
        dkwx: "",
        gjj: "",
        dkgs: ""
        // qtjq: "",
        // njts: "",
        // cdkk: "",
        // bjts: "",
        // bjkk: "",
        // sjts: "",
        // sjkk: "",
        // kkhj: "",
        // yftz: "",
        // byyf: "",
        // sfgz: ""
      };
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