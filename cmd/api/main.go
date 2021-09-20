package main

import (
	"github.com/SongOf/edge-storage-agent/internal/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.DefaultHandler)
	http.HandleFunc("/health", handler.HealthHandler)
	http.HandleFunc("/system/stat", handler.SystemHandler)
	http.HandleFunc("/volume/stat", handler.VolumeHandler)

	err := http.ListenAndServe(":36000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
