import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from '../views/Index'
import Login from '../views/Login'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/index',
    name: 'index',
    component: Index
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach(async (to, from, next) => {
  const token = window.sessionStorage.getItem('token')
  // 如果不携带token并且访问非登录页面则强制跳转到登录页面
  if (!token && to.path !== '/login') {
    next('/login')
  } else {
    next()
  }
})

export default router
