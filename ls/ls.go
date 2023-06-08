package ls

import (
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
	file, err := os.Open(lister.path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 50, 0, '\t', 0)
	defer w.Flush()

	err = writeDirContent(w, file)

	return err
}

func writeDirContent(w *tabwriter.Writer, file *os.File) error {
	entries, err := file.ReadDir(0)
	if err != nil {
		return err
	}

	info, err := file.Stat()
	if err != nil {
		return err
	}
	writeFileInfo(w, info)

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return err
		}
		writeFileInfo(w, info)
	}

	return nil
}

func writeFileInfo(w *tabwriter.Writer, info fs.FileInfo) {
	infoTime := info.ModTime()

	switch time.Now().Year() {
	case infoTime.Year():
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%02d:%02d\t%v\t\n",
			info.Mode(), info.Size(), infoTime.Month(), infoTime.Day(), infoTime.Hour(), infoTime.Minute(), info.Name())
	default:
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t\n",
			info.Mode(), info.Size(), infoTime.Month(), infoTime.Day(), infoTime.Year(), info.Name())
	}
}
