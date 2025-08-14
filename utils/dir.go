package utils

import (
	"os"
	"path/filepath"
)

func getSubdirectories(root string) ([]string, error) {
	var subdirs []string

	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subdirs = append(subdirs, filepath.Join(root, entry.Name()))
		}
	}

	return subdirs, nil
}
