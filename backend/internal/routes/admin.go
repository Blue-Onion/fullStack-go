package routes

import "net/http"

func AdminRoutes(mux *http.ServeMux){
	authMux:=http.NewServeMux()
	authMux.HandleFunc("/hello",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	mux.Handle("/auth/",http.StripPrefix("/auth",authMux))
}