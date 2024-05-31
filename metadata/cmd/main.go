package cmd

import (
	"log"
	"moviesexample.com/metadata/internal/controller/metadata"
	httphandler "moviesexample.com/metadata/internal/handler/http"
	"moviesexample.com/metadata/internal/repository/memory"
	"net/http"
)

func main() {
	log.Println("Starting the movie metadata service")
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
