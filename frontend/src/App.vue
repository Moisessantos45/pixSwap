<template>
    <div
        :class="{ dark: darkMode }"
        class="bg-gray-50 dark:bg-gray-900 transition-colors duration-300"
    >
        <!-- Header -->
        <header
            class="bg-white dark:bg-gray-800 shadow-sm border-b border-gray-200 dark:border-gray-700"
        >
            <div class="max-w-7xl mx-auto px-6 py-6">
                <div class="flex justify-between items-center">
                    <div>
                        <h1
                            class="text-3xl font-bold text-gray-900 dark:text-white"
                        >
                            PixSwap
                        </h1>
                        <p class="text-gray-600 dark:text-gray-400 mt-2">
                            Convertir imágenes a diferentes formatos de manera
                            rápida y sencilla.
                        </p>
                    </div>
                    <button
                        @click="toggleDarkMode"
                        class="p-3 rounded-lg bg-gray-100 dark:bg-gray-700 text-gray-600 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-600 transition-colors duration-200"
                    >
                        <component
                            :is="darkMode ? SunIcon : MoonIcon"
                            class="w-5 h-5"
                        />
                    </button>
                </div>
            </div>
        </header>

        <!-- Main Content -->
        <main class="max-w-7xl mx-auto px-6 py-8">
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
                <!-- Left Column - Image Selection -->
                <div class="space-y-6">
                    <ImageSelector
                        @files-selected="handleFilesSelected"
                        :is-processing="isProcessing"
                    />

                    <ImageList
                        :images="selectedImages"
                        @remove-image="removeImage"
                        @clear-images="clearImages"
                    />
                </div>

                <!-- Right Column - Conversion Options -->
                <div class="space-y-6">
                    <ConversionOptions
                        v-model:output-format="outputFormat"
                        v-model:quality="quality"
                        :disabled="isProcessing"
                        v-model:directory="directory"
                    />

                    <ConversionButton
                        :disabled="selectedImages.length === 0"
                        :is-processing="isProcessing"
                        @convert="convertImages"
                    />

                    <ProgressIndicator
                        v-if="isProcessing"
                        :progress="conversionProgress"
                        :is-processing="isProcessing"
                    />
                </div>
            </div>
        </main>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { SunIcon, MoonIcon } from "@heroicons/vue/24/outline";
import ImageSelector from "./components/ImageSelector.vue";
import ImageList from "./components/ImageList.vue";
import ConversionOptions from "./components/ConversionOptions.vue";
import ConversionButton from "./components/ConversionButton.vue";
import ProgressIndicator from "./components/ProgressIndicator.vue";
import type { ImageFile } from "./types";
import { GetFilesAsBase64, ConvertToFormatImg } from "../wailsjs/go/main/App";
import showNotification from "./utils/notification";

const darkMode = ref(false);
const selectedImages = ref<ImageFile[]>([]);
const outputFormat = ref("png");
const quality = ref(60);
const directory = ref("");
const isProcessing = ref(false);
const conversionProgress = ref(0);

onMounted(() => {
    const saved = localStorage.getItem("darkMode");
    darkMode.value = saved ? JSON.parse(saved) : false;
});

const toggleDarkMode = () => {
    darkMode.value = !darkMode.value;
    localStorage.setItem("darkMode", JSON.stringify(darkMode.value));
};

const handleFilesSelected = async (files: string[]) => {
    try {
        const results = await GetFilesAsBase64(files);
        selectedImages.value = [...results];
    } catch (error) {
        console.log("Error handling files:", error);
    }
};

const removeImage = (id: number) => {
    const index = selectedImages.value.filter((img) => img.id != id);
    selectedImages.value = [...index];
    showNotification("Imagen eliminada.", "success");
};

const clearImages = () => {
    selectedImages.value = [];
    showNotification("Todas las imágenes han sido eliminadas.", "success");
};

const convertImages = async () => {
    if (selectedImages.value.length === 0) return;

    isProcessing.value = true;
    conversionProgress.value = 0;

    const filesPaths = selectedImages.value.map((img) => img.path);
    try {
        const result = await ConvertToFormatImg(
            filesPaths,
            outputFormat.value,
            quality.value,
            directory.value,
        );
        showNotification(result, "success");
        selectedImages.value = [];
    } catch (error) {
        showNotification("Error during conversion. Please try again.", "error");
    } finally {
        isProcessing.value = false;
        conversionProgress.value = 100;
    }
};
</script>

<style>
#app {
    scrollbar-width: none;
}
</style>
