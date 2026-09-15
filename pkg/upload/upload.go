package upload

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SaveFile(c *fiber.Ctx, file *multipart.FileHeader, folder string) (string, error) {
	ext := filepath.Ext(file.Filename)
	newName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	path := fmt.Sprintf("uploads/%s/%s", folder, newName)
	if err := c.SaveFile(file, path); err != nil {
		return "", err
	}
	return path, nil
}
