<template>
  <div>
    <el-row>
      <el-col :xs="24" :sm="10" :md="10" :lg="10">
        <!-- 左边编辑器 -->
        <EditBox v-on:freshResume="loadResume" v-bind:resume="resume" />
      </el-col>
      <el-col :xs="24" :sm="14" :md="14" :lg="14">
        <!-- 右边预览 -->
        <ShowBox v-bind:resume="resume" class="showBox" />
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { findJobUserById, findUsersResumeInfo } from "@/api/jobuser";
import EditBox from "./components/EditBox";
import ShowBox from "./components/ShowBox";
import html2canvas from "html2canvas";
import jspdf from "jspdf";
export default {
  data() {
    return {
      resume: {
        openid: "",
        user_base_info: {
          openid: "",
          userName: "",
          genderindex: "",
          worksYearindex: "",
          edulevelindex: "",
          cityindex: "",
          contact: "",
          email: "",
          birthday: "",
          edu_level: {
            name: ""
          },
          job_work_expire: {
            name: ""
          },
          city_name: {
            name: ""
          }
        },
        //声明工作经历数据
        user_works: [
        ],
        // 声明教育信息
        user_edus: [
          {
            openid: "",
            schoolname: "",
            graduation: "",
            graduationTime: "",
            edulevelindex: "",
            profession: "",
            edu_level: {
              name: ""
            }
          }
        ],
        user_dreams: {
          openid: "",
          dreamposi: "",
          workTypeindex: "",
          cityindex: "",
          salaryindex: "",
          dutyTimeindex: "",
          job_duty_time: {
            name: ""
          },
          job_salary: {
            name: ""
          },
          job_work_type: {
            name: ""
          },
          city_name: {
            name: ""
          }
        }
      }
    };
  },
  methods: {
    download() {
      var opts = { useCORS: true };
      html2canvas(document.querySelector(".showBox"), opts).then(canvas => {
        let imgData = canvas.toDataURL("image/JPEG");
        let imgWidth = 210;
        let pageHeight = 295;
        let imgHeight = (canvas.height * imgWidth) / canvas.width;
        let heightLeft = imgHeight;
        let doc = new jspdf("p", "mm");
        let position = 0;
        doc.addImage(imgData, "JPEG", 0, position, imgWidth, imgHeight);
        heightLeft -= pageHeight;
        while (heightLeft > 0) {
          position = heightLeft - imgHeight;
          doc.addPage();
          doc.addImage(imgData, "JPEG", 0, position, imgWidth, imgHeight);
          heightLeft -= pageHeight;
        }
        doc.save("Resume.pdf");
      });
      // */
    },
    loadResume() {
      const row = {
        openid: this.resume.openid
      };
      findUsersResumeInfo(row).then(res => {
        if (res.success) {
          if (res.data.user != null) {
            this.resume.user = res.data.user;
          }
          if (res.data.base != null) {
            this.resume.user_base_info = res.data.base;
          }
          if (res.data.works != null) {
            this.resume.user_works = res.data.works;
          }
          if (res.data.edus != null) {
            this.resume.user_edus = res.data.edus;
          }
          if (res.data.dream != null) {
            this.resume.user_dreams = res.data.dream;
          }
        }
      });
    }
  },
  components: {
    EditBox,
    ShowBox
  },
  created() {
    if (this.$route.query && this.$route.query.id) {
      const id = this.$route.query && this.$route.query.id;
      const row = {
        ID: Number(id)
      };
      findJobUserById(row).then(res => {
        if (res.success) {
          this.resume.openid = res.data.reusers.openid;
          this.loadResume();
        }
      });
    }
  }
};
</script>
