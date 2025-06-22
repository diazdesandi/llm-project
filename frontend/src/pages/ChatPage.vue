<script setup lang="ts">
import { Send, LoaderCircle } from 'lucide-vue-next'
import { computed, ref } from 'vue'
import { cn } from '@/lib/utils'
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { questionOllama } from '@/clients/llmClient'
import type { Message } from '@/types'


// TODO: Breakdown the entire page into components

const input = ref('')
const inputLength = computed(() => input.value.trim().length)
const model: string = 'tinyllama'
const description: string = "The TinyLlama project is an open endeavor to train a compact 1.1B Llama model on 3 trillion tokens."

const messages = ref<Message[]>([])

const isLoading = ref(false)

const ollamaRequest = async (input: string) => {
  isLoading.value = true;


  const resp = await questionOllama(
    {
      model: "tinyllama",
      prompt: input,
      stream: false,
    })

  isLoading.value = false;

  messages.value.push({
    role: 'agent',
    content: resp.response
  })
}
</script>

<template>
  <Card>
    <CardHeader class="flex flex-row items-center justify-between">
      <div class="flex items-center space-x-4">
        <Avatar>
          <AvatarImage src="../src/assets/images/tinyllama.png" alt="Image" />
          <AvatarFallback>tl</AvatarFallback>
        </Avatar>
        <div>
          <p class="font-medium leading-none">
            {{ model }}
          </p>
          <p class="text-sm text-muted-foreground">
            {{ description }}
          </p>
        </div>
      </div>
    </CardHeader>
    <CardContent>
      <div class="space-y-4">
        <div v-for="(message, index) in messages" :key="index" :class="cn(
          'flex w-max max-w-[75%] flex-col gap-2 rounded-lg px-3 py-2 text-sm',
          message.role === 'user' ? 'ml-auto bg-primary text-primary-foreground' : 'bg-muted',
        )">
          {{ message.content }}
        </div>
        <div v-if="isLoading" class="flex w-max max-w-[75%] items-center gap-2 rounded-lg px-3 py-2 text-sm bg-muted">
          <LoaderCircle class="animate-spin w-4 h-4" />
          <span class="text-muted-foreground">Thinking...</span>
        </div>
      </div>
    </CardContent>
    <CardFooter>
      <form class="flex w-full items-center space-x-2" @submit.prevent="() => {
        if (inputLength === 0) return
        messages.push({
          role: 'user',
          content: input,
        })
        input = ''
      }">
        <Input v-model="input" placeholder="Type a message..." class="flex-1" />
        <Button class="p-2.5 flex items-center justify-center" :disabled="inputLength === 0"
          @click="ollamaRequest(input)">
          <Send class="w-4 h-4" />
          <span class="sr-only">Send</span>
        </Button>
      </form>
    </CardFooter>
  </Card>
</template>
