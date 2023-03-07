package txbuf_file_test

import (
	"io/fs"
	"testing"
	"time"

	"github.com/sosukesuzuki/txbuf/internal/txbuf_file"
)

const layout = "20060102"

func TestParse(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantDate    time.Time
		wantVersion uint
		wantError   bool
	}{
		{
			name:        "バージョンが1桁のtxbufファイル名をパース",
			input:       "20220301-1.txt",
			wantDate:    time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC),
			wantVersion: 1,
			wantError:   false,
		},
		{
			name:        "バージョンが2桁のtxbufファイル名をパース",
			input:       "20220301-14.txt",
			wantDate:    time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC),
			wantVersion: 14,
			wantError:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := txbuf_file.Parse(tt.input)
			if err != nil {
				if !tt.wantError {
					t.Errorf("予期せぬエラー %v", err)
					return
				}
				return
			}
			if !parsed.Date.Equal(tt.wantDate) {
				t.Errorf("予期せぬ値 期待する値: %v, 実際の値: %v", tt.wantDate, parsed.Date)
				return
			}
			if parsed.Version != tt.wantVersion {
				t.Errorf("予期せぬ値 期待する値: %v, 実際の値: %v", tt.wantVersion, parsed.Date)
				return
			}
		})
	}
}

type DummyFileInfo struct{ name string }

func (f DummyFileInfo) Name() string       { return f.name }
func (f DummyFileInfo) Size() int64        { return 0 }
func (f DummyFileInfo) Mode() fs.FileMode  { return 0 }
func (f DummyFileInfo) ModTime() time.Time { return time.Now() }
func (f DummyFileInfo) IsDir() bool        { return false }
func (f DummyFileInfo) Sys() any           { return nil }

func NewDummyFileInfo(name string) fs.FileInfo {
	return DummyFileInfo{name}
}
func TestLatest(t *testing.T) {
	tests := []struct {
		name       string
		inputNames []string
		wantName   string
	}{
		{
			name: "異なる日付の場合、最新の日付のファイルを取得する",
			inputNames: []string{
				"20220301-1.txt", "20220302-1.txt", "20220303-1.txt",
			},
			wantName: "20220303-1.txt",
		},
		{
			name: "同一の日付の場合、最新のバージョンのファイルを取得する",
			inputNames: []string{
				"20220301-3.txt", "20220301-1.txt", "20220301-2.txt",
			},
			wantName: "20220301-3.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files := make([]fs.FileInfo, len(tt.inputNames), len(tt.inputNames))
			for i, name := range tt.inputNames {
				files[i] = NewDummyFileInfo(name)
			}
			l := txbuf_file.Latest(files)
			if l.Name() != tt.wantName {
				t.Errorf("予期せぬ値 期待する値: %s, 実際の値: %s", tt.wantName, l.Name())
			}
		})
	}
}
