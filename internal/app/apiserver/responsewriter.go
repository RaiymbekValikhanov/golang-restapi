package apiserver

import (
	"net/http"
)

type responsewriter struct {
	http.ResponseWriter
	code int
}

func (w*responsewriter) WriteHeader(code int) {
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}