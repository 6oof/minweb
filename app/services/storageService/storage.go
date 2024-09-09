package storageService

import (
	"fmt"
	"os"
	"path/filepath"
)

// StorageInterface defines the methods for interacting with file storage systems.

type LocalStorage struct {
	rootPath string
}

func Make(path string) *LocalStorage {
	storage := &LocalStorage{}
	storage.Init(path)
	return storage
}

// Init initializes the local storage with the root directory.
func (l *LocalStorage) Init(rootPath string) {
	l.rootPath = rootPath
}

// Put stores a file at the given path.
func (l *LocalStorage) Put(filePath string, content []byte) error {
	fullPath := filepath.Join(l.rootPath, filePath)
	dir := filepath.Dir(fullPath)

	// Ensure the directory exists
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	// Write the file
	return os.WriteFile(fullPath, content, os.ModePerm)
}

// Get retrieves the content of a file.
func (l *LocalStorage) Get(filePath string) ([]byte, error) {
	fullPath := filepath.Join(l.rootPath, filePath)
	return os.ReadFile(fullPath)
}

// Delete removes a file.
func (l *LocalStorage) Delete(filePath string) error {
	fullPath := filepath.Join(l.rootPath, filePath)
	return os.Remove(fullPath)
}

// Exists checks if a file exists.
func (l *LocalStorage) Exists(filePath string) (bool, error) {
	fullPath := filepath.Join(l.rootPath, filePath)
	_, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}
