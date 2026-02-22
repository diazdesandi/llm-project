import type { Message, Conversation } from '@/types'
import { defineStore } from 'pinia'
import { questionOllama } from '@/clients/modelClient'
import { nanoid } from 'nanoid'

// TODO: Refactor Store
export const useModelStore = defineStore('model', {
  state: () => ({
    conversations: [] as Conversation[],
    activeConversationId: null as string | null,
    isLoading: false,
    // TODO: Replace for DB data
    modelInfo: {
      model: 'tinyllama',
      description:
        'The TinyLlama project is an open endeavor to train a compact 1.1B Llama model on 3 trillion tokens.',
      fallback: 'tl',
    },
  }),
  getters: {
    getActiveConversation(state): Conversation | null {
      return state.conversations.find((c) => c.id === state.activeConversationId) || null
    },
  },
  actions: {
    selectConversation(id: string) {
      this.activeConversationId = id
    },

    async newMessage(input: string) {
      let convo = this.getActiveConversation

      try {
        // Refactor for DB usage
        const message: Message = {
          id: nanoid(),
          role: 'user',
          content: input,
          conversationId: '',
          createdAt: new Date().toISOString(),
        }

        this.isLoading = true

        if (!convo) {
          this.createConversation(message)
          convo = this.conversations[0]
        } else {
          convo.messages.push(message)
          convo.updatedAt = new Date().toISOString()
        }

        const { response } = await questionOllama({
          model: 'tinyllama',
          prompt: input,
          stream: false,
        })

        this.isLoading = false

        const modelResponse: Message = {
          id: '1212',
          role: 'agent',
          content: response,
          conversationId: '12',
          createdAt: new Date().toISOString(),
        }

        convo.messages.push(modelResponse)
      } catch (error) {
        this.isLoading = false
        throw new Error(`${error}`)
      }
    },

    createConversation(initialMessage: Message) {
      const datetime = new Date().toISOString()

      const newConversation: Conversation = {
        id: '',
        title: 'Convo 1',
        messages: [initialMessage],
        createdAt: datetime,
        updatedAt: datetime,
        model: '',
        isArchived: false,
      }

      this.conversations.unshift(newConversation)
      this.activeConversationId = newConversation.id
    },
  },
})
