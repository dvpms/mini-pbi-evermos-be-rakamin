package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	MaxUploadSize = 5 * 1024 * 1024 // 5MB
	UploadDir     = "./uploads"
)

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

// SaveUploadedFile validates and saves an uploaded multipart file to destDir
func SaveUploadedFile(file *multipart.FileHeader, destDir string) (string, error) {
	if file == nil {
		return "", errors.New("file is required")
	}

	if file.Size > MaxUploadSize {
		return "", errors.New("file size exceeds the 5MB limit")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		return "", errors.New("invalid file type. Only JPG, JPEG, and PNG are allowed")
	}

	if destDir == "" {
		destDir = UploadDir
	}

	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate safe unique filename
	base := filepath.Base(file.Filename)
	sanitized := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '.' || r == '-' || r == '_' {
			return r
		}
		return '_'
	}, base)

	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), sanitized)
	targetPath := filepath.Join(destDir, filename)

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(targetPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to copy file contents: %w", err)
	}

	return filename, nil
}
