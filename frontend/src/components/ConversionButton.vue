<template>
  <div class="card p-6">
    <button
      @click="$emit('convert')"
      :disabled="disabled || isProcessing"
      class="w-full btn-primary flex items-center justify-center space-x-3 text-lg py-4"
    >
      <component
        :is="isProcessing ? 'div' : ArrowDownTrayIcon"
        :class="isProcessing ? 'animate-spin w-5 h-5 border-2 border-white border-t-transparent rounded-full' : 'w-6 h-6'"
      />
      <span>
        {{ isProcessing ? 'Converting Images...' : 'Convert Images' }}
      </span>
    </button>
    
    <p class="text-sm text-gray-500 dark:text-gray-400 text-center mt-3">
      {{ disabled ? 'Select images to start conversion' : 'All images will be processed with the selected options' }}
    </p>
  </div>
</template>

<script setup lang="ts">
import { ArrowDownTrayIcon } from '@heroicons/vue/24/outline'

defineEmits<{
  convert: []
}>()

defineProps<{
  disabled: boolean
  isProcessing: boolean
}>()
</script>