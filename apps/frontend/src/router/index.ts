import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: "root",
      redirect: { name: "Dashboard" }
    },
    {
      path: '/',
      component: () => import ('../layouts/App.layout.vue'),
      children: [
        {
          path: 'dashboard',
          name: 'Dashboard',
          component: () => import ('../views/Dashboard.vue')
        }
      ]
    },
    {
      path: '/',
      component: () => import('../layouts/Login.layout.vue'),
      children: [
        {
          path: '/login',
          name: 'Login',
          component: () => import('../views/Login.vue')
        }
      ]
    }
  ]
})

export default router
