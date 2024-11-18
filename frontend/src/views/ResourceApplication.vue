<!-- 资源申请 -->
<template>
  <section class="sheet-container">
    <el-tabs v-model="activeName" class="m-tabs" @tab-click="handleClick">
      <el-tab-pane label="气象基础设施资源池资源申请表" name="first">
        <sheet1Vue></sheet1Vue>
      </el-tab-pane>
      <el-tab-pane label="资源需求预估表" name="second">
        <sheet2Vue></sheet2Vue>
      </el-tab-pane>
      <el-tab-pane label="气象基础设施资池资源变更表" name="third">
        <sheet3Vue></sheet3Vue>
      </el-tab-pane>
    </el-tabs>
  </section>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import sheet1Vue from "../components/sheet/sheet1.vue";
import sheet2Vue from "../components/sheet/sheet2.vue";
import sheet3Vue from "../components/sheet/sheet3.vue";

const activeName = ref('first')

const handleClick = (tab: any, event: Event) => {
  console.log(tab, event)
}
</script>

<style>
.m-tabs{
  margin-left: 10px;
}
</style>