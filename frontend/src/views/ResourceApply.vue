<!-- 资源审批 -->
<template>
  <section class="application-container">
    <el-table :data="tableData" style="width: 100%">
      <el-table-column prop="id" label="序号" />
      <el-table-column prop="type" label="申请类型" />
      <el-table-column prop="title" label="标题" />
      <el-table-column prop="updateTime" label="办理时间" />
      <el-table-column prop="createUser" label="拟稿人" />
      <el-table-column prop="createTime" label="拟稿时间" />
      <el-table-column prop="status" label="审批状态" />
      <el-table-column fixed="right" label="操作">
        <template #default>
          <el-button link type="primary" size="small" @click="handleClick">查看详情</el-button>
          <!-- <el-button link type="success" size="small" @click="handleClick">
          通过
        </el-button>
        <el-button link type="danger" size="small">驳回</el-button> -->
        </template>
      </el-table-column>
    </el-table>
    <section class="pagination-container">
      显示第1到第4条记录，总共4条记录
      <el-pagination v-model:page-size="pageSize2" :page-sizes="[10, 20, 30, 50]" background layout="prev, pager, next"
        :total="10" />
    </section>

    <section class="m-dailog">
      <el-dialog v-model="dialogTableVisible" title="查看详情" width="1600">
        <!-- 操作 -->
        <section class="action-container">
          <el-button size="small">下载申请表</el-button>
        </section>

        <!-- 流程 -->
        <section>
          <p style="padding: 15px 0;">
            <span style="font-weight: 600;">当前状态：</span>
            <span style="color: rgb(70, 140, 220);">当前节点：用户单位局级领导，当前办理人：结束流程</span>
          </p>
          <el-steps style="max-width: 1200px" :active="5" finish-status="success" align-center>
            <el-step title="申请人" description="2024-08-15 18:05:47" />
            <el-step title="申请单位意见" description="2024-08-16 09:28:34" />
            <el-step title="信息保障中心意见" description="2024-08-16 10:03:34" />
            <el-step title="省局主管部门意见" description="2024-08-16 11:05:27" />
            <el-step title="资源池运维人员 结束流程" description="2024-08-16 16:35:11" />
          </el-steps>
        </section>

        <!-- 意见 -->
        <section style="margin-top: 30px;">
          <span style="font-weight: 600; padding: 15px 0;">审批意见</span>
          <el-tabs v-model="activeName1" class="demo-tabs" @tab-click="handleClick">
            <el-tab-pane label="申请单位意见" name="first">同意</el-tab-pane>
            <el-tab-pane label="信息保障中心意见" name="second">同意</el-tab-pane>
            <el-tab-pane label="省局主管部门意见" name="third">同意</el-tab-pane>
            <el-tab-pane label="资源池运维人员" name="four">{{value4}}</el-tab-pane>
          </el-tabs>
        </section>

        <!-- 基本信息 -->
        <section style="margin-top: 30px;">
          <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
            <el-tab-pane label="基本信息" name="first">
              <el-descriptions :column="2" border>
                <el-descriptions-item label="具体联系人" label-align="center" align="center">
                    <el-input v-model="value2" readonly disabled/>
                </el-descriptions-item>
                <el-descriptions-item label="电话" label-align="center" align="center">
                    <section class="number-container">
                        <section>
                            <span>办公</span>
                            <el-input v-model="value3" readonly disabled />
                        </section>
                        <section>
                            <span>手机</span>
                            <el-input v-model="value3" readonly disabled />
                        </section>
                    </section>
                </el-descriptions-item>
                <el-descriptions-item label="EMail" label-align="center" align="center">
                    <el-input v-model="value6" readonly disabled />
                </el-descriptions-item>
                <el-descriptions-item label="单位" label-align="center" align="center">
                  <el-input v-model="value7" readonly disabled />
                </el-descriptions-item>
                <el-descriptions-item label="用户性质" span="2" label-align="center" align="center">
                  <el-input v-model="value8" readonly disabled />
                </el-descriptions-item>
                <el-descriptions-item label="应用描述" span="2" label-align="center" align="center">
                    <el-input v-model="value5" :rows="2" type="textarea" readonly disabled/>
                </el-descriptions-item>
                <el-descriptions-item label="是否属于互联网应用" span="2" label-align="center" align="center">
                  <el-input v-model="value9" readonly disabled />
                </el-descriptions-item>
                <el-descriptions-item label="操作系统" span="2" label-align="center" align="center">
                  <el-input v-model="value10" readonly disabled />
                </el-descriptions-item>
                <el-descriptions-item label="防火墙设置" span="2" label-align="center" align="center">
                    <el-input v-model="input" readonly disabled/>
                </el-descriptions-item>
                <el-descriptions-item label="存储资源（GB）" span="2" label-align="center" align="center">
                    <el-input v-model="value11" readonly disabled/>
                </el-descriptions-item>
                <el-descriptions-item label="内存（GB）" span="2" label-align="center" align="center">
                  <el-input v-model="value12" readonly disabled/>
                </el-descriptions-item>
                <el-descriptions-item label="计算资源(CPU内核数)" label-align="center" align="center">
                  <el-input v-model="value12" readonly disabled/>
                </el-descriptions-item>
                <el-descriptions-item label="网络配置需求" label-align="center" align="center">
                  <el-input readonly disabled/>
                </el-descriptions-item>
            </el-descriptions>
            </el-tab-pane>
          </el-tabs>
        </section>

        <!-- 责任人信息 -->
        <!-- <section style="margin-top: 30px;">
          <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
            <el-tab-pane label="责任人信息" name="first">

            </el-tab-pane>
          </el-tabs>
        </section> -->
      </el-dialog>
    </section>
  </section>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const dialogTableVisible = ref(false)
