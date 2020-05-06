package fs

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//Base: /Users/scottxiong/Desktop/ then /Users/scottxiong/Desktop/img/1.jpg  => img/1.jpg
func appendFiles(filename string, zipw *zip.Writer, Base string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Failed to open %s: %s", filename, err)
	}
	defer file.Close()

	if strings.Contains(filename, Base) {
		filename = strings.TrimPrefix(filename, Base)
	}
	wr, err := zipw.Create(filename)
	if err != nil {
		msg := "Failed to create entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("Failed to write %s to zip: %s", filename, err)
	}

	return nil
}

func Zip(zipName string, Base string, files []string) {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(zipName, flags, 0644)
	if err != nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}
	defer file.Close()
	zipw := zip.NewWriter(file)
	defer zipw.Close()
	for _, filename := range files {
		if err := appendFiles(filename, zipw, Base); err != nil {
			log.Fatalf("Failed to add file %s to zip: %s", filename, err)
		}
	}
}

func (f *F) Read(b []byte) (int, error) {

	return f.Size, nil
}

func ZipNotify(zipName string, Base string, files []string, chan_zip chan int) {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(zipName, flags, 0644)
	if err != nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}
	defer file.Close()
	zipw := zip.NewWriter(file)
	defer zipw.Close()
	for _, filename := range files {
		if err := appendFiles(filename, zipw, Base); err != nil {
			log.Fatalf("Failed to add file %s to zip: %s", filename, err)
		}
		chan_zip <- 1
	}
}
