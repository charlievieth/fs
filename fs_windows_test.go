// +build windows

package fs

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func longPathName() string {
	var buf bytes.Buffer
	for i := 0; i < 2; i++ {
		for i := byte('A'); i <= 'Z'; i++ {
			buf.Write(bytes.Repeat([]byte{i}, 4))
			buf.WriteRune(filepath.Separator)
		}
	}
	return filepath.Clean(buf.String())
}

func TestRemoveAll(t *testing.T) {
	name := longPathName()
	temp := os.TempDir()
	path := filepath.Join(temp, name)
	target := filepath.Join(temp, strings.Split(name, string(os.PathSeparator))[0])

	err := MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("TestRemoveAll: %s", err)
	}
	defer os.RemoveAll(`\\?\` + target)

	// TODO: cleanup on failure
	if err := RemoveAll(target); err != nil {
		t.Fatalf("TestRemoveAll: %s", err)
	}
	if _, err := Stat(path); err == nil {
		t.Fatalf("TestRemoveAll: failed to remove directory: %s", path)
	}
	if _, err := Stat(target); err == nil {
		t.Fatalf("TestRemoveAll: failed to remove directory: %s", target)
	}
}

func TestMkdirAll(t *testing.T) {
	name := longPathName()
	temp := os.TempDir()
	path := filepath.Join(temp, name)
	target := filepath.Join(temp, strings.Split(name, string(os.PathSeparator))[0])

	err := MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("TestMkdirAll: %s", err)
	}
	defer os.RemoveAll(`\\?\` + target)

	if _, err := Stat(path); err != nil {
		t.Fatalf("TestMkdirAll: Stat failed %s", err)
	}
	// Make sure the handling of long paths is case-insensitive
	if _, err := Stat(strings.ToLower(path)); err != nil {
		t.Fatalf("TestMkdirAll: Stat failed %s", err)
	}
	if err := os.RemoveAll(path); err != nil {
		t.Fatalf("TestMkdirAll: RemoveAll %s", err)
	}
}
