<template>
  <div>
    <div>
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <el-input placeholder="请输入位置" v-model="address"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSearch" type="primary">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div id="container" style="width:600px;height:500px;"></div>
  </div>
</template>

<script>
var geocoder,
  map,
  markerIcon = null;
export default {
  name: "news",
  data() {
    return {
      address: "",
      longitude: "39.916527", //经度
      latitude: "116.397128" //纬度
    };
  },
  methods: {
    //第一部分
    //定位获得当前位置信息
    getMyLocation() {
      var geolocation = new qq.maps.Geolocation(
        "R5TBZ-XD7WR-U64WM-WH2RS-4UD57-J4FQG",
        "gogoflying"
      );
      //geolocation.getIpLocation(this.showPosition, this.showErr);
      geolocation.getLocation(this.showPosition, this.showErr); //或者用getLocation精确度比较高
    },
    showPosition(position) {
      this.latitude = position.lat;
      this.longitude = position.lng;
      this.setMap();
    },
    showErr() {
      console.log("定位失败");
      // this.getMyLocation(); //定位失败再请求定位，测试使用
    },
    setMap() {
      //步骤：定义map变量 调用 qq.maps.Map() 构造函数   获取地图显示容器
      //设置地图中心点
      var myLatlng = new qq.maps.LatLng(this.latitude, this.longitude);
      map.setCenter(myLatlng);
      //第三部分
      //给定位的位置添加图片标注
      var marker = new qq.maps.Marker({
        position: myLatlng,
        map: map
      });
      marker.setIcon(markerIcon);
      //alert("经度:"+this.longitude+","+"纬度:"+this.latitude);
    },
    onSearch() {
      geocoder.getLocation(this.address);
    },
    init() {
      //步骤：定义map变量 调用 qq.maps.Map() 构造函数   获取地图显示容器
      //设置地图中心点
      var myLatlng = new qq.maps.LatLng(this.longitude, this.latitude);
      var that = this;
      var myOptions = {
        zoom: 13, //设置地图缩放级别
        center: myLatlng, //设置中心点样式
        mapTypeId: qq.maps.MapTypeId.ROADMAP //设置地图样式详情参见MapType
      };
      //获取dom元素添加地图信息
      map = new qq.maps.Map(document.getElementById("container"), myOptions);
      markerIcon = new qq.maps.MarkerImage(
        "https://3gimg.qq.com/lightmap/api_v2/2/4/99/theme/default/imgs/marker.png",
        new qq.maps.Size(42, 68),
        new qq.maps.Point(0, 0),
        new qq.maps.Point(0, 39)
      );
      //给地图添加点击事件
      //并获取鼠标点击的经纬度
      qq.maps.event.addListener(map, "click", function(event) {
        that.latitude = event.latLng.getLat();
        that.longitude = event.latLng.getLng();
        that.setMap();
      });
      //调用地址解析类
      geocoder = new qq.maps.Geocoder({
        complete: function(result) {
          that.latitude = result.detail.location.lat;
          that.longitude = result.detail.location.lng;
          that.setMap();
        }
      });
    }
  },
  mounted() {
    this.init();
    this.getMyLocation();
  }
};
</script>