package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Blue-Onion/fullStack-go/internal/routes"
)

func main() {
	log.Println("Server Started")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	server := http.Server{
		Addr:    "localhost:6160",
		Handler: mux,
	}
	routes.AuthRoutes(mux)
	routes.AdminRoutes(mux)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Println(err.Error())
		}
	}()
	<-done
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Server Shutdown Gracefully")
}
