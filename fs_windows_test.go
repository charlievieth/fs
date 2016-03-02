// +build windows

package fs

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func MakePath() string {
	volume := os.TempDir() + string(filepath.Separator)
	buf := bytes.NewBufferString(volume)
	for i := 0; i < 2; i++ {
		for i := byte('A'); i <= 'Z'; i++ {
			buf.Write(bytes.Repeat([]byte{i}, 4))
			buf.WriteRune(filepath.Separator)
		}
	}
	return filepath.Clean(buf.String())
}

func TestMkdirAll(t *testing.T) {
	path := MakePath()
	defer func() {
		if err := os.RemoveAll(path); err != nil {
			t.Fatalf("TestMkdirAll: RemoveAll %s", err)
		}
	}()
	err := MkdirAll(MakePath(), 0755)
	if err != nil {
		t.Fatalf("TestMkdirAll: %s", err)
	}
}
