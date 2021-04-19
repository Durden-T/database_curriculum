<template>
  <div>
    <el-form
      ref="elForm"
      :model="formData"
      :rules="rules"
      :inline="true"
      class="demo-form-inline"
    >
      <el-form-item label="小区名称" prop="sectorName">
        <el-select
          v-model="formData.sectorName"
          placeholder="请选择或输入小区名称"
          filterable
          clearable
        >
          <el-option
            v-for="(item, index) in sectorNameOptions"
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

      <el-form-item label="日期范围" prop="time">
        <el-date-picker
          type="daterange"
          v-model="formData.time"
          format="yyyy-MM-dd"
          value-format="yyyy-MM-dd"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          range-separator="至"
          clearable
        ></el-date-picker>
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

import { getTbkpiChartValue, getAllSectorName } from "@/api/tbkpi";
export default {
  data() {
    return {
      formData: {
        sectorName: "",
        type: "",
        time: ["2020-07-17 00:00:00", "2020-07-19 00:00:00"],
      },
      rules: {
        sectorName: [
          {
            required: true,
            message: "请选择或输入小区名称",
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
        time: [
          {
            required: true,
            message: "日期范围不能为空",
            trigger: "change",
          },
        ],
      },
      sectorNameOptions: [],
      typeOptions: [
        {
          label: "RCC连接建立完成次数",
          value: "rcc_conn_succ",
        },
        {
          label: "RCC连接请求次数",
          value: "rcc_conn_att",
        },
        {
          label: "RCC连接成功率",
          value: "rcc_conn_rate",
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
          name: "日期", // x轴名称
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
      let res = await getTbkpiChartValue({
        type: param.type,
        name: param.sectorName,
        startTime: param.time[0],
        endTime: param.time[1],
      });
      if (res.code != 0) {
        this.$message.error("获取失败");
        return;
      }

      this.$message.success("获取成功");

      this.echartsOptions.xAxis.data = res.data.Xaxis;
      this.echartsOptions.yAxis.name = param.type;
      this.echartsOptions.series[0].data = res.data.Yaxis;
      this.echartsOptions.series[0].name = param.sectorName;

      this.myChart.setOption(this.echartsOptions);
    },
  },

  async created() {
    let res = await getAllSectorName();
    if (res.code != 0) {
      return;
    }
    this.sectorNameOptions = res.data.map((value) => {
      return {
        label: value,
        value: value,
      };
    });

    this.myChart = echarts.init(document.getElementById("chartLineBox"));
  },
};
</script>
<style>
</style>
