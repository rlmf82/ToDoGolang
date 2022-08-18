package controller

import "net/http"

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Execute())
	mux.HandleFunc("/ping", Ping())
	return mux
}
