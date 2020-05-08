<template>
  <div>
    <div class="edit_container">
      <el-form :rules="rules" ref="jobForm" :model="jobmanagerinfo">
        <el-form-item label="工作名称" label-width="120px" prop="p_name">
          <el-input v-model="jobmanagerinfo.p_name" placeholder="请输入工作名称" style="width:50%;"></el-input>
        </el-form-item>
        <el-form-item v-show="enPriseId == 0" label="所属公司" label-width="120px" prop="p_salary_high">
          <el-select
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
          </el-select>
        </el-form-item>
        <el-form-item label="薪资上限" label-width="120px" prop="p_salary_high">
          <el-input
            v-model.number="jobmanagerinfo.p_salary_high"
            placeholder="请输入薪资上限"
            style="width:50%;"
          ></el-input>
        </el-form-item>
        <el-form-item label="薪资下限" label-width="120px" prop="p_salary_low">
          <el-input
            v-model.number="jobmanagerinfo.p_salary_low"
            placeholder="请输入薪资下限"
            style="width:50%;"
          ></el-input>
        </el-form-item>
        <el-form-item label="工作地点" label-width="120px" prop="p_address">
          <el-input v-model="jobmanagerinfo.p_address" placeholder="请输入工作地点" style="width:45%;"></el-input>
          <el-button size="mini" icon="el-icon-map-location" @click="openMap"></el-button>
        </el-form-item>
        <el-form-item label="工作地点纬度" label-width="120px" prop="p_latitude" hidden>
          <el-input v-model="jobmanagerinfo.p_latitude" placeholder="请输入工作地点纬度" style="width:50%;"></el-input>
        </el-form-item>
        <el-form-item label="工作地点经度" label-width="120px" prop="p_longitude" hidden>
          <el-input v-model="jobmanagerinfo.p_longitude" placeholder="请输入工作地点经度" style="width:50%;"></el-input>
        </el-form-item>
        <el-form-item label="工作城市" label-width="120px" prop="p_city_id">
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
        <el-form-item label="工作年限" label-width="120px" prop="p_edujy_id">
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
        <el-form-item label="最低学历" label-width="120px" prop="p_education_id">
          <el-select
            @change="selectJobedu"
            placeholder="请选择最低学历"
            v-model="jobmanagerinfo.p_education_id"
          >
            <el-option
              :key="jobedu.name"
              :label="jobedu.name"
              :value="jobedu.id"
              v-for="jobedu in jobedus"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工作类型" label-width="120px" prop="p_type_id">
          <el-select
            @change="selectJobtype"
            placeholder="请选择工作类型"
            v-model.number="jobmanagerinfo.p_type_id"
          >
            <el-option
              :key="jobtype.name"
              :label="jobtype.name"
              :value="jobtype.id"
              v-for="jobtype in jobtypes"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工作描述" label-width="120px" prop="p_desc">
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
        <el-row type="flex" justify="center">
          <el-col :span="13">
            <el-button @click="closeAddjobDialog">取 消</el-button>
            <el-button @click="enterAddjobDialog" type="primary">确 定</el-button>
          </el-col>
        </el-row>
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
      jobmanagerinfo: {
        p_name: "",
        p_salary_high: "",
        p_salary_low: "",
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
        enterprise_id: "",
        enterprise_name: "",
        enterprise_logo: ""
      },
      cityinfo: [
        {
          id: 1,
          name: "北京"
        },
        {
          id: 2,
          name: "上海"
        },
        {
          id: 3,
          name: "广州"
        }
      ],
      jobyears: [
        {
          id: 0,
          name: "1年"
        },
        {
          id: 1,
          name: "3年"
        },
        {
          id: 2,
          name: "5年"
        }
      ],
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
      rules: {
        p_name: [
          {
            required: true,
            message: "请输入工作名称",
            trigger: "blur"
          }
        ],
        p_salary_high: [
          { required: true, message: "请输入薪资上限", trigger: "blur" }
        ],
        p_salary_low: [
          { required: true, message: "请输入薪资下限", trigger: "blur" }
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
        return item.id === val;
      });
      this.jobmanagerinfo.p_type = selectedItem.name;
    },
    selectJobedu(val) {
      var selectedItem = {};
      selectedItem = this.jobedus.find(item => {
        return item.id === val;
      });
      this.jobmanagerinfo.p_education = selectedItem.name;
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
        }
      });
      this.$router.push({ name: "jobmanagerinfo" });
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
  created() {
    if (this.$route.query && this.$route.query.id) {
      const id = this.$route.query && this.$route.query.id;
      const row = {
        ID: Number(id)
      };
      findJoblist(row).then(res => {
        if (res.success) {
          this.isEdit = true;
          this.jobmanagerinfo = res.data.rejl;
        } else {
          this.isEdit = false;
          this.jobmanagerinfo = {};
        }
      });
    }
    if (this.enPriseId == 0) {
      getjoblistOptions().then(res => {
        if (res.success) {
          this.enterpriseInfo = res.data.rep;
          this.jobyears = res.data.jbwe;
          this.cityinfo = res.data.cityinfo;
        } else {
          this.enterpriseInfo = [];
          this.jobyears = [];
          this.cityinfo = [];
        }
      });
    }
  }
};
</script>
<style>
.ql-editor strong {
  font-style: normal;
  font-weight: bold;
}
.ql-editor em {
  font-style: italic;
}
</style>
