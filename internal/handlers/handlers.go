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

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func Upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Не удалось получить файл: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Не удалось прочитать файл: "+err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	filename := time.Now().UTC().Format("2006-01-02_15-04-05.000") + ext

	outFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Не удалось создать файл: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	if _, err := outFile.WriteString(result); err != nil {
		http.Error(w, "Не удалось записать результат: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, result)
}
