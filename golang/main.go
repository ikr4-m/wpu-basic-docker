package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"time"

	_ "embed"
)

//go:embed assets/video.mp4
var videoFile []byte

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("REQUEST: %v\n", r.URL.Path)
		http.ServeContent(w, r, "video.mp4", time.Now(), bytes.NewReader(videoFile))
	})

	log.Printf("Listening to 0.0.0.0:%v\n", port)
	http.ListenAndServe(":"+port, nil)
}
