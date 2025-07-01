import type { Message } from '@/types'
import { defineStore } from 'pinia'
import { questionOllama } from '@/clients/modelClient'

export const useModelStore = defineStore('model', {
  state: () => ({
    messages: [] as Message[],
    // message: {} as Message,
    isLoading: false,
    modelInfo: {
      model: "tinyllama",
      description: "The TinyLlama project is an open endeavor to train a compact 1.1B Llama model on 3 trillion tokens.",
      fallback: 'tl'
    }
  }),

  actions: {
    async newMessage(input: string) {
      try {

        this.messages.push({
          role: 'user',
          content: input,
        })

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
      } catch (error) {
        this.isLoading = false
        throw new Error(`${error}`)
      }
    },
  },
})