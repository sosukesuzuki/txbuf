package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sosukesuzuki/txbuf/internal/editor"
	"github.com/sosukesuzuki/txbuf/internal/peco"
	"github.com/sosukesuzuki/txbuf/internal/txbuf_dir"
)

func Query() {
	txbufDir, err := txbuf_dir.GetTxbufDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
	out, err := peco.OpenPeco(txbufDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
	if out != "" {
		// 末尾の文字を削除
		o := out[:len(out)-1]
		editor.OpenVim(filepath.Join(txbufDir, o))
	}
}
