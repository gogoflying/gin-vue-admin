<template>
  <!-- <el-dialog
    :title="base_info.name+'的简历'"
    :visible.sync="visible"
    size="tiny"
    :close-on-click-modal="false"
  >-->
  <div>
    <!-- <h2>基本信息</h2>
    <el-form :model="use_base_info" label-width="100px">
      <el-form-item label="姓名" prop="model_id">
        <el-input v-model="use_base_info.user_name"></el-input>
      </el-form-item>
      <el-form-item label="性别" prop="model_id">
        <el-input v-model="use_base_info.gender"></el-input>
      </el-form-item>
      <el-form-item label="最高学历" prop="model_id">
        <el-input v-model="use_base_info.edulevel"></el-input>
      </el-form-item>
      <el-form-item label="工作年限" prop="model_id">
        <el-input v-model="use_base_info.worksYear"></el-input>
      </el-form-item>
      <el-form-item label="所在城市" prop="model_id">
        <el-input v-model="use_base_info.city"></el-input>
      </el-form-item>
      <el-form-item label="联系电话" prop="client_version">
        <el-input v-model="use_base_info.mobile"></el-input>
      </el-form-item>
      <el-form-item label="联系邮箱" prop="client_version">
        <el-input v-model="use_base_info.email"></el-input>
      </el-form-item>
    </el-form>-->
    <section>
      <h2>基本信息</h2>
      <div class="message">
        <div class="content">
          <span>姓名：{{user_base_info.user_name}}</span>
          <span>性别：{{user_base_info.gender}}</span>
          <span>最高学历：{{user_base_info.edulevel}}</span>
          <span>工作年限：{{user_base_info.worksYear}}</span>
          <span>所在城市：{{user_base_info.city}}</span>
          <span>联系电话：{{user_base_info.mobile}}</span>
          <span>联系邮箱：{{user_base_info.email}}</span>
        </div>
      </div>
    </section>
    <section v-if="filter(user_works).length > 0 ">
      <h3>工作经历</h3>
      <ul v-for="(userwork,index) in filter(user_works)">
        <li>起止日期 {{userwork.join}}-{{userwork.leave}}</li>
        <li>公司名称 {{userwork.company}}</li>
        <li>工作地点 {{userwork.city}}</li>
        <li>工作职位 {{userwork.department}}</li>
        <li>工作描述 {{userwork.workContent}}</li>
      </ul>
    </section>
    <section v-if="filter(user_edus).length > 0 ">
      <h3>教育经历</h3>
      <ul v-for="useredu in filter(user_edus)">
        <li>起止日期 {{useredu.joinTime}}-{{useredu.graduationTime}}</li>
        <li>学校名称 {{useredu.schoolname}}</li>
        <li>学历 {{useredu.edulevel}}</li>
        <li>专业 {{useredu.profession}}</li>
      </ul>
    </section>
    <section v-if="filter(user_edus).length > 0 ">
      <h3>期望职位</h3>
      <ul v-for="useredu in filter(user_edus)">
        <li>起止日期 {{useredu.joinTime}}-{{useredu.graduationTime}}</li>
        <li>学校名称 {{useredu.schoolname}}</li>
        <li>学历 {{useredu.edulevel}}</li>
        <li>专业 {{useredu.profession}}</li>
      </ul>
    </section>

    <section>
      <h3>期望职位</h3>
      <div class="message">
        <div class="content">
          <span>期望职位：{{user_dreams.dreamposi}}</span>
          <span>工作类型：{{user_dreams.workType}}</span>
          <span>期望薪资：{{user_dreams.salary}}</span>
          <span>期望城市：{{user_dreams.city}}</span>
          <span>到岗时间：{{user_dreams.dutyTime}}</span>
        </div>
        <div class="print">
          <Imges />
        </div>
      </div>
    </section>
  </div>
  <!-- </el-dialog> -->
</template>

<script>
export default {
  name: "viewresume",
  data() {
    return {
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
  methods: {
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
    }
  }
};
</script>
