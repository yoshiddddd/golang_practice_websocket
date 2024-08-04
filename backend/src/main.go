package main


import (
	"fmt"
	"net/http"
	"log"
	"github.com/yoshiddddd/golang_practice_websocket/backend/src/handler"
)

func main(){
	http.HandleFunc("/ws", handler.NewWebsocketHandler().Handle)
	port := 8080;
	log.Printf("Listening on port %d", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}