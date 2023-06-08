package main

import (
	"flag"
	"fmt"

	"github.com/marcocab/gocommander/command"
	"github.com/marcocab/gocommander/ls"
)

var cmds []command.Command

func init() {
	flag.Func("ls", "List the files of a given path", func(s string) error {
		fmt.Println("ls: ", s)
		cmds = append(cmds, ls.NewLister(s))
		return nil
	})
	flag.Func("which", "Show the location of an executable", func(s string) error {
		fmt.Println("which: ", s)
		return nil
	})
	flag.Func("cat", "Print the content of a file", func(s string) error {
		fmt.Println("cat: ", s)
		return nil
	})
}

func main() {

	flag.Parse()

	for _, cmd := range cmds {
		cmd.Run()
	}
}
