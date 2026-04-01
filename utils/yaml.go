// Package utils provides generic file and data parsing mechanisms for Flamingo GitOps Engine.
package utils

import (
	"io/fs"
	"os"
	"path/filepath"
)

// ReadYAMLFiles recursively scans the requested root directory and reads any file matching
// the `.yaml` or `.yml` extension. Returns a list of the byte contents of all discovered files.
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
