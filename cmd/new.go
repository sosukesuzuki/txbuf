package cmd

import (
	"fmt"
	"os"

	"github.com/sosukesuzuki/txbuf/internal/editor"
	"github.com/sosukesuzuki/txbuf/internal/txbuf_dir"
	"github.com/sosukesuzuki/txbuf/internal/txbuf_file"
)

func New() {
	txbufFiles, err := txbuf_dir.Files()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
	txbufDir, err := txbuf_dir.GetTxbufDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
	path, err := txbuf_file.Create(txbufFiles, txbufDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
	editor.OpenVim(path)
}
