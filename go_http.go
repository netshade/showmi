package go_http

import (
	"io"
	"net/http"
)

type Request struct {
	Method string
	Path   string
	Body   []byte
}

type Response struct {
	Status int64
	Body   []byte
}

func Serve(addr string, handle func(Request) Response) error {
	return http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		resp := handle(Request{Method: r.Method, Path: r.URL.Path, Body: body})
		w.WriteHeader(int(resp.Status))
		w.Write(resp.Body)
	}))
}
