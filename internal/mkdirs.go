package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/karrick/godirwalk"
)

type mkProject struct {
	dirs chan string
}

func NewMkProject() (*mkProject, chan error) {
	mkp := &mkProject{
		dirs: make(chan string),
	}

	done := make(chan error)
	go mkp.mockDirs(done)

	return mkp, done
}

func (m *mkProject) QueueDirs() chan<- string {
	return m.dirs
}

func (m *mkProject) QueueClose() {
	close(m.dirs)
}

func (m *mkProject) mockDirs(done chan<- error) {
	defer close(done)

	mocker := make(chan int, runtime.NumCPU()) // concurrent mocker number is 512
	defer close(mocker)
	for dir := range m.dirs {
		err := godirwalk.Walk(dir, &godirwalk.Options{
			Callback: func(osPathname string, de *godirwalk.Dirent) error {
				if err := m.ignore(de, osPathname); err != nil {
					switch {
					case isIsDirError(err), isExtNotGoError(err):
						fmt.Printf("ignore: %s\n", err)
						return nil
					default:
						fmt.Printf("ignore: %s\n", err)
						return godirwalk.SkipThis
					}
				}

				mocker <- 1
				go func() {
					if err := m.mockFile(osPathname); err != nil {
						fmt.Printf("mock file: %s, %s\n", osPathname, err)
					} else {
						fmt.Printf("mock file: %s\n", osPathname)
					}
					<-mocker
				}()
				return nil
			},
			Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
		})

		for { // wait mocker finish
			if len(mocker) == 0 {
				break
			}
			time.Sleep(time.Second)
		}

		if err != nil {
			done <- fmt.Errorf("godirwalk.Walk: %s", err)
		} else {
			done <- nil
		}

		return
	}

	done <- nil
	return
}

func (m *mkProject) mockFile(osPathname string) error {
	mockFilePath := filepath.Join(MockDir, osPathname)
	if err := os.MkdirAll(filepath.Dir(mockFilePath), os.ModePerm); err != nil {
		return fmt.Errorf("os.MkdirAll: %s", err)
	}

	mockgen := filepath.Join(os.Getenv("GOPATH"), `bin`, `mockgen`)
	flag0 := `-source`
	flag1 := `-destination`
	cmd := exec.Command(mockgen, flag0, osPathname, flag1, mockFilePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("cmd.CombinedOutput: %s, command: %s, output: %s", err, cmd.String(), string(output))
	}

	return nil
}

type errIsDir struct {
	osPathname string
}

func (e errIsDir) Error() string {
	return fmt.Sprintf("%s: do not mock dir", e.osPathname)
}

type errExtNotGo struct {
	osPathname string
	suffix     string
}

func (e errExtNotGo) Error() string {
	msg := fmt.Sprintf("%s: do not mock %s", e.osPathname, e.suffix)
	if len(e.suffix) == 0 {
		msg = fmt.Sprintf("%s: do not mock normal file", e.osPathname)
	}
	return msg
}

func isIsDirError(err error) bool {
	_, ok := err.(errIsDir)
	return ok
}

func isExtNotGoError(err error) bool {
	_, ok := err.(errExtNotGo)
	return ok
}

func (m *mkProject) ignore(de *godirwalk.Dirent, osPathname string) error {
	const ignoreMsgFmt string = "%s: skipped dir"

	if filepath.Base(osPathname) == `.` {
		return errIsDir{osPathname: osPathname}
	}
	if filepath.Base(osPathname) == `.git` {
		return fmt.Errorf(ignoreMsgFmt, osPathname)
	}
	if filepath.Base(osPathname) == `.idea` {
		return fmt.Errorf(ignoreMsgFmt, osPathname)
	}
	if filepath.Base(osPathname) == `vendor` {
		return fmt.Errorf(ignoreMsgFmt, osPathname)
	}
	if filepath.Base(osPathname) == `test` {
		return fmt.Errorf(ignoreMsgFmt, osPathname)
	}
	if filepath.Base(osPathname) == `.fwconfig` {
		return fmt.Errorf(ignoreMsgFmt, osPathname)
	}
	if filepath.Base(osPathname) == `.tools` {
		return fmt.Errorf(ignoreMsgFmt, osPathname)
	}
	if filepath.Base(osPathname) == `build` {
		return fmt.Errorf(ignoreMsgFmt, osPathname)
	}

	if de.IsDir() {
		return errIsDir{osPathname: osPathname}
	}

	if strings.Contains(osPathname, "_test.go") {
		return errExtNotGo{osPathname: osPathname, suffix: "_test.go"}
	}

	if filepath.Ext(osPathname) != ".go" {
		return errExtNotGo{osPathname: osPathname, suffix: filepath.Ext(osPathname)}
	}

	return nil
}
