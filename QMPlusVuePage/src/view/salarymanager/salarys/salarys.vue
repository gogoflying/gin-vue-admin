<template>
  <div>
    <div class="search-term">
      <el-row>
        <el-col :span="4">
          <!-- <el-select v-model="enterprise_id" placeholder="请选择企业">
          <el-option
            v-for="ep in enterpriseInfo"
            :key="ep.enterprise_name"
            :value="ep.ID"
            :label="ep.enterprise_name"
          ></el-option>
          </el-select>-->
          <el-select
            v-show="enPriseId == 0"
            v-model="enterprise_id"
            filterable
            remote
            reserve-keyword
            placeholder="请选择企业"
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
        </el-col>
        <el-col :span="20">
          <el-upload
            :disabled="importDataDisabled"
            style="display: inline-flex;margin-right: 8px;"
            :action="`${path}/un/importsalarys?id=${enterprise_id}`"
            :before-upload="beforeUpload"
            :headers="{'x-token':token}"
            :on-error="onError"
            :on-success="onSuccess"
            :show-file-list="false"
          >
            <el-tooltip :disabled="enPriseId != 0" class="item" effect="dark" content="请先选择企业" placement="top-start">
              <el-button type="success" :icon="importDataBtnIcon">{{importDataBtnText}}</el-button>
            </el-tooltip>
          </el-upload>
          <el-button type="primary" @click="downSalarytemplate" icon="el-icon-download">下载模板</el-button>
          <el-button @click="addSalaryDetail" type="primary" icon="el-icon-plus">新增薪资</el-button>
        </el-col>
      </el-row>
    </div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="姓名">
          <el-input placeholder="姓名" style="width:120px" clearable v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="年">
          <el-input placeholder="年" style="width:80px" clearable v-model.number="searchInfo.year"></el-input>
        </el-form-item>
        <el-form-item label="月">
          <el-input placeholder="月" style="width:80px" clearable v-model.number="searchInfo.month"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary" icon="el-icon-search">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="60" prop="ID" v-if="enPriseId == 0"></el-table-column>
      <el-table-column label="姓名" min-width="100" prop="name"></el-table-column>
      <el-table-column label="年" min-width="60" prop="year"></el-table-column>
      <el-table-column label="月" min-width="40" prop="month"></el-table-column>
      <el-table-column label="所属企业" min-width="120" prop="enterprise" v-if="enPriseId == 0"></el-table-column>
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
      <el-table-column label="假期明细" align="center">
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
      </el-table-column>
      <el-table-column label="五险一金" align="center">
        <el-table-column label="代扣五险" min-width="80" prop="dkwx"></el-table-column>
        <el-table-column label="代扣住房公积金" min-width="80" prop="gjj"></el-table-column>
        <el-table-column label="代扣个人所得税" min-width="80" prop="dkgs"></el-table-column>
      </el-table-column>
      <el-table-column label="应发调整" align="center">
        <el-table-column label="应发工资" min-width="80" prop="byyf"></el-table-column>
        <el-table-column label="应发调整" min-width="80" prop="yftz"></el-table-column>
        <el-table-column label="实发工资" min-width="80" prop="sfgz"></el-table-column>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="180">
        <template slot-scope="scope">
          <el-button
            @click="editSalaryDetail(scope.row)"
            icon="el-icon-edit"
            size="small"
            type="text"
          >编辑</el-button>
          <el-button
            @click="deleteSalaryDetail(scope.row)"
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
      :visible.sync="addSalaryDetailDialog"
      custom-class="user-dialog"
      @close="closeAddSalaryDetailDialog"
      title="新增薪资"
      width="800px"
    >
      <el-form :rules="rules" ref="salarydetailForm" :model="salarydetailinfo">
        <section>
          <h3>基本信息</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item
                v-show="enPriseId == 0"
                label="所属企业"
                label-width="80px"
                prop="enterprise_id"
              >
                <!-- <el-select
                  @change="selectEp"
                  placeholder="请选择企业"
                  v-model="salarydetailinfo.enterprise_id"
                >
                  <el-option
                    :key="industry.enterprise_name"
                    :label="industry.enterprise_name"
                    :value="industry.ID"
                    v-for="industry in enterpriseInfo"
                  ></el-option>
                </el-select>-->
                <el-select
                  @change="selectEp"
                  v-model="salarydetailinfo.enterprise_id"
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
            <el-col :span="6">
              <el-form-item label="薪酬模板" label-width="80px">
                <el-select
                  @change="selectSalaryTemp"
                  v-model="salarytemp_id"
                  filterable
                  remote
                  reserve-keyword
                  placeholder="请选择模板"
                  :remote-method="remoteSearchST"
                  :loading="loading"
                >
                  <el-option
                    v-for="st in salarytemplates"
                    :key="st.gwtemp_name"
                    :value="st.ID"
                    :label="st.gwtemp_name"
                  ></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="姓名" label-width="50px" prop="name">
                <el-input v-model="salarydetailinfo.name"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="3">
              <el-form-item label="年" label-width="40px" prop="year">
                <el-input v-model.number="salarydetailinfo.year"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="3">
              <el-form-item label="月" label-width="40px" prop="month">
                <el-input v-model.number="salarydetailinfo.month"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>薪资构成</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="基本工资" label-width="80px" prop="base">
                <el-input v-model="salarydetailinfo.base"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="岗位工资" label-width="80px" prop="gangwei">
                <el-input v-model="salarydetailinfo.gangwei"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="业绩提成" label-width="80px" prop="yjtc">
                <el-input v-model="salarydetailinfo.yjtc"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="奖金基数" label-width="80px" prop="jjjs">
                <el-input v-model="salarydetailinfo.jjjs"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row style="margin-top:-20px">
            <el-col :span="6">
              <el-form-item label="浮动系数" label-width="80px" prop="fdxs">
                <el-input v-model="salarydetailinfo.fdxs"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="月度奖金" label-width="80px" prop="ydjj">
                <el-input v-model="salarydetailinfo.ydjj"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="工作天数" label-width="80px" prop="gzts">
                <el-input v-model="salarydetailinfo.gzts"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="加班费" label-width="80px" prop="jbf">
                <el-input v-model="salarydetailinfo.jbf"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row style="margin-top:-20px">
            <el-col :span="12">
              <el-form-item label="薪资合计" label-width="80px" prop="xzhj">
                <el-input style="200px" v-model="salarydetailinfo.xzhj"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>补贴明细</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="通讯补贴" label-width="80px" prop="txbt">
                <el-input v-model="salarydetailinfo.txbt"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="餐食补贴" label-width="80px" prop="csbt">
                <el-input v-model="salarydetailinfo.csbt"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="交通补贴" label-width="80px" prop="jtbt">
                <el-input v-model="salarydetailinfo.jtbt"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="其他补贴" label-width="80px" prop="qtbt">
                <el-input v-model="salarydetailinfo.qtbt"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row style="margin-top:-20px">
            <el-col :span="12">
              <el-form-item label="补贴合计" label-width="80px" prop="bthj">
                <el-input style="200px" v-model="salarydetailinfo.bthj"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>假期明细</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="年假天数" label-width="80px" prop="njts">
                <el-input v-model="salarydetailinfo.njts"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="病假天数" label-width="80px" prop="bjts">
                <el-input v-model="salarydetailinfo.bjts"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="事假天数" label-width="80px" prop="sjts">
                <el-input v-model="salarydetailinfo.sjts"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="其他假期" label-width="80px" prop="qtjq">
                <el-input v-model="salarydetailinfo.qtjq"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>扣款明细</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="迟到扣款" label-width="80px" prop="cdkk">
                <el-input v-model="salarydetailinfo.cdkk"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="病假扣款" label-width="80px" prop="bjkk">
                <el-input v-model="salarydetailinfo.bjkk"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="事假扣款" label-width="80px" prop="sjkk">
                <el-input v-model="salarydetailinfo.sjkk"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="扣款合计" label-width="80px" prop="kkhj">
                <el-input v-model="salarydetailinfo.kkhj"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>五险一金</h3>
          <el-row>
            <el-col :span="8">
              <el-form-item label="代扣五险" label-width="80px" prop="dkwx">
                <el-input v-model="salarydetailinfo.dkwx"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="代扣住房公积金" label-width="120px" prop="gjj">
                <el-input v-model="salarydetailinfo.gjj"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="代扣个人所得税" label-width="120px" prop="dkgs">
                <el-input v-model="salarydetailinfo.dkgs"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>应发调整</h3>
          <el-row>
            <el-col :span="8">
              <el-form-item label="应发工资" label-width="80px" prop="byyf">
                <el-input v-model="salarydetailinfo.byyf"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="应发调整" label-width="80px" prop="yftz">
                <el-input v-model="salarydetailinfo.yftz"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="实发工资" label-width="80px" prop="sfgz">
                <el-input v-model="salarydetailinfo.sfgz"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
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
import { getEnterpriseListBySearch } from "@/api/enterpriseinfo";
import { getSalaryTemplatesListBySearch } from "@/api/salarytemplate";
import { mapGetters } from "vuex";
import infoList from "@/components/mixins/infoList";
export default {
  name: "Salarys",
  mixins: [infoList],
  data() {
    return {
      listApi: getSalarysList,
      listKey: "userSalaryList",
      path: path,
      enterprise_id: "",
      salarytemp_id: "",
      addSalaryDetailDialog: false,
      isEdit: false,
      loading: false,
      title: "",
      importDataBtnText: "导入数据",
      importDataBtnIcon: "el-icon-upload2",
      importDataDisabled: true,
      enterpriseInfo: [],
      salarytemplates: [],
      salarydetailinfo: {
        name: "",
        openid: "",
        year: "",
        month: "",
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
      rules: {}
    };
  },
  computed: {
    ...mapGetters("user", ["enPriseId", "token"])
  },
  watch: {
    enterprise_id(val) {
      if (val != null) {
        this.importDataDisabled = false;
      }
    }
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
    remoteSearchST(query) {
      if (query !== "") {
        this.loading = true;
        setTimeout(() => {
          this.loading = false;
          const params = {
            query: query,
            id: this.salarydetailinfo.enterprise_id
          };
          getSalaryTemplatesListBySearch(params).then(res => {
            if (res.success) {
              this.salarytemplates = res.data.result;
            } else {
              this.salarytemplates = [];
            }
          });
        }, 200);
      } else {
        this.salarytemplates = [];
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
      this.salarydetailinfo.enterprise = selectedItem.enterprise_name;
    },
    selectSalaryTemp(val) {
      var selectedItem = {};
      selectedItem = this.salarytemplates.find(item => {
        return item.ID === val;
      });
      this.salarydetailinfo.base = selectedItem.base;
      this.salarydetailinfo.gangwei = selectedItem.gangwei;
      this.salarydetailinfo.xzhj = selectedItem.xzhj;
      this.salarydetailinfo.yjtc = selectedItem.yjtc;
      this.salarydetailinfo.jjjs = selectedItem.jjjs;
      this.salarydetailinfo.fdxs = selectedItem.fdxs;
      this.salarydetailinfo.ydjj = selectedItem.ydjj;
      this.salarydetailinfo.gzts = selectedItem.gzts;
      this.salarydetailinfo.jbf = selectedItem.jbf;
      this.salarydetailinfo.txbt = selectedItem.txbt;
      this.salarydetailinfo.csbt = selectedItem.csbt;
      this.salarydetailinfo.jtbt = selectedItem.jtbt;
      this.salarydetailinfo.qtbt = selectedItem.qtbt;
      this.salarydetailinfo.bthj = selectedItem.bthj;
      this.salarydetailinfo.dkwx = selectedItem.dkwx;
      this.salarydetailinfo.gjj = selectedItem.gjj;
      this.salarydetailinfo.dkgs = selectedItem.dkgs;
    },
    downSalarytemplate() {
      const link = document.createElement("a");
      const url = `${path}/un/template/salarys.xlsx`;
      link.setAttribute("href", url);
      link.setAttribute("download", "薪资模板");
      link.click();
    },
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
        name: "",
        openid: "",
        year: "",
        month: "",
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
      this.enterpriseInfo = [];
      this.salarytemplates = [];
      this.addSalaryDetailDialog = true;
    },
    //编辑企业
    async editSalaryDetail(row) {
      this.title = "编辑薪资";
      this.enterpriseInfo = [];
      this.salarytemplates = [];
      const res = await findSalarys(row);
      this.salarydetailinfo = res.data.reun;
      if (this.salarydetailinfo != null) {
        this.enterpriseInfo.push({
          ID: this.salarydetailinfo.enterprise_id,
          enterprise_name: this.salarydetailinfo.enterprise
        });
      }
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
    },
    onError() {
      this.importDataBtnText = "导入数据";
      this.importDataBtnIcon = "el-icon-upload2";
      this.importDataDisabled = false;
      this.fullscreenLoading = false;
    },
    onSuccess(res) {
      if (res.success) {
        this.$message({
          type: "success",
          message: "上传成功"
        });
        this.getTableData();
      }
      this.importDataBtnText = "导入数据";
      this.importDataBtnIcon = "el-icon-upload2";
      this.importDataDisabled = false;
      this.fullscreenLoading = false;
    },
    beforeUpload() {
      this.importDataBtnText = "正在导入";
      this.importDataBtnIcon = "el-icon-loading";
      this.importDataDisabled = true;
      this.fullscreenLoading = true;
    }
  },
  created() {
    this.importDataDisabled = this.enPriseId == 0 ? true : false;
    // getEnterpriseAllInfo().then(res => {
    //   if (res.success) {
    //     this.enterpriseInfo = res.data.result;
    //   } else {
    //     this.enterpriseInfo = [];
    //   }
    // });
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