import { createRouter, createWebHistory } from 'vue-router'
import HomeLayout from '@/layouts/HomeLayout.vue'
import ChatPage from '@/pages/ChatPage.vue';
import SettingsPage from '@/pages/SettingsPage.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
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
    }
  ],
})

export default router
