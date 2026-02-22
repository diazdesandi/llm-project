<script setup lang="ts">
import { ChatCard, ChatHeader, InputArea } from '@/components/common'
import { Card } from '@/components/ui'
</script>

<template>
  <div class="h-full p-6">
    <Card class="h-full flex flex-col">
      <ChatHeader />
      <ChatCard class="flex-1 min-h-0" />
      <InputArea />
    </Card>
  </div>
</template>