package utils

import (
	"io"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

// From: https://github.com/gofiber/fiber/pull/1557
func SaveFileToStorage(fileheader *multipart.FileHeader, path string, storage fiber.Storage) error {
	file, err := fileheader.Open()
	if err != nil {
		return err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return storage.Set(path, content, 0)
}
