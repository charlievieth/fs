package fs

import "os"

func MkdirAll(path string, perm os.FileMode) error {
	return mkdirAll(path, perm)
}

func Open(name string) (*os.File, error) {
	return open(name)
}

func RemoveAll(path string) error {
	return removeAll(path)
}

func Symlink(oldname, newname string) error {
	return symlink(oldname, newname)
}

func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return openFile(name, flag, perm)
}
