package main

import (
	_ "crypto/rand"
	_ "fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/scott-x/gutils/fs"
	"io"
	"io/ioutil"
)

func main() {
	run()
}

// type Task io.Writer

// func (t *Task) Write(p []byte) (n int, err error) {
// 	return <-zip_chan, nil
// }

func run() {
	tmpl := `{{ red "当前进度:" }} {{ bar . "[" "=" (cycle . ">" ) "." "]"}} {{percent .}}`

	var limit int64 = 1024
	// we will copy 200 Mb from /dev/rand to /dev/null

	reader := io.LimitReader(&fs.F{}, limit)

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

// package main

// import (
// 	"fmt"
// 	"github.com/scott-x/gutils/fs"
// )

// func main() {
// 	_, sum, _ := fs.ListAll1("/Users/scottxiong/go/src/github.com/scott-x/gutils", []string{".git"})
// 	fmt.Println(sum)
// 	// for _, v := range *ff {
// 	// 	fmt.Println(v.Size)
// 	// 	fmt.Println(v.Path)
// 	// }
// }
