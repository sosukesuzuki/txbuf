package txbuf_file_test

import (
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
