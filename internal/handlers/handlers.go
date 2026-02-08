package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Ошибка разбора формы", http.StatusBadRequest)
		return
	}

	// Исправлено: используем "myFile" вместо "file"
	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка получения файла: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	content := string(fileBytes)

	convertedString, err := service.AutoDetectAndConvert(content)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка обработки файла: %v", err), http.StatusInternalServerError)
		return
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000Z")
	originalExt := filepath.Ext(fileHeader.Filename)
	resultFilename := fmt.Sprintf("результат_%s%s", timestamp, originalExt)

	resultDir := "результаты"
	if _, err := os.Stat(resultDir); os.IsNotExist(err) {
		err = os.Mkdir(resultDir, 0755)
		if err != nil {
			http.Error(w, "Ошибка создания директории", http.StatusInternalServerError)
			return
		}
	}

	resultPath := filepath.Join(resultDir, resultFilename)
	resultFile, err := os.Create(resultPath)
	if err != nil {
		http.Error(w, "Ошибка создания файла результата", http.StatusInternalServerError)
		return
	}
	defer resultFile.Close()

	_, err = resultFile.WriteString(convertedString)
	if err != nil {
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Файл успешно обработан!\n\n")
	fmt.Fprintf(w, "Исходный файл: %s\n", fileHeader.Filename)
	fmt.Fprintf(w, "Результат сохранен в: %s\n\n", resultPath)
	fmt.Fprintf(w, "Конвертированное содержимое:\n%s", convertedString)
}
