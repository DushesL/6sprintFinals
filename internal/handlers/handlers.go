package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HtmlHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./index.html")
}

func UploadHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20)

	file, header, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "error while getting file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "error while reading file", http.StatusInternalServerError)
		return
	}

	converted, err := service.Convert(string(content))
	if err != nil {
		http.Error(res, "error while converting", http.StatusInternalServerError)
		return
	}

	fileName := time.Now().UTC().Format("20060102150405") + filepath.Ext(header.Filename)
	newFile, err := os.Create(fileName)
	if err != nil {
		http.Error(res, "error while creating file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(converted)
	if err != nil {
		http.Error(res, "error while writing in file", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(converted))
}
