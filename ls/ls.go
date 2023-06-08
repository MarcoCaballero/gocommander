package ls

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"text/tabwriter"
	"time"
)

type Lister struct {
	path string
}

func NewLister(path string) *Lister {
	return &Lister{
		path: path,
	}
}

func (lister *Lister) Run() error {
	if !fs.ValidPath(lister.path) {
		return errors.New("Invalid path")
	}

	if _, err := os.Stat(lister.path); err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 50, 0, '\t', 0)
	defer w.Flush()
	err := fs.WalkDir(os.DirFS(lister.path), ".", func(path string, d fs.DirEntry, err error) error {
		info, infoErr := d.Info()
		if infoErr == nil {
			printInfoLine(w, info)
		}
		return infoErr
	})
	return err
}

func printInfoLine(w *tabwriter.Writer, info fs.FileInfo) {
	infoTime := info.ModTime()

	switch time.Now().Year() {
	case infoTime.Year():
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%02d:%02d\t%v\t\n", info.Mode(), info.Size(), infoTime.Month(), infoTime.Day(), infoTime.Hour(), infoTime.Minute(), info.Name())
	default:
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t\n", info.Mode(), info.Size(), infoTime.Month(), infoTime.Day(), infoTime.Year(), info.Name())
	}
}
