<template>
  <div>
    <div>
      <el-form :rules="rules" ref="jobForm" :model="jobmanagerinfo" label-width="80px">
        <el-form-item label="工作名称" prop="p_name">
          <el-input v-model="jobmanagerinfo.p_name" placeholder="请输入工作名称" style="width:50%;"></el-input>
        </el-form-item>
        <el-form-item v-show="enPriseId == 0" label="所属公司" prop="enterprise_id">
          <!-- <el-select
            @change="selectEp"
            placeholder="请选择所属公司"
            v-model="jobmanagerinfo.enterprise_id"
          >
            <el-option
              :key="ep.enterprise_name"
              :label="ep.enterprise_name"
              :value="ep.ID"
              v-for="ep in enterpriseInfo"
            ></el-option>
          </el-select>-->
          <el-select
            @change="selectEp"
            v-model="jobmanagerinfo.enterprise_id"
            filterable
            remote
            reserve-keyword
            placeholder="请选择所属公司"
            :remote-method="remoteMethod"
            :loading="loading"
          >
            <el-option
              v-for="ep in enterpriseInfo"
              :key="ep.enterprise_name"
              :value="ep.ID"
              :label="ep.enterprise_name"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="薪资范围" prop="p_salary_id">
          <el-select
            @change="selectJobsalary"
            placeholder="请选择薪资范围"
            v-model.number="jobmanagerinfo.p_salary_id"
          >
            <el-option :key="js.name" :label="js.name" :value="js.ID" v-for="js in jobsalary"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工作地点" prop="p_address">
          <el-input v-model="jobmanagerinfo.p_address" placeholder="请输入工作地点" style="width:45%;"></el-input>
          <el-button size="mini" icon="el-icon-map-location" @click="openMap"></el-button>
        </el-form-item>
        <el-form-item label="工作地点纬度" prop="p_latitude" hidden>
          <el-input v-model="jobmanagerinfo.p_latitude" placeholder="请输入工作地点纬度" style="width:50%;"></el-input>
        </el-form-item>
        <el-form-item label="工作地点经度" prop="p_longitude" hidden>
          <el-input v-model="jobmanagerinfo.p_longitude" placeholder="请输入工作地点经度" style="width:50%;"></el-input>
        </el-form-item>
        <el-form-item label="工作城市" prop="p_city_id">
          <el-select
            @change="selectCityName"
            placeholder="请选择工作城市"
            v-model.number="jobmanagerinfo.p_city_id"
          >
            <el-option
              :key="city.name"
              :label="city.name"
              :value="city.ID"
              v-for="city in cityinfo"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工作年限" prop="p_edujy_id">
          <el-select
            @change="selectJobyear"
            placeholder="请选择工作年限"
            v-model="jobmanagerinfo.p_edujy_id"
          >
            <el-option
              :key="jobyear.name"
              :label="jobyear.name"
              :value="jobyear.ID"
              v-for="jobyear in jobyears"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="最低学历" prop="p_education_id">
          <el-select
            @change="selectJobedu"
            placeholder="请选择最低学历"
            v-model="jobmanagerinfo.p_education_id"
          >
            <el-option
              :key="jobedu.name"
              :label="jobedu.name"
              :value="jobedu.ID"
              v-for="jobedu in jobedus"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工作类型" prop="p_type_id">
          <el-select
            @change="selectJobtype"
            placeholder="请选择工作类型"
            v-model.number="jobmanagerinfo.p_type_id"
          >
            <el-option
              :key="jobtype.name"
              :label="jobtype.name"
              :value="jobtype.ID"
              v-for="jobtype in jobtypes"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="p_status">
          <el-radio-group v-model="jobmanagerinfo.p_status">
            <el-radio :label="0">普通</el-radio>
            <el-radio :label="1">急招</el-radio>
            <el-radio :label="2">热门</el-radio>
            <el-radio :label="3">下线</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="招聘人数" prop="p_count">
          <el-input
            type="p_count"
            v-model.number="jobmanagerinfo.p_count"
            placeholder="请输入招聘人数"
            style="width:50%;"
          ></el-input>
        </el-form-item>
        <el-form-item label="联系人" prop="p_contact">
          <el-input v-model="jobmanagerinfo.p_contact" placeholder="请输入联系人" style="width:50%;"></el-input>
        </el-form-item>
        <el-form-item label="面试地点" prop="p_interview">
          <el-input v-model="jobmanagerinfo.p_interview" placeholder="请输入面试地点" style="width:50%;"></el-input>
        </el-form-item>
        <el-form-item label="失效日期" prop="p_outdate">
          <el-date-picker
            placeholder="选择结束时间"
            type="date"
            :editable="false"
            format="yyyy-MM-dd"
            value-format="timestamp"
            style="width:50%;"
            v-model="jobmanagerinfo.p_outdate"
          ></el-date-picker>
        </el-form-item>
        <el-form-item label="工作描述" prop="p_desc">
          <quill-editor
            class="ql-editor"
            style="width:80%;"
            :options="editorOption"
            @blur="onEditorBlur($event)"
            @change="onEditorChange($event)"
            @focus="onEditorFocus($event)"
            ref="myQuillEditor"
            v-model="jobmanagerinfo.p_desc"
          ></quill-editor>
        </el-form-item>
        <el-form-item label-width="100px">
          <el-row justify="center">
            <el-col>
              <el-button @click="closeAddjobDialog" type="primary" round>返 回</el-button>
              <el-button @click="enterAddjobDialog" type="primary" round>确 定</el-button>
            </el-col>
          </el-row>
        </el-form-item>
      </el-form>
    </div>
    <el-dialog :visible.sync="addMapDialog" custom-class="user-dialog">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-input placeholder="请输入位置" v-model="searchaddress"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSearch" type="primary">查询</el-button>
        </el-form-item>
      </el-form>
      <div id="container" style="height:400px;"></div>
    </el-dialog>
  </div>
