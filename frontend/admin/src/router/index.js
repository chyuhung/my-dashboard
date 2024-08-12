import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from '../views/Index'
import Login from '../views/Login'
import Servers from '../views/Servers'

// 使用 VueRouter
Vue.use(VueRouter)

// 定义路由
const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/index',
    name: 'index',
    component: Index,
    meta: { requiresAuth: true }
  },
  {
    path: '/servers',
    name: 'servers',
    component: Servers,
    meta: { requiresAuth: true }
  }
]

// 创建路由实例
const router = new VueRouter({
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem('token')

  // 允许访问登录页面
  if (to.path === '/login') return next()

  // 如果需要认证且未登录，重定向到登录页
  if (to.matched.some((record) => record.meta.requiresAuth) && !token) {
    return next('/login')
  }

  next()
})

export default router
