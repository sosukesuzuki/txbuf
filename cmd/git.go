package cmd

import (
	"fmt"
	"os"

	"github.com/sosukesuzuki/txbuf/internal/git"
	"github.com/sosukesuzuki/txbuf/internal/txbuf_dir"
)

func Git(args []string) {
	txbufDir, err := txbuf_dir.GetTxbufDir()
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
	o, err := git.RunGit(txbufDir, args)
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)	
	}
	fmt.Print(o)
}
