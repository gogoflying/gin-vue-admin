<template>
  <div>
    <h2>
      期望职位
      <el-button
        @click="addDream"
        size="small"
        type="text"
        style="padding-left:60%;font-size: 15px;"
      >保存</el-button>
    </h2>
    <el-form>
      <el-form-item label="期望职位" label-width="80px">
        <el-input v-model="user_dreams.dreamposi" placeholder="请输入期望职位"></el-input>
      </el-form-item>
      <el-form-item label="工作类型" label-width="80px">
        <el-select @change="selectWorkType" placeholder="工作类型" v-model="user_dreams.workTypeindex">
          <el-option :key="item.name" :label="item.name" :value="item.ID" v-for="item in option.wt"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="期望薪资" label-width="80px">
        <el-select @change="selectSalary" placeholder="请输入期望薪资" v-model="user_dreams.salaryindex">
          <el-option :key="item.name" :label="item.name" :value="item.ID" v-for="item in option.js"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="期望城市" label-width="80px">
        <el-select @change="selectCityName" placeholder="请输入期望薪资" v-model="user_dreams.cityindex">
          <el-option
            :key="item.name"
            :label="item.name"
            :value="item.ID"
            v-for="item in option.cityinfo"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="到岗时间" label-width="80px">
        <el-select
          @change="selectJobDuty"
          placeholder="请输入到岗时间"
          v-model="user_dreams.dutyTimeindex"
        >
          <el-option
            :key="item.name"
            :label="item.name"
            :value="item.ID"
            v-for="item in option.jdt"
          ></el-option>
        </el-select>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { createUserDream, updateUserDream } from "@/api/jobuser";
export default {
  props: ["user_dreams", "option", "openid"],
  methods: {
    selectWorkType(val) {
      var selectedItem = {};
      selectedItem = this.option.wt.find(item => {
        return item.ID === val;
      });
      this.user_dreams.job_work_type = selectedItem;
    },
    selectSalary(val) {
      var selectedItem = {};
      selectedItem = this.option.js.find(item => {
        return item.ID === val;
      });
      this.user_dreams.job_salary = selectedItem;
    },
    selectCityName(val) {
      var selectedItem = {};
      selectedItem = this.option.cityinfo.find(item => {
        return item.ID === val;
      });
      this.user_dreams.city_name = selectedItem;
    },
    selectJobDuty(val) {
      var selectedItem = {};
      selectedItem = this.option.jdt.find(item => {
        return item.ID === val;
      });
      this.user_dreams.job_duty_time = selectedItem;
    },
    async addDream() {
      let res;
      if (this.user_dreams.ID == null) {
        this.user_dreams.openid = this.openid;
        res = await createUserDream(this.user_dreams);
      } else {
        res = await updateUserDream(this.user_dreams);
      }
      if (res.success) {
        this.$message({ type: "success", message: "创建成功" });
      } else {
        this.$message({ type: "error", message: "添加失败!" });
      }
      this.$emit("freshResume");
    }
  }
};
</script>
<style lang="scss">
.el-form-item {
  margin-bottom: 3px;
}
</style>
