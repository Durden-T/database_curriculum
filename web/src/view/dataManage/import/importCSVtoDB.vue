<template>
  <div>
    <el-form
      ref="elForm"
      :model="formData"
      :inline="true"
      class="demo-form-inline"
    >
      <el-form-item label="数据库名称">
        <el-input
          @blur="getTable"
          placeholder="请输入数据库名称"
          v-model="formData.databaseName"
        ></el-input>
      </el-form-item>

      <el-form-item label="表格名称" prop="tableName">
        <el-select
          v-model="formData.tableName"
          placeholder="请选择或输入表格名称"
          filterable
          clearable
        >
          <el-option
            v-for="item in tableOptions"
            :key="item.tableName"
            :label="item.tableName"
            :value="item.tableName"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="csv文件位置">
        <el-input
          placeholder="请输入csv文件位置"
          v-model="formData.csvPath"
        ></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { importCSVtoDB } from "@/api/importCSVtoDB";
import { getTable } from "@/api/autoCode";

export default {
  data() {
    return {
      formData: {
        databaseName: "",
        sectorName: "",
        csvPath: "",
      },
      tableOptions: [],
    };
  },
  methods: {
    onSubmit() {
      this.$refs["elForm"].validate((valid) => {
        if (!valid) return;
        this.importCSV();
      });
    },

    async importCSV() {
      let param = this.formData;
      let res = await importCSVtoDB({
        tableName: param.tableName,
        csvPath: param.csvPath,
      });
      if (res.code == 0) {
        this.$message.success("导入成功");
      }
    },
    async getTable() {
      console.log('fuck')
      const res = await getTable({ dbName: this.formData.databaseName });
      if (res.code == 0) {
        this.tableOptions = res.data.tables;
      }
    },
  },
};
</script>
<style>
</style>
