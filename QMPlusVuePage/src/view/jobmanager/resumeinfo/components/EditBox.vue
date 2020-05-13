<template>
  <div id="EditBox">
    <el-tabs @tab-click="tabClick" tab-position="left">
      <el-tab-pane label="基本信息"></el-tab-pane>
      <el-tab-pane label="工作经历"></el-tab-pane>
      <el-tab-pane label="教育信息"></el-tab-pane>
      <el-tab-pane label="期望职位"></el-tab-pane>
    </el-tabs>
    <ol class="write">
      <li v-bind:class="{active: currentTab === 0}">
        <EditPersonal
          v-on:freshResume="freshResume"
          v-bind:option="option"
          v-bind:user_base_info="resume.user_base_info"
          v-bind:openid="resume.openid"
        />
      </li>
      <li v-bind:class="{active: currentTab === 1}">
        <EditWork
          v-on:freshResume="freshResume"
          v-bind:option="option"
          v-bind:workExperience="resume.user_works"
          v-bind:openid="resume.openid"
        />
      </li>
      <li v-bind:class="{active: currentTab === 2}">
        <EditEdu
          v-on:freshResume="freshResume"
          v-bind:option="option"
          v-bind:eduExperience="resume.user_edus"
          v-bind:openid="resume.openid"
        />
      </li>
      <li v-bind:class="{active: currentTab === 3}">
        <EditDream
          v-on:freshResume="freshResume"
          v-bind:option="option"
          v-bind:user_dreams="resume.user_dreams"
          v-bind:openid="resume.openid"
        />
      </li>
    </ol>
  </div>
</template>

<script>
import EditPersonal from "./EditPersonal";
import EditWork from "./EditWork";
import EditEdu from "./EditEdu";
import EditDream from "./EditDream";
import { getUserOptions } from "@/api/resumestatus";
export default {
  components: {
    EditPersonal,
    EditWork,
    EditEdu,
    EditDream
  },
  props: ["resume"],
  data() {
    return {
      currentTab: 0,
      option: {}
    };
  },
  methods: {
    tabClick(tab) {
      this.currentTab = parseInt(tab.index);
    },
    freshResume() {
      this.$emit("freshResume");
    }
  },
  async created() {
    const res = await getUserOptions();
    if (res.success) {
      this.option = res.data;
    }
  }
};
</script>

<style lang="scss">
#EditBox {
  display: flex;
  > .write {
    overflow: hidden;
    flex: 1;
    .useless {
      position: relative;
      > i {
        font-size: 20px;
      }
      > .el-icon-delete {
        position: absolute;
        right: 0;
        top: 12px;
      }
      > .el-icon-check {
        position: absolute;
        right: 10%;
        top: 12px;
      }
    }
    > li {
      display: none;
      padding: 32px;
      height: 100%;
      overflow: auto;
    }
    & .active {
      display: block;
    }
  }
}
</style>
