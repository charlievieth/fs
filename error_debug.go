// +build FSDEBUG

package fs

import (
	"fmt"
	"runtime"
)

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
