package fs

import "testing"

func TestGetEnv(t *testing.T) {
	res := GetEnv("HOME")
	if res != "/Users/scottxiong" {
		t.Errorf("GetEnv(\"HOME\") failed, Got %s, expect \"/Users/scottxiong\"", res)
	}
}
