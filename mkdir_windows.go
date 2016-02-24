package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func mkdirAll(path string, perm os.FileMode) error {
	p, wd, err := longPath(path)
	if err != nil {
		return newPathError(path, p, wd, err)
	}
	if err := os.MkdirAll(p, perm); err != nil {
		return newPathError(path, p, wd, err)
	}
	return nil
}

func open(name string) (*os.File, error) {
	p, wd, err := longPath(name)
	if err != nil {
		return nil, newPathError(name, p, wd, err)
	}
	f, err := os.Open(name)
	if err != nil {
		return nil, newPathError(name, p, wd, err)
	}
	return f, nil
}

func removeAll(path string) error {
	p, wd, err := longPath(path)
	if err != nil {
		return newPathError(path, p, wd, err)
	}
	if err := os.RemoveAll(p); err != nil {
		return newPathError(path, p, wd, err)
	}
	return nil
}

func symlink(oldname, newname string) error {
	op, wd, err := longPath(oldname)
	if err != nil {
		return newPathError(oldname, op, wd, err)
	}
	np, _, err := longPath(newname)
	if err != nil {
		return newPathError(newname, np, wd, err)
	}
	if err := os.Symlink(op, np); err != nil {
		return newPathError(op, np, wd, err)
	}
	return nil
}

func openFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	p, wd, err := longPath(name)
	if err != nil {
		return nil, newPathError(name, p, wd, err)
	}
	f, err := os.OpenFile(p, flag, perm)
	if err != nil {
		return nil, newPathError(name, p, wd, err)
	}
	return f, nil
}

func longPath(p string) (path, wd string, err error) {
	path, wd, err = absPath(p)
	if err != nil {
		return
	}
	if len(path) > 245 {
		path = `\\?\` + path
	}
	return
}

func absPath(path string) (abs, wd string, err error) {
	if filepath.IsAbs(path) {
		abs = filepath.Clean(path)
	} else {
		wd, err = os.Getwd()
		abs = filepath.Join(wd, path)
	}
	return
}

func newPathError(old, new, wd string, err error) error {
	var caller string
	_, file, line, ok := runtime.Caller(2)
	if ok {
		caller = fmt.Sprintf("%s:#%d", file, line)
	}
	var function string
	_, file, line, ok = runtime.Caller(1)
	if ok {
		function = fmt.Sprintf("%s:#%d", file, line)
	}
	return &pathError{
		err:      err,
		oldPath:  old,
		newPath:  new,
		wd:       wd,
		caller:   caller,
		function: function,
	}
}

type pathError struct {
	err      error
	oldPath  string
	newPath  string
	wd       string
	caller   string
	function string
}

func (e *pathError) Error() string {
	const format = "WindowsPathError:\n" +
		"  Error: %s\n" +
		"  Original Path: %s\n" +
		"  Modified Path: %s\n" +
		"  Working Directory: %s\n" +
		"  Caller: %s\n" +
		"  Function: %s"
	return fmt.Sprintf(format, e.err, e.oldPath, e.newPath, e.wd, e.caller)
}
