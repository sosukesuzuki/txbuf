package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sosukesuzuki/txbuf/internal/editor"
	"github.com/sosukesuzuki/txbuf/internal/txbuf_dir"
	"github.com/sosukesuzuki/txbuf/internal/txbuf_file"
)

func Root() {
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
	var path string
	latest := txbuf_file.Latest(txbufFiles)
	if latest == nil {
		path, err = txbuf_file.Create(txbufFiles, txbufDir)
	} else {
		path = filepath.Join(txbufDir, latest.Name())
	}
	editor.OpenVim(path)
}
