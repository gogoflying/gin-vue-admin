<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="姓名">
          <el-input placeholder="姓名" clearable v-model="searchInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="状态" prop="cityindex">
          <el-select
            placeholder="请选择订单状态"
            v-model="searchInfo.status"
            clearable
            @clear="clearOrderstatus"
          >
            <el-option :key="ot.name" :label="ot.name" :value="ot.id" v-for="ot in orderstatus"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary" icon="el-icon-search">查询</el-button>
        </el-form-item>

        <!-- <el-form-item>
          <el-button @click="addSI" type="primary" icon="el-icon-plus">新增社保</el-button>
        </el-form-item>-->
      </el-form>
    </div>

    <el-table :data="tableData" border stripe>
      <el-table-column label="ID" min-width="60" type="index" :index="indexMethod"></el-table-column>
      <el-table-column label="订单ID" min-width="100" prop="order_id" align="center"></el-table-column>
      <el-table-column label="参保人姓名" min-width="100" prop="name" align="center"></el-table-column>
      <el-table-column label="参保人身份证" min-width="150" prop="id_card" align="center"></el-table-column>
      <el-table-column label="参保类型" min-width="80" prop="insured_type" :formatter="itFormat" align="center"></el-table-column>
      <el-table-column label="订单状态" min-width="120" prop="status" align="center">
        <template slot-scope="scope">
          <el-select
            @change="changeOrderStatus(scope.row)"
            placeholder="请选择"
            v-model="scope.row.status"
          >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.id"
              v-for="item in orderstatus"
            ></el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column label="开始时间" min-width="100" prop="start_time" align="center"></el-table-column>
      <el-table-column label="结束时间" min-width="100" prop="end_time" align="center"></el-table-column>
      <el-table-column label="参保时长(月)" min-width="120" prop="duration" align="center"></el-table-column>
      <el-table-column label="缴费总金额" min-width="80" prop="pay_amount"></el-table-column>
      <el-table-column label="单月明细" align="center">
        <el-table-column label="五险" align="center">
          <el-table-column label="是否缴纳" min-width="80" prop="is_ins" :formatter="isInsFormat"></el-table-column>
          <el-table-column label="缴费基数" min-width="100" prop="ins_base"></el-table-column>
          <el-table-column label="养老" min-width="80" prop="endowment_ins"></el-table-column>
          <el-table-column label="医疗" min-width="80" prop="medical_ins"></el-table-column>
          <el-table-column label="生育" min-width="80" prop="maternity_ins"></el-table-column>
          <el-table-column label="失业" min-width="80" prop="unemployment_ins"></el-table-column>
          <el-table-column label="工伤" min-width="80" prop="employment_injury_ins"></el-table-column>
          <el-table-column label="五险合计" min-width="80" prop="insurance_fee"></el-table-column>
        </el-table-column>
        <el-table-column label="公积金" align="center">
          <el-table-column label="是否缴纳" min-width="80" prop="is_gjj" :formatter="isGjjFormat"></el-table-column>
          <el-table-column label="缴费基数" min-width="100" prop="ins_base"></el-table-column>
          <el-table-column label="公积金合计" min-width="80" prop="gjj_fee"></el-table-column>
        </el-table-column>
        <el-table-column label="服务费" min-width="80" prop="serve_fee"></el-table-column>
        <el-table-column label="单月合计" min-width="80" prop="total_fee"></el-table-column>
      </el-table-column>

      <el-table-column label="所属社保" min-width="100" align="center">
        <template slot-scope="scope">
          <el-tooltip placement="right" v-if="scope.row.si_info">
            <div slot="content">
              <table>
                <tr>
                  <td>养老[</td>
                  <td>缴费下限:</td>
                  <td>{{scope.row.si_info.yanglao_lower_limit}}</td>
                  <td>缴费上限:</td>
                  <td>{{scope.row.si_info.yanglao_upper_limit}}</td>
                  <td>个人缴费比例:</td>
                  <td>{{scope.row.si_info.yanglao_person_ratio}}</td>
                  <td>单位缴费比例:</td>
                  <td>{{scope.row.si_info.yanglao_company_ratio}}</td>
                  <td>]</td>
                </tr>
                <tr>
                  <td>医疗[</td>
                  <td>缴费下限:</td>
                  <td>{{scope.row.si_info.yiliao_lower_limit}}</td>
                  <td>缴费上限:</td>
                  <td>{{scope.row.si_info.yiliao_upper_limit}}</td>
                  <td>个人缴费比例:</td>
                  <td>{{scope.row.si_info.yiliao_person_ratio}}</td>
                  <td>单位缴费比例:</td>
                  <td>{{scope.row.si_info.yiliao_company_ratio}}</td>
                  <td>]</td>
                </tr>
                <tr>
                  <td>生育[</td>
                  <td>缴费下限:</td>
                  <td>{{scope.row.si_info.shengyu_lower_limit}}</td>
                  <td>缴费上限:</td>
                  <td>{{scope.row.si_info.shengyu_upper_limit}}</td>
                  <td>个人缴费比例:</td>
                  <td>{{scope.row.si_info.shengyu_person_ratio}}</td>
                  <td>单位缴费比例:</td>
                  <td>{{scope.row.si_info.shengyu_company_ratio}}</td>
                  <td>]</td>
                </tr>
                <tr>
                  <td>失业[</td>
                  <td>缴费下限:</td>
                  <td>{{scope.row.si_info.shiye_lower_limit}}</td>
                  <td>缴费上限:</td>
                  <td>{{scope.row.si_info.shiye_upper_limit}}</td>
                  <td>个人缴费比例:</td>
                  <td>{{scope.row.si_info.shiye_person_ratio}}</td>
                  <td>单位缴费比例:</td>
                  <td>{{scope.row.si_info.shiye_company_ratio}}</td>
                  <td>]</td>
                </tr>
                <tr>
                  <td>工伤[</td>
                  <td>缴费下限:</td>
                  <td>{{scope.row.si_info.gongshang_lower_limit}}</td>
                  <td>缴费上限:</td>
                  <td>{{scope.row.si_info.gongshang_upper_limit}}</td>
                  <td>个人缴费比例:</td>
                  <td>{{scope.row.si_info.gongshang_person_ratio}}</td>
                  <td>单位缴费比例:</td>
                  <td>{{scope.row.si_info.gongshang_company_ratio}}</td>
                  <td>]</td>
                </tr>
                <tr>
                  <td>公积金[</td>
                  <td>缴费下限:</td>
                  <td>{{scope.row.si_info.gjj_lower_limit}}</td>
                  <td>缴费上限:</td>
                  <td>{{scope.row.si_info.gjj_upper_limit}}</td>
                  <td>个人缴费比例:</td>
                  <td>{{scope.row.si_info.gjj_person_ratio}}</td>
                  <td>单位缴费比例:</td>
                  <td>{{scope.row.si_info.gjj_company_ratio}}</td>
                  <td>]</td>
                </tr>
              </table>
            </div>
            <el-tag>社保详情</el-tag>
          </el-tooltip>
          <el-tag v-else>暂未设置</el-tag>
        </template>
      </el-table-column>

      <el-table-column label="操作" width="100">
        <template slot-scope="scope">
          <el-button @click="deleteSO(scope.row)" icon="el-icon-delete" size="small" type="text">删除</el-button>
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
import { getSocialOrderList, updateSocialOrder, deleteSocialOrder } from "@/api/socialorder";
import infoList from "@/components/mixins/infoList";
export default {
  name: "SocialOrder",
  mixins: [infoList],
  data() {
    return {
      listApi: getSocialOrderList,
      listKey: "socialOrderList",
      path: path,
      orderstatus: [
        {
          id: 1,
          name: "待支付"
        },
        {
          id: 2,
          name: "待激活"
        },
        {
          id: 3,
          name: "已完成"
        }
      ],
      insuredtypes: [
        {
          id: 1,
          name: "本地城镇"
        },
        {
          id: 2,
          name: "本地农村"
        },
        {
          id: 3,
          name: "外地城镇"
        },
        {
          id: 4,
          name: "外地农村"
        },
        {
          id: 5,
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
    isInsFormat(row) {
      if (row.is_ins == 1) {
        return "是";
      } else {
        return "否";
      }
    },
    isGjjFormat(row) {
      if (row.is_gjj == 1) {
        return "是";
      } else {
        return "否";
      }
    },
    clearOrderstatus() {
      this.searchInfo.status = null;
    },
    async changeOrderStatus(row) {
      const res = await updateSocialOrder(row);
      if (res.success) {
        this.$message({ type: "success", message: "状态设置成功" });
      }
    },
    //删除企业
    deleteSO(row) {
      this.$confirm("此操作将永久删除该订单信息, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(async () => {
          const res = await deleteSocialOrder(row);
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