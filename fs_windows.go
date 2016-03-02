package fs

import (
	"os"
	"path/filepath"
)

// Max is 255 but use 245 to be safe.
const winMaxLength = 245

func absPath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return filepath.Clean(path), nil
	}
	wd, err := os.Getwd()
	return filepath.Join(wd, path), err
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

func newPathError(op, path string, err error) error {
	return &os.PathError{
		Op:   "fs: " + op,
		Path: path,
		Err:  err,
	}
}

func newLinkError(op, oldname, newname string, err error) error {
	return &os.LinkError{
		Op:  "fs: " + op,
		Old: oldname,
		New: newname,
		Err: err,
	}
}
