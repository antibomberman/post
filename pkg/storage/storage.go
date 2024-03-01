package storage

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func randomFileName(ext string) string {
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
}
func UploadImage(r *http.Request) (string, error) {
	err := r.ParseMultipartForm(32 << 20) // Ограничить размер памяти 32 МБ
	if err != nil {
		return "", err
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		return "", err
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	if ext != ".png" && ext != ".jpg" {
		return "", errors.New("unsupported file type")
	}

	filename := fmt.Sprintf("./storage/images/%s", randomFileName(ext))
	tempFile, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return "", err
	}

	return filename, nil
}
func Delete(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
