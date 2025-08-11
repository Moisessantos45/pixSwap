<template>
  <div class="fixed top-4 right-4 z-50 space-y-2">
    <div
      v-for="toast in toasts"
      :key="toast.id"
      :class="[
        'p-4 rounded-lg shadow-lg max-w-sm transform transition-all duration-300 animate-slide-up',
        toastClasses[toast.type]
      ]"
    >
      <div class="flex items-start space-x-3">
        <component :is="toastIcons[toast.type]" class="w-5 h-5 flex-shrink-0 mt-0.5" />
        <div class="flex-1">
          <p class="text-sm font-medium">{{ toast.message }}</p>
        </div>
        <button
          @click="$emit('remove-toast', toast.id)"
          class="flex-shrink-0 text-gray-400 hover:text-gray-600 transition-colors duration-200"
        >
          <XMarkIcon class="w-4 h-4" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { 
  CheckCircleIcon, 
  ExclamationTriangleIcon, 
  XCircleIcon, 
  InformationCircleIcon,
  XMarkIcon 
} from '@heroicons/vue/24/outline'
import type { Toast } from '../types'

defineEmits<{
  'remove-toast': [id: number]
}>()

defineProps<{
  toasts: Toast[]
}>()

const toastClasses = {
  success: 'bg-green-50 border border-green-200 text-green-800',
  error: 'bg-red-50 border border-red-200 text-red-800',
  warning: 'bg-yellow-50 border border-yellow-200 text-yellow-800',
  info: 'bg-blue-50 border border-blue-200 text-blue-800'
}

const toastIcons = {
  success: CheckCircleIcon,
  error: XCircleIcon,
  warning: ExclamationTriangleIcon,
  info: InformationCircleIcon
}
</script>