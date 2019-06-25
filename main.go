package main

import (
	"fmt"
	"net/http"
)

func logRequest(w http.ResponseWriter, r *http.Request) {

	s := fmt.Sprintf("%v: %v %v", r.Method, r.URL.Path, r.Proto)
	s = fmt.Sprintf("%v\nHost: %v", s, r.Host)
	s = fmt.Sprintf("%v\nRemoteAddr: %v", s, r.RemoteAddr)
	s = fmt.Sprintf("%v\nRemoteURI: %v", s, r.RequestURI)
	for k, v := range r.Header {
		s = fmt.Sprintf("%v\nHeader: %v -> %v", s, k, v)
	}
	s = fmt.Sprintf("%v\nHeaders: %v", s, r.Header)
	s = fmt.Sprintf("%v\nUserAgent: %v", s, r.UserAgent())
	s = fmt.Sprintf("%v\nTransfer Encoding: %v", s, r.TransferEncoding)

	// b, _ := httputil.DumpRequest(r, true)

	fmt.Println("------------------------\n" + string(s) + "\n---------------------")

	w.WriteHeader(200)
	w.Write([]byte(s))
}
func main() {
	http.HandleFunc("/", logRequest)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}