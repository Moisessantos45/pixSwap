package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"image"

	"image/gif"
	"image/jpeg"
	"image/png"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	_ "image/gif"

	_ "image/jpeg"

	_ "image/png"
)

// App struct
type App struct {
	ctx context.Context
}

var counter int64

// FileData representa los datos de un archivo procesado
type FileData struct {
	ID       int    `json:"id"`       // Cambiado a mayúscula para exportación
	Base64   string `json:"base64"`   // Cambiado a mayúscula para exportación
	MimeType string `json:"mimeType"` // Cambiado a mayúscula para exportación
	Path     string `json:"path"`     // Cambiado a mayúscula para exportación
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.LogInfo(ctx, "App started and context initialized")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GenerateUniqueInt genera un ID entero único
func GenerateUniqueInt() int {
	now := time.Now().UnixNano() / 1000 // Reduce precisión para caber en int
	unique := atomic.AddInt64(&counter, 1)
	return int(now + unique)
}

// GetFilesAsBase64 procesa archivos y devuelve sus datos en Base64
func (a *App) GetFilesAsBase64(files []string) ([]FileData, error) {
	if a.ctx == nil {
		return nil, fmt.Errorf("app context not initialized")
	}

	var results []FileData

	for _, file := range files {
		// Leer el contenido del archivo
		fileContent, err := os.ReadFile(file)
		if err != nil {
			runtime.LogError(a.ctx, fmt.Sprintf("Error reading file %s: %v", file, err))
			continue
		}

		// Obtener información del archivo
		ext, mimeType := getFileInfo(file)
		if ext == "" {
			runtime.LogError(a.ctx, fmt.Sprintf("File %s has no extension", file))
			continue
		}

		// Codificar a Base64
		encodedString := base64.StdEncoding.EncodeToString(fileContent)
		if len(strings.TrimSpace(encodedString)) == 0 {
			runtime.LogError(a.ctx, fmt.Sprintf("Error encoding file %s to Base64", file))
			continue
		}

		// Crear objeto FileData
		fileData := FileData{
			ID:       GenerateUniqueInt(),
			Base64:   fmt.Sprintf("data:%s;base64,%s", mimeType, encodedString),
			MimeType: mimeType,
			Path:     file,
		}

		results = append(results, fileData)
	}

	return results, nil
}

// getFileInfo obtiene la extensión y tipo MIME de un archivo
func getFileInfo(file string) (string, string) {
	ext := strings.ToLower(filepath.Ext(file))
	if ext == "" {
		return "", ""
	}

	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		switch ext {
		case ".jpg", ".jpeg":
			mimeType = "image/jpeg"
		case ".png":
			mimeType = "image/png"
		case ".mp3":
			mimeType = "audio/mpeg"
		case ".wav":
			mimeType = "audio/wav"
		default:
			mimeType = "application/octet-stream"
		}
	}

	return ext, mimeType
}

// OpenDialog opens a directory dialog and returns the selected path
func (a *App) OpenDirectoryDialog() (string, error) {
	if a.ctx == nil {
		return "", fmt.Errorf("app context not initialized")
	}

	selection, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a directory",
	})

	if err != nil {
		return "", err
	}

	return selection, nil
}

func (a *App) OpenFilesDialog() ([]string, error) {
	if a.ctx == nil {
		return nil, fmt.Errorf("app context not initialized")
	}

	selection, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg)",
				Pattern:     "*.png;*.jpg",
			}, {
				DisplayName: "Videos (*.mov;*.mp4)",
				Pattern:     "*.mov;*.mp4",
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return selection, nil
}
func (a *App) ConvertToFormatImg(files []string, format string, quality int64, directory string) (string, error) {

	// Se recomienda una validación básica del formato para evitar errores en el switch.
	validFormats := map[string]bool{"jpeg": true, "jpg": true, "png": true, "gif": true}
	if !validFormats[strings.ToLower(format)] {
		return "", fmt.Errorf("formato no soportado: %s", format)
	}

	// Un slice para guardar los errores que ocurran durante el bucle.
	var conversionErrors []error

	for _, file := range files {
		// Abre el archivo de imagen de entrada.
		inputFile, err := os.Open(file)
		if err != nil {
			// Si hay un error, lo registramos y continuamos con el siguiente archivo.
			conversionErrors = append(conversionErrors, fmt.Errorf("error al abrir el archivo %s: %v", file, err))
			continue
		}

		// image.Decode decodifica la imagen sin importar su formato de origen.
		img, _, err := image.Decode(inputFile)
		if err != nil {
			inputFile.Close()
			conversionErrors = append(conversionErrors, fmt.Errorf("error al decodificar el archivo %s: %v", file, err))
			continue
		}
		inputFile.Close() // Cierra el archivo de entrada tan pronto como sea posible.

		// Determina el nombre del archivo de salida.
		baseFileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		outputFileName := baseFileName + "." + strings.ToLower(format)

		// Une el directorio de salida al nombre del archivo.
		var outputPath string
		if directory != "" {
			// Crea el directorio de salida si no existe.
			if _, err := os.Stat(directory); os.IsNotExist(err) {
				os.MkdirAll(directory, 0755)
			}
			outputPath = filepath.Join(directory, outputFileName)
		} else {
			outputPath = filepath.Join(filepath.Dir(file), outputFileName)
		}

		// Crea el archivo de salida.
		outputFile, err := os.Create(outputPath)
		if err != nil {
			conversionErrors = append(conversionErrors, fmt.Errorf("error al crear el archivo de salida %s: %v", outputPath, err))
			continue
		}

		// Codifica la imagen según el formato de salida.
		switch strings.ToLower(format) {
		case "jpeg", "jpg":
			options := &jpeg.Options{Quality: int(quality)}
			err = jpeg.Encode(outputFile, img, options)
		case "png":
			// El paquete png.Encode no tiene un parámetro de calidad,
			// simplemente codifica la imagen.
			err = png.Encode(outputFile, img)
		case "gif":
			// El paquete gif.Encode no tiene un parámetro de calidad,
			// simplemente codifica la imagen.
			// La compresión de GIFs es diferente.
			err = gif.Encode(outputFile, img, nil)
		}

		// Cierra el archivo de salida.
		outputFile.Close()

		if err != nil {
			conversionErrors = append(conversionErrors, fmt.Errorf("error al codificar la imagen %s a %s: %v", file, format, err))
		}
	}

	// Verifica si ocurrieron errores durante el bucle y devuelve un error concatenado.
	if len(conversionErrors) > 0 {
		return "", fmt.Errorf("errores durante la conversión: %v", conversionErrors)
	}

	return "Conversión completada con éxito.", nil
}
