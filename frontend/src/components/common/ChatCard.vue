<script setup lang="ts">
import { LoaderCircle } from 'lucide-vue-next'
import { cn } from '@/lib/utils'
import { useModelStore } from '@/stores/useModelStore'

const store = useModelStore()
</script>
<template>
  <div
    ref="messagesContainer"
    class="flex-1 p-4 space-y-4 overflow-y-auto"
    style="max-height: none"
  >
    <div
      v-for="(message, index) in store.messages"
      :key="index"
      :class="
        cn(
          // Cambia aquí el max-w
          'flex w-max max-w-xl flex-col gap-2 rounded-lg px-3 py-2 text-sm break-words',
          message.role === 'user' ? 'ml-auto bg-primary text-primary-foreground' : 'bg-muted',
        )
      "
    >
      <span class="whitespace-pre-wrap">{{ message.content }}</span>
    </div>
    <div
      v-if="store.isLoading"
      class="flex w-max max-w-xl items-center gap-2 rounded-lg px-3 py-2 text-sm bg-muted"
    >
      <LoaderCircle class="animate-spin w-4 h-4" />
      <span class="text-muted-foreground">Thinking...</span>
    </div>
  </div>
</template>
