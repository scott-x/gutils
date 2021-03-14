package fs

import (
	"regexp"
	"io/ioutil"
	"path"
)

//get expected file or folder, which match re, only return the 1st one
func GetExpectedPath(folder,re string) string {
	var result string
	r := regexp.MustCompile(re) //`^[UCBP]2.*\.xlsx?`
	fls, err := ioutil.ReadDir(folder)
	if err != nil {
		panic(err)
	}
	for _, f := range fls {
		res := r.FindString(f.Name())
		if len(res) > 0 {
			result = path.Join(folder, f.Name())
			break
		}
	}
	return result
}