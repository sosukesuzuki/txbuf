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
	latest := txbuf_file.Latest(txbufFiles)
	if latest == nil {
		// ファイルの作成にフォールバックする
		fmt.Fprintf(os.Stdout, "ファイルなし")
		os.Exit(0)
	}
}
