import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin/index.vue'
import InterfaceInfo from '../views/Admin/InterfaceInfo/index.vue'
import Dashboard from '../views/Admin/Dashboard/index.vue'
import User from '../views/User/index.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: Login
    },
    {
      path: '/admin',
      name: 'admin',
      component: Admin,
      meta: {
        title: "个人中心"
      },
      children: [
        {
          path: "",
          name: "dashboard",
          component: Dashboard,
          meta: {
            title: "首页"
          },
        },
        {
          path: "interface",
          name: "interface",
          component: InterfaceInfo,
          meta: {
            title: "接口管理"
          },
        }
      ]
    },
    {
      path: '/user',
      name: 'user',
      component: User,
      meta: {
        title: "个人中心"
      },
      children: [
        {
          path: "",
          name: "dashboard",
          component: Dashboard,
          meta: {
            title: "首页"
          },
        },
        {
          path: "interface",
          name: "interface",
          component: InterfaceInfo,
          meta: {
            title: "接口管理"
          },
        }
      ]
    }
  ]
})

export default router
