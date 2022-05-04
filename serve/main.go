package main

import (
	"esp_webrtc/routers"
	"net/http"
)

func main() {
	endPoint := "127.0.0.1:8000"
	routers := routers.InitRouter()
	server := &http.Server{
		Addr:    endPoint,
		Handler: routers,
	}
	server.ListenAndServe()
}
