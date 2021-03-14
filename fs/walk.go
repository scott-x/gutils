package fs

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//exected file type
const (
	FILE = iota
	FOLDER
	MIX
)

type Task struct {
	Dir string
	*Expect
}

//constraint
type Expect struct {
	T             int // YOUR TARGET FILE TYPE, only FILE,FOLDER,MIX can be used here
	Match         string //regexp string, used to filter the data
	IgnoreFolders []string //ignore folders contains that defined the []string
	Depth         int 
	/* If your file type is folder, then you can set as 1, when found the target folder, 
	it won't walk the target folder; if not set, will walk all folder*/
}

func NewTasks(pths []string, e *Expect) *[]Task {
	tasks := make([]Task, 0)
	for _, pth := range pths {
		task := &Task{
			Dir:    pth,
			Expect: e,
		}
		tasks = append(tasks, *task)
	}
	return &tasks
}

func (t *Task) Walk(c chan string, done chan bool) error {
	err := filepath.Walk(t.Dir, func(pth string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//ignore folders
		ignores := t.IgnoreFolders
		expect_file_type := t.T
		re := regexp.MustCompile(t.Match)
		res := re.FindString(info.Name())
		if info.IsDir() {
			if expect_file_type == FOLDER || expect_file_type == MIX {

				if len(res) > 0 {
					go func() {
						c <- pth
					}()
					if t.Depth == 1 {
						//should_ignore = append(should_ignore, pth)
						return filepath.SkipDir
					}
				} else {
					if len(ignores) > 0 {
						//ignored := false
						for _, ignore := range ignores {
							if strings.Contains(info.Name(), ignore) {
								//fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
								return filepath.SkipDir
							}
						}
					}
				}
			} else {

			}
		} else { // if file
			if expect_file_type == FILE || expect_file_type == MIX {
				if len(res) > 0 {
					go func() {
						c <- pth
					}()
				}
			}
		}
		//fmt.Printf("visited file: %q\n", pth)
		return nil
	})
	if err != nil {
		return err
	}
	done <- true
	return nil
}
