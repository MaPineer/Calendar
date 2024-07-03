package utils

import (
	"Calendar/models"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var clients = make(map[*websocket.Conn]string) // 维护websocket与对应用户的映射
var Broadcast = make(chan models.Reminder)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// HandleConnections 对所有的websocket连接进行注册(绑定到map)
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Error upgrading HTTP connection to WebSocket: %v", err)
	}
	defer ws.Close()

	userID := r.URL.Query().Get("creator_id")
	if userID == "" {
		log.Printf("userID is required")
		ws.Close()
		return
	}
	if err != nil {
		log.Printf("Invalid userID")
		ws.Close()
		return
	}

	mu.Lock()
	clients[ws] = userID
	mu.Unlock()

	for {
		var reminder models.Reminder
		err := ws.ReadJSON(&reminder)
		if err != nil {
			log.Printf("Error reading JSON from WebSocket: %v", err)
			mu.Lock()
			delete(clients, ws)
			mu.Unlock()
			break
		}
	}
}

// HandleMessages 进行处理所有的websocket
func HandleMessages() {
	for {
		reminder := <-Broadcast
		mu.Lock()
		for client, userID := range clients {
			if userID == reminder.CreatorID {
				err := client.WriteJSON(reminder)
				if err != nil {
					log.Printf("Error writing JSON to WebSocket: %v", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
		mu.Unlock()
	}
}

func Start() {
	go HandleMessages()
}
