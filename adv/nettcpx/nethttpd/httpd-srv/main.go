package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"x-learn/adv/nettcpx/nethttpd"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome Simple HTTP: %v", time.Now())
	})

	s := &http.Server{
		Addr:           nethttpd.NetHttpdAddr,
		Handler:        mux,
		ReadTimeout:    time.Second,
		WriteTimeout:   time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Infof("listen on:%v", nethttpd.NetHttpdAddr)
	log.Fatal(s.ListenAndServe())
}
