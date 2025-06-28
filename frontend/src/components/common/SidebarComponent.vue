<script setup lang="ts">
import { GalleryVerticalEnd } from 'lucide-vue-next'

import {
    Sidebar,
    SidebarContent,
    SidebarGroup,
    SidebarHeader,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
    SidebarMenuSub,
    SidebarMenuSubButton,
    SidebarMenuSubItem,
    type SidebarProps,
} from '@/components/ui/sidebar'

const props = withDefaults(defineProps<SidebarProps>(), {
    variant: 'floating',
})

// This is sample data.
const data = {
    navMain: [
        {
            title: 'Getting Started',
            url: '#',
            items: [
                {
                    title: 'Installation',
                    url: '#',
                },
                {
                    title: 'Project Structure',
                    url: '#',
                },
            ],
        },
        {
            title: 'Building Your Application',
            url: '#',
            items: [
                {
                    title: 'Routing',
                    url: '#',
                },
                {
                    title: 'Data Fetching',
                    url: '#',
                    isActive: true,
                },
            ],
        },
    ],
}
</script>

<template>
    <Sidebar v-bind="props">
        <SidebarHeader>
            <SidebarMenu>
                <SidebarMenuItem>
                    <SidebarMenuButton size="lg" as-child>
                        <a href="#">
                            <div
                                class="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground">
                                <GalleryVerticalEnd class="size-4" />
                            </div>
                            <div class="flex flex-col gap-0.5 leading-none">
                                <span class="font-semibold">Documentation</span>
                                <span class="">v1.0.0</span>
                            </div>
                        </a>
                    </SidebarMenuButton>
                </SidebarMenuItem>
            </SidebarMenu>
        </SidebarHeader>
        <SidebarContent>
            <SidebarGroup>
                <SidebarMenu class="gap-2">
                    <SidebarMenuItem v-for="item in data.navMain" :key="item.title">
                        <SidebarMenuButton as-child>
                            <a :href="item.url" class="font-medium">
                                {{ item.title }}
                            </a>
                        </SidebarMenuButton>
                        <SidebarMenuSub v-if="item.items.length">
                            <SidebarMenuSubItem v-for="childItem in item.items" :key="childItem.title">
                                <SidebarMenuSubButton as-child :is-active="childItem.isActive">
                                    <a :href="childItem.url">{{ childItem.title }}</a>
                                </SidebarMenuSubButton>
                            </SidebarMenuSubItem>
                        </SidebarMenuSub>
                    </SidebarMenuItem>
                </SidebarMenu>
            </SidebarGroup>
        </SidebarContent>
    </Sidebar>
</template>
