package main

import (
	"fmt"
	"io"
	"net/http"
)

type Handler int8

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)

	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	s := &http.Server{
		Addr:    ":8080",
		Handler: new(Handler),
	}
	s.ListenAndServe()
}
