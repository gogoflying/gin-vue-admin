<template>
  <div id="ShowBox">
    <!-- <div class="motto">座右铭：你与名企之间的距离就差一份简历。</div> -->
    <div id="show">
      <section>
        <h3>基本信息</h3>
        <div class="message">
          <div class="content">
            <ul>
              <li>姓名：{{resume.user_base_info.userName || '你的名字'}}</li>
              <li>性别：{{resume.user_base_info.genderindex == 0 ? '男':'女'}}</li>
              <li
                v-if="resume.user_base_info.edu_level"
              >最高学历：{{resume.user_base_info.edu_level.name || '最高学历'}}</li>
              <li
                v-if="resume.user_base_info.job_work_expire"
              >工作年限：{{resume.user_base_info.job_work_expire.name || '工作年限'}}</li>
              <li
                v-if="resume.user_base_info.city_name"
              >所在城市：{{resume.user_base_info.city_name.name || '所在城市'}}</li>
              <li>联系电话：{{resume.user_base_info.contact || '联系电话'}}</li>
              <li>联系邮箱：{{resume.user_base_info.email || '联系邮箱'}}</li>
              <li>生日：{{resume.user_base_info.birthday || '生日'}}</li>
            </ul>
          </div>
          <div class="print" v-if="resume.user_base_info.avatarUrl != '' && resume.user_base_info.avatarUrl != null">
            <img
              :src="resume.user_base_info.avatarUrl + '?' + new Date().getTime()"
              crossorigin="anonymous"
              height="100"
              width="100"
            />
          </div>
        </div>
      </section>
      <section v-if="filter(resume.user_works).length > 0 ">
        <h3>工作经历</h3>
        <ul v-for="(work,index) in filter(resume.user_works)" :key="index">
          <li>起止日期：{{work.join || '开始日期'}}至{{work.leave || '结束日期'}}</li>
          <li>公司名称：{{work.companyname || '公司名称'}}</li>
          <li>工作职位：{{work.department || '职位'}}</li>
          <li>
            <div style="display:inline">
            <div style="float:left;width:15%">工作描述：</div>
            <div
              style="float:left;width:85%"
              v-html="work.workContent ? work.workContent.replace(/\n/g, '<br/>') : '工作描述'"
            ></div>
            </div>
          </li>
        </ul>
      </section>
      <section v-if="filter(resume.user_edus).length > 0 ">
        <h3>教育信息</h3>
        <ul v-for="(edu,index) in filter(resume.user_edus)" :key="index">
          <li>起止日期：{{edu.graduation || '开始日期'}}至{{edu.graduationTime || '结束日期'}}</li>
          <li>学校名称：{{edu.schoolname || '学校名称'}}</li>
          <li v-if="edu.edu_level">学历：{{edu.edu_level.name || '学历'}}</li>
          <li>专业：{{edu.profession || '专业'}}</li>
        </ul>
      </section>
      <section>
        <h3>期望职位</h3>
        <ul>
          <li>期望职位：{{resume.user_dreams.dreamposi || '期望职位'}}</li>
          <li
            v-if="resume.user_dreams.job_work_type"
          >工作类型：{{resume.user_dreams.job_work_type.name || '工作类型'}}</li>
          <li
            v-if="resume.user_dreams.job_salary"
          >期望薪资：{{resume.user_dreams.job_salary.name || '期望薪资'}}</li>
          <li
            v-if="resume.user_dreams.city_name"
          >期望城市：{{resume.user_dreams.city_name.name || '期望城市'}}</li>
          <li
            v-if="resume.user_dreams.job_duty_time"
          >到岗时间：{{resume.user_dreams.job_duty_time.name || '到岗时间'}}</li>
        </ul>
      </section>
    </div>
  </div>
</template>

<style lang="scss">
#ShowBox {
  font-size: 16px;
  padding-top: 1px;
  color: rgba(0, 0, 0, 0.6);
  > .motto {
    background: rgb(199, 192, 236);
    color: #fff;
    padding: 8px 32px;
  }
  > #show {
    padding: 0 32px 32px 32px;
    // height: 450px;//与下面一起设置滚动条
    overflow: auto;
    > section {
      overflow: auto;
      > .message {
        display: flex;
        padding: 8px;
        > .content {
          flex-grow: 2;
          > ul {
            padding: 8px;
            width: 95%;
            overflow: hidden;
            font-size: 16px;
            > li {
              padding: 5px;
              line-height: 20px;
            }
          }
        }
        > .print {
          flex-grow: 1;
          display: flex;
          justify-content: center;
          align-items: center;
        }
      }
      > h3 {
        color: rgba(0, 0, 0, 0.7);
        text-align: left;
        background: rgba(0, 0, 0, 0.1);
        // border-left: 3px solid #000;
        padding-left: 4px;
        margin: 20px 0 10px 0;
        padding: 5px;
      }
      > ul {
        padding: 8px;
        width: 95%;
        overflow: hidden;
        font-size: 16px;
        > li {
          padding: 5px;
          line-height: 20px;
        }
      }
    }
  }
}
</style>

<script>
export default {
  props: ["resume"],
  methods: {
    filter(array) {
      //找出非空对象
      return array.filter(item => !this.isEmpty(item));
    },
    isEmpty(object) {
      //只要有一个value不是falsy,就返回false
      let empty = true;
      for (let key in object) {
        if (object[key]) {
          empty = false;
          break;
        }
      }
      return empty;
    }
  }
};
</script>
