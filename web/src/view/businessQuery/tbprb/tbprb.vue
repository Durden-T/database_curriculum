<template>
  <div>
    <el-form
      ref="elForm"
      :model="formData"
      :rules="rules"
      :inline="true"
      class="demo-form-inline"
    >
      <el-form-item label="网元名称" prop="enobedName">
        <el-select
          v-model="formData.enobedName"
          placeholder="请选择或输入网元名称"
          filterable
          clearable
        >
          <el-option
            v-for="(item, index) in enobedNameOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="查询属性" prop="type">
        <el-select
          v-model="formData.type"
          placeholder="请选择或输入想查询的属性"
          filterable
          clearable
        >
          <el-option
            v-for="(item, index) in typeOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="开始时间" prop="startTime">
        <el-input
          placeholder="格式为2021-03-27 21:33:00"
          v-model="formData.startTime"
        ></el-input>
      </el-form-item>
      <el-form-item label="结束时间" prop="endTime">
        <el-input
          placeholder="格式为2021-03-27 21:33:00"
          v-model="formData.endTime"
        ></el-input>
      </el-form-item>

      <el-form-item label="时间间隔" prop="interval">
        <el-select
          v-model="formData.interval"
          placeholder="请选择或输入时间间隔"
          filterable
          clearable
        >
          <el-option
            v-for="(item, index) in intervalOptions"
            :key="index"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
      </el-form-item>
    </el-form>

    <div id="chartLineBox" style="width: 90%; height: 75vh"></div>
  </div>
</template>
<script>
let echarts = require("echarts");
require("echarts/lib/chart/line");
require("echarts/lib/component/tooltip"); // tooltip组件

import { getTbprbChartValue, getAllEnodebName } from "@/api/tbprb";
export default {
  data() {
    return {
      formData: {
        enobedName: "",
        type: "",
        startTime: "2020-07-17 00:00:00",
        endTime: "2020-07-19 00:00:00",
        interval: 0,
      },
      rules: {
        enobedName: [
          {
            required: true,
            message: "请选择或输入网元名称",
            trigger: "change",
          },
        ],
        type: [
          {
            required: true,
            message: "请选择或输入想查询的属性",
            trigger: "change",
          },
        ],
        startTime: [
          {
            required: true,
            message: "开始时间不能为空",
            trigger: "change",
          },
        ],
        endTime: [
          {
            required: true,
            message: "结束时间不能为空",
            trigger: "change",
          },
        ],
        interval: [
          {
            required: true,
            message: "时间间隔不能为空",
            trigger: "change",
          },
        ],
      },
      enobedNameOptions: [],
      typeOptions: [],
      intervalOptions: [
        {
          label: "小时级",
          value: 0,
        },
        {
          label: "分钟级",
          value: 1,
        },
      ],
      mychart: null,
      echartsOptions: {
        toolbox: {
          show: true,
          feature: {
            saveAsImage: {
              show: true,
              excludeComponents: ["toolbox", "dataZoom"],
              pixelRatio: 2,
            },
          },
        },
        dataZoom: [
          //滑动条
          {
            xAxisIndex: 0, //这里是从X轴的0刻度开始
            show: true, //是否显示滑动条，不影响使用
            type: "slider", // 这个 dataZoom 组件是 slider 型 dataZoom 组件
            startValue: 0, // 从头开始。
            endValue: 5, // 一次性展示5个。
          },
        ],
        xAxis: {
          type: "category",
          data: [], // x轴数据
          name: "时间", // x轴名称
        },
        yAxis: {
          type: "value",
          name: "",
          scale: true,
        },
        tooltip: {
          trigger: "axis", // axis   item   none三个值
        },
        series: [
          {
            name: "",
            data: [],
            type: "line",
            smooth: true,
          },
        ],
      },
    };
  },
  methods: {
    onSubmit() {
      this.$refs["elForm"].validate((valid) => {
        if (!valid) return;
        this.drawChart();
      });
    },

    async drawChart() {
      let param = this.formData;
      let res = await getTbprbChartValue({
        type: param.type,
        name: param.enobedName,
        startTime: param.startTime,
        endTime: param.endTime,
        interval: param.interval,
      });
      if (res.code != 0) {
        this.$message.error("获取失败");
        return;
      }
      this.$message.success("获取成功");

      this.echartsOptions.xAxis.data = res.data.Xaxis;
      this.echartsOptions.yAxis.name = param.type;
      this.echartsOptions.series[0].data = res.data.Yaxis;
      this.echartsOptions.series[0].name = param.enobedName;

      this.myChart.setOption(this.echartsOptions);
    },
  },

  async created() {
    let res = await getAllEnodebName();
    if (res.code != 0) {
      return;
    }
    this.enobedNameOptions = res.data.map((value) => {
      return {
        label: value,
        value: value,
      };
    });

    for (let i = 0; i < 100; ++i) {
      let val = "prb" + i;
      this.typeOptions.push({
        label: val,
        value: val,
      });
    }
    this.myChart = echarts.init(document.getElementById("chartLineBox"));
  },
};
</script>
<style>
</style>
