package grafana

import (
	"testing"
)

type DisplayName struct {
	Matcher struct {
		Id      string `json:"id"`
		Options string `json:"options"`
	} `json:"matcher"`
	Properties []struct {
		Id    string `json:"id"`
		Value string `json:"value"`
	} `json:"properties"`
}

type overWrite []*DisplayName

func TestGenDisplayName(t *testing.T) {

}
