<template>
  <el-container class="layout-cont">
    <el-container :class="[isSider?'openside':'hideside',isMobile ? 'mobile': '']">
      <el-row :class="[isShadowBg?'shadowBg':'']" @click.native="changeShadow()"></el-row>
      <el-aside class="main-cont main-left">
        <Aside class="aside" />
      </el-aside>
      <!-- 分块滑动功能 -->
      <el-main class="main-cont main-right">
        <el-header class="header-cont">
          <div @click="totalCollapse" class="menu-total">
            <i class="el-icon-s-unfold" v-if="isCollapse"></i>
            <i class="el-icon-s-fold" v-else></i>
          </div>
          <el-breadcrumb class="breadcrumb" separator-class="el-icon-arrow-right">
            <el-breadcrumb-item
              :key="item.path"
              v-for="item in matched.slice(1,matched.length)"
            >{{item.meta.title}}</el-breadcrumb-item>
          </el-breadcrumb>
          <div class="fl-right right-box">
            <el-dropdown>
              <span class="el-dropdown-link">
                <img :src="userInfo.headerImg" height="30" width="30" />
                {{userInfo.nickName}}
                <i class="el-icon-arrow-down"></i>
              </span>
              <el-dropdown-menu class="dropdown-group" slot="dropdown">
                <el-dropdown-item>
                  <span>
                    更多信息
                    <el-badge is-dot />
                  </span>
                </el-dropdown-item>
                <el-dropdown-item @click.native="showPassword=true" icon="el-icon-s-custom">修改密码</el-dropdown-item>
                <el-dropdown-item @click.native="toPerson" icon="el-icon-s-custom">个人信息</el-dropdown-item>
                <el-dropdown-item @click.native="toCompanyAuth" icon="el-icon-s-check">企业认证</el-dropdown-item>
                <el-dropdown-item @click.native="LoginOut" icon="el-icon-table-lamp">登 出</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
            <el-dialog title="修改密码" :visible.sync="showPassword" @close="clearPassword">
              <el-form ref="modifyPwdForm" :model="pwdModify" :rules="rules" label-width="80px">
                <el-form-item prop="password" :minlength="6" label="原密码">
                  <el-input v-model="pwdModify.password" show-password></el-input>
                </el-form-item>
                <el-form-item prop="newPassword" :minlength="6" label="新密码">
                  <el-input v-model="pwdModify.newPassword" show-password></el-input>
                </el-form-item>
                <el-form-item prop="confirmPassword" :minlength="6" label="确认密码">
                  <el-input v-model="pwdModify.confirmPassword" show-password></el-input>
                </el-form-item>
              </el-form>
              <div slot="footer" class="dialog-footer">
                <el-button @click="showPassword=false">取 消</el-button>
                <el-button type="primary" @click="savePassword">确 定</el-button>
                <el-button type="primary" @click="download">导 出</el-button>
              </div>
              <div id="ShowBox">
                <div class="motto">座右铭：细心从每一个小细节开始。</div>
                <div id="show">
                  <section>
                    <h3>基本信息</h3>
                    <ul>
                      <li>姓名：{{user_base_info.user_name}}</li>
                      <br />
                      <li>性别：{{user_base_info.gender}}</li>
                      <br />
                      <li>最高学历：{{user_base_info.edulevel}}</li>
                      <br />
                      <li>工作年限：{{user_base_info.worksYear}}</li>
                      <br />
                      <li>所在城市：{{user_base_info.city}}</li>
                      <br />
                      <li>联系电话：{{user_base_info.mobile}}</li>
                      <br />
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
                      <br />
                      <li>公司名称：{{userwork.companyname}}</li>
                      <br />
                      <li>工作地点：{{userwork.city}}</li>
                      <br />
                      <li>工作职位：{{userwork.department}}</li>
                      <br />
                      <li>工作描述：{{userwork.workContent}}</li>
                    </ul>
                  </section>
                  <section v-if="user_edus.length > 0 ">
                    <h3>教育经历</h3>
                    <ul v-for="useredu in user_edus" :key="useredu">
                      <li>起止日期：{{useredu.joinTime}}-{{useredu.graduationTime}}</li>
                      <br />
                      <li>学校名称：{{useredu.schoolname}}</li>
                      <br />
                      <li>学历：{{useredu.edulevel}}</li>
                      <br />
                      <li>专业：{{useredu.profession}}</li>
                    </ul>
                  </section>
                  <section>
                    <h3>期望职位</h3>
                    <ul>
                      <li>期望职位：{{user_dreams.dreamposi}}</li>
                      <br />
                      <li>工作类型：{{user_dreams.workType}}</li>
                      <br />
                      <li>期望薪资：{{user_dreams.salary}}</li>
                      <br />
                      <li>期望城市：{{user_dreams.city}}</li>
                      <br />
                      <li>到岗时间：{{user_dreams.dutyTime}}</li>
                    </ul>
                  </section>
                </div>
              </div>
            </el-dialog>
          </div>
        </el-header>
        <!-- 当前面包屑用路由自动生成可根据需求修改 -->
        <!-- 
        :to="{ path: item.path }" 暂时注释不用-->
        <HistoryComponent />
        <transition mode="out-in" name="el-fade-in-linear">
          <router-view class="admin-box"></router-view>
        </transition>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import Aside from "@/view/layout/aside";
