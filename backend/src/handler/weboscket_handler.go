// package handler

// import (
// 	"log"
// 	"net/http"
// 	"github.com/gorilla/websocket"
// )

// type WebsocketHandler struct {}

// func NewWebsocketHandler() *WebsocketHandler{
// 	return &WebsocketHandler{}
// }

// func (h *WebsocketHandler) Handle(w http.ResponseWriter, r *http.Request){
// 	upgrader := &websocket.Upgrader{
// 		CheckOrigin: func(r *http.Request) bool {
// 			 return true
// 		},
// 	}
// 	_, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Printf("Failed to upgrade connection: %v", err)
// 		return
// 	}
// }

package handler

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

type WebsocketHandler struct{}

func NewWebsocketHandler() *WebsocketHandler {
	return &WebsocketHandler{}
}

func (h *WebsocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		log.Printf("Received message: %s", message)

		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Printf("Error writing message: %v", err)
			break
		}
	}
}
