package routes

import "net/http"

func AuthRoutes(mux *http.ServeMux){
	authMux:=http.NewServeMux()
	authMux.HandleFunc("/hello",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	mux.Handle("/admin/",http.StripPrefix("/admin",authMux))
}