import HistoryComponent from "@/view/layout/aside/historyComponent/history";

import { mapGetters, mapActions } from "vuex";
import { changePassword } from "@/api/user";
import html2canvas from "html2canvas";
import jspdf from "jspdf";
export default {
  name: "Layout",
  data() {
    return {
      isCollapse: false,
      isSider: true,
      isMobile: false,
      isShadowBg: false,
      showPassword: false,
      pwdModify: {},
      rules: {
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          { min: 6, message: "最少6个字符", trigger: "blur" }
        ],
        newPassword: [
          { required: true, message: "请输入新密码", trigger: "blur" },
          { min: 6, message: "最少6个字符", trigger: "blur" }
        ],
        confirmPassword: [
          { required: true, message: "请输入确认密码", trigger: "blur" },
          { min: 6, message: "最少6个字符", trigger: "blur" },
          {
            validator: (rule, value, callback) => {
              if (value !== this.pwdModify.newPassword) {
                callback(new Error("两次密码不一致"));
              } else {
                callback();
              }
            },
            trigger: "blur"
          }
        ]
      },
      visible: false,
      base_info: {},
      name: "zhenghao",
      user_base_info: {
        user_name: "zhenghao",
        gender: "男",
        edulevel: "本科",
        worksYear: "3年一下",
        city: "北京",
        mobile: "13521367746",
        email: "540214845"
      },
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
      }
    };
  },
  components: {
    Aside,
    HistoryComponent
  },
  created() {
    let screenWidth = document.body.clientWidth;
    if (screenWidth < 1000) {
      this.isMobile = true;
      this.isSider = false;
      this.isCollapse = !this.isCollapse;
    } else {
      this.isMobile = false;
    }
  },
  methods: {
    ...mapActions("user", ["LoginOut"]),
    totalCollapse() {
      this.isCollapse = !this.isCollapse;
      this.isSider = !this.isCollapse;
      this.isShadowBg = !this.isCollapse;
      this.$bus.emit("totalCollapse");
    },
    toPerson() {
      this.$router.push({ name: "person" });
    },
    toCompanyAuth() {
      this.$router.push({ name: "companyauth" });
    },
    changeShadow() {
      this.isShadowBg = !this.isShadowBg;
      this.isSider = !!this.isCollapse;
      this.totalCollapse();
    },
    savePassword() {
      this.$refs.modifyPwdForm.validate(valid => {
        if (valid) {
          changePassword({
            username: this.userInfo.userName,
            password: this.pwdModify.password,
            newPassword: this.pwdModify.newPassword
          }).then(() => {
            this.$message.success("修改密码成功！");
            this.showPassword = false;
          });
        } else {
          return false;
        }
      });
    },
    clearPassword() {
      this.pwdModify = {
        password: "",
        newPassword: "",
        confirmPassword: ""
      };
      this.$refs.modifyPwdForm.clearValidate();
    },
    open(base_info) {
      this.visible = true;
      this.base_info = base_info;

      this.fetchresume();
    },
    fetchresume() {
      //获取基本信息
      //获取工作经历
      //获取教育经历
      //获取期望职位
    },
    filter(array) {
      //找出非空对象
      return array.filter(item => !this.isEmpty(item));
    },
    download() {
      let resume = document.getElementById("ShowBox");
      html2canvas(resume).then(canvas => {
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
    }
  },
  computed: {
    ...mapGetters("user", ["userInfo"]),
    ...mapGetters("history", ["historys", "activeValue"]),
    title() {
      return this.$route.meta.title || "当前页面";
    },
    matched() {
      return this.$route.matched;
    }
  },
  mounted() {
    window.onresize = () => {
      return (() => {
        let screenWidth = document.body.clientWidth;
        if (!this.screenWidth && this.isSider) {
          if (screenWidth < 1000) {
            this.isMobile = true;
            this.isSider = false;
            this.isCollapse = true;
            this.$bus.emit("collapse", this.isCollapse);
          }
        } else {
          if (screenWidth < 1000) {
            this.isMobile = true;
            this.isSider = false;
            this.isCollapse = true;
          } else {
            this.isMobile = false;
          }
        }
      })();
    };
  }
};
</script>

<style lang="scss">
$headerHigh: 52px;
$mainHight: 100vh;
.el-dropdown-link {
  cursor: pointer;
}
.dropdown-group {
  min-width: 100px;
}
.el-scrollbar__wrap {
  padding-bottom: 17px;
}
.layout-cont {
  .right-box {
    text-align: center;
    vertical-align: middle;
    img {
      vertical-align: middle;
      border: 1px solid #ccc;
      border-radius: 6px;
    }
  }
  .menu-contorl {
    line-height: 52px;
    font-size: 20px;
    color: #eee;
    display: table-cell;
    vertical-align: middle;
  }
  .header-cont {
    height: $headerHigh !important;
    background: #fff;
    box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
    line-height: $headerHigh;
  }
  .main-cont {
    .breadcrumb {
      line-height: 48px;
      display: inline-block;
      padding: 0 24px;
      // padding: 6px;
      // border-bottom: 1px solid #eee;
    }
    .router-history {
      background: #fff;
      margin-top: 1px;
      padding: 0 6px;
    }
    &.el-main {
      overflow: auto;
      background: #fff;
      // padding: 0px 10px;
      // background: #fff;
    }
    height: $mainHight !important;
    overflow: visible;
    position: relative;
    .menu-total {
      // z-index: 5;
      // position: absolute;
      // top: 10px;
      // right: -35px;
      margin-left: -10px;
      float: left;
      margin-top: 10px;
      width: 30px;
      height: 30px;
      line-height: 30px;
      font-size: 30px;
      // border: 0 solid #ffffff;
      // border-radius: 50%;
      // background: #fff;
    }
    .aside {
      overflow: auto;
      // background: #fff;
      &::-webkit-scrollbar {
        display: none;
      }
    }

    .el-menu-vertical {
      height: 100vh !important;
      visibility: auto;
      &:not(.el-menu--collapse) {
        width: 220px;
      }
    }

    &::-webkit-scrollbar {
      display: none;
    }
    &.main-left {
      width: auto !important;
    }
    &.main-right {
      .admin-title {
        float: left;
        font-size: 16px;
        vertical-align: middle;
        margin-left: 20px;
        img {
          vertical-align: middle;
        }
        &.collapse {
          width: 53px;
        }
      }
    }
  }
}

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
        padding: 8px;
        width: 100%;
        overflow: hidden;
        font-size: 16px;
        > li {
          padding: 10px;
          line-height: 30px;
          float: left;
        }
      }
    }
  }
}
</style>