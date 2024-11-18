<!-- 监视告警 -->
<template>
  <section class="view-of-monitoring-alarm">
    <span class="title">告警列表</span>
    <el-alert title="" type="info" description="显示系统所有的告警信息，帮助用户及时发现与快速定位故障，单击告警名称，查看告警处理指导。同时提供导出功能，便于用户定期收集与备份。"
      show-icon />
    <section>
      <section class="action-container">
        <section class="left-container">
          <el-button type="primary">清除</el-button>
          <el-button>导出</el-button>
        </section>

        <section class="right-container">
          <el-select v-model="value" placeholder="请选择筛选级别" style="width: 240px">
          </el-select>

          <div class="block">
            <span class="demonstration" style="margin-right: 10px;">产生时间:</span>
            <el-date-picker v-model="value1" type="datetimerange" range-separator="至" start-placeholder="开始时间"
              end-placeholder="结束时间" />
          </div>

          <el-input style="width: 200px;" placeholder="请输入告警名称或告警对象"></el-input>
          <el-button style="margin-left: 10px;" type="primary">搜索</el-button>
        </section>
      </section>
      <el-table ref="multipleTableRef" @selection-change="handleSelectionChange" :data="tableData" style="width: 100%">
        <el-table-column type="selection" width="55" />
        <el-table-column prop="type" label="级别">
          <template #default="scope">
            <el-tag v-if="scope.row.type === '次要'" type="warning">{{ scope.row.type }}</el-tag>
            <el-tag v-else-if="scope.row.type === '重要'" type="danger">{{ scope.row.type }}</el-tag>
            <el-tag v-else type="primary">{{ scope.row.type }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="name" label="告警名称" />
        <el-table-column prop="object" label="告警对象" />
        <!-- <el-table-column prop="objectType" label="对象类型" /> -->
        <!-- <el-table-column prop="partType" label="部件类型" /> -->
        <el-table-column prop="createTime" label="产生时间" />
        <!-- <el-table-column prop="updateTime" label="清除时间" />
        <el-table-column prop="updateType" label="清除类型" /> -->
        <el-table-column fixed="right" label="操作">
          <template #default>
            <el-button link type="primary" size="small" @click="handleClick">
              清除
            </el-button>
            <el-button link type="primary" size="small">屏蔽</el-button>
          </template>
        </el-table-column>
      </el-table>

      <section class="pagination-container">
        显示第1到第4条记录，总共4条记录
        <el-pagination v-model:page-size="pageSize2" :page-sizes="[10, 20, 30, 50]" background layout="prev, pager, next"
          :total="10" />
      </section>
    </section>
  </section>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElTable } from 'element-plus'

const multipleTableRef = ref<InstanceType<typeof ElTable>>()
const multipleSelection = ref<any[]>([])

const toggleSelection = (rows?: any[]) => {
  if (rows) {
    rows.forEach((row) => {
      multipleTableRef.value!.toggleRowSelection(row, undefined)
    })
  } else {
    multipleTableRef.value!.clearSelection()
  }
}
const handleSelectionChange = (val: any[]) => {
  multipleSelection.value = val
}
const tableData = [{
  type: '次要',
  name: 'CPU占用已达75%',
  object: '华三超融合资源池',
  objectType: 'host',
  partType: 'OpenStack',
  createTime: '2024-03-01 14:34:23',
  updateTime: '-',
  updateType: '-'
}, {
  type: '次要',
  name: '存储占用已达70%',
  object: '华为资源池',
  objectType: 'host',
  partType: 'OpenStack',
  createTime: '2024-04-03 15:33:12',
  updateTime: '-',
  updateType: '-'
}, {
  type: '重要',
  name: 'CPU占用已达85%',
  object: '华三超融合资源池',
  objectType: 'host',
  partType: 'OpenStack',
  createTime: '2024-06-03 14:33:11',
  updateTime: '-',
  updateType: '-'
}, {
  type: '重要',
  name: '存储占用已达80%',
  object: '华为资源池',
  objectType: 'host',
  partType: 'OpenStack',
  createTime: '2024-08-06 09:33:14',
  updateTime: '-',
  updateType: '-'
}]
</script>

<style scoped>
.view-of-monitoring-alarm {
  padding: 10px;
}

.title {
  font-size: 20px;
  font-weight: 600;
}

.pagination-container {
  margin-top: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.action-container {
  display: flex;
  padding: 20px 10px;
  align-items: center;
  justify-content: space-between;
}

.right-container {
  display: flex;
  align-items: center;
}

.block {
  margin: 0 20px;
}
</style>