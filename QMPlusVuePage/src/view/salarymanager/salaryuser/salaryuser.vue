<template>
  <div v-loading.fullscreen.lock="fullscreenLoading">
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
            :action="`${path}/un/importsalaryuser?id=${enterprise_id}`"
            :before-upload="beforeUpload"
            :headers="{'x-token':token}"
            :on-error="onError"
            :on-success="onSuccess"
            :show-file-list="false"
          >
            <el-tooltip class="item" effect="dark" content="请先选择企业" placement="top-start">
              <el-button type="success" :icon="importDataBtnIcon">{{importDataBtnText}}</el-button>
            </el-tooltip>
          </el-upload>
          <el-button type="primary" @click="downSalarytemplate" icon="el-icon-download">下载模板</el-button>
          <el-button type="primary" @click="addSalaryUser" icon="el-icon-plus">新增员工</el-button>
        </el-col>
      </el-row>
    </div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" label-width="60px">
        <el-form-item label="姓名">
          <el-input placeholder="姓名" v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="身份证">
          <el-input placeholder="身份证" v-model="searchInfo.card"></el-input>
        </el-form-item>
        <el-form-item label="手机号">
          <el-input placeholder="手机号" v-model="searchInfo.mobile"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary" icon="el-icon-search">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column type="expand">
        <template slot-scope="scope">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="身份证正反面:">
              <el-row v-if="scope.row.cards != null ">
                <el-col :span="12" v-for="cd in scope.row.cards" :key="cd">
                  <img :src="cd" height="100%" width="100%" />
                </el-col>
              </el-row>
              <span v-else style="color:red">身份证未上传</span>
            </el-form-item>
            <el-form-item label="银行卡正反面:">
              <el-row v-if="scope.row.banks != null ">
                <el-col :span="12" v-for="cd in scope.row.banks" :key="cd">
                  <img :src="cd" height="100%" width="100%" />
                </el-col>
              </el-row>
              <span v-else style="color:red">银行卡未上传</span>
            </el-form-item>
            <el-form-item label="户口本照片:">
              <el-row v-if="scope.row.hukous != null ">
                <el-col :span="12" v-for="cd in scope.row.hukous" :key="cd">
                  <img :src="cd" height="100%" width="100%" />
                </el-col>
              </el-row>
              <span v-else style="color:red">户口本未上传</span>
            </el-form-item>
            <el-form-item label="本人照片:">
              <el-row v-if="scope.row.photos != null ">
                <el-col :span="12" v-for="cd in scope.row.photos" :key="cd">
                  <img :src="cd" height="100%" width="100%" />
                </el-col>
              </el-row>
              <span v-else style="color:red">本人照片未上传</span>
            </el-form-item>
            <el-form-item label="体检报告合格照片:">
              <el-row v-if="scope.row.checkups != null ">
                <el-col :span="12" v-for="cd in scope.row.checkups" :key="cd">
                  <img :src="cd" height="100%" width="100%" />
                </el-col>
              </el-row>
              <span v-else style="color:red">体检报告未上传</span>
            </el-form-item>
            <el-form-item label="离职证明照片:">
              <span v-if="scope.row.leaves != null ">
                <ul v-for="cd in scope.row.leaves" :key="cd">
                  <li>
                    <img :src="cd" height="500" width="500" />
                  </li>
                </ul>
              </span>
              <span v-else style="color:red">离职证明未上传</span>
            </el-form-item>
            <el-form-item label="毕业证照片:">
              <el-row v-if="scope.row.diplomas != null ">
                <el-col :span="12" v-for="cd in scope.row.diplomas" :key="cd">
                  <img :src="cd" height="100%" width="100%" />
                </el-col>
              </el-row>
              <span v-else style="color:red">毕业证未上传</span>
            </el-form-item>
            <el-form-item label="其他:">
              <el-row v-if="scope.row.others != null ">
                <el-col :span="12" v-for="cd in scope.row.others" :key="cd">
                  <img :src="cd" height="100%" width="100%" />
                </el-col>
              </el-row>
              <span v-else style="color:red">未上传</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column label="id" min-width="60" prop="ID"></el-table-column>
      <el-table-column label="姓名" min-width="80" prop="name"></el-table-column>
      <!-- <el-table-column label="手机" min-width="110" prop="mobile"></el-table-column> -->
      <el-table-column label="岗位名称" min-width="100" prop="job_name"></el-table-column>
      <el-table-column label="合同期限(月)" min-width="120" prop="contract_date" align="center"></el-table-column>
      <!-- <el-table-column label="邮箱" min-width="100" prop="email"></el-table-column> -->
      <el-table-column label="入职企业" min-width="100" prop="enterprise" v-if="enPriseId == 0"></el-table-column>
      <el-table-column label="入职进度" min-width="100" prop="enter_step" :formatter="EnstepFormat"></el-table-column>
      <el-table-column label="离职进度" min-width="100" prop="leave_step" :formatter="LestepFormat"></el-table-column>
      <!-- <el-table-column label="工资" min-width="100" prop="salary"></el-table-column>
      <el-table-column label="身份证号" min-width="100" prop="card" :show-overflow-tooltip="true"></el-table-column>
      <el-table-column label="入职日期" min-width="100" prop="enter_time"></el-table-column>
      <el-table-column label="离职日期" min-width="100" prop="leave_time"></el-table-column>
      <el-table-column label="入职进度" min-width="150" prop="enter_step">
        <template slot-scope="scope">
          <el-select
            @change="handleRowEdit(scope.$index, scope.row)"
            placeholder="请选择"
            v-model="scope.row.enter_step"
          >
            <el-option :key="es.name" :label="es.name" :value="es.id" v-for="es in enterSteps"></el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column label="离职进度" min-width="150" prop="leave_step">
        <template slot-scope="scope">
          <el-select
            @change="handleRowEdit(scope.$index, scope.row)"
            placeholder="请选择"
            v-model="scope.row.leave_step"
          >
            <el-option :key="ls.name" :label="ls.name" :value="ls.id" v-for="ls in leaveSteps"></el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column label="备注" min-width="100" prop="errors">
        <template slot-scope="scope">
          <el-input
            size="small"
            v-model="scope.row.errors"
            placeholder="请输入备注"
            @change="handleRowEdit(scope.$index, scope.row)"
          ></el-input>
          <span>{{scope.row.date}}</span>
        </template>
      </el-table-column>-->
      <el-table-column label="操作" width="350px">
        <template slot-scope="scope">
          <el-button
            @click="editSalaryUser(scope.row)"
            icon="el-icon-edit"
            size="small"
            type="text"
          >编辑</el-button>

          <el-upload
            :disabled="scope.row.enter_step < 2"
            :action="`${path}/un/importUserContract?openid=${scope.row.openid}`"
            style="display: inline-block;"
            :headers="{'x-token':token}"
            :on-error="onErrorUc"
            :on-success="onSuccessUc"
            :before-upload="beforeUploadPdf"
            :show-file-list="false"
          >
            <el-button
              icon="el-icon-upload"
              size="small"
              type="text"
              :disabled="scope.row.enter_step < 2"
            >上传合同</el-button>
          </el-upload>
          <el-button
            @click="viewContract(scope.row)"
            icon="el-icon-view"
            size="small"
            type="text"
            :disabled="scope.row.enter_step < 3 "
          >查看合同</el-button>
          <el-button
            @click="deleteSalaryUser(scope.row)"
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
      :visible.sync="addSalaryUserDialog"
      custom-class="user-dialog"
      @close="closeAddSalaryUserDialog"
      title="新增员工"
      width="600px"
    >
      <el-form :rules="rules" ref="salaryuserForm" :model="salaryuserinfo">
        <el-form-item label="姓名" label-width="80px" prop="name">
          <el-input v-model="salaryuserinfo.name" style="width:90%"></el-input>
        </el-form-item>
        <el-form-item label="手机" label-width="80px" prop="mobile">
          <el-input v-model="salaryuserinfo.mobile" style="width:90%"></el-input>
        </el-form-item>
        <el-form-item label="身份证号" label-width="80px" prop="card">
          <el-input v-model="salaryuserinfo.card" style="width:90%"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" label-width="80px" prop="email">
          <el-input v-model="salaryuserinfo.email" style="width:90%"></el-input>
        </el-form-item>
        <el-form-item label="岗位名称" label-width="80px" prop="job_name">
          <el-input v-model="salaryuserinfo.job_name" style="width:90%"></el-input>
        </el-form-item>
        <el-form-item label="岗位工资" label-width="80px" prop="salary">
          <el-input v-model.number="salaryuserinfo.salary" style="width:90%"></el-input>
        </el-form-item>
        <el-form-item label="合同期限" label-width="80px" prop="contract_date">
          <el-input v-model.number="salaryuserinfo.contract_date" style="width:90%"></el-input>
          <span>月</span>
        </el-form-item>
        <el-form-item v-show="enPriseId == 0" label="入职企业" label-width="90px" prop="enterprise_id">
          <!-- <el-select @change="selectEp" placeholder="请选择企业" v-model="salaryuserinfo.enterprise_id">
            <el-option
              :key="industry.enterprise_name"
              :label="industry.enterprise_name"
              :value="industry.ID"
              v-for="industry in enterpriseInfo"
            ></el-option>
          </el-select>-->
          <el-select
            @change="selectEp"
            v-model="salaryuserinfo.enterprise_id"
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
        <el-form-item label="入职日期" label-width="80px" prop="enter_time">
          <el-date-picker
            placeholder="选择入职日期"
            type="date"
            :editable="false"
            format="yyyy 年 MM 月 dd 日"
            value-format="yyyy-MM-dd"
            v-model="salaryuserinfo.enter_time"
          ></el-date-picker>
        </el-form-item>
        <el-form-item label="离职日期" label-width="80px" prop="leave_time">
          <el-date-picker
            placeholder="选择离职日期"
            type="date"
            :editable="false"
            format="yyyy 年 MM 月 dd 日"
            value-format="yyyy-MM-dd"
            v-model="salaryuserinfo.leave_time"
          ></el-date-picker>
        </el-form-item>
        <el-form-item label="入职进度" label-width="80px" prop="enter_step">
          <el-select placeholder="请选择入职进度" v-model="salaryuserinfo.enter_step">
            <el-option :key="es.name" :label="es.name" :value="es.id" v-for="es in enterSteps"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="离职进度" label-width="80px" prop="leave_step">
          <el-select placeholder="请选择离职进度" v-model="salaryuserinfo.leave_step">
            <el-option :key="ls.name" :label="ls.name" :value="ls.id" v-for="ls in leaveSteps"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-show="isShow" label="备注" label-width="80px" prop="errors">
          <el-input v-model="salaryuserinfo.errors"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddSalaryUserDialog">取 消</el-button>
        <el-button @click="enterAddSalaryUserDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="viewContractDialog" custom-class="user-dialog" title="查看合同">
      <div id="ShowBox">
        <div id="show">
          <section v-if="user_cons.length > 0 ">
            <ul v-for="usercon in user_cons" :key="usercon">
              <li>
                <img
                  :src="usercon + '?' + new Date().getTime()"
                  style="max-width:100%;"
                  crossorigin="anonymous"
                />
              </li>
            </ul>
          </section>
        </div>
      </div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="download" type="primary">导 出</el-button>
        <el-button @click="viewContractDialog=false">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {
  getSalaryUsersList,
  createSalaryUsers,
  updateSalaryUsers,
  findSalaryUsers,
  deleteSalaryUsers,
  uploadUserContract,
  downloadContractList
} from "@/api/salaryuser";
import { getEnterpriseListBySearch } from "@/api/enterpriseinfo";
import infoList from "@/components/mixins/infoList";
import { mapGetters } from "vuex";
import html2canvas from "html2canvas";
import jspdf from "jspdf";
export default {
  name: "Salaryuser",
  mixins: [infoList],
  data() {
    return {
      fullscreenLoading: false,
      listApi: getSalaryUsersList,
      listKey: "userSalaryList",
      path: path,
      enterprise_id: "",
      addSalaryUserDialog: false,
      viewContractDialog: false,
      user_cons: [],
      isEdit: false,
      isShow: false,
      title: "",
      loading: false,
      importDataBtnText: "导入数据",
      importDataBtnIcon: "el-icon-upload2",
      importDataDisabled: true,
      enterpriseInfo: [],
      salaryuserinfo: {
        openid: "",
        name: "",
        mobile: "",
        card: "",
        email: "",
        job_name: "",
        salary: 0,
        contract_date: "",
        enterprise: "",
        enterprise_id: "",
        enter_time: "",
        leave_time: "",
        enter_step: "",
        leave_step: "",
        errors: ""
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
        enter_time: [
          { required: true, message: "请输入入职日期", trigger: "blur" }
        ]
      },
      enterSteps: [
        {
          id: 0,
          name: "已录入"
        },
        {
          id: 1,
          name: "待补充材料"
        },
        {
          id: 2,
          name: "待审核"
        },
        {
          id: 3,
          name: "已上传合同"
        },
        {
          id: 4,
          name: "已入职"
        }
      ],
      leaveSteps: [
        {
          id: 0,
          name: "在职"
        },
        {
          id: 1,
          name: "离职申请"
        },
        {
          id: 2,
          name: "已审批"
        },
        {
          id: 3,
          name: "已离职"
        }
      ]
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
    },
    "salaryuserinfo.enter_step": "changeData",
    "salaryuserinfo.leave_step": "changeData"
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
      this.salaryuserinfo.enterprise = selectedItem.enterprise_name;
    },
    changeData(curval, oldval) {
      if (curval < oldval) {
        this.isShow = true;
      }
    },
    EnstepFormat(row) {
      var selectedItem = {};
      selectedItem = this.enterSteps.find(item => {
        return item.id === row.enter_step;
      });
      return selectedItem.name;
    },
    LestepFormat(row) {
      var selectedItem = {};
      selectedItem = this.leaveSteps.find(item => {
        return item.id === row.leave_step;
      });
      return selectedItem.name;
    },
    downSalarytemplate() {
      const link = document.createElement("a");
      const url = `${path}/un/template/salaryuser.xlsx`;
      link.setAttribute("href", url);
      link.setAttribute("download", "员工模板");
      link.click();
    },
    async handleRowEdit(index, row) {
      const res = await updateSalaryUsers(row);
      if (res.success) {
        this.$message({ type: "success", message: "修改成功" });
      }
    },
    async enterAddSalaryUserDialog() {
      this.$refs.salaryuserForm.validate(async valid => {
        if (valid) {
          let res;
          if (this.isEdit) {
            res = await updateSalaryUsers(this.salaryuserinfo);
          } else {
            res = await createSalaryUsers(this.salaryuserinfo);
          }
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          } else {
            this.$message({ type: "error", message: "添加失败!" });
          }
          await this.getTableData();
          this.closeAddSalaryUserDialog();
        }
      });
    },
    closeAddSalaryUserDialog() {
      this.salaryuserinfo = {
        openid: "",
        name: "",
        mobile: "",
        card: "",
        email: "",
        job_name: "",
        salary: "",
        contract_date: "",
        enterprise: "",
        enterprise_id: "",
        enter_time: "",
        leave_time: "",
        enter_step: "",
        leave_step: "",
        errors: ""
      };
      this.addSalaryUserDialog = false;
    },
    //新增企业
    addSalaryUser() {
      this.title = "新增员工";
      this.isEdit = false;
      this.enterpriseInfo = [];
      this.addSalaryUserDialog = true;
    },
    //编辑企业
    async editSalaryUser(row) {
      this.title = "编辑员工";
      this.enterpriseInfo = [];
      const res = await findSalaryUsers(row);
      this.salaryuserinfo = res.data.reun;
      if (this.salaryuserinfo != null) {
        this.enterpriseInfo.push({
          ID: this.salaryuserinfo.enterprise_id,
          enterprise_name: this.salaryuserinfo.enterprise
        });
      }
      this.isEdit = true;
      this.addSalaryUserDialog = true;
    },
    //删除企业
    deleteSalaryUser(row) {
      this.$confirm("此操作将永久删除该员工信息, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteSalaryUsers(row);
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
    },
    beforeUploadPdf(file) {
      var testmsg = file.name.substring(file.name.lastIndexOf(".") + 1);
      const extension = testmsg === "pdf";
      const isLt1MB = file.size / 1024 / 1024 < 1;
      if (!extension) {
        this.$message({
          message: "上传文件只能是 jpg、png格式!",
          type: "warning"
        });
        return false;
      }
      if (!isLt1MB) {
        this.$message({
          message: "上传文件大小不能超过 1MB!",
          type: "warning"
        });
        return false;
      }
      return extension && isLt1MB;
    },
    onErrorUc() {
      this.$message({
        type: "error",
        message: "上传合同成功"
      });
    },
    onSuccessUc(res) {
      if (res.success) {
        this.$message({
          type: "success",
          message: "上传合同成功"
        });
        uploadUserContract(res.data).then(res => {
          if (res.success) {
            this.$message({
              type: "success",
              message: "上传合同写入数据库成功"
            });
          } else {
            this.$message({
              type: "error",
              message: "上传合同写入数据库失败"
            });
          }
        });
      } else {
        this.$message({
          type: "error",
          message: "上传合同失败"
        });
      }
    },
    async viewContract(row) {
      const res = await downloadContractList({ openid: row.openid });
      if (res.success) {
        this.user_cons = res.data.jpgList;
        this.viewContractDialog = true;
      }
    },
    closeviewContractDialog() {
      this.user_cons = [];
      this.viewContractDialog = false;
    },
    onCreated() {
      this.importDataDisabled = this.enPriseId == 0 ? true:false;
      // getEnterpriseAllInfo().then(res => {
      //   if (res.success) {
      //     this.enterpriseInfo = res.data.result;
      //   } else {
      //     this.enterpriseInfo = [];
      //   }
      // });
    },
    download() {
      //let resume = document.getElementById("ShowBox");
      //let c = document.createElement("canvas");
      let opts = {
        useCORS: true
        // scale: 2,
        // canvas: c,
        // logging: true,
        // width: document.getElementById("ShowBox").clientWidth,
        // height: document.getElementById("ShowBox").clientHeight
      };
      // c.width = document.getElementById("ShowBox").clientWidth * 2;
      // c.height = document.getElementById("ShowBox").clientHeight * 2;
      // c.getContext("2d").scale(2, 2);
      html2canvas(document.querySelector("#ShowBox"), opts).then(canvas => {
        var contentWidth = canvas.width;
        var contentHeight = canvas.height;
        //一页pdf显示html页面生成的canvas高度;
        var pageHeight = (contentWidth / 592.28) * 841.89;
        //未生成pdf的html页面高度
        var leftHeight = contentHeight;
        //页面偏移
        var position = 0;
        //a4纸的尺寸[595.28,841.89]，html页面生成的canvas在pdf中图片的宽高
        var imgWidth = 595.28;
        var imgHeight = (592.28 / contentWidth) * contentHeight;

        var pageData = canvas.toDataURL("image/jpeg", 1.0);

        var pdf = new jspdf("", "pt", "a4");

        //有两个高度需要区分，一个是html页面的实际高度，和生成pdf的页面高度(841.89)
        //当内容未超过pdf一页显示的范围，无需分页
        if (leftHeight < pageHeight) {
          pdf.addImage(pageData, "JPEG", 0, 0, imgWidth, imgHeight);
        } else {
          while (leftHeight > 0) {
            pdf.addImage(pageData, "JPEG", 0, position, imgWidth, imgHeight);
            leftHeight -= pageHeight;
            position -= 841.89;
            //避免添加空白页
            if (leftHeight > 0) {
              pdf.addPage();
            }
          }
        }
        pdf.save("contract .pdf");
      });
    }
  },
  created() {
    this.onCreated();
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

.demo-table-expand {
  font-size: 0;
}
.demo-table-expand label {
  width: 150px;
  color: #4e5868;
}
.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 100%;
}
</style>