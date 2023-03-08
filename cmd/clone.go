package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/sosukesuzuki/txbuf/internal/editor"
	"github.com/sosukesuzuki/txbuf/internal/txbuf_dir"
	"github.com/sosukesuzuki/txbuf/internal/txbuf_file"
)

func Clone() {
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
		fmt.Fprintf(os.Stdout, "[INFO] 最新のファイルが見つかりませんでした")
		os.Exit(0)
	}
	path = filepath.Join(txbufDir, latest.Name())

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}

	newFileName, err := txbuf_file.NewFileName(time.Now(), latest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}

	newPath := filepath.Join(txbufDir, newFileName)

	f, err := os.Create(newPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	err = ioutil.WriteFile(newPath, data, fs.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
		os.Exit(1)
	}

	editor.OpenVim(newPath)
}
