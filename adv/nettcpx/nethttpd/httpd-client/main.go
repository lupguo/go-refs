package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"x-learn/adv/nettcpx/nethttpd"
)

func main() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    5 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second,
	}

	resp, err := client.Get(fmt.Sprintf("http://localhost%v", nethttpd.NetHttpdAddr))
	if err != nil {
		log.Errorf("err: %v", err)
		return
	}
	log.Infof("http resp: %v", resp)
}
