<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="工作名称">
          <el-input placeholder="工作名称" v-model="searchInfo.p_name"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="工作名称" min-width="250" prop="job_info.p_name"></el-table-column>
      <el-table-column label="用户名" min-width="150" prop="user_info.userName"></el-table-column>
      <el-table-column label="简历状态" min-width="150">
        <template slot-scope="scope">
          <el-select
            @change="changeResumeStatus(scope.row)"
            placeholder="请选择"
            v-model="scope.row.resume_status"
          >
            <el-option
              :key="item.name"
              :label="item.name"
              :value="item.id"
              v-for="item in resume_status_infos"
            ></el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="viewResume(scope.row)" size="small" type="text">查看</el-button>
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

    <el-dialog :visible.sync="viewResumeInfoDialog" custom-class="user-dialog" title="查看简历">
      <div id="ShowBox">
        <div class="motto">座右铭：细心从每一个小细节开始。</div>
        <div id="show">
          <section>
            <h3>基本信息</h3>
            <ul>
              <li>姓名：{{user_base_info.user_name}}</li>
              <li>性别：{{user_base_info.gender}}</li>
              <li>最高学历：{{user_base_info.edulevel}}</li>
              <li>工作年限：{{user_base_info.worksYear}}</li>
              <li>所在城市：{{user_base_info.city}}</li>
              <li>联系电话：{{user_base_info.mobile}}</li>
              <li>联系邮箱：{{user_base_info.email}}</li>
            </ul>
          </section>
          <!-- <section>
                  <h3>工作经历</h3>
                  <ul>
                    <li v-for="userwork in user_works" :key="userwork">
                      <span>起止日期 {{userwork.join}}-{{userwork.leave}}</span>
                      <span>公司名称 {{userwork.company}}</span>
                      <span>工作地点 {{userwork.city}}</span>
                      <span>工作职位 {{userwork.department}}</span>
                      <span>工作描述 {{userwork.workContent}}</span>
                      <span>工作描述 {{userwork.workContent}}</span>
                    </li>
                  </ul>
          </section>-->
          <section v-if="user_works.length > 0 ">
            <h3>工作经历</h3>
            <ul v-for="userwork in user_works" :key="userwork">
              <li>起止日期：{{userwork.join}}-{{userwork.leave}}</li>
              <li>公司名称：{{userwork.companyname}}</li>
              <li>工作地点：{{userwork.city}}</li>
              <li>工作职位：{{userwork.department}}</li>
              <li>工作描述：{{userwork.workContent}}</li>
            </ul>
          </section>
          <section v-if="user_edus.length > 0 ">
            <h3>教育经历</h3>
            <ul v-for="useredu in user_edus" :key="useredu">
              <li>起止日期：{{useredu.joinTime}}-{{useredu.graduationTime}}</li>
              <li>学校名称：{{useredu.schoolname}}</li>
              <li>学历：{{useredu.edulevel}}</li>
              <li>专业：{{useredu.profession}}</li>
            </ul>
          </section>
          <section>
            <h3>期望职位</h3>
            <ul>
              <li>期望职位：{{user_dreams.dreamposi}}</li>
              <li>工作类型：{{user_dreams.workType}}</li>
              <li>期望薪资：{{user_dreams.salary}}</li>
              <li>期望城市：{{user_dreams.city}}</li>
              <li>到岗时间：{{user_dreams.dutyTime}}</li>
            </ul>
          </section>
        </div>
      </div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="download" type="primary">导 出</el-button>
        <el-button @click="viewResumeInfoDialog=false">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import { getResumeStatusList, updateResumeStatus } from "@/api/resumestatus";
