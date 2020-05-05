package main

import (
	"crypto/rand"
	"io"
	"io/ioutil"

	"github.com/cheggaaa/pb/v3"
)

func main() {

}

func run() {
	tmpl := `{{ red "当前进度:" }} {{ bar . "[" "=" (cycle . ">" ) "." "]"}} {{percent .}}`

	var limit int64 = 1024 * 1024 * 500
	// we will copy 200 Mb from /dev/rand to /dev/null
	reader := io.LimitReader(rand.Reader, limit)
	writer := ioutil.Discard
	// start new bar
	// start bar based on our template
	bar := pb.ProgressBarTemplate(tmpl).Start64(limit)
	// set values for string elements
	// create proxy reader
	barReader := bar.NewProxyReader(reader)
	// copy from proxy reader
	io.Copy(writer, barReader)
	// finish bar
	bar.Finish()
}
