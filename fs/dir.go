package fs

import (
	"log"
	"os"
)

func Dir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
