package session

import (
	"io"
	"crypto/rand"
	"encoding/base64"
)

//generate session id
func New() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
