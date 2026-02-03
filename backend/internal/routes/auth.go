package routes

import "net/http"

func AuthRoutes(mux *http.ServeMux){
	authMux:=http.NewServeMux()
	authMux.HandleFunc("/login",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	authMux.HandleFunc("/regsiter",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	authMux.HandleFunc("/signOut",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	
	mux.Handle("/auth/",http.StripPrefix("/auth",authMux))
}