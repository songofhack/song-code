package main

import (
	"fmt"
	http "net/http"
	"strconv"
	"strings"
	"time"
)

type MyHandler struct{}

func main() {

	srv := &http.Server{
		Addr:         "192.168.21.45:80",
		Handler:      &MyHandler{},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
func (*MyHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	localip := "192.168.21.45"
	localipduan := strings.Split(localip, ".")
	for i := 0; i < 4; i++ {
		localipduan[i] = strings.Trim(localipduan[i], " ")

	}
	remoteip := req.RemoteAddr
	remoteipduan := strings.Split(remoteip, ".")
	for i := 0; i < 4; i++ {
		remoteipduan[i] = strings.Trim(remoteipduan[i], " ")
	}
	if strings.EqualFold(localipduan[0], remoteipduan[0]) && strings.EqualFold(localipduan[1], remoteipduan[1]) && strings.EqualFold(localipduan[2], remoteipduan[2]) {
		fmt.Fprintf(rw, req.RemoteAddr+"\n")
		fmt.Fprintf(rw, req.RequestURI+"\n")
		fenge := strings.Split(remoteipduan[3], ":")
		fenge_int, _ := strconv.Atoi(strings.Trim(fenge[0], " "))
		fmt.Println(fenge_int)
		fmt.Fprintf(rw, req.Host)

	} else {
		fmt.Fprintf(rw, "no! failed")

	}
	fmt.Println(req.RemoteAddr)

	return

}
