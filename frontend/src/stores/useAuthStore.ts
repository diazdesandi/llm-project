import { defineStore } from 'pinia'

// Auth Store Placeholder
export const useAuthStore = defineStore('authStore', {

    state: () => ({
        user: {},
        token: 'JWT Token Placeholder'
    }),
    getters: {},
    actions: {
        async login() {
            try {

            } catch (error) {
                
            }
        },
        async signup() {
            // Signup
        }
    }
})