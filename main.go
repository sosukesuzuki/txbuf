package main

import (
	"fmt"
	"os"

	"github.com/sosukesuzuki/txbuf/cmd"
)

func main() {
	if len(os.Args) == 1 {
		cmd.Root()
	} else {
		switch os.Args[1] {
		case "new":
			cmd.New()
		case "query":
			cmd.Query()
		case "clone":
			cmd.Clone()
		case "git":
			cmd.Git(os.Args[2:])
		default:
			fmt.Fprintf(os.Stderr, "[ERROR] txbuf doesn't support `%s`\n", os.Args[1])
			os.Exit(1)
		}
	}
}
