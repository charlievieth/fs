package fs

import (
	"os"
	"path/filepath"
)

// WARN (CEV): Investigate handling of environment variables and path expansion.

// Max is 255 but use 245 to be safe.
const winMaxLength = 245

func longPath(p string) (string, error) {
	path, err := absPath(p)
	if err == nil && len(path) >= winMaxLength {
		path = `\\?\` + path
	}
	return path, err
}

func absPath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return filepath.Clean(path), nil
	}
	wd, err := os.Getwd()
	return filepath.Join(wd, path), err
}

func osPath(path string) (string, error) {
	p, err := absPath(path)
	if err != nil {
		return "", err
	}
	if len(p) >= winMaxLength {
		p = `\\?\` + p
	}
	return p, nil
}

func Path(path string) (string, error) {
	p, err := absPath(path)
	if err != nil {
		return "", err
	}
	if len(p) >= winMaxLength {
		p = `\\?\` + p
	}
	return p, nil
}
