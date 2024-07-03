package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"testing"
	"time"
)

func TestWebSocket(t *testing.T) {
	// WebSocket服务器地址
	wsURL := "ws://localhost:8080/ws?creator_id=user2"

	// 连接到WebSocket服务器
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer ws.Close()

	// 读取WebSocket服务器的响应
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			t.Fatalf("Failed to read message: %v", err)
		}
		receivedMessage := string(message)
		fmt.Println(receivedMessage)

		// 等待一秒，确保所有消息都被处理
		time.Sleep(1 * time.Second)
	}

}