</template>

<script>
const path = process.env.VUE_APP_BASE_API;
import {
  createJoblist,
  updateJoblist,
  findJoblist,
  getjoblistOptions
} from "@/api/jobmanagerinfo";
import { getEnterpriseListBySearch } from "@/api/enterpriseinfo";
import { mapGetters } from "vuex";
export default {
  name: "NewJobInfo",
  data() {
    return {
      path: path,
      addMapDialog: false,
      isEdit: false,
      editorOption: {},
      searchaddress: "",
      enterpriseInfo: [],
      loading: false,
      jobmanagerinfo: {
        p_name: "",
        p_salary_id: "",
        p_salary: "",
        p_latitude: "39.916527",
        p_longitude: "116.397128",
        p_address: "",
        p_city: "",
        p_city_id: "",
        p_edujy: "",
        p_edujy_id: "",
        p_education: "",
        p_education_id: "",
        p_type: "",
        p_type_id: "",
        p_desc: "",
        p_contact: "",
        p_interview: "",
        p_status: 0,
        p_count: "",
        p_outdate: "",
        enterprise_id: "",
        enterprise_name: "",
        enterprise_logo: ""
      },
      cityinfo: [
        // {
        //   id: 1,
        //   name: "北京"
        // },
        // {
        //   id: 2,
        //   name: "上海"
        // },
        // {
        //   id: 3,
        //   name: "广州"
        // }
      ],
      jobsalary: [],
      jobyears: [
        // {
        //   id: 0,
        //   name: "1年"
        // },
        // {
        //   id: 1,
        //   name: "3年"
        // },
        // {
        //   id: 2,
        //   name: "5年"
        // }
      ],
      jobtypes: [
        // {
        //   id: 0,
        //   name: "全职"
        // },
        // {
        //   id: 1,
        //   name: "兼职"
        // },
        // {
        //   id: 2,
        //   name: "实习"
        // },
        // {
        //   id: 3,
        //   name: "志愿者"
        // }
      ],
      jobedus: [
        // {
        //   id: -1,
        //   name: "其他"
        // },
        // {
        //   id: 0,
        //   name: "不限"
        // },
        // {
        //   id: 1,
        //   name: "初中"
        // },
        // {
        //   id: 2,
        //   name: "高中"
        // },
        // {
        //   id: 3,
        //   name: "中专"
        // },
        // {
        //   id: 4,
        //   name: "大专"
        // },
        // {
        //   id: 5,
        //   name: "本科"
        // },
        // {
        //   id: 6,
        //   name: "硕士"
        // },
        // {
        //   id: 7,
        //   name: "博士"
        // }
      ],
      rules: {
        p_name: [
          {
            required: true,
            message: "请输入工作名称",
            trigger: "blur"
          }
        ],
        p_salary_id: [
          { required: true, message: "请输入薪资范围", trigger: "blur" }
        ],
        p_latitude: [
          { required: true, message: "请输入工作地点纬度", trigger: "blur" },
          {
            pattern: /^\d+(\.\d+)?$/,
            message: "请输入浮点数",
            trigger: "blur"
          }
        ],
        p_longitude: [
          { required: true, message: "请输入工作地点经度", trigger: "blur" },
          {
            pattern: /^\d+(\.\d+)?$/,
            message: "请输入浮点数",
            trigger: "blur"
          }
        ],
        p_address: [
          { required: true, message: "请输入工作地点", trigger: "blur" }
        ],
        p_city_id: [
          { required: true, message: "请输入工作城市", trigger: "blur" }
        ],
        p_edujy_id: [
          { required: true, message: "请输入工作年限", trigger: "blur" }
        ],
        p_education_id: [
          { required: true, message: "请输入最低学历", trigger: "blur" }
        ],
        p_type_id: [
          { required: true, message: "请输入工作类型", trigger: "blur" }
        ],
        p_count: [
          { required: true, message: "请输入招聘人数", trigger: "blur" },
          { type: "number", message: "人数必须为数字", trigger: "blur" },
          {
            type: "number",
            min: 0,
            message: "人数必须为大于等于零",
            trigger: "blur"
          }
        ],
        p_contact: [
          { required: true, message: "请输入联系人", trigger: "blur" }
        ],
        p_interview: [
          { required: true, message: "请输入面试地点", trigger: "blur" }
        ],
        p_desc: [{ required: true, message: "请输入工作描述", trigger: "blur" }]
      },
      geocoder: null,
      map: "",
      mapLister: "",
      marker: "",
      markerIcon: null
    };
  },
  computed: {
    ...mapGetters("user", ["enPriseId", "token"]),
    editor() {
      return this.$refs.myQuillEditor.quill;
    }
  },
  methods: {
    remoteMethod(query) {
      if (query !== "") {
        this.loading = true;
        setTimeout(() => {
          this.loading = false;
          const params = {
            query: query
          };
          getEnterpriseListBySearch(params).then(res => {
            if (res.success) {
              this.enterpriseInfo = res.data.result;
            } else {
              this.enterpriseInfo = [];
            }
          });
        }, 200);
      } else {
        this.enterpriseInfo = [];
      }
    },
    selectEp(val) {
      var selectedItem = {};
      selectedItem = this.enterpriseInfo.find(item => {
        return item.ID === val;
      });
      this.jobmanagerinfo.enterprise_name = selectedItem.enterprise_name;
      this.jobmanagerinfo.enterprise_logo = selectedItem.enterprise_logo;
    },
    selectCityName(val) {
      var selectedItem = {};
      selectedItem = this.cityinfo.find(item => {
        return item.ID === val;
      });
      this.jobmanagerinfo.p_city = selectedItem.name;
    },
    selectJobyear(val) {
      var selectedItem = {};
      selectedItem = this.jobyears.find(item => {
        return item.ID === val;
      });
      this.jobmanagerinfo.p_edujy = selectedItem.name;
    },
    selectJobtype(val) {
      var selectedItem = {};
      selectedItem = this.jobtypes.find(item => {
        return item.ID === val;
      });
      this.jobmanagerinfo.p_type = selectedItem.name;
    },
    selectJobedu(val) {
      var selectedItem = {};
      selectedItem = this.jobedus.find(item => {
        return item.ID === val;
      });
      this.jobmanagerinfo.p_education = selectedItem.name;
    },
    selectJobsalary(val) {
      var selectedItem = {};
      selectedItem = this.jobsalary.find(item => {
        return item.ID === val;
      });
      this.jobmanagerinfo.p_salary = selectedItem.name;
    },
    // onEditorReady(editor) { // 准备编辑器

    // },
    onEditorBlur() {}, // 失去焦点事件
    onEditorFocus() {}, // 获得焦点事件
    onEditorChange() {}, // 内容改变事件
    async enterAddjobDialog() {
      this.$refs.jobForm.validate(async valid => {
        if (valid) {
          let res;
          this.jobmanagerinfo.p_latitude = parseFloat(
            this.jobmanagerinfo.p_latitude
          );
          this.jobmanagerinfo.p_longitude = parseFloat(
            this.jobmanagerinfo.p_longitude
          );
          if (this.isEdit) {
            res = await updateJoblist(this.jobmanagerinfo);
          } else {
            res = await createJoblist(this.jobmanagerinfo);
          }
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          } else {
            this.$message({ type: "error", message: "添加失败!" });
          }
          this.$router.push({ name: "jobmanagerinfo" });
        }
      });
    },
    closeAddjobDialog() {
      this.$router.push({ name: "jobmanagerinfo" });
    },
    onSearch() {
      this.geocoder.getLocation(this.searchaddress);
    },
    init() {
      //步骤：定义map变量 调用 qq.maps.Map() 构造函数   获取地图显示容器
      //设置地图中心点
      var self = this;
      var myLatlng = new window.qq.maps.LatLng(
        self.jobmanagerinfo.p_latitude,
        self.jobmanagerinfo.p_longitude
      );
      var myOptions = {
        zoom: 15, //设置地图缩放级别
        center: myLatlng //设置中心点样式
        // mapTypeId: qq.maps.MapTypeId.ROADMAP //设置地图样式详情参见MapType
      };
      //获取dom元素添加地图信息
      this.map = new window.qq.maps.Map(
        document.getElementById("container"),
        myOptions
      );
      //创建marker
      if (self.marker == null || self.marker == "") {
        self.marker = new window.qq.maps.Marker({
          position: myLatlng,
          map: self.map
        });
      } else {
        self.marker = new window.qq.maps.Marker({
          position: myLatlng,
          map: self.map
        });
      }
      //给地图添加点击事件
      //并获取鼠标点击的经纬度
      this.mapLister = window.qq.maps.event.addListener(
        self.map,
        "click",
        function(event) {
          self.jobmanagerinfo.p_latitude = event.latLng.getLat();
          self.jobmanagerinfo.p_longitude = event.latLng.getLng();
          self.marker.setPosition(event.latLng);
          self.map.setCenter(event.latLng);
        }
      );
      //调用地址解析类
      this.geocoder = new window.qq.maps.Geocoder({
        complete: function(result) {
          self.jobmanagerinfo.p_latitude = result.detail.location.lat;
          self.jobmanagerinfo.p_longitude = result.detail.location.lng;
          self.marker.setPosition(result.detail.location);
          self.map.setCenter(result.detail.location);
        }
      });
    },
    openMap() {
      this.addMapDialog = true;
      //在这里使用$nextTick延时加载地图
      this.$nextTick(() => {
        this.init();
      });
    }
  },
  async created() {
    if (this.$route.query && this.$route.query.id) {
      const id = this.$route.query && this.$route.query.id;
      const row = {
        ID: Number(id)
      };
      const res = await findJoblist(row);
      if (res.success) {
        this.isEdit = true;
        this.jobmanagerinfo = res.data.rejl;
      } else {
        this.isEdit = false;
        this.jobmanagerinfo = {};
      }
    }
    if (this.enPriseId == 0) {
      if (this.jobmanagerinfo != null) {
        //this.remoteMethod(this.jobmanagerinfo.enterprise_name);
        this.enterpriseInfo.push({
          ID: this.jobmanagerinfo.enterprise_id,
          enterprise_name: this.jobmanagerinfo.enterprise_name,
          enterprise_logo: this.jobmanagerinfo.enterprise_logo
        });
      }
      // getEnterpriseAllInfo().then(res => {
      //   if (res.success) {
      //     this.enterpriseInfo = res.data.result;
      //   } else {
      //     this.enterpriseInfo = [];
      //   }
      // });
    }
    getjoblistOptions().then(res => {
      if (res.success) {
        this.jobyears = res.data.jbwe;
        this.jobsalary = res.data.js;
        this.jobtypes = res.data.jwt;
        this.jobedus = res.data.el;
        this.cityinfo = res.data.cityinfo;
      } else {
        this.jobyears = [];
        this.jobsalary = [];
        this.jobtypes = [];
        this.jobedus = [];
        this.cityinfo = [];
      }
    });
  }
};
</script>
<style lang="scss">
.ql-container {
  height: 200px;
}
.ql-editor strong {
  font-style: normal;
  font-weight: bold;
}
.ql-editor em {
  font-style: italic;
}
</style>
