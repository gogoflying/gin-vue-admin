<template>
  <div>
    <div>
      <el-form :rules="rules" ref="newForm" :model="user_news" label-width="80px">
        <el-form-item label="新闻标题" label-width="80px" prop="title">
          <el-input v-model="user_news.title" style="width:50%;" placeholder="请输入新闻标题"></el-input>
        </el-form-item>
        <el-form-item label="副标题" prop="sub_title">
          <el-input v-model="user_news.sub_title" style="width:50%;" placeholder="请输入新闻副标题"></el-input>
        </el-form-item>
        <el-form-item label="排序" prop="order">
          <el-input v-model.number="user_news.order" style="width:25%;" placeholder="请输入顺序"></el-input>
          <label>【数值越大,越靠前】</label>
        </el-form-item>
        <el-form-item label="新闻类型" prop="news_type">
          <el-select placeholder="请选择新闻类型" v-model.number="user_news.news_type">
            <el-option :key="nt.name" :label="nt.name" :value="nt.id" v-for="nt in newtypes"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="新闻图片" prop="img">
          <el-upload
            :headers="{'x-token':token}"
            :on-success="handleAvatarSuccess"
            :show-file-list="false"
            :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
            class="avatar-uploader"
            name="file"
            :before-upload="beforeAvatarUpload"
          >
            <img :src="user_news.img" class="avatar" v-if="user_news.img" height="300" width="300" />
            <i class="el-icon-plus avatar-uploader-icon" v-else></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="新闻内容" prop="content">
          <quill-editor
            class="ql-editor"
            :options="editorOption"
            @blur="onEditorBlur($event)"
            @change="onEditorChange($event)"
            @focus="onEditorFocus($event)"
            ref="myQuillEditor"
            v-model="user_news.content"
          ></quill-editor>
        </el-form-item>
        <el-row type="flex" justify="center">
          <el-col :span="13">
            <el-button plain @click="saveDrft">存入草稿</el-button>
            <el-button plain @click="savePublish">发布</el-button>
          </el-col>
        </el-row>
      </el-form>
    </div>
  </div>
</template>

<script>
const path = process.env.VUE_APP_BASE_API;
import { createUserNews, updateUserNews, findUserNews } from "@/api/newmanager";
export default {
  name: "NewDetail",
  data() {
    return {
      isEdit: false,
      editorOption: {},
      path: path,
      user_news: {
        title: "",
        sub_title: "",
        content: "",
        order: "",
        news_type: "",
        img: "",
        count: "",
        status: ""
      },
      newtypes: [
        {
          id: 1,
          name: "社保"
        },
        {
          id: 2,
          name: "个税"
        }
      ],
      rules: {
        title: [{ required: true, message: "请输入标题", trigger: "blur" }],
        news_type: [{ required: true, message: "请输入类型", trigger: "blur" }],
        content: [{ required: true, message: "请输入内容", trigger: "blur" }]
      }
    };
  },
  computed: {
    editor() {
      return this.$refs.myQuillEditor.quill;
    }
  },
  methods: {
    handleAvatarSuccess(res) {
      this.user_news.img = res.data.file.url;
    },
    beforeAvatarUpload(file) {
      var testmsg = file.name.substring(file.name.lastIndexOf(".") + 1);
      const extension = testmsg === "jpg";
      const extension2 = testmsg === "png";
      const isLt50KB = file.size / 1024 < 50;
      if (!extension && !extension2) {
        this.$message({
          message: "上传文件只能是 jpg、png格式!",
          type: "warning"
        });
        return false;
      }
      if (!isLt50KB) {
        this.$message({
          message: "上传文件大小不能超过 50KB!",
          type: "warning"
        });
        return false;
      }
      return (extension || extension2) && isLt50KB;
    },
    // onEditorReady(editor) { // 准备编辑器

    // },
    onEditorBlur() {}, // 失去焦点事件
    onEditorFocus() {}, // 获得焦点事件
    onEditorChange() {}, // 内容改变事件
    saveDrft() {
      this.$refs.newForm.validate(async valid => {
        if (valid) {
          let res;
          this.user_news.status = 1;
          if (this.isEdit) {
            res = await updateUserNews(this.user_news);
          } else {
            res = await createUserNews(this.user_news);
          }
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          } else {
            this.$message({ type: "error", message: "添加失败!" });
          }
        }
      });
    }, // 保存方法
    savePublish() {
      this.$refs.newForm.validate(async valid => {
        if (valid) {
          let res;
          this.user_news.status = 2;
          if (this.isEdit) {
            res = await updateUserNews(this.user_news);
          } else {
            res = await createUserNews(this.user_news);
          }
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          } else {
            this.$message({ type: "error", message: "添加失败!" });
          }
          this.$router.push({ name: "newlist" });
        }
      });
    }
  },
  created() {
    if (this.$route.query && this.$route.query.id) {
      const id = this.$route.query && this.$route.query.id;
      const row = {
        ID: Number(id)
      };
      findUserNews(row).then(res => {
        if (res.success) {
          this.isEdit = true;
          this.user_news = res.data.reun;
        } else {
          this.isEdit = false;
          this.user_news = {};
        }
      });
    }
  }
};
</script>
<style lang="scss">
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  border: 1px dashed #d9d9d9 !important;
  border-radius: 6px;
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
.ql-container {
  height: 500px;
}
.ql-editor strong {
  font-style: normal;
  font-weight: bold;
}
.ql-editor em {
  font-style: italic;
}
</style>
