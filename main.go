package main

import (
	"fmt"
	"os"

	"github.com/sosukesuzuki/txbuf/internal/txbuf_dir"
	"github.com/sosukesuzuki/txbuf/internal/txbuf_file"
)

func main() {
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
	err = txbuf_file.Create(txbufFiles, txbufDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
}
