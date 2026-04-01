package utils

import (
	"io/fs"
	"os"
	"path/filepath"
)

func ReadYAMLFiles(root string) ([][]byte, error) {
	var files [][]byte

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml" {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			files = append(files, data)
		}
		return nil
	})

	return files, err
}