import infoList from "@/components/mixins/infoList";
import html2canvas from "html2canvas";
import jspdf from "jspdf";
export default {
  name: "Api",
  mixins: [infoList],
  data() {
    return {
      listApi: getResumeStatusList,
      listKey: "resumeStatusList",
      path: path,
      viewResumeInfoDialog: false,
      name: "zhenghao",
      urlImg:
        "http://vinustseng.oss-cn-beijing.aliyuncs.com/1587350295head.jpg",
      user_base_info: {
        user_name: "zhenghao",
        gender: "男",
        edulevel: "本科",
        worksYear: "3年一下",
        city: "北京",
        mobile: "13521367746",
        email: "540214845"
      },
      jobtypes: [
        {
          id: 0,
          name: "全职"
        },
        {
          id: 1,
          name: "兼职"
        },
        {
          id: 2,
          name: "实习"
        },
        {
          id: 3,
          name: "志愿者"
        }
      ],
      jobedus: [
        {
          id: -1,
          name: "其他"
        },
        {
          id: 0,
          name: "不限"
        },
        {
          id: 1,
          name: "初中"
        },
        {
          id: 2,
          name: "高中"
        },
        {
          id: 3,
          name: "中专"
        },
        {
          id: 4,
          name: "大专"
        },
        {
          id: 5,
          name: "本科"
        },
        {
          id: 6,
          name: "硕士"
        },
        {
          id: 7,
          name: "博士"
        }
      ],
      user_works: [
        {
          join: "2017-09",
          leave: "2018-05",
          companyname: "爱奇艺",
          department: "测试",
          city: "北京",
          workContent: "吃饭睡觉打豆豆"
        },
        {
          join: "2018-05",
          leave: "2019-10",
          companyname: "爱奇艺",
          department: "测试",
          city: "北京",
          workContent: "吃饭睡觉打豆豆"
        },
        {
          join: "2019-10",
          leave: "至今",
          companyname: "爱奇艺",
          department: "测试",
          city: "北京",
          workContent: "吃饭睡觉打豆豆"
        }
      ],
      user_edus: [
        {
          joinTime: "2013-09",
          graduationTime: "2017-09",
          schoolname: "清华",
          edulevel: "本科",
          profession: "计算机"
        },
        {
          joinTime: "2010-09",
          graduationTime: "2013-09",
          schoolname: "人大附中",
          edulevel: "高中",
          profession: "数学"
        },
        {
          joinTime: "2007-09",
          graduationTime: "2010-09",
          schoolname: "北大附中",
          edulevel: "初中",
          profession: "美术"
        }
      ],
      user_dreams: {
        dreamposi: "科学家",
        workType: "全职",
        salary: "10k-15k",
        dutyTime: "一周到岗",
        city: "北京"
      },
      resume_status_infos: [
        {
          id: 0,
          name: "未读"
        },
        {
          id: 1,
          name: "已读"
        },
        {
          id: 2,
          name: "有意向"
        },
        {
          id: 3,
          name: "约面试"
        },
        {
          id: 4,
          name: "不合适"
        }
      ]
    };
  },
  methods: {
    viewResume(row) {
      this.viewResumeInfoDialog = true;
      this.base_info = row;
      this.fetchresume();
    },
    fetchresume() {
      //获取基本信息
      //const res = await getAuthorityList({ page: 1, pageSize: 999 });
      //this.authOptions = res.data.list;
      //获取工作经历
      //获取教育经历
      //获取期望职位
    },
    closeviewResumeDialog() {
      this.user_base_info = {};
      this.user_works = [];
      this.user_edus = [];
      this.user_dreams = {};
      this.viewResumeInfoDialog = false;
    },
    async changeResumeStatus(row) {
      const res = await updateResumeStatus(row);
      if (res.success) {
        this.$message({ type: "success", message: "状态设置成功" });
      }
    },
    filter(array) {
      //找出非空对象
      return array.filter(item => !this.isEmpty(item));
    },
    download() {
      let resume = document.getElementById("ShowBox");
      html2canvas(resume, { useCORS: true }).then(canvas => {
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
      // html2canvas(resume, {
      //   useCORS: true,
      //   onrendered: function(canvas) {
      //     //这是文件名
      //     var fileName =
      //       "${submitExperimentFile.studentSubmitter.majorName}${submitExperimentFile.studentSubmitter.className}-${submitExperimentFile.studentSubmitter.username}${submitExperimentFile.studentSubmitter.name}" +
      //       "《${submitExperimentFile.experiment.experimentName}》";
      //     var contentWidth = canvas.width;
      //     var contentHeight = canvas.height;
      //     //一页pdf显示html页面生成的canvas高度;
      //     var pageHeight = (contentWidth / 592.28) * 841.89;
      //     //未生成pdf的html页面高度
      //     var leftHeight = contentHeight;
      //     //页面偏移
      //     var position = 0;
      //     //a4纸的尺寸[595.28,841.89]，html页面生成的canvas在pdf中图片的宽高
      //     var imgWidth = 595.28;
      //     var imgHeight = (592.28 / contentWidth) * contentHeight;

      //     var pageData = canvas.toDataURL("image/jpeg", 1.0);

      //     var pdf = new jspdf("", "pt", "a4");

      //     //有两个高度需要区分，一个是html页面的实际高度，和生成pdf的页面高度(841.89)
      //     //当内容未超过pdf一页显示的范围，无需分页
      //     if (leftHeight < pageHeight) {
      //       pdf.addImage(pageData, "JPEG", 0, 0, imgWidth, imgHeight);
      //     } else {
      //       while (leftHeight > 0) {
      //         pdf.addImage(pageData, "JPEG", 0, position, imgWidth, imgHeight);
      //         leftHeight -= pageHeight;
      //         position -= 841.89;
      //         //避免添加空白页
      //         if (leftHeight > 0) {
      //           pdf.addPage();
      //         }
      //       }
      //     }
      //     pdf.save(fileName + ".pdf");
      //   }
      // });
    }
  }
};
</script>
<style scoped lang="scss">
#ShowBox {
  font-size: 16px;
  padding-top: 30px;
  color: rgba(0, 0, 0, 0.6);
  > .motto {
    background: #000;
    color: #fff;
    padding: 8px 32px;
  }
  > #show {
    padding: 0 32px 32px 32px;
    overflow: auto;
    > section {
      overflow: auto;
      > h3 {
        color: rgba(0, 0, 0, 0.7);
        text-align: left;
        background: rgba(0, 0, 0, 0.1);
        border-left: 3px solid #000;
        padding-left: 4px;
        margin: 20px 0 10px 0;
        padding: 5px;
      }
      > ul {
        padding: 15px;
        width: 100%;
        overflow: hidden;
        font-size: 16px;
        > li {
          padding: 10px;
          line-height: 10px;
        }
      }
    }
  }
}
</style>