// +build !windows

package fs

func Path(path string) (string, error) {
	return path, nil
}

func newPathError(_, _ string, err error) error {
	return err
}

func newLinkError(_, _, _ string, err error) error {
	return err
}
