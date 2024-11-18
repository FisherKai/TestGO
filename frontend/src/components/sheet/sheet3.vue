<!-- 气象基础设施资池资源变更表 -->
<template>
    <section class="sheet-container">
        <p class="title">气象基础设施资池资源变更表</p>
        <section class="m-container">
            <el-descriptions title="基本信息（填写变更信息即可）" :column="2" border>
                <el-descriptions-item label="变更事由" span="2" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入变更事由" />
                </el-descriptions-item>
                <el-descriptions-item label="联系人" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入联系人" />
                </el-descriptions-item>
                <el-descriptions-item label="电话" label-align="center" align="center">
                    <section class="number-container">
                        <section>
                            <span>办公:</span>
                            <el-input v-model="input" style="width: 150px" placeholder="请输入办公电话" />
                        </section>
                        <section>
                            <span>手机:</span>
                            <el-input v-model="input" style="width: 150px" placeholder="请输入手机号码" />
                        </section>
                    </section>
                </el-descriptions-item>
                <el-descriptions-item label="EMail" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入EMail" />
                </el-descriptions-item>
                <el-descriptions-item label="单位" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入单位名称" />
                </el-descriptions-item>
                <el-descriptions-item label="用户性质" span="2" label-align="center" align="center">
                    <el-radio-group v-model="radio1">
                        <el-radio :value="1">实时业务</el-radio>
                        <el-radio :value="2">政务</el-radio>
                        <el-radio :value="3">一般业务</el-radio>
                        <el-radio :value="4">科研</el-radio>
                        <el-radio :value="5">培训</el-radio>
                    </el-radio-group>
                </el-descriptions-item>
                <el-descriptions-item label="使用期限" span="2" label-align="center" align="center">
                    <el-radio-group v-model="radio1">
                        <el-radio :value="1">长期</el-radio>
                        <el-radio :value="2">
                            <section>
                                <span>临时</span>
                                <el-input style="width: 100px" />年<el-input style="width: 100px"/>月-<el-input style="width: 100px"/>年<el-input style="width: 100px"/>月
                            </section>
                        </el-radio>
                    </el-radio-group>
                </el-descriptions-item>
                <el-descriptions-item label="CPU变更需求" span="2" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入CPU变更需求" />
                </el-descriptions-item>
                <el-descriptions-item label="存储变更需求" span="2" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入存储变更需求" />
                </el-descriptions-item>
                <el-descriptions-item label="网络变更需求" span="2" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入网络变更需求" />
                </el-descriptions-item>
                <el-descriptions-item label="内存变更需求" span="2" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入内存变更需求" />
                </el-descriptions-item>
                <el-descriptions-item label="其他变更需求" span="2" label-align="center" align="center">
                    <el-input v-model="input" placeholder="请输入其他变更需求" />
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
                <el-descriptions-item label="信息保障中心意见" span="2" label-align="center" align="center">
                    <section>
                        <el-input v-model="textarea" :rows="3" type="textarea" placeholder="" />
                        <section class="user">
                            <p>单位负责人：</p>
                            <p class="min-right">日期：</p>
                        </section>
                    </section>
                </el-descriptions-item>
                <el-descriptions-item label="省局主管部门意见" span="2" label-align="center" align="center">
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