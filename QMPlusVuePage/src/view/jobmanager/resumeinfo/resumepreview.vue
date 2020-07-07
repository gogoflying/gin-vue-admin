<template>
  <div>
    <div>
      <el-row type="flex" justify="center">
        <el-col :span="22">
          <!-- <el-button @click="back" type="primary" round>返 回</el-button> -->
          <el-button @click="download" type="success" plain round>导 出</el-button>
        </el-col>
      </el-row>
    </div>
    <ShowBox v-bind:resume="resume" class="showBox" />
  </div>
</template>

<script>
import { findResumeStatus, findUsersByOpenid } from "@/api/resumestatus";
import ShowBox from "./components/ShowBox";
import html2canvas from "html2canvas";
import jspdf from "jspdf";
export default {
  data() {
    return {
      resume: {
        openid: "",
        user_base_info: {
          userName: "",
          genderindex: "",
          edulevel: "",
          worksYear: "",
          city: "",
          mobile: "",
          email: "",
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
          {
            companyname: "",
            join: "",
            leave: "",
            department: "",
            workContent: ""
          }
        ],
        // 声明教育信息
        user_edus: [
          {
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
    back() {
      this.$router.push({ name: "resumeinfo" });
    },
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
      findUsersByOpenid(row).then(res => {
        if (res.success) {
          this.resume.user = res.data.user;
          this.resume.user_base_info = res.data.base;
          this.resume.user_works = res.data.works;
          this.resume.user_edus = res.data.edus;
          this.resume.user_dreams = res.data.dream;
        }
      });
    }
  },
  components: {
    ShowBox
  },
  created() {
    if (this.$route.query && this.$route.query.id) {
      const id = this.$route.query && this.$route.query.id;
      const row = {
        ID: Number(id)
      };
      findResumeStatus(row).then(res => {
        if (res.success) {
          this.resume.openid = res.data.rers.openid;
          this.loadResume();
        }
      });
    }
  }
};
</script>
