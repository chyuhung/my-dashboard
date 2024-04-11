import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from '../views/Index'
import Login from '../views/Login'

// 页面组件
import Instance from '../components/Instance'
import Volume from '../components/Volume'

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
    component: Index,
    children: [
      { path: 'instance', component: Instance },
      { path: 'volume', component: Volume }
    ]
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
  } else if (token && to.path == '/login') {
    // 如果携带token访问登录页面则跳转首页
    next('/index')
  } else {
    next()
  }
})

export default router
