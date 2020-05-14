<template>
  <div>
    <h2>
      教育信息
      <el-button
        @click="addEdu"
        size="small"
        type="text"
        style="padding-left:60%;font-size: 15px;"
        :disabled="enable"
      >新增</el-button>
    </h2>
    <el-table :data="eduExperience" border stripe style="width: 100%">
      <el-table-column prop="schoolname" label="学校名称" min-width="5"></el-table-column>
      <el-table-column label="操作" min-width="5">
        <template slot-scope="scope">
          <el-button @click="editEdu(scope.row)" size="small" type="text">编辑</el-button>
          <el-button @click="deleteEdu(scope.row)" size="small" type="text">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-form :model="edu_info" v-show="isShow">
      <el-form-item label="学校名称" label-width="80px">
        <el-input v-model="edu_info.schoolname" placeholder="请输入学校"></el-input>
      </el-form-item>
      <el-form-item label="开始时间" label-width="80px">
        <el-date-picker
          placeholder="选择开始时间"
          type="date"
          :editable="false"
          format="yyyy/MM"
          value-format="yyyy/MM"
          style="width:80%;"
          v-model="edu_info.graduation"
        ></el-date-picker>
      </el-form-item>
      <el-form-item label="结束时间" label-width="80px">
        <el-date-picker
          placeholder="选择结束时间"
          type="date"
          :editable="false"
          format="yyyy/MM"
          value-format="yyyy/MM"
          style="width:80%;"
          v-model="edu_info.graduationTime"
        ></el-date-picker>
      </el-form-item>
      <el-form-item label="学历" label-width="80px">
        <el-select @change="selectJobedu" placeholder="请输入学历" v-model="edu_info.edulevelindex">
          <el-option :key="item.name" :label="item.name" :value="item.ID" v-for="item in option.el"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="专业" label-width="80px">
        <el-input v-model="edu_info.profession" placeholder="请输入专业"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button style="margin-left:50px;margin-top:20px;" @click="save" size="small" round>保存</el-button>
        <el-button style="margin-left:50px;margin-top:20px;" @click="cancle" size="small" round>取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import {
  createUsereducation,
  deleteUsereducation,
  updateUsereducation
} from "@/api/jobuser";
export default {
  props: ["eduExperience", "option", "openid"],
  data() {
    return {
      enable: false,
      isShow: false,
      isEdit: false,
      edu_info: {}
    };
  },
  computed: {
    keys() {
      return Object.keys(this.eduExperience[0]);
    }
  },
  methods: {
    filter(array) {
      //找出非空对象
      return array.filter(item => !this.isEmpty(item));
    },
    selectJobedu(val) {
      var selectedItem = {};
      selectedItem = this.option.el.find(item => {
        return item.ID === val;
      });
      this.edu_info.edu_level = selectedItem;
    },
    addEdu() {
      this.isShow = true;
      this.edu_info = {};
      this.eduExperience.push(this.edu_info);
      this.enable = true;
      this.isEdit = false;
    },
    editEdu(row) {
      this.isShow = true;
      this.edu_info = row;
      this.isEdit = true;
    },
    deleteEdu(row) {
      if (row.ID == null) {
        this.cancle();
        return;
      } else {
        this.$confirm("删除该教育经历, 是否继续?", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        })
          .then(async () => {
            const res = await deleteUsereducation(row);
            if (res.success) {
              this.$message({
                type: "success",
                message: "删除成功!"
              });
              this.$emit("freshResume");
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
    async save() {
      let res;
      if (this.isEdit) {
        res = await updateUsereducation(this.edu_info);
      } else {
        this.edu_info.openid = this.openid;
        res = await createUsereducation(this.edu_info);
      }
      if (res.success) {
        this.$message({ type: "success", message: "创建成功" });
      } else {
        this.$message({ type: "error", message: "添加失败!" });
      }
      this.$emit("freshResume");
    },
    cancle() {
      if (this.edu_info.ID == null) {
        this.eduExperience.pop();
      }
      this.isShow = false;
      this.edu_info = {};
      this.enable = false;
    }
  }
};
</script>

<style lang="scss">
</style>
