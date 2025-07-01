<script setup lang="ts">
import { Plus } from 'lucide-vue-next'
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
  Button,
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui'
import { useModelStore } from '@/stores/useModelStore'

const store = useModelStore()

const createNewConversation = () => {
  // TODO: New chat
  console.log('New conversation with the model')
}
</script>

<template>
  <div class="flex flex-row items-center justify-between px-3">
    <div class="flex items-center space-x-4">
      <Avatar>
        <AvatarImage src="../src/assets/images/tinyllama.png" alt="Image" />
        <AvatarFallback>{{ store.modelInfo.fallback }}</AvatarFallback>
      </Avatar>
      <div>
        <p class="text-sm font-medium leading-none">
          {{ store.modelInfo.model }}
        </p>
        <p class="text-sm text-muted-foreground">
          {{ store.modelInfo.description }}
        </p>
      </div>
    </div>
    <TooltipProvider>
      <Tooltip :delay-duration="200">
        <TooltipTrigger as-child>
          <Button
            variant="outline"
            class="rounded-full p-2.5 flex items-center justify-center"
            @click="createNewConversation"
          >
            <Plus class="w-4 h-4" />
          </Button>
        </TooltipTrigger>
        <TooltipContent :side-offset="10"> New conversation </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  </div>
</template>
