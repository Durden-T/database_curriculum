<template>
  <div>
    <el-form
      ref="elForm"
      :model="formData"
      :inline="true"
      class="demo-form-inline"
    >
      <el-form-item label="阈值参数" prop="threshold">
        <el-input
          placeholder="请输入阈值参数"
          v-model="formData.threshold"
        ></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { sectorTripleAnalysis } from "@/api/tbc2i3";
export default {
  data() {
    return {
      formData: {
        threshold: 0.0,
      },
    };
  },
  methods: {
    onSubmit() {
      this.sendRequest();
    },

    async sendRequest() {
      let res = await sectorTripleAnalysis({
        threshold: this.formData.threshold,
      });
      if (res.code != 0) {
        this.$message.error("获取失败");
        return;
      }
      this.$message.success("获取成功");
    },
  },
};
</script>
<style>
</style>
