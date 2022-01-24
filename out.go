package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func doRequest(rw http.ResponseWriter, r *http.Request) {
	l := len(r.RequestURI)
	if l > 80 {
		l = 80
	}
	log.Print("IN->", r.RemoteAddr, r.RequestURI[0:l])

	if strings.Contains(r.RequestURI, "/nioproxy/statik/images") {
		// if strings.Contains(r.RequestURI, "/nioproxy/statik/images/down2-icon.png") {
		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "image/png")
		rw.Write(bufImage)
		return
	}

	if strings.Contains(r.RequestURI, "/cert/check") {
		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "text/html")
		rw.Write(bufHtml)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "text/html")
	rw.Write([]byte(""))
}

var bufImage, bufHtml []byte

func main() {
	log.SetFlags(log.Ltime)
	log.Println("Begin....")

	bufImage, _ = ioutil.ReadFile("out.png")
	bufHtml, _ = ioutil.ReadFile("out.html")

	http.HandleFunc("/", doRequest)

	http.ListenAndServe("0.0.0.0:2281", nil)
}
