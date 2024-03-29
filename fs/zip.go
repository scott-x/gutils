package fs

import (
	"archive/zip"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

//zip info
type ZIP struct {
	Folders []string //完整的folder
	Where   string 
	Base    string
}

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

func appendFilesIncludingHeader(filename string, zipw *zip.Writer, Base string, header *zip.FileHeader) error {
	_filename := filename
	file_info, err := os.Stat(filename)
	if err != nil {
		return fmt.Errorf("Failed to open %s: %s", filename, err)
	}
	// fmt.Println(filename)
	if strings.Contains(filename, Base) {
		filename = strings.TrimPrefix(filename, Base)
		if filename[0]=='/'{
			filename=filename[1:]
		}
	}
	//update header name https://golang.org/pkg/archive/zip/#FileHeader
	header.Name = filename
	// log.Println("filename:",filename)
	//fix bug after the file was ziped, the time is slower 8 hours than normal
	header.SetModTime(file_info.ModTime())
	w, err := zipw.CreateHeader(header)
	if err != nil {
		msg := "Failed to create entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	n := FileType(_filename)
	//if folder
	if n != 1 {
		return nil
	}
	// fmt.Println(_filename)
	//if file
	file, err := os.Open(_filename)
	if err != nil {
		return fmt.Errorf("Failed to open %s: %s", filename, err)
	}
	defer file.Close()

	if _, err := io.Copy(w, file); err != nil {
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

func zip_with_bar(zipName string, Base string, infos *INFOS, bar *pb.ProgressBar) {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(zipName, flags, 0644)
	if err != nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}
	defer file.Close()
	zipw := zip.NewWriter(file)
	defer zipw.Close()
	for _, v := range *infos {
		for i := 0; i < int(v.Size); i++ {
			bar.Increment()
		}
		time.Sleep(200 * time.Microsecond)
		if err := appendFilesIncludingHeader(v.Path, zipw, Base, v.Header); err != nil {
			log.Fatalf("Failed to add file %s to zip: %s", v.Path, err)
		}
	}
}

//func ListAllWithFileHeaders(folders []string) (*INFOS, int64, error)
func ZipWithBar(z *ZIP) {
	infos, count, _ := ListAllWithFileHeaders(z.Folders)
	// create and start new bar
	myTemplate := `{{ red "当前进度:" }} {{ bar . "[" "=" (cycle . ">" ) "." "]"}} {{percent .}} {{string . "my_green_string" | green}} {{string . "my_blue_string" | blue}}`
	bar := pb.StartNew(int(count))
	bar.SetTemplateString(myTemplate)

	zip_with_bar(z.Where, z.Base, infos, bar)
	bar.Finish()
}