const activeName = ref('first')
const activeName1 = ref('four')
const value2 = ref('韩书倩')
const value3 = ref('18375183097')
const value4 = ref('虚拟机IP地址：10.28.105.138')
const value5 = ref('极端天气气候事件监测服务系统(省级)-气象灾害识别、气象灾害算法规则管理、气象灾害序列产品加工及可视化应用')
const value6 = ref('1079860031@qq.com')
const value7 = ref('湖北省气象信息与技术保障中心')
const value8 = ref('实时业务')
const value9 = ref('否')
const value10 = ref('Linux')
const value11 = ref('500')
const value12 = ref('64')


const pageSize2 = ref(10)
const handleClick = () => {
  console.log('click')
  dialogTableVisible.value = true
}

const tableData = [
  {
    id: 1,
    type: "资源申请",
    title: "韩书倩对气象基础设施资源池的申请",
    updateTime: "2024-06-21 17.05:47",
    createUser: "韩书倩",
    createTime: "2024-06-17 14:25:31",
    status: "办结",
  },
  {
    id: 2,
    type: "资源申请",
    title: "张萧对气象基础设施资源池的申请",
    updateTime: "2024-06-24 15.55:36",
    createUser: "张萧",
    createTime: "2024-06-19 18:05:47",
    status: "办结",
  },
  {
    id: 3,
    type: "资源变更",
    title: "张萧对气象基础设施资源池的变更",
    updateTime: "2024-07-14 11.08:45",
    createUser: "张萧",
    createTime: "2024-07-11 09:44:23",
    status: "办结",
  },
  {
    id: 4,
    type: "资源变更",
    title: "韩书倩对气象基础设施资源池的变更",
    updateTime: "2024-08-07 18.05:47",
    createUser: "韩书倩",
    createTime: "2024-08-02 14:05:23",
    status: "办结",
  }
]
</script>

<style>
.application-container {
  margin: 20px 30px;
}

.pagination-container {
  margin-top: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>