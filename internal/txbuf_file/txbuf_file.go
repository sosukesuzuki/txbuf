package txbuf_file

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const layout = "20060102"

type TxbufFileError struct {
	msg string
	err error
}

func (e *TxbufFileError) Error() string {
	return fmt.Sprintf("Error from texbuf_file: %s (%s)", e.msg, e.err.Error())
}
func (e *TxbufFileError) Unwrap() error {
	return e.err
}

type ParsedTxbufName struct {
	date    time.Time
	version uint
}

// 20230301-1 のような形のファイル名をパースする
func Parse(n string) (*ParsedTxbufName, error) {
	raws := strings.Split(n, "-")
	if len(raws) != 2 {
		return nil, &TxbufFileError{msg: "ファイル名のパースに失敗"}
	}

	t, err := time.Parse(layout, raws[0])
	if err != nil {
		return nil, &TxbufFileError{msg: "ファイル名の日付のパースに失敗", err: err}
	}

	v, err := strconv.Atoi(strings.Split(raws[1], ".")[0])
	if err != nil {
		return nil, &TxbufFileError{msg: "ファイル名のバージョンのパースに失敗", err: err}
	}

	return &ParsedTxbufName{
		date:    t,
		version: uint(v),
	}, nil
}

type ByTxbufName []fs.FileInfo

func (f ByTxbufName) Len() int      { return len(f) }
func (f ByTxbufName) Swap(i, j int) { f[i], f[j] = f[j], f[i] }
func (f ByTxbufName) Less(i, j int) bool {
	a, err := Parse(f[i].Name())
	if err != nil {
		fmt.Fprintf(os.Stdout, "[INFO] %v\n", err)
		return false
	}
	b, err := Parse(f[j].Name())
	if err != nil {
		fmt.Fprintf(os.Stdout, "[INFO] %v\n", err)
		return false
	}
	if a.date.Before(b.date) {
		return true
	}
	return a.version < b.version
}

// 与えられたファイルの配列から、命名規則に従ってもっとも最近作られたファイルを返す
func Latest(files []fs.FileInfo) fs.FileInfo {
	sort.Sort(ByTxbufName(files))
	if len(files) == 0 {
		return nil
	}
	return files[len(files)-1]
}

// 新しい .txbuf ファイルの名前を生成する
func NewFileName(t time.Time, latest fs.FileInfo) (string, error) {
	d := t.Format(layout)
	var version uint
	if latest == nil {
		version = 1
	} else {
		p, err := Parse(latest.Name())
		if err != nil {
			return "", &TxbufFileError{msg: "最新のファイルのパースに失敗", err: err}
		}
		version = p.version + 1
	}
	newFile := fmt.Sprintf("%s-%d.txt", d, version)
	return newFile, nil
}

func Create(files []fs.FileInfo, txbufDir string) (string, error) {
	latest := Latest(files)
	newFile, err := NewFileName(time.Now(), latest)
	if err != nil {
		return "", &TxbufFileError{msg: "新しいファイルの作成に失敗", err: err}
	}
	newAbsFile := filepath.Join(txbufDir, newFile)
	f, err := os.Create(newAbsFile)
	if err != nil {
		return "", &TxbufFileError{msg: "新しいファイルの作成に失敗", err: err}
	}
	defer f.Close()
	return newAbsFile, nil
}
