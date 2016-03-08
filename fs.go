package fs

import (
	"os"
	"time"
)

// TODO:
// 	 - PathError: make sure that returned error type is the same as the
// 	   error type returned from the standard library
// 	 - Test 255 char MAX_PATH and adjust the cutoff that we use (245)

// Chdir changes the current working directory to the named directory.
// If there is an error, it will be of type *PathError.
func Chdir(dir string) error {
	p, err := Path(dir)
	if err != nil {
		return newPathError("chdir", dir, err)
	}
	return os.Chdir(p)
}

// Chmod changes the mode of the named file to mode.
// If the file is a symbolic link, it changes the mode of the link's target.
// If there is an error, it will be of type *PathError.
func Chmod(name string, mode os.FileMode) error {
	p, err := Path(name)
	if err != nil {
		return newPathError("chmod", name, err)
	}
	return os.Chmod(p, mode)
}

// Chown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link's target.
// If there is an error, it will be of type *PathError.
func Chown(name string, uid, gid int) error {
	p, err := Path(name)
	if err != nil {
		return newPathError("chown", name, err)
	}
	return os.Chown(p, uid, gid)
}

// Chtimes changes the access and modification times of the named
// file, similar to the Unix utime() or utimes() functions.
//
// The underlying filesystem may truncate or round the values to a
// less precise time unit.
// If there is an error, it will be of type *PathError.
func Chtimes(name string, atime time.Time, mtime time.Time) error {
	p, err := Path(name)
	if err != nil {
		return newPathError("chtimes", name, err)
	}
	return os.Chtimes(p, atime, mtime)
}

// Lchown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link itself.
// If there is an error, it will be of type *PathError.
func Lchown(name string, uid, gid int) error {
	p, err := Path(name)
	if err != nil {
		return newPathError("lchown", name, err)
	}
	return os.Lchown(p, uid, gid)
}

// Link creates newname as a hard link to the oldname file.
// If there is an error, it will be of type *LinkError.
func Link(oldname, newname string) error {
	op, err := Path(oldname)
	if err != nil {
		return newLinkError("link", oldname, newname, err)
	}
	np, err := Path(newname)
	if err != nil {
		return newLinkError("link", oldname, newname, err)
	}
	return os.Link(op, np)
}

// Mkdir creates a new directory with the specified name and permission bits.
// If there is an error, it will be of type *PathError.
func Mkdir(name string, perm os.FileMode) error {
	p, err := Path(name)
	if err != nil {
		return newPathError("mkdir", name, err)
	}
	return os.Mkdir(p, perm)
}

// MkdirAll creates a directory named path,
// along with any necessary parents, and returns nil,
// or else returns an error.
// The permission bits perm are used for all
// directories that MkdirAll creates.
// If path is already a directory, MkdirAll does nothing
// and returns nil.
func MkdirAll(path string, perm os.FileMode) error {
	p, err := Path(path)
	if err != nil {
		return err
	}
	return os.MkdirAll(p, perm)
}

// Readlink returns the destination of the named symbolic link.
// If there is an error, it will be of type *PathError.
func Readlink(name string) (string, error) {
	p, err := Path(name)
	if err != nil {
		return "", newPathError("readlink", name, err)
	}
	return os.Readlink(p)
}

// Remove removes the named file or directory.
// If there is an error, it will be of type *PathError.
func Remove(name string) error {
	p, err := Path(name)
	if err != nil {
		return newPathError("remove", name, err)
	}
	return os.Remove(p)
}

// RemoveAll removes path and any children it contains.
// It removes everything it can but returns the first error
// it encounters.  If the path does not exist, RemoveAll
// returns nil (no error).
func RemoveAll(path string) error {
	p, err := Path(path)
	if err != nil {
		return err
	}
	return os.RemoveAll(p)
}

// Rename renames (moves) a file. OS-specific restrictions might apply.
// If there is an error, it will be of type *LinkError.
func Rename(oldpath, newpath string) error {
	op, err := Path(oldpath)
	if err != nil {
		return newLinkError("rename", oldpath, newpath, err)
	}
	np, err := Path(newpath)
	if err != nil {
		return newLinkError("rename", oldpath, newpath, err)
	}
	return os.Rename(op, np)
}

// Symlink creates newname as a symbolic link to oldname.
// If there is an error, it will be of type *LinkError.
func Symlink(oldname, newname string) error {
	op, err := Path(oldname)
	if err != nil {
		return newLinkError("symlink", oldname, newname, err)
	}
	np, err := Path(newname)
	if err != nil {
		return newLinkError("symlink", oldname, newname, err)
	}
	return os.Symlink(op, np)
}

// File

// Create creates the named file with mode 0666 (before umask), truncating
// it if it already exists. If successful, methods on the returned
// File can be used for I/O; the associated file descriptor has mode
// O_RDWR.
// If there is an error, it will be of type *PathError.
func Create(name string) (*os.File, error) {
	p, err := Path(name)
	if err != nil {
		return nil, newPathError("create", name, err)
	}
	return os.Create(p)
}

// NewFile returns a new File with the given file descriptor and name.
func NewFile(fd uintptr, name string) *os.File {
	p, err := Path(name)
	if err != nil {
		return os.NewFile(fd, name)
	}
	return os.NewFile(fd, p)
}

// Open opens the named file for reading.  If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.
func Open(name string) (*os.File, error) {
	p, err := Path(name)
	if err != nil {
		return nil, newPathError("open", name, err)
	}
	return os.Open(p)
}

// OpenFile is the generalized open call; most users will use Open
// or Create instead.  It opens the named file with specified flag
// (O_RDONLY etc.) and perm, (0666 etc.) if applicable.  If successful,
// methods on the returned File can be used for I/O.
// If there is an error, it will be of type *PathError.
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	p, err := Path(name)
	if err != nil {
		return nil, newPathError("openfile", name, err)
	}
	return os.OpenFile(p, flag, perm)
}

// FileInfo

// Lstat returns a FileInfo describing the named file.
// If the file is a symbolic link, the returned FileInfo
// describes the symbolic link.  Lstat makes no attempt to follow the link.
// If there is an error, it will be of type *PathError.
func Lstat(name string) (os.FileInfo, error) {
	p, err := Path(name)
	if err != nil {
		return nil, newPathError("lstat", name, err)
	}
	return os.Lstat(p)
}

// Stat returns a FileInfo describing the named file.
// If there is an error, it will be of type *PathError.
func Stat(name string) (os.FileInfo, error) {
	p, err := Path(name)
	if err != nil {
		return nil, newPathError("stat", name, err)
	}
	return os.Stat(p)
}
