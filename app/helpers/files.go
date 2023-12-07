// helpers/filestorage.go
package helpers

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

// StoreFilePublic saves the file to a public static folder.
// subfolder and filename are optional.
func StoreFilePublic(file io.Reader, subfolder, filename string) (string, error) {
	// Specify the public static folder
	publicFolder := "./static/uploads"

	// Add subfolder to the path if provided
	if subfolder != "" {
		publicFolder = filepath.Join(publicFolder, subfolder)
	}

	// Ensure the public folder exists
	if err := os.MkdirAll(publicFolder, os.ModePerm); err != nil {
		return "", err
	}

	// Detect the MIME type of the file
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}
	fileType := http.DetectContentType(buffer)

	// Derive the file extension from the MIME type
	fileExt, err := mimeToExtension(fileType)
	if err != nil {
		// Handle error if unable to determine the file extension
		return "", err
	}

	// Append the extension to the provided filename (if not already present)
	if filepath.Ext(filename) == "" {
		filename += fileExt
	}

	// Create the file on the server
	filePath := filepath.Join(publicFolder, filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := dst.Close(); err != nil {
			// Handle error closing file if necessary
		}
	}()

	// Rewind the file reader back to the beginning
	_, err = file.(io.Seeker).Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	// Copy the file to the destination
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// StoreFileProtected saves the file to a protected storage location.
// subfolder and filename are optional.
func StoreFileProtected(file io.Reader, subfolder, filename string) (string, error) {
	// Specify the protected storage folder
	protectedFolder := "./storage"

	// Add subfolder to the path if provided
	if subfolder != "" {
		protectedFolder = filepath.Join(protectedFolder, subfolder)
	}

	// Ensure the protected folder exists
	if err := os.MkdirAll(protectedFolder, os.ModePerm); err != nil {
		return "", err
	}

	// Create the file on the server
	filePath := filepath.Join(protectedFolder, filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := dst.Close(); err != nil {
			// Handle error closing file if necessary
		}
	}()

	// Rewind the file reader back to the beginning
	_, err = file.(io.Seeker).Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	// Copy the file to the destination
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

// mimeToExtension returns the file extension for a given MIME type.
func mimeToExtension(mimeType string) (string, error) {
	extensions, err := mime.ExtensionsByType(mimeType)
	if err != nil || len(extensions) == 0 {
		return "", fmt.Errorf("unable to determine file extension for MIME type %s", mimeType)
	}
	return extensions[0], nil
}
