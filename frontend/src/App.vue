<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router'
import {ref, watch} from 'vue';
import headerVue from "./components/header.vue"
import { menuList } from "./configs/baseConf.js"
const router = useRouter()

const currentRoutePath = ref(router.currentRoute.value.path);
let isShowNormal = ref(!location.pathname.includes('login'))

watch(() => router.currentRoute.value.path, (newPath, oldPath) => {
  currentRoutePath.value = newPath;
  isShowNormal.value = !currentRoutePath.value.includes('login')
}, { immediate: true });

const menuClk = (item: any) => {
  console.log(item);

  router.push({
    path: item.index
  })
}
</script>

<template>
  <section class="all-container">
    <template v-show="isShowNormal">
      <headerVue></headerVue>
    </template>
    <section class="m-container">
      <!-- 菜单 -->
      <section v-show="isShowNormal" class="m-menu">
        <el-menu active-text-color="#ffd04b" background-color="#468cdc" class="el-menu-vertical-demo" text-color="#fff">
          <template v-for="item in menuList" :key="item.name">
            <el-sub-menu v-if="item.children && item.children.length > 0">
              <template #title>
                <span>{{ item.nameZh }}</span>
              </template>

              <el-menu-item v-for="itm in item.children" :key="itm.name" :index="itm.path" @click="menuClk">
                <span>{{ itm.nameZh }}</span>
              </el-menu-item>
            </el-sub-menu>
            <el-menu-item v-else :key="item.name" :index="item.path" @click="menuClk">
              <span>{{ item.nameZh }}</span>
            </el-menu-item>
          </template>
        </el-menu>
      </section>
      <section class="routerview-container">
        <RouterView />
      </section>
    </section>
  </section>
</template> 

<style scoped>
.routerview-container {
  width: 100%;
  height: 94vh;
  overflow: auto;
}

.m-container {
  display: flex;

}

.m-menu {
  width: 11%;
  height: 94vh;
}

.m-menu span {
  font-size: 20px;
}

.el-menu {
  height: 94vh;
}
</style>
