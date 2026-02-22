<script setup lang="ts">
import { AppSidebar } from '@/components/common'
import { SidebarInset, SidebarProvider } from '@/components/ui'
</script>

<template>
    <SidebarProvider :style="{ '--sidebar-width': '19rem' }">
        <AppSidebar />
        <SidebarInset class="rounded-lg">
            <div class="flex-1 overflow-auto">
                <router-view />
            </div>
        </SidebarInset>
    </SidebarProvider>
</template>

