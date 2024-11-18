<!-- 资源需求预估表 -->
<template>
    <section class="sheet-container">
        <p class="title">资源需求预估表</p>
        <section class="m-container">
            <el-descriptions title="基本信息（如实完整填写）" :column="2" border>
                <el-descriptions-item label="应用系统名称" span="2" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入应用系统名称" />
                </el-descriptions-item>
                <el-descriptions-item label="应用描述" span="2" label-align="center" align="center">
                    <el-input v-model="textarea" :rows="2" type="textarea" placeholder="请输入应用描述" />
                </el-descriptions-item>
                <el-descriptions-item label="应用类型" label-align="center" align="center">
                    <el-radio-group v-model="radio1">
                        <el-radio :value="1">数据处理</el-radio>
                        <el-radio :value="2">WEB服务</el-radio>
                        <el-radio :value="3">数据库</el-radio>
                        <el-radio :value="4">其他</el-radio>
                    </el-radio-group>
                </el-descriptions-item>
                <el-descriptions-item label="是否对接天擎系统" label-align="center" align="center">
                    <el-radio-group v-model="radio2">
                        <el-radio :value="1">是</el-radio>
                        <el-radio :value="2">否</el-radio>
                    </el-radio-group>
                </el-descriptions-item>
                <el-descriptions-item label="存储数据种类（数据库填写）" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入存储数据种类" />
                </el-descriptions-item>
                <el-descriptions-item label="存储方式" label-align="center" align="center">
                    <el-radio-group v-model="radio2">
                        <el-radio :value="1">数据库</el-radio>
                        <el-radio :value="2">文件库</el-radio>
                    </el-radio-group>
                </el-descriptions-item>
                <el-descriptions-item label="WEB服务器" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入WEB服务器" />
                </el-descriptions-item>
                <el-descriptions-item label="支撑用户" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入支撑用户" />
                </el-descriptions-item>
                <el-descriptions-item label="计算资源预估" span="2" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入计算资源" />
                </el-descriptions-item>
                <el-descriptions-item label="内存预估" span="2" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入内存" />
                </el-descriptions-item>
                <el-descriptions-item label="数据规模预估" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入数据规模预估" />
                </el-descriptions-item>
                <el-descriptions-item label="保存期限" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入保存期限" />
                </el-descriptions-item>
                <el-descriptions-item label="预计存储资源增长情况" span="2" label-align="center" align="center">
                    <el-input v-model="input" placeholder="" />
                </el-descriptions-item>
                <el-descriptions-item label="申请单位意见" span="2" label-align="center" align="center">
                    <section>
                        <el-input v-model="textarea" :rows="3" type="textarea" placeholder="" />
                        <section class="user">
                            <p>单位负责人：</p>
                            <p class="min-right">日期：</p>
                        </section>
                    </section>
                </el-descriptions-item>
            </el-descriptions>
        </section>
    </section>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const radio1 = ref(1)
</script>

<style scoped>
.sheet-container {}

.title {
    text-align: center;
    font-size: 20px;
    font-weight: 600;
    margin-top: 50px;
}

.m-container {
    padding: 80px 200px;
}

.number-container {
    display: flex;
    justify-content: space-around;
}

.user{
    display: flex;
    justify-content: space-between;
    text-align: left;
    padding: 10px 0;
    font: 22px;
    line-height: 1;
}
.min-right{
    padding-right: 200px;
}
</style>