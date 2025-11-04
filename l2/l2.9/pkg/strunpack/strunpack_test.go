package strunpack

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"simple repeat", "a4bc2d5e", "aaaabccddddde", false},
		{"no digits", "abcd", "abcd", false},
		{"only digits", "45", "", true},
		{"empty string", "", "", false},
		{"escaped digits", `qwe\4\5`, "qwe45", false},
		{"escaped single digit", `qwe\45`, "qwe44444", false},
		{"ends with slash", `abc\`, "", true},
		{"starts with digit", "2abc", "", true},
		{"zero in", "a0b2", "bb", false},
		{"utf-8", "ф3в4", "фффвввв", false},
		{"10", "a10b2", "aaaaaaaaaabb", false},
		{"punctuations", "a^$b2", "", true},
		{"escaped digits #2", `\2`, "2", false},
		{"escaped letter", `\a2`, "aa", false},
		{"digit 1", `ab1d2`, "abdd", false},
		{"one digit at the end", `abc5`, "abccccc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StrUnpack(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("StrUnpack(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StrUnpack(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
