package token

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

//new token
func New() string {
	curtime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(curtime, 10))
	return fmt.Sprintf("%x", h.Sum(nil))
}

//generate token with salt
func NewWithSalt(salt string) string {
	h := md5.New()
	io.WriteString(h, salt+time.Now().String())
	return fmt.Sprintf("%x", h.Sum(nil))
}