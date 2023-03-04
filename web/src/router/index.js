import { createRouter, createWebHistory } from 'vue-router'
import AppIndex from '@/components/Index/Index'
import PostData from '@/components/Index/PostData'
import UserData from '@/components/Index/UserData'

const routerHistory = createWebHistory()

const constantRoutes = [
  {
    path: '/',
    name: 'index',
    component: AppIndex
  },
  {
    path: '/postdata',
    name: 'postdata',
    component: PostData
  },
  {
    path: '/userdata',
    name: 'userdata',
    component: UserData
  }
]


let router = createRouter({
  history: routerHistory,
  routes: constantRoutes,
})

export default router