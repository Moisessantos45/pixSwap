<template>
    <div v-if="images.length > 0" class="card p-6">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">
            Selected Images ({{ images.length }})
        </h3>

        <button
            @click="$emit('clear-images')"
            class="mb-4 px-3 py-2 text-sm text-gray-700 bg-gray-200 dark:bg-gray-600 rounded-lg hover:bg-gray-300 dark:hover:bg-gray-700 transition-colors duration-200 dark:text-gray-300"
            title="Clear all images"
        >
            Clear All
        </button>

        <div class="space-y-3 max-h-96 overflow-y-auto">
            <div
                v-for="image in images"
                :key="image.id"
                class="flex items-center space-x-4 p-3 bg-gray-50 dark:bg-gray-700/50 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors duration-200 animate-slide-up"
            >
                <img
                    :src="image.base64"
                    :alt="image.path"
                    class="w-12 h-12 object-cover rounded-lg flex-shrink-0"
                />

                <div class="flex-1 min-w-0">
                    <p
                        class="text-sm font-medium text-gray-900 dark:text-white truncate"
                    >
                        {{ image.path.split("/").pop() }}
                    </p>
                    <p class="text-xs text-gray-500 dark:text-gray-400">
                        {{ image.mimeType }}
                    </p>
                </div>

                <button
                    @click="$emit('remove-image', image.id)"
                    class="flex-shrink-0 p-2 text-gray-400 hover:text-red-500 transition-colors duration-200"
                    title="Remove image"
                >
                    <XMarkIcon class="w-4 h-4" />
                </button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { XMarkIcon } from "@heroicons/vue/24/outline";
import type { ImageFile } from "../types";

defineEmits<{
    "remove-image": [id: number];
    "clear-images": [];
}>();

defineProps<{
    images: ImageFile[];
}>();
</script>
