import { createRouter, createWebHistory } from 'vue-router'
import AuthLayout from '@/layouts/AuthLayout.vue'
import HomeLayout from '@/layouts/HomeLayout.vue'
import ChatPage from '@/pages/ChatPage.vue';
import { LoginForm, SignupForm } from "@/components/forms";
import SettingsPage from '@/pages/SettingsPage.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/auth',
      component: AuthLayout,
      children: [
        {
          path: 'login',
          name: 'login',
          component: LoginForm
        },
        {
          path: 'signup',
          name: 'signup',
          component: SignupForm
        }
      ]
    },
    {
      path: '/',
      component: HomeLayout,
      children: [
        {
          path: '',
          name: 'chat',
          component: ChatPage
        },
        {
          path: 'settings',
          name: 'settings',
          component: SettingsPage
        }
      ]
    },
  ],
})

export default router
