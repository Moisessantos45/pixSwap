# PixSwap

Convierte imágenes por lotes a múltiples formatos de forma rápida, con vista en tiempo real, control de calidad y modo oscuro.

<p align="left">
  <img alt="Wails" src="https://img.shields.io/badge/Wails-v2-24A1C1?logo=go&logoColor=white" />
  <img alt="Vue" src="https://img.shields.io/badge/Vue-3-42b883?logo=vue.js&logoColor=white" />
  <img alt="Vite" src="https://img.shields.io/badge/Vite-7-646CFF?logo=vite&logoColor=white" />
  <img alt="TailwindCSS" src="https://img.shields.io/badge/TailwindCSS-4-06B6D4?logo=tailwindcss&logoColor=white" />
  <img alt="Go" src="https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white" />
</p>

## ✨ Características
- Conversión de imágenes a PNG, JPEG, WEBP, GIF y BMP
- Control de calidad para JPEG/WEBP
- Selección de carpeta de destino desde el sistema
- Progreso de conversión y estado de procesamiento
- Modo oscuro con interruptor en la UI
- Interfaz moderna con Tailwind CSS

## 📦 Tecnologías
- Wails (Go + Frontend moderno)
- Vue 3 + Vite
- Tailwind CSS 4
- Heroicons (iconos)

## 🚀 Empezar

### Requisitos
- Go 1.22 o superior
- Node.js 18+ (o compatible)
- Wails CLI instalado: https://wails.io/docs/gettingstarted/installation

### Desarrollo en vivo
```bash
wails dev
```
- Recarga rápida del frontend (Vite)
- Dev server para pruebas del frontend y acceso a métodos Go

### Compilación
```bash
wails build
```
Genera el ejecutable de producción en `build/bin/`.

## 🧭 Uso
1. Arrastra y suelta imágenes o selecciónalas con el selector
2. Elige el formato de salida
3. (Opcional) Ajusta la calidad si el formato lo permite
4. Selecciona la carpeta de destino
5. Inicia la conversión y sigue el progreso

## 📁 Estructura (resumen)
```
/ (Wails)
├─ app.go, main.go, wails.json
├─ frontend/
│  ├─ src/
│  │  ├─ App.vue (UI principal)
│  │  ├─ components/ (selector, lista, opciones, progreso)
│  │  └─ utils/
│  └─ index.html, vite.config.ts, style.css
└─ build/
```

## 👤 Autor
- Desarrollador: Moisés Santos  
- GitHub: https://github.com/Moisessantos45

## 📜 Licencia
Este proyecto se distribuye con fines educativos/demostrativos. Ajusta la licencia según tus necesidades.

## 📝 Notas
- La configuración de proyecto se encuentra en `wails.json`.
- Documentación Wails: https://wails.io/docs
