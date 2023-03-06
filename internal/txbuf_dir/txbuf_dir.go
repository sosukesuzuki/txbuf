package txbuf_dir

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

type TxbufDirError struct {
	msg string
	err error
}

func (e *TxbufDirError) Error() string {
	return fmt.Sprintf("Error from texbuf: %s (%s)", e.msg, e.err.Error())
}

func (e *TxbufDirError) Unwrap() error {
	return e.err
}

// .txtbuf ディレクトリのパスを取得する
func GetTxbufDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", &TxbufDirError{msg: ".txbufの取得に失敗", err: err}
	}
	txbufDir := filepath.Join(home, ".txbuf")
	return txbufDir, nil
}

// 与えられたパスのディレクトリが存在するか確認する
func existsDir(d string) bool {
	if f, err := os.Stat(d); os.IsNotExist(err) || !f.IsDir() {
		return false
	} else {
		return true
	}
}

// .txbuf ディレクトリに含まれているすべてのファイルを返す
func Files() ([]fs.FileInfo, error) {
	txbufDir, err := GetTxbufDir()
	if err != nil {
		return nil, err
	}
	ok := existsDir(txbufDir)
	if !ok {
		err := os.Mkdir(txbufDir, 0777)
		if err != nil {
			return nil, &TxbufDirError{msg: ".txbufの作成に失敗", err: err}
		}
	}
	files, err := ioutil.ReadDir(txbufDir)
	if err != nil {
		return nil, &TxbufDirError{msg: ".txbufの中身の取得に失敗", err: err}
	}
	return files, nil
}
