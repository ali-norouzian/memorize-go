package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindProjectRoot() (*string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	for {
		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {

			return &cwd, nil
		}

		parent := filepath.Dir(cwd)
		if parent == cwd {
			return nil, fmt.Errorf("go.mod not found in any parent directories")
		}

		cwd = parent
	}
}
