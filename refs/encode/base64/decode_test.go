package base64

import (
	"encoding/base64"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		encs string
		want string
	}{
		{"t1", "Y29udGFjdHNBZG1pbjpEZW1vQDEyMwo=", "contactsAdmin:Demo@123"}, // 多了一个\n
		{"t2", "Y29udGFjdHNBZG1pbjpEZW1vQDEyMw==", "contactsAdmin:Demo@123"}, // 正常的
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := base64.StdEncoding.DecodeString(tt.encs)
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != tt.want {
				t.Fatalf("base64 decode(%s) got=%s, but want=%s", tt.encs, got, tt.want)
			}
			t.Logf("decode str ok: %s", got)
		})
	}

}
