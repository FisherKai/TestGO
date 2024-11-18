import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import { menuList } from "../configs/baseConf"

const routerList = [];

menuList.forEach(item => {
  if (item.children && item.children.length > 0) {
    item.children.forEach(itm => {
      routerList.push({
        path: itm.path,
        name: itm.name,
        component: itm.component
      })
    });
  } else {
    routerList.push({
      path: item.path,
      name: item.name,
      component: item.component
    })
  }
});

let router: any;
router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [{
    path: '/',
    name: 'home',
    component: HomeView
  }, {
    path: '/login',
    name: 'login',
    component: LoginView
  }, ...routerList]
});

router.beforeEach((to, from, next) => {
  console.log('路由从', from.path, '变化到', to.path);
  if(from.path.includes("login") && to.path === "/"){
    to.meta.isLogin = true
  }
  next();
});

export default router