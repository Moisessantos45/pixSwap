<template>
  <div class="card p-6">
    <div class="flex items-center justify-between mb-2">
      <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
        {{ isProcessing ? 'Converting...' : 'Conversion Complete' }}
      </h3>
      <span class="text-sm font-medium text-gray-600 dark:text-gray-400">
        {{ Math.round(progress) }}%
      </span>
    </div>
    
    <div class="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-3 overflow-hidden">
      <div
        class="h-full bg-gradient-to-r from-primary-500 to-primary-600 rounded-full transition-all duration-500 ease-out"
        :style="{ width: `${progress}%` }"
      ></div>
    </div>
    
    <div class="flex items-center justify-center mt-4">
      <div v-if="isProcessing" class="flex items-center space-x-2 text-primary-600 dark:text-primary-400">
        <div class="animate-spin w-4 h-4 border-2 border-primary-600 dark:border-primary-400 border-t-transparent rounded-full"></div>
        <span class="text-sm font-medium">Processing images...</span>
      </div>
      <div v-else class="flex items-center space-x-2 text-green-600 dark:text-green-400">
        <CheckCircleIcon class="w-5 h-5" />
        <span class="text-sm font-medium">All images converted successfully!</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { CheckCircleIcon } from '@heroicons/vue/24/outline'

defineProps<{
  progress: number
  isProcessing: boolean
}>()
</script>