<!-- 个人信息 -->
<template>
  <div>
    <h2>
      基本信息
      <el-button
        @click="addPerson"
        size="small"
        type="text"
        style="padding-left:60%;font-size: 15px;"
      >保存</el-button>
    </h2>
    <el-form>
      <el-form-item label="姓名" label-width="80px">
        <el-input v-model="user_base_info.userName" placeholder="请输入姓名"></el-input>
      </el-form-item>
      <el-form-item label="性别" label-width="80px">
        <el-radio-group v-model="user_base_info.genderindex">
          <el-radio :label="0">男</el-radio>
          <el-radio :label="1">女</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="最高学历" label-width="80px">
        <el-select
          @change="selectJobedu"
          placeholder="请输入学历"
          v-model="user_base_info.edulevelindex"
        >
          <el-option :key="item.name" :label="item.name" :value="item.ID" v-for="item in option.el"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="工作年限" label-width="80px">
        <el-select
          @change="selectWorkYear"
          placeholder="请输入工作年限"
          v-model="user_base_info.worksYearindex"
        >
          <el-option
            :key="item.name"
            :label="item.name"
            :value="item.ID"
            v-for="item in option.jwe"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="所在城市" label-width="80px">
        <el-select @change="selectCityName" placeholder="所在城市" v-model="user_base_info.cityindex">
          <el-option
            :key="item.name"
            :label="item.name"
            :value="item.ID"
            v-for="item in option.cityinfo"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="联系电话" label-width="80px">
        <el-input v-model="user_base_info.contact" placeholder="请输入联系电话"></el-input>
      </el-form-item>
      <el-form-item label="联系邮箱" label-width="80px">
        <el-input v-model="user_base_info.email" placeholder="请输入联系邮箱"></el-input>
      </el-form-item>
      <el-form-item label="生日" label-width="80px">
        <el-date-picker
          placeholder="选择结束时间"
          type="date"
          :editable="false"
          format="yyyy-MM-dd"
          value-format="yyyy-MM-dd"
          style="width:80%;"
          v-model="user_base_info.birthday"
        ></el-date-picker>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { createUserBaseInfo, updateUserBaseInfo } from "@/api/jobuser";
export default {
  props: ["user_base_info", "option", "openid"],
  methods: {
    selectCityName(val) {
      var selectedItem = {};
      selectedItem = this.option.cityinfo.find(item => {
        return item.ID === val;
      });
      this.user_base_info.city_name = selectedItem;
    },
    selectWorkYear(val) {
      var selectedItem = {};
      selectedItem = this.option.jwe.find(item => {
        return item.ID === val;
      });
      this.user_base_info.job_work_expire = selectedItem;
    },
    selectJobedu(val) {
      var selectedItem = {};
      selectedItem = this.option.el.find(item => {
        return item.ID === val;
      });
      this.user_base_info.edu_level = selectedItem;
    },
    async addPerson() {
      let res;
      if (this.user_base_info.ID == null) {
        this.user_base_info.openid = this.openid;
        res = await createUserBaseInfo(this.user_base_info);
      } else {
        res = await updateUserBaseInfo(this.user_base_info);
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



