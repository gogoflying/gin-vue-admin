<template>
  <div>
    <h2>
      工作经历
      <el-button
        @click="addWork"
        size="small"
        type="text"
        style="margin-left:50%;font-size: 15px;"
        :disabled="enable"
      >新增</el-button>
    </h2>
    <el-table :data="workExperience" border stripe style="width: 100%">
      <el-table-column prop="companyname" label="公司名称" min-width="5"></el-table-column>
      <el-table-column label="操作" min-width="5">
        <template slot-scope="scope">
          <el-button @click="editWork(scope.row)" size="small" type="text">编辑</el-button>
          <el-button @click="deleteWork(scope.row)" size="small" type="text">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-form :rules="rules" ref="workForm" label-width="80px" :model="work_info" v-show="isShow">
      <el-form-item label="公司名称" prop="companyname">
        <el-input v-model="work_info.companyname" placeholder="请输入公司"></el-input>
      </el-form-item>
      <el-form-item label="开始时间" prop="join">
        <el-date-picker
          placeholder="选择开始时间"
          type="month"
          :editable="false"
          format="yyyy-MM"
          value-format="yyyy-MM"
          style="width:80%;"
          v-model="work_info.join"
        ></el-date-picker>
      </el-form-item>
      <el-form-item label="结束时间" prop="leave">
        <el-date-picker
          placeholder="选择结束时间"
          type="month"
          :editable="false"
          format="yyyy-MM"
          value-format="yyyy-MM"
          style="width:80%;"
          v-model="work_info.leave"
        ></el-date-picker>
      </el-form-item>
      <el-form-item label="部门" prop="department">
        <el-input v-model="work_info.department" placeholder="请输入部门"></el-input>
      </el-form-item>
      <el-form-item label="工作内容" prop="workContent">
        <el-input type="textarea" rows="8" v-model="work_info.workContent" placeholder="请输入工作内容"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button style="margin-top:20px;" @click="save" size="small" round>保存</el-button>
        <el-button style="margin-top:20px;" @click="cancle" size="small" round>取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { createUserWork, deleteUserWork, updateUserWork } from "@/api/jobuser";
export default {
  props: ["workExperience", "option", "openid"],
  data() {
    const checkEndTime = (rule, value, callback) => {
      if (value <= this.work_info.join) {
        return callback(new Error("结束时间<=开始时间"));
      } else {
        callback();
      }
    };
    return {
      enable: false,
      isShow: false,
      isEdit: false,
      work_info: {},
      rules: {
        companyname: [
          {
            required: true,
            message: "请输入公司名称",
            trigger: "blur"
          }
        ],
        join: [{ required: true, message: "请选择开始时间", trigger: "blur" }],
        leave: [
          { required: true, message: "请选择结束时间", trigger: "blur" },
          {
            validator: checkEndTime,
            trigger: "blur"
          }
        ],
        department: [
          { required: true, message: "请输入部门", trigger: "blur" }
        ],
        workContent: [
          { required: true, message: "请输入工作内容", trigger: "blur" }
        ]
      }
    };
  },
  computed: {
    keys() {
      return Object.keys(this.workExperience[0]);
    }
  },
  methods: {
    addWork() {
      this.isShow = true;
      this.work_info = {};
      this.workExperience.push(this.work_info);
      this.enable = true;
      this.isEdit = false;
    },
    editWork(row) {
      this.isShow = true;
      this.enable = true;
      this.work_info = row;
      this.isEdit = true;
    },
    deleteWork(row) {
      if (row.ID == null) {
        this.cancle();
        return;
      } else {
        this.$confirm("删除工作经历, 是否继续?", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        })
          .then(async () => {
            const res = await deleteUserWork(row);
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
      this.$refs.workForm.validate(async valid => {
        if (valid) {
          let res;
          if (this.isEdit) {
            res = await updateUserWork(this.work_info);
          } else {
            this.work_info.openid = this.openid;
            res = await createUserWork(this.work_info);
          }
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          } else {
            this.$message({ type: "error", message: "添加失败!" });
          }
          this.$emit("freshResume");
          this.isShow = false;
          this.enable = false;
        }
      });
    },
    cancle() {
      if (this.work_info.ID == null) {
        this.workExperience.pop();
      }
      this.isShow = false;
      this.work_info = {};
      this.enable = false;
    }
  }
};
</script>

<style lang="scss">
// .el-form-item {
//   margin-bottom: 3px;
//   margin-top:10px;
// }
</style>
