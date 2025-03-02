package delivery

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"homework_ipl/internal/config"
	"homework_ipl/utils/errors"
	"homework_ipl/utils/logger"
)

type FileHandler struct{}

type FileResponse struct {
	Path string `json:"path"`
}

var (
	maxFileSizeMB    = 1
	maxFileSizeBytes = int64(maxFileSizeMB) * 1024 * 1024
	magicTable       = map[string]string{
		"\xff\xd8\xff":      "image/jpeg",
		"\x89PNG\r\n\x1a\n": "image/png",
		"GIF87a":            "image/gif",
		"GIF89a":            "image/gif",
	}
	errUploadFile = errors.HttpError{
		Code:    http.StatusInternalServerError,
		Message: "failed upload file",
	}
	errFileExtension = errors.HttpError{
		Code:    http.StatusBadRequest,
		Message: "bad extension",
	}
	errFileSize = errors.HttpError{
		Code:    http.StatusBadRequest,
		Message: "file size exceeds the limit of " + strconv.Itoa(maxFileSizeMB) + " MB",
	}
)

func DetectType(b []byte) bool {
	flag := false
	s := string(b)
	for key, val := range magicTable {
		if strings.HasPrefix(s, key) {
			fmt.Println(val)
			flag = true
		}
	}
	return flag
}

func ValidateFileExtension(file multipart.File) bool {
	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		return false
	}

	val := DetectType(buff)

	return val
}

func ValidateFileSize(handler *multipart.FileHeader) bool {
	// Get file size
	fileSize := handler.Size

	// Check if file size exceeds a certain limit
	return fileSize <= maxFileSizeBytes
}

func SaveFile(r *http.Request, id string) (string, error) {
	logger := logger.Logger()

	// Разбор формы
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		logger.Error("Ошибка при разборе формы:", "error", err)
		return "", err
	}

	// Извлечение файла
	file, handler, err := r.FormFile("file")
	if err != nil {
		logger.Error("Ошибка при извлечении файла:", "error", err)
		return "", err
	}
	defer file.Close()

	// Логирование имени и размера загружаемого файла
	logger.Info("Загружается файл:", "name", handler.Filename, "size", handler.Size)

	// Проверка расширения файла
	if !ValidateFileExtension(file) {
		logger.Error("Неверный формат файла:", "name", handler.Filename)
		return "", errFileExtension
	}
	logger.Info("Формат файла успешно проверен:", "name", handler.Filename)

	// Сброс указателя файла после проверки расширения
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		logger.Error("Ошибка сброса указателя файла:", "name", handler.Filename, "error", err)
		return "", err
	}

	// Проверка размера файла
	if !ValidateFileSize(handler) {
		logger.Error("Файл превышает допустимый размер:", "name", handler.Filename, "size", handler.Size)
		return "", errFileSize
	}
	logger.Info("Размер файла успешно проверен:", "name", handler.Filename)

	// Создание файла на сервере
	cfg, _ := config.LoadConfig()
	targetFilePath := cfg.FileUploadPath + id + "_" + handler.Filename
	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		logger.Error("Ошибка создания файла на сервере:", "path", targetFilePath, "error", err)
		return "", err
	}
	defer targetFile.Close()

	// Копирование содержимого файла
	_, err = io.Copy(targetFile, file)
	if err != nil {
		logger.Error("Ошибка при сохранении файла:", "path", targetFilePath, "error", err)
		return "", err
	}

	// Логирование успешного сохранения файла
	logger.Info("Файл успешно загружен и сохранен:", "path", targetFilePath)

	return targetFilePath, nil
}
