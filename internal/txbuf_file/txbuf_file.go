package txbuf_file

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

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

	layout := "20060102"
	t, err := time.Parse(layout, raws[0])
	if err != nil {
		return nil, &TxbufFileError{msg: "ファイル名の日付のパースに失敗", err: err}
	}

	v, err := strconv.Atoi(raws[1])
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
	return files[0]
}
