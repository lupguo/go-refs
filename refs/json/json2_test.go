package json

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	cases := []interface{}{
		"",
		"{}",
		true,
		false,
		0,
		1,
		1.1,
		1.2,
		[]int{100, 200},
		[]float64{33.01, 22.20},
		struct {
			id   uint64
			User string
		}{201, "clark"},
	}
	for _, c := range cases {
		marshal, err := json.Marshal(c)
		assert.Nil(t, err)
		t.Logf("marshal(%v)=%s", c, marshal)
	}
}

func TestUnmarshalCases(t *testing.T) {
	cases := []interface{}{
		"",
		"{}",
		true,
		false,
		0,
		1,
		1.1,
		1.2,
		[]int{100, 200},
		[]float64{33.01, 22.20},
		struct {
			id   uint64
			User string
		}{201, "clark"},
	}
	for _, c := range cases {
		v := ""
		err := json.Unmarshal(c.([]byte), &v)
		assert.Nil(t, err)
		t.Logf("marshal(%v)=%s", c, v)
	}
}
