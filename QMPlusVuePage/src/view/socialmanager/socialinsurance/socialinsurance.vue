<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="社保城市">
          <el-select
            placeholder="请选择社保城市"
            v-model="searchInfo.cityindex"
            style="width:92%"
            clearable
            @clear="clearCitystatus"
          >
            <el-option
              :key="city.name"
              :label="city.name"
              :value="city.ID"
              v-for="city in cityinfo"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary" icon="el-icon-search">查询</el-button>
        </el-form-item>

        <el-form-item>
          <el-button @click="addSI" type="primary" icon="el-icon-plus">新增社保</el-button>
        </el-form-item>
      </el-form>
    </div>

    <el-table :data="tableData" border stripe>
      <el-table-column label="ID" min-width="60" type="index" :index="indexMethod"></el-table-column>
      <el-table-column label="参保城市" min-width="80" prop="cityindex" :formatter="cityFormat"></el-table-column>
      <el-table-column label="参保类型" min-width="80" prop="insured_type" :formatter="itFormat"></el-table-column>
      <el-table-column label="养老保险" align="center">
        <el-table-column label="缴费基数" align="center">
          <el-table-column label="下限" min-width="80" prop="yanglao_lower_limit"></el-table-column>
          <el-table-column label="上限" min-width="80" prop="yanglao_upper_limit"></el-table-column>
        </el-table-column>
        <el-table-column label="缴费比例" align="center">
          <el-table-column label="个人" min-width="80" prop="yanglao_person_ratio"></el-table-column>
          <el-table-column label="单位" min-width="80" prop="yanglao_company_ratio"></el-table-column>
        </el-table-column>
      </el-table-column>
      <el-table-column label="医疗保险" align="center">
        <el-table-column label="缴费基数" align="center">
          <el-table-column label="下限" min-width="100" prop="yiliao_lower_limit"></el-table-column>
          <el-table-column label="上限" min-width="100" prop="yiliao_upper_limit"></el-table-column>
        </el-table-column>
        <el-table-column label="缴费比例" align="center">
          <el-table-column label="个人" min-width="80" prop="yiliao_person_ratio"></el-table-column>
          <el-table-column label="单位" min-width="80" prop="yiliao_company_ratio"></el-table-column>
        </el-table-column>
      </el-table-column>
      <el-table-column label="生育保险" align="center">
        <el-table-column label="缴费基数" align="center">
          <el-table-column label="下限" min-width="80" prop="shengyu_lower_limit"></el-table-column>
          <el-table-column label="上限" min-width="80" prop="shengyu_upper_limit"></el-table-column>
        </el-table-column>
        <el-table-column label="缴费比例" align="center">
          <el-table-column label="个人" min-width="80" prop="shengyu_person_ratio"></el-table-column>
          <el-table-column label="单位" min-width="80" prop="shengyu_company_ratio"></el-table-column>
        </el-table-column>
      </el-table-column>
      <el-table-column label="失业保险" align="center">
        <el-table-column label="缴费基数" align="center">
          <el-table-column label="下限" min-width="80" prop="shiye_lower_limit"></el-table-column>
          <el-table-column label="上限" min-width="80" prop="shiye_upper_limit"></el-table-column>
        </el-table-column>
        <el-table-column label="缴费比例" align="center">
          <el-table-column label="个人" min-width="80" prop="shiye_person_ratio"></el-table-column>
          <el-table-column label="单位" min-width="80" prop="shiye_company_ratio"></el-table-column>
        </el-table-column>
      </el-table-column>
      <el-table-column label="工伤保险" align="center">
        <el-table-column label="缴费基数" align="center">
          <el-table-column label="下限" min-width="80" prop="gongshang_lower_limit"></el-table-column>
          <el-table-column label="上限" min-width="80" prop="gongshang_upper_limit"></el-table-column>
        </el-table-column>
        <el-table-column label="缴费比例" align="center">
          <el-table-column label="个人" min-width="80" prop="gongshang_person_ratio"></el-table-column>
          <el-table-column label="单位" min-width="80" prop="gongshang_company_ratio"></el-table-column>
        </el-table-column>
      </el-table-column>
      <el-table-column label="住房公积金" align="center">
        <el-table-column label="缴费基数" align="center">
          <el-table-column label="下限" min-width="80" prop="gjj_lower_limit"></el-table-column>
          <el-table-column label="上限" min-width="80" prop="gjj_upper_limit"></el-table-column>
        </el-table-column>
        <el-table-column label="缴费比例" align="center">
          <el-table-column label="个人" min-width="80" prop="gjj_person_ratio"></el-table-column>
          <el-table-column label="单位" min-width="80" prop="gjj_company_ratio"></el-table-column>
        </el-table-column>
      </el-table-column>

      <el-table-column label="操作" width="180" align="center">
        <template slot-scope="scope">
          <el-button @click="editSI(scope.row)" icon="el-icon-edit" size="small" type="text">编辑</el-button>
          <el-button @click="deleteSI(scope.row)" icon="el-icon-delete" size="small" type="text">删除</el-button>
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
      :visible.sync="addSocialInsuranceDialog"
      custom-class="user-dialog"
      @close="closeaddSocialInsuranceDialog"
      :title="title"
      width="800px"
    >
      <el-form :rules="rules" ref="socialinsuranceForm" :model="socialinsuranceInfo">
        <section>
          <h3>参保信息</h3>
          <el-row>
            <el-col :span="12">
              <el-form-item label="社保城市" label-width="80px" prop="cityindex">
                <el-select placeholder="请选择社保城市" v-model="socialinsuranceInfo.cityindex">
                  <el-option
                    :key="city.name"
                    :label="city.name"
                    :value="city.ID"
                    v-for="city in cityinfo"
                  ></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="社保类型" label-width="80px" prop="insured_type">
                <el-select placeholder="请选择社保类型" v-model="socialinsuranceInfo.insured_type">
                  <el-option
                    :key="it.name"
                    :label="it.name"
                    :value="it.id"
                    v-for="it in insuredtypes"
                  ></el-option>
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>养老保险</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="缴费下限" label-width="80px" prop="yanglao_lower_limit">
                <el-input v-model="socialinsuranceInfo.yanglao_lower_limit"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="缴费上限" label-width="80px" prop="yanglao_upper_limit">
                <el-input v-model="socialinsuranceInfo.yanglao_upper_limit"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="个人比例" label-width="80px" prop="yanglao_person_ratio">
                <el-input v-model="socialinsuranceInfo.yanglao_person_ratio"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="单位比例" label-width="80px" prop="yanglao_company_ratio">
                <el-input v-model="socialinsuranceInfo.yanglao_company_ratio"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>医疗保险</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="缴费下限" label-width="80px" prop="yiliao_lower_limit">
                <el-input v-model="socialinsuranceInfo.yiliao_lower_limit"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="缴费上限" label-width="80px" prop="yiliaoo_upper_limit">
                <el-input v-model="socialinsuranceInfo.yiliao_upper_limit"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="个人比例" label-width="80px" prop="yiliao_person_ratio">
                <el-input v-model="socialinsuranceInfo.yiliao_person_ratio"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="单位比例" label-width="80px" prop="yiliao_company_ratio">
                <el-input v-model="socialinsuranceInfo.yiliao_company_ratio"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>生育保险</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="缴费下限" label-width="80px" prop="shengyu_lower_limit">
                <el-input v-model="socialinsuranceInfo.shengyu_lower_limit"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="缴费上限" label-width="80px" prop="shengyu_upper_limit">
                <el-input v-model="socialinsuranceInfo.shengyu_upper_limit"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="个人比例" label-width="80px" prop="shengyu_person_ratio">
                <el-input v-model="socialinsuranceInfo.shengyu_person_ratio"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="单位比例" label-width="80px" prop="shengyu_company_ratio">
                <el-input v-model="socialinsuranceInfo.shengyu_company_ratio"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>失业保险</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="缴费下限" label-width="80px" prop="shiye_lower_limit">
                <el-input v-model="socialinsuranceInfo.shiye_lower_limit"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="缴费上限" label-width="80px" prop="shiye_upper_limit">
                <el-input v-model="socialinsuranceInfo.shiye_upper_limit"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="个人比例" label-width="80px" prop="shiye_person_ratio">
                <el-input v-model="socialinsuranceInfo.shiye_person_ratio"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="单位比例" label-width="80px" prop="shiye_company_ratio">
                <el-input v-model="socialinsuranceInfo.shiye_company_ratio"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>工伤保险</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="缴费下限" label-width="80px" prop="gongshang_lower_limit">
                <el-input v-model="socialinsuranceInfo.gongshang_lower_limit"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="缴费上限" label-width="80px" prop="gongshang_upper_limit">
                <el-input v-model="socialinsuranceInfo.gongshang_upper_limit"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="个人比例" label-width="80px" prop="gongshang_person_ratio">
                <el-input v-model="socialinsuranceInfo.gongshang_person_ratio"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="单位比例" label-width="80px" prop="gongshang_company_ratio">
                <el-input v-model="socialinsuranceInfo.gongshang_company_ratio"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
        <section>
          <h3>住房公积金</h3>
          <el-row>
            <el-col :span="6">
              <el-form-item label="缴费下限" label-width="80px" prop="gjj_lower_limit">
                <el-input v-model="socialinsuranceInfo.gjj_lower_limit"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="缴费上限" label-width="80px" prop="gjj_upper_limit">
                <el-input v-model="socialinsuranceInfo.gjj_upper_limit"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="个人比例" label-width="80px" prop="gjj_person_ratio">
                <el-input v-model="socialinsuranceInfo.gjj_person_ratio"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="单位比例" label-width="80px" prop="gjj_company_ratio">
                <el-input v-model="socialinsuranceInfo.gjj_company_ratio"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </section>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeaddSocialInsuranceDialog">取 消</el-button>
        <el-button @click="enteraddSocialInsuranceDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {
  getSocialInsuranceList,
  createSocialInsurance,
  updateSocialInsurance,
  findSocialInsurance,
  deleteSocialInsurance,
  getSocialInsuranceOptions
} from "@/api/socialinsurance";
import infoList from "@/components/mixins/infoList";
export default {
  name: "SocialInsurance",
  mixins: [infoList],
  data() {
    const isFloatVlidator = (rule, value, callback) => {
      var pattern = /^\d+.?\d{0,4}$/;
      if (value.length <= 0) {
        return callback(new Error("不能为空"));
      }
      if (!pattern.test(value)) {
        return callback(new Error("请输入非负数并且最多四位小数"));
      } else {
        return callback();
      }
    };
    return {
      listApi: getSocialInsuranceList,
      listKey: "socialInsuranceList",
      path: path,
      addSocialInsuranceDialog: false,
      isEdit: false,
      title: "",
      cityinfo: [],
      socialinsuranceInfo: {
        cityindex: "",
        insured_type: "",
        yanglao_lower_limit: "",
        yanglao_upper_limit: "",
        yanglao_person_ratio: "",
        yanglao_company_ratio: "",
        yiliao_lower_limit: "",
        yiliao_upper_limit: "",
        yiliao_person_ratio: "",
        yiliao_company_ratio: "",
        shengyu_lower_limit: "",
        shengyu_upper_limit: "",
        shengyu_person_ratio: "",
        shengyu_company_ratio: "",
        shiye_lower_limit: "",
        shiye_upper_limit: "",
        shiye_person_ratio: "",
        shiye_company_ratio: "",
        gongshang_lower_limit: "",
        gongshang_upper_limit: "",
        gongshang_person_ratio: "",
        gongshang_company_ratio: "",
        gjj_lower_limit: "",
        gjj_upper_limit: "",
        gjj_person_ratio: "",
        gjj_company_ratio: ""
      },
      rules: {
        cityindex: [
          {
            required: true,
            message: "请输入参保城市",
            trigger: "blur"
          }
        ],
        insured_type: [
          { required: true, message: "请输入参保类型", trigger: "blur" }
        ],
        yanglao_lower_limit: [{ validator: isFloatVlidator, trigger: "blur" }],
        yanglao_upper_limit: [{ validator: isFloatVlidator, trigger: "blur" }],
        yanglao_person_ratio: [{ validator: isFloatVlidator, trigger: "blur" }],
        yanglao_company_ratio: [
          { validator: isFloatVlidator, trigger: "blur" }
        ],
        yiliao_lower_limit: [{ validator: isFloatVlidator, trigger: "blur" }],
        yiliao_upper_limit: [{ validator: isFloatVlidator, trigger: "blur" }],
        yiliao_person_ratio: [{ validator: isFloatVlidator, trigger: "blur" }],
        yiliao_company_ratio: [{ validator: isFloatVlidator, trigger: "blur" }],
        shengyu_lower_limit: [{ validator: isFloatVlidator, trigger: "blur" }],
        shengyu_upper_limit: [{ validator: isFloatVlidator, trigger: "blur" }],
        shengyu_person_ratio: [{ validator: isFloatVlidator, trigger: "blur" }],
        shengyu_company_ratio: [
          { validator: isFloatVlidator, trigger: "blur" }
        ],
        shiye_lower_limit: [{ validator: isFloatVlidator, trigger: "blur" }],
        shiye_upper_limit: [{ validator: isFloatVlidator, trigger: "blur" }],
        shiye_person_ratio: [{ validator: isFloatVlidator, trigger: "blur" }],
        shiye_company_ratio: [{ validator: isFloatVlidator, trigger: "blur" }],
        gongshang_lower_limit: [
          { validator: isFloatVlidator, trigger: "blur" }
        ],
        gongshang_upper_limit: [
          { validator: isFloatVlidator, trigger: "blur" }
        ],
        gongshang_person_ratio: [
          { validator: isFloatVlidator, trigger: "blur" }
        ],
        gongshang_company_ratio: [
          { validator: isFloatVlidator, trigger: "blur" }
        ],
        gjj_lower_limit: [{ validator: isFloatVlidator, trigger: "blur" }],
        gjj_upper_limit: [{ validator: isFloatVlidator, trigger: "blur" }],
        gjj_person_ratio: [{ validator: isFloatVlidator, trigger: "blur" }],
        gjj_company_ratio: [{ validator: isFloatVlidator, trigger: "blur" }]
      },

      insuredtypes: [
        {
          id: 1,
          name: "城镇社保"
        },
        {
          id: 2,
          name: "农村社保"
        },
        {
          id: 3,
          name: "其他"
        }
      ]
    };
  },
  methods: {
    indexMethod(index) {
      return index + 1 + (this.page - 1) * this.pageSize;
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
    cityFormat(row) {
      if (row.cityindex == null || this.cityinfo == null) {
        return "";
      } else {
        var selectedItem = null;
        selectedItem = this.cityinfo.find(item => {
          return item.ID === row.cityindex;
        });
        return selectedItem == null ? "" : selectedItem.name;
      }
    },
    itFormat(row) {
      if (row.insured_type == null || this.insuredtypes == null) {
        return "";
      } else {
        var selectedItem = null;
        selectedItem = this.insuredtypes.find(item => {
          return item.id === row.insured_type;
        });
        return selectedItem == null ? "" : selectedItem.name;
      }
    },
    selectCityName(val) {
      var selectedItem = {};
      selectedItem = this.cityinfo.find(item => {
        return item.ID === val;
      });
      this.socialinsuranceInfo.enterprise_address = selectedItem.name;
    },
    async enteraddSocialInsuranceDialog() {
      this.$refs.socialinsuranceForm.validate(async valid => {
        if (valid) {
          this.socialinsuranceInfo.yanglao_lower_limit = parseFloat(
            this.socialinsuranceInfo.yanglao_lower_limit
          );
          this.socialinsuranceInfo.yanglao_upper_limit = parseFloat(
            this.socialinsuranceInfo.yanglao_upper_limit
          );
          this.socialinsuranceInfo.yanglao_person_ratio = parseFloat(
            this.socialinsuranceInfo.yanglao_person_ratio
          );
          this.socialinsuranceInfo.yanglao_company_ratio = parseFloat(
            this.socialinsuranceInfo.yanglao_company_ratio
          );
          this.socialinsuranceInfo.yiliao_lower_limit = parseFloat(
            this.socialinsuranceInfo.yiliao_lower_limit
          );
          this.socialinsuranceInfo.yiliao_upper_limit = parseFloat(
            this.socialinsuranceInfo.yiliao_upper_limit
          );
          this.socialinsuranceInfo.yiliao_person_ratio = parseFloat(
            this.socialinsuranceInfo.yiliao_person_ratio
          );
          this.socialinsuranceInfo.yiliao_company_ratio = parseFloat(
            this.socialinsuranceInfo.yiliao_company_ratio
          );
          this.socialinsuranceInfo.shengyu_lower_limit = parseFloat(
            this.socialinsuranceInfo.shengyu_lower_limit
          );
          this.socialinsuranceInfo.shengyu_upper_limit = parseFloat(
            this.socialinsuranceInfo.shengyu_upper_limit
          );
          this.socialinsuranceInfo.shengyu_person_ratio = parseFloat(
            this.socialinsuranceInfo.shengyu_person_ratio
          );
          this.socialinsuranceInfo.shengyu_company_ratio = parseFloat(
            this.socialinsuranceInfo.shengyu_company_ratio
          );
          this.socialinsuranceInfo.shiye_lower_limit = parseFloat(
            this.socialinsuranceInfo.shiye_lower_limit
          );
          this.socialinsuranceInfo.shiye_upper_limit = parseFloat(
            this.socialinsuranceInfo.shiye_upper_limit
          );
          this.socialinsuranceInfo.shiye_person_ratio = parseFloat(
            this.socialinsuranceInfo.shiye_person_ratio
          );
          this.socialinsuranceInfo.shiye_company_ratio = parseFloat(
            this.socialinsuranceInfo.shiye_company_ratio
          );
          this.socialinsuranceInfo.gongshang_lower_limit = parseFloat(
            this.socialinsuranceInfo.gongshang_lower_limit
          );
          this.socialinsuranceInfo.gongshang_upper_limit = parseFloat(
            this.socialinsuranceInfo.gongshang_upper_limit
          );
          this.socialinsuranceInfo.gongshang_person_ratio = parseFloat(
            this.socialinsuranceInfo.gongshang_person_ratio
          );
          this.socialinsuranceInfo.gongshang_company_ratio = parseFloat(
            this.socialinsuranceInfo.gongshang_company_ratio
          );
          this.socialinsuranceInfo.gjj_lower_limit = parseFloat(
            this.socialinsuranceInfo.gjj_lower_limit
          );
          this.socialinsuranceInfo.gjj_upper_limit = parseFloat(
            this.socialinsuranceInfo.gjj_upper_limit
          );
          this.socialinsuranceInfo.gjj_person_ratio = parseFloat(
            this.socialinsuranceInfo.gjj_person_ratio
          );
          this.socialinsuranceInfo.gjj_company_ratio = parseFloat(
            this.socialinsuranceInfo.gjj_company_ratio
          );
          let res;
          if (this.isEdit) {
            res = await updateSocialInsurance(this.socialinsuranceInfo);
          } else {
            res = await createSocialInsurance(this.socialinsuranceInfo);
          }
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          } else {
            this.$message({ type: "error", message: "添加失败!" });
          }
          await this.getTableData();
          this.closeaddSocialInsuranceDialog();
        }
      });
    },
    closeaddSocialInsuranceDialog() {
      this.socialinsuranceInfo = {
        cityindex: "",
        insured_type: "",
        yanglao_lower_limit: "",
        yanglao_upper_limit: "",
        yanglao_person_ratio: "",
        yanglao_company_ratio: "",
        yiliao_lower_limit: "",
        yiliao_upper_limit: "",
        yiliao_person_ratio: "",
        yiliao_company_ratio: "",
        shengyu_lower_limit: "",
        shengyu_upper_limit: "",
        shengyu_person_ratio: "",
        shengyu_company_ratio: "",
        shiye_lower_limit: "",
        shiye_upper_limit: "",
        shiye_person_ratio: "",
        shiye_company_ratio: "",
        gongshang_lower_limit: "",
        gongshang_upper_limit: "",
        gongshang_person_ratio: "",
        gongshang_company_ratio: "",
        gjj_lower_limit: "",
        gjj_upper_limit: "",
        gjj_person_ratio: "",
        gjj_company_ratio: ""
      };
      this.$refs.socialinsuranceForm.resetFields();
      this.addSocialInsuranceDialog = false;
    },
    clearCitystatus() {
      this.searchInfo.city_id = null;
    },
    //新增企业
    addSI() {
      this.title = "新增社保";
      this.isEdit = false;
      this.addSocialInsuranceDialog = true;
    },
    //编辑企业
    async editSI(row) {
      this.title = "编辑社保";
      const res = await findSocialInsurance(row);
      this.socialinsuranceInfo = res.data.reet;
      this.isEdit = true;
      this.addSocialInsuranceDialog = true;
    },
    //删除企业
    deleteSI(row) {
      this.$confirm("此操作将永久删除该社保信息, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteSocialInsurance(row);
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
    const res = await getSocialInsuranceOptions();
    this.cityinfo = res.data.cityinfo;
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