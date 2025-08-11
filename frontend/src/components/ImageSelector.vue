<template>
    <div class="card p-6">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">
            Select Images
        </h2>

        <div :class="[
            'border-2 border-dashed rounded-xl p-8 text-center transition-all duration-300',
            isProcessing
                ? 'opacity-50 pointer-events-none'
                : 'hover:border-primary-400 hover:bg-gray-50 dark:hover:bg-gray-700/50',
        ]">
            <div class="space-y-4">
                <div
                    class="mx-auto w-16 h-16 bg-primary-100 dark:bg-primary-900/30 rounded-full flex items-center justify-center">
                    <PhotoIcon class="w-8 h-8 text-primary-600 dark:text-primary-400" />
                </div>

                <div>
                    <p class="text-gray-600 dark:text-gray-400 mt-1">
                        Support for PNG, JPG, WEBP, GIF, BMP formats
                    </p>
                </div>

                <button @click="handleFileChange" class="btn-primary inline-flex items-center space-x-2">
                    <PlusIcon class="w-5 h-5" />
                    <span>Seleccionar imagenes</span>
                </button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { PhotoIcon, PlusIcon } from "@heroicons/vue/24/outline";
import {
    OpenFilesDialog
} from "../../wailsjs/go/main/App";

const emit = defineEmits<{
    "files-selected": [string[]];
    "file-selector-opened": [string];
}>();

defineProps<{
    isProcessing: boolean;
}>();

const handleFileChange = async () => {
    const files = await OpenFilesDialog();
    console.log("Selected files:", files);
    if (!files || files.length === 0) {
        emit("files-selected", []);
        return;
    }
    emit("files-selected", files);
};


</script>
