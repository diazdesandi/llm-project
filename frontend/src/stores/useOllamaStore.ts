import type { Message } from '@/types'
import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { questionOllama } from '../clients/llmClient'

export const useModelStore = defineStore('model', {
  state: () => ({
    messages: [] as Message[],
    message: {} as Message,
    isLoading: false,
    modelInfo: {
      model: "tinyllama",
      description: "The TinyLlama project is an open endeavor to train a compact 1.1B Llama model on 3 trillion tokens."
    }
  }),

  actions: {
    async newMessage(input: string) {
      try {
        this.isLoading = true

        const { response } = await questionOllama({
          model: 'tinyllama',
          prompt: input,
          stream: false,
        })

        this.isLoading = false

        this.messages.push({
          role: 'agent',
          content: response,
        })
      } catch (error) { }
    },
  },
})