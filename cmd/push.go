package cmd

import (
	"fmt"
	"os"

	"github.com/sosukesuzuki/txbuf/internal/git"
	"github.com/sosukesuzuki/txbuf/internal/txbuf_dir"
)

func Push() {
	txbuf_dir, err := txbuf_dir.GetTxbufDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
	o1, err := git.RunGit(txbuf_dir, []string{"add", "."})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
	fmt.Print(o1)
	o2, err := git.RunGit(txbuf_dir, []string{"commit", "-m", "\"Update\""})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
	fmt.Print(o2)
	o3, err := git.RunGit(txbuf_dir, []string{"push", "origin", "HEAD"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
	fmt.Print(o3)
}
