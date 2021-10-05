package main

import (
	"bytes"
	"fmt"
	"lesson_practice/lesson_4/util"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
)

type ResponseWithRecorder struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (rec *ResponseWithRecorder) WriteHeader(statusCode int) {
	rec.ResponseWriter.WriteHeader(statusCode)
	rec.statusCode = statusCode
}

func (rec *ResponseWithRecorder) Write(d []byte) (n int, err error) {
	n, err = rec.ResponseWriter.Write(d)
	if err != nil {
		return
	}
	rec.body.Write(d)
	return
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	version := os.Getenv("VERSION")
	wc := &ResponseWithRecorder{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		body:           bytes.Buffer{},
	}
	for k, v := range header {
		w.Header().Add("Response_Header_"+k, fmt.Sprintf("%s", v))
	}
	wc.Header().Add("Response_Header_VERSION", version)
	ip := util.GetIP(r)
	glog.V(0).Info(fmt.Sprintf("IP: %s, statuscode: %d", ip, wc.statusCode))

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

}

func main() {
	glog.V(0).Info("Starting http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthHandler)
	err := http.ListenAndServe(":81", nil)
	if err != nil {
		log.Fatal(err)
	}

}
