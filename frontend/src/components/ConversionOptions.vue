<template>
  <div class="card p-6">
    <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-6">Conversion Options</h3>

    <div class="space-y-6">
      <!-- Output Format -->
      <div>
        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
          Output Format
        </label>
        <select :value="outputFormat" @input="$emit('update:output-format', ($event.target as HTMLSelectElement).value)"
          :disabled="disabled" class="input-field">
          <option value="png">PNG</option>
          <option value="jpeg">JPEG</option>
          <option value="gif">GIF</option>
        </select>
      </div>

      <!-- Quality Slider -->
      <div v-if="showQualitySlider">
        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
          Quality: {{ quality }}%
        </label>
        <div class="flex items-center space-x-4">
          <span class="text-sm text-gray-500 dark:text-gray-400">1%</span>
          <input type="range" min="1" max="100" :value="quality"
            @input="$emit('update:quality', parseInt(($event.target as HTMLInputElement).value))" :disabled="disabled"
            class="flex-1 h-2 bg-gray-200 dark:bg-gray-700 rounded-lg appearance-none cursor-pointer" />
          <span class="text-sm text-gray-500 dark:text-gray-400">100%</span>
        </div>
        <div class="mt-2">
          <input type="number" min="1" max="100" :value="quality"
            @input="$emit('update:quality', parseInt(($event.target as HTMLInputElement).value))" :disabled="disabled"
            class="w-20 px-2 py-1 text-sm border border-gray-300 dark:border-gray-600 rounded bg-white dark:bg-gray-800 text-gray-900 dark:text-gray-100" />
          <span class="text-sm text-gray-500 dark:text-gray-400 ml-1">%</span>
        </div>
      </div>

      <!-- Destination Directory -->
      <div>
        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
          Output Folder
        </label>
        <div class="flex items-center gap-3">
          <button type="button" class="btn-secondary" @click="openFileSelector" :disabled="disabled">
            Seleccionar carpeta
          </button>
          <span class="text-sm text-gray-600 dark:text-gray-300 truncate" :title="directory">
            {{ directory ? directory : 'Ninguna carpeta seleccionada' }}
          </span>
        </div>
      </div>

      <!-- Format Info -->
      <div class="bg-blue-50 dark:bg-blue-900/20 p-4 rounded-lg">
        <h4 class="text-sm font-medium text-blue-900 dark:text-blue-200 mb-2">Format Information</h4>
        <p class="text-xs text-blue-800 dark:text-blue-300">
          {{ formatInfo }}
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, toRefs } from 'vue'
import { OpenDirectoryDialog } from "../../wailsjs/go/main/App";

const props = defineProps<{
  outputFormat: string
  quality: number
  disabled: boolean
  directory: string
}>()

const { outputFormat, quality, disabled, directory } = toRefs(props)

const emit = defineEmits<{
  'update:output-format': [value: string]
  'update:quality': [value: number]
  'update:directory': [value: string]
  'file-selector-opened': [directory: string]
}>()

const showQualitySlider = computed(() => {
  return ['jpeg', 'webp'].includes(outputFormat.value)
})

const formatInfo = computed(() => {
  const formats: Record<string, string> = {
    png: 'PNG supports transparency and is best for graphics with sharp edges.',
    jpeg: 'JPEG is ideal for photos with adjustable quality and smaller file sizes.',
    // webp: 'WebP offers superior compression with quality options, best for web use.',
    gif: 'GIF supports animation and transparency, limited to 256 colors.',
   // bmp: 'BMP provides uncompressed images with larger file sizes but perfect quality.'
  }
  return formats[outputFormat.value] || 'Select a format to see information.'
})

const openFileSelector = async () => {
  const dir = await OpenDirectoryDialog();
  console.log("Selected directory:", dir);
  if (dir) {
    emit('update:directory', dir);
    emit('file-selector-opened', dir);
  }
};
</script>