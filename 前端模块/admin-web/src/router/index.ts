import { createRouter, createWebHistory } from 'vue-router'
import BasicLayout from '../layouts/BasicLayout.vue'
import LoginView from '../views/LoginView.vue'
import DashboardView from '../views/dashboard/DashboardView.vue'
import BannersView from '../views/content/BannersView.vue'
import NoticesView from '../views/content/NoticesView.vue'
import PackagesView from '../views/content/PackagesView.vue'
import OrdersView from '../views/orders/OrdersView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      component: LoginView
    },
    {
      path: '/',
      component: BasicLayout,
      redirect: '/dashboard',
      children: [
        { path: 'dashboard', component: DashboardView },
        { path: 'banners', component: BannersView },
        { path: 'notices', component: NoticesView },
        { path: 'packages', component: PackagesView },
        { path: 'orders', component: OrdersView }
      ]
    }
  ]
})

export default router
