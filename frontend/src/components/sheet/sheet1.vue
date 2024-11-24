<!-- 气象基础设施资源池资源申请表 -->
<template>
  <section class="sheet-container">
    <el-alert title="提示" type="warning" description="注:计算资源、存储资源、内存资源应有充足的测算依据，请以最小满足业务需求为准,按需申请使用。信息保障中心会对资源使用情况开展定期评估,并根据评估结果对资源进行
适当调整。" show-icon/>
    <p class="title">气象基础设施资源池资源申请表</p>
    <section class="m-container">
      <el-descriptions title="基本信息（如实完整填写）" :column="2" border>
        <el-descriptions-item label="具体联系人" label-align="center" align="center">
          <el-input v-model="formData.userName" placeholder="请输入具体联系人"/>
        </el-descriptions-item>
        <el-descriptions-item label="电话" label-align="center" align="center">
          <section class="number-container">
            <section>
              <span>办公:</span>
              <el-input v-model="formData.callNum" style="width: 150px" placeholder="请输入办公电话"/>
            </section>
            <section>
              <span>手机:</span>
              <el-input v-model="formData.phoneNum" style="width: 150px" placeholder="请输入手机号码"/>
            </section>
          </section>
        </el-descriptions-item>
        <el-descriptions-item label="EMail" label-align="center" align="center">
          <el-input v-model="formData.Email" placeholder="请输入EMail"/>
        </el-descriptions-item>
        <el-descriptions-item label="单位" label-align="center" align="center">
          <el-select v-model="formData.unitValue" placeholder="请选择单位">
            <el-option v-for="item in optionsUnit" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </el-descriptions-item>
        <el-descriptions-item label="用户性质" span="2" label-align="center" align="center">
          <el-radio-group v-model="formData.userType">
            <el-radio :value="1">实时业务</el-radio>
            <el-radio :value="2">政务</el-radio>
            <el-radio :value="3">一般业务</el-radio>
            <el-radio :value="4">科研</el-radio>
            <el-radio :value="5">培训</el-radio>
          </el-radio-group>
        </el-descriptions-item>
        <el-descriptions-item label="应用描述" span="2" label-align="center" align="center">
          <el-input v-model="formData.applicationDescription" :rows="2" type="textarea" placeholder="请输入应用描述"/>
        </el-descriptions-item>
        <el-descriptions-item label="是否属于互联网应用" span="2" label-align="center" align="center">
          <el-radio-group v-model="formData.isInternet">
            <el-radio :value="1">是</el-radio>
            <el-radio :value="2">否</el-radio>
          </el-radio-group>
        </el-descriptions-item>
        <el-descriptions-item label="操作系统" span="2" label-align="center" align="center">
          <el-select v-model="formData.OS" placeholder="请选择操作系统">
            <el-option v-for="item in optionsOSList" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </el-descriptions-item>
        <el-descriptions-item label="防火墙设置" span="2" label-align="center" align="center">
          <el-input v-model="formData.firewallSetting" placeholder="请输入防火墙设置"/>
        </el-descriptions-item>
        <el-descriptions-item label="存储资源（GB）" span="2" label-align="center" align="center">
          <el-input v-model="formData.memoryGB" placeholder="请输入存储资源（GB）"/>
        </el-descriptions-item>
        <el-descriptions-item label="内存（GB）" span="2" label-align="center" align="center">
          <el-input v-model="formData.RAMGB" placeholder="请输入内存（GB）"/>
        </el-descriptions-item>
        <el-descriptions-item label="计算资源(CPU内核数)" label-align="center" align="center">
          <el-input v-model="formData.cpuNum" placeholder="请输入计算资源(CPU内核数)"/>
        </el-descriptions-item>
        <el-descriptions-item label="网络配置需求" label-align="center" align="center">
          <el-select v-model="formData.networkSetting" placeholder="请选择内网/DMZ区">
            <el-option v-for="item in optionsNetworkList" :key="item.value" :label="item.label" :value="item.value"/>
          </el-select>
        </el-descriptions-item>
        <!-- <el-descriptions-item label="申请单位意见" span="2" label-align="center" align="center">
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
        </el-descriptions-item> -->
      </el-descriptions>

      <section class="m-action">
        <el-button type="primary" size="large" style="width: 200px;" @click="submitSheet()">提交</el-button>
        <el-button size="large" style="width: 200px;" @click="saveSheet()">保存</el-button>
      </section>
    </section>
  </section>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import api from "../../api/api.js"
import { CommonVerification } from "../../utils/common.js"
import { ElMessage } from 'element-plus'
import {Sheet1Name,OptionsUnit,OptionsOSList,OptionsNetworkList} from "../../appconst/appconst.js"
import tips from "../../appconst/appconst.js"

const sheetName = Sheet1Name

const optionsUnit = ref(OptionsUnit)

const optionsOSList = ref(OptionsOSList)

const optionsNetworkList = ref(OptionsNetworkList)

const formData = ref({
  userName: "",
  callNum: "",
  phoneNum: "",
  Email: "",
  unitValue: 1,
  userType: 1,
  applicationDescription: "",
  isInternet: 1,
  OS: 1,
  firewallSetting: "",
  memoryGB: "",
  RAMGB: "",
  cpuNum: "",
  networkSetting: ""
})

const submitSheet = async () => {
  const valid = CommonVerification(formData.value)
  if(!valid){
    ElMessage.error(tips.ErrorTips.Incomplete_Of_Information)
    return;
  }
  const resp = await api.createContent({
    name: sheetName,
    content: JSON.stringify(formData.value),
    resource_type: "1"
  }, "1")

  console.log('resp',resp)
}

const saveSheet = async () => {
  const resp = await api.createContent({
    name: sheetName,
    content: formData,
    resource_type: "1"
  })
  if(resp && resp.data && resp.data.id){
      ElMessage.success("创建审批单成功")
  }
}


</script>

<style scoped>
.sheet-container {
}

.m-action {
  display: flex;
  justify-content: center;
  margin-top: 30px;
}

.title {
  text-align: center;
  font-size: 25px;
  font-weight: 600;
  margin-top: 20px;
}

.m-container {
  padding: 40px 30px;
}

.number-container {
  display: flex;
  justify-content: space-around;
}

.user {
  display: flex;
  justify-content: space-between;
  text-align: left;
  padding: 10px 0;
  font: 22px;
  line-height: 1;
}

.min-right {
  padding-right: 200px;
}
</style>