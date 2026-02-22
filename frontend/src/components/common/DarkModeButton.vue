<script setup lang="ts">
import { useColorMode } from '@vueuse/core'
import { Icon } from '@iconify/vue'
import { Button } from '@/components/ui'

const mode = useColorMode({
  attribute: 'class',
  disableTransition: true,
  selector: 'html',
  storageKey: 'vueuse-color-scheme',
})

const toggleMode = () => {
  mode.value = mode.value === 'dark' ? 'light' : 'dark'
}

</script>
<template>
  <Button variant="ghost" @click="toggleMode">
    <Icon
      icon="radix-icons:moon"
      class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
      @click="mode = 'dark'"
    />
    <Icon
      icon="radix-icons:sun"
      class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
      @click="mode = 'light'"
    />
  </Button>
</template>
