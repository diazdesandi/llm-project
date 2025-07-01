<script setup lang="ts">
import { ref, computed } from 'vue'
import { useModelStore } from '@/stores/useModelStore'
import { Send } from 'lucide-vue-next'
import { Button, Input } from '@/components/ui'

const store = useModelStore()

const input = ref('')
const inputLength = computed(() => input.value.trim().length)

const onSubmit = () => {
  if (inputLength.value === 0) return
  store.newMessage(input.value)
  input.value = ''
}
</script>
<template>
  <div class="p-4 rounded-lg">
    <form class="flex w-full items-center space-x-2" @submit.prevent="onSubmit">
      <Input v-model="input" placeholder="Type a message..." class="flex-1 shadow-0" />
      <Button
        class="p-2.5 flex items-center justify-center"
        :disabled="inputLength === 0"
        type="submit"
      >
        <Send class="w-4 h-4" />
        <span class="sr-only">Send</span>
      </Button>
    </form>
  </div>
</template>
