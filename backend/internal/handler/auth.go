package handler

import "net/http"

func Registering() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello"))
	}

}
