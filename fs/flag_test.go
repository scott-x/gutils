package fs

import (
	"testing"
)

// 0 dir, 1 file, -1 error
func TestFileType(t *testing.T) {
	res := FileType("flag.go")
	if res != 1 {
		t.Errorf("Test FileType(\"flag.go\") failed, expect 1, but got %d", res)
	}
}
