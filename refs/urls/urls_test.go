package urls

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

func TestUrlParams(t *testing.T) {
	m, err := url.ParseQuery(`x=1&y=2&y=3;z`)
	if err != nil {
		log.Fatal(err)
	}
	m.Add("status", "1")
	m.Add("status", "2")
	m.Add("status", "3")

	for _, val := range []string{"1", "2", "3"} {
		fmt.Printf("val=>%s", val)
	}

	m.Add("xx", "a=10&a=20")
	t.Logf("encode: %s", m.Encode())

	envs := []string{"test", "dev", "prod"}
	m["envs"] = append(m["envs"], envs...)
	t.Logf("encode: %s", m.Encode())

}
