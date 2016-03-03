package fs

import (
	"os"
	"syscall"
	"time"
)

// For testing.
func atime(fi os.FileInfo) time.Time {
	return time.Unix(0, fi.Sys().(*syscall.Win32FileAttributeData).LastAccessTime.Nanoseconds())
}
