// +build !FSDEBUG

package fs

func newPathError(old, new, wd string, err error) error {
	return err
}
