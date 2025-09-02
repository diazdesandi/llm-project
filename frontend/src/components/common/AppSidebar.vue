<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router';
import { Command, Settings, Inbox, } from 'lucide-vue-next'
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarHeader,
  SidebarInput,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  type SidebarProps,
} from '../ui'
import { useModelStore } from '@/stores/useModelStore';
import NavUser from './NavUser.vue';
import DarkModeButton from './DarkModeButton.vue';

const props = withDefaults(defineProps<SidebarProps>(), {
  collapsible: 'icon',
})

const router = useRouter()

// This is sample data
const data = {
  user: {
    name: 'shadcn',
    email: 'm@example.com',
    avatar: '/avatars/shadcn.jpg',
  },
  navMain: [
    {
      title: 'Conversations',
      url: '/',
      icon: Inbox,
      isActive: true,
    },
    {
      title: 'Settings',
      url: '/settings',
      icon: Settings,
      isActive: false,
    },
  ],
}

const activeItem = ref(data.navMain[0])
const store = useModelStore()
// const { setOpen } = useSidebar()

const navigateTo = (url: string) => {
  router.push(url)
}

</script>

<template>
  <Sidebar class="overflow-hidden [&>[data-sidebar=sidebar]]:flex-row" v-bind="props">
    <Sidebar collapsible="none" class="!w-[calc(var(--sidebar-width-icon)_+_1px)] border-r">
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" as-child class="md:h-8 md:p-0">
              <a href="#">
                <div
                  class="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground">
                  <Command class="size-4" />
                </div>
                <div class="grid flex-1 text-left text-sm leading-tight">
                  <span class="truncate font-semibold">Acme Inc</span>
                  <span class="truncate text-xs">Enterprise</span>
                </div>
              </a>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupContent class="px-1.5 md:px-0">
            <SidebarMenu>
              <SidebarMenuItem v-for="item in data.navMain" :key="item.title">
                <SidebarMenuButton :is-active="$route.path === item.url" class="px-2.5 md:px-2"
                  @click="navigateTo(item.url)">
                  <component :is="item.icon" />
                  <span>{{ item.title }}</span>
                </SidebarMenuButton>
              </SidebarMenuItem>
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <NavUser :user="data.user" />
      </SidebarFooter>
    </Sidebar>

    <Sidebar collapsible="none" class="hidden flex-1 md:flex">
      <SidebarHeader class="gap-3.5 border-b p-4">
        <div class="flex w-full items-center justify-between">
          <div class="text-base font-medium text-foreground">
            {{ activeItem.title }}
          </div>
          <DarkModeButton />
        </div>
        <SidebarInput placeholder="Type to search..." />
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup class="px-0">
          <SidebarGroupContent>
            <a v-for="conversation in store.conversations" :key="conversation.id" href="#"
              class="flex flex-col items-start gap-2 whitespace-nowrap border-b p-4 text-sm leading-tight last:border-b-0 hover:bg-sidebar-accent hover:text-sidebar-accent-foreground">
              <span class="font-medium">{{ conversation.title }}</span>
              <span class="line-clamp-2 w-[260px] whitespace-break-spaces text-xs">
                {{ conversation.messages[0].content }}
              </span>
            </a>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  </Sidebar>
</template>
