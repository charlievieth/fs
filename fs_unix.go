// +build !windows

package fs

import "os"

func mkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func open(name string) (*os.File, error) {
	return os.Open(name)
}

func removeAll(path string) error {
	return os.RemoveAll(path)
}

func symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}

func openFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}
