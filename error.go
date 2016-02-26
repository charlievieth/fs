// +build !FSDEBUG

package fs

func newPathError(old, new string, err error) error {
	return err
}
