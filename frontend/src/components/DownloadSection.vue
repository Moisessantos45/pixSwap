<template>
  <div class="card p-6">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
        Converted Images ({{ convertedImages.length }})
      </h3>
      <button
        @click="$emit('download-all')"
        class="btn-secondary flex items-center space-x-2"
      >
        <ArchiveBoxIcon class="w-4 h-4" />
        <span>Download ZIP</span>
      </button>
    </div>
    
    <div class="space-y-3 max-h-80 overflow-y-auto">
      <div
        v-for="image in convertedImages"
        :key="image.id"
        class="flex items-center space-x-4 p-3 bg-gray-50 dark:bg-gray-700/50 rounded-lg animate-fade-in"
      >
        <img
          :src="image.url"
          :alt="image.name"
          class="w-12 h-12 object-cover rounded-lg flex-shrink-0"
        />
        
        <div class="flex-1 min-w-0">
          <p class="text-sm font-medium text-gray-900 dark:text-white truncate">
            {{ image.name }}
          </p>
          <p class="text-xs text-gray-500 dark:text-gray-400">
            {{ formatFileSize(image.size) }} • {{ image.format.toUpperCase() }}
          </p>
        </div>
        
        <div class="flex items-center space-x-2">
          <span class="text-xs px-2 py-1 bg-green-100 dark:bg-green-900/30 text-green-800 dark:text-green-200 rounded-full">
            Converted
          </span>
          <button
            @click="downloadImage(image)"
            class="p-2 text-gray-400 hover:text-primary-500 transition-colors duration-200"
            title="Download image"
          >
            <ArrowDownTrayIcon class="w-4 h-4" />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ArrowDownTrayIcon, ArchiveBoxIcon } from '@heroicons/vue/24/outline'
import type { ConvertedImage } from '../types'

defineEmits<{
  'download-all': []
}>()

defineProps<{
  convertedImages: ConvertedImage[]
}>()

const downloadImage = (image: ConvertedImage) => {
  const link = document.createElement('a')
  link.href = image.url
  link.download = image.name
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}
</script>