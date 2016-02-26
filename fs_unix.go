// +build !windows

package fs

func Path(path string) (string, error) {
	return path, nil
}
