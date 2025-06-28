<script setup lang="ts">
import {
    Card,
    CardContent,
    CardFooter,
    CardHeader,
} from '@/components/ui/card'
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { Send, LoaderCircle } from 'lucide-vue-next'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { cn } from '@/lib/utils';
import { computed, ref } from 'vue';
import { useModelStore } from '@/stores/useOllamaStore';


const store = useModelStore()

const input = ref('')
const inputLength = computed(() => input.value.trim().length)

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
                        {{ store.modelInfo.model }}
                    </p>
                    <p class="text-sm text-muted-foreground">
                        {{ store.modelInfo.description }}
                    </p>
                </div>
            </div>
        </CardHeader>
        <CardContent>
            <div class="space-y-4">
                <div v-for="(message, index) in store.messages" :key="index" :class="cn(
                    'flex w-max max-w-[75%] flex-col gap-2 rounded-lg px-3 py-2 text-sm',
                    message.role === 'user' ? 'ml-auto bg-primary text-primary-foreground' : 'bg-muted',
                )">
                    {{ message.content }}
                </div>
                <div v-if="store.isLoading"
                    class="flex w-max max-w-[75%] items-center gap-2 rounded-lg px-3 py-2 text-sm bg-muted">
                    <LoaderCircle class="animate-spin w-4 h-4" />
                    <span class="text-muted-foreground">Thinking...</span>
                </div>
            </div>
        </CardContent>
        <CardFooter>
            <form class="flex w-full items-center space-x-2" @submit.prevent="() => {
                if (inputLength === 0) return
                store.messages.push({
                    role: 'user',
                    content: input,
                })
                input = ''
            }">
                <Input v-model="input" placeholder="Type a message..." class="flex-1" />
                <Button class="p-2.5 flex items-center justify-center" :disabled="inputLength === 0"
                    @click="store.newMessage(input)">
                    <Send class="w-4 h-4" />
                    <span class="sr-only">Send</span>
                </Button>
            </form>
        </CardFooter>
    </Card>
</template>