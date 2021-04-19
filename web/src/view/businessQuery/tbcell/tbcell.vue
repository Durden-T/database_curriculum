<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="小区ID">
          <el-input placeholder="" v-model="searchInfo.sectorId"></el-input>
        </el-form-item>
        <el-form-item label="小区名称">
          <el-input placeholder="" v-model="searchInfo.sectorName"></el-input>
        </el-form-item>
        <el-form-item label="基站ID">
          <el-input placeholder="" v-model="searchInfo.enodebId"></el-input>
        </el-form-item>
        <el-form-item label="基站名称">
          <el-input placeholder="" v-model="searchInfo.enodebName"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
      <el-table-column label="城市" prop="city" width="120"></el-table-column>

      <el-table-column
        label="小区ID"
        prop="sectorId"
        width="120"
      ></el-table-column>

      <el-table-column
        label="小区名称"
        prop="sectorName"
        width="120"
      ></el-table-column>

      <el-table-column
        label="基站ID"
        prop="enodebId"
        width="120"
      ></el-table-column>

      <el-table-column
        label="基站名称"
        prop="enodebName"
        width="120"
      ></el-table-column>

      <el-table-column
        label="频点编号"
        prop="earfcn"
        width="120"
      ></el-table-column>

      <el-table-column
        label="物理小区标识"
        prop="pci"
        width="120"
      ></el-table-column>

      <el-table-column
        label="主同步信号标识"
        prop="pss"
        width="120"
      ></el-table-column>

      <el-table-column
        label="辅同步信号标识"
        prop="sss"
        width="120"
      ></el-table-column>

      <el-table-column
        label="跟踪区编码"
        prop="tac"
        width="120"
      ></el-table-column>

      <el-table-column
        label="设备厂家"
        prop="vendor"
        width="120"
      ></el-table-column>

      <el-table-column
        label="经度"
        prop="longitude"
        width="120"
      ></el-table-column>

      <el-table-column
        label="纬度"
        prop="latitude"
        width="120"
      ></el-table-column>

      <el-table-column
        label="基站类型"
        prop="style"
        width="120"
      ></el-table-column>

      <el-table-column
        label="小区天线方位角"
        prop="azimuth"
        width="120"
      ></el-table-column>

      <el-table-column
        label="小区天线高度"
        prop="height"
        width="120"
      ></el-table-column>

      <el-table-column
        label="小区天线电下倾角"
        prop="electtilt"
        width="120"
      ></el-table-column>

      <el-table-column
        label="小区天线机械下倾角"
        prop="mechtilt"
        width="120"
      ></el-table-column>

      <el-table-column
        label="总下倾角"
        prop="totletilt"
        width="120"
      ></el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{ float: 'right', padding: '20px' }"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
  </div>
</template>

<script>
import { getTbcellList } from "@/api/tbcell"; //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Tbcell",
  mixins: [infoList],
  data() {
    return {
      listApi: getTbcellList,
      formData: {
        city: "",
        sectorId: "",
        sectorName: "",
        enodebId: 0,
        enodebName: "",
        earfcn: 0,
        pci: 0,
        pss: 0,
        sss: 0,
        tac: 0,
        vendor: "",
        longitude: 0,
        latitude: 0,
        style: "",
        azimuth: 0,
        height: 0,
        electtilt: 0,
        mechtilt: 0,
        totletilt: 0,
      },
    };
  },
  methods: {
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.getTableData();
    },
  },
  async created() {
    await this.getTableData();
  },
};
</script>

<style>
</style>
