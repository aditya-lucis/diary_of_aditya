import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from '@/views/LandingPage.vue'
import LoginPage from '@/views/LoginPage.vue'
import RegisterPage from '@/views/RegisterPage.vue'
import DashboardLayout from '@/layouts/DashboardLayout.vue'
import DashboardHome from '@/views/DashboardHome.vue'

const routes = [
  {
    path: '/',
    name: 'Landing',
    component: LandingPage
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginPage
  },
  {
    path: '/register',
    name: 'Register',
    component: RegisterPage
  },
  {
    path: '/dashboard',
    component: DashboardLayout,
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: DashboardHome
      }
      // Add more dashboard routes here
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router