package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var mutex = &sync.Mutex{}
var messagesFile = "messages.json"

type Message struct {
	ID          int    `json:"id"`
	UserID      string `json:"userId"`
	Timestamp   int64  `json:"timestamp"`
	MessageType int    `json:"messageType"`
	Content     string `json:"content"`
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Register the new client
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	// Send existing messages to the new client
	messages, err := loadMessages()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, msg := range messages {
		msgJSON, err := json.Marshal(msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = conn.WriteMessage(websocket.TextMessage, msgJSON)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Handle WebSocket messages here
	for {
		messageType, content, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			break
		}
		fmt.Printf("In server: Received message: %s\n", content)

		// Generate a random ID for the message
		messageID := rand.Int()

		// Get the current timestamp
		timestamp := time.Now().Unix()

		// Create the message object
		msg := Message{
			ID:          messageID,
			UserID:      fmt.Sprintf("%p", conn),
			Timestamp:   timestamp,
			MessageType: messageType,
			Content:     string(content),
		}

		// Send the message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		// Save the message to the file
		saveMessage(msg)

		// Send it out to every client that is currently connected
		mutex.Lock()
		for client := range clients {
			msgJSON, err := json.Marshal(msg)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = client.WriteMessage(websocket.TextMessage, msgJSON)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}

func saveMessage(msg Message) {
	mutex.Lock()
	defer mutex.Unlock()

	// Read existing messages from the file
	messages, err := loadMessages()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Append the new message
	messages = append(messages, msg)

	// Write the updated messages to the file
	content, err := json.Marshal(messages)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(messagesFile, content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func loadMessages() ([]Message, error) {
	content, err := os.ReadFile(messagesFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Message{}, nil
		}
		return nil, err
	}

	var messages []Message
	err = json.Unmarshal(content, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func main() {
	// Load existing messages from the file
	messages, err := loadMessages()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the loaded messages to the broadcast channel
	go func() {
		for _, msg := range messages {
			broadcast <- msg
		}
	}()

	http.HandleFunc("/ws", handleWebSocket)
	go handleMessages()
	fmt.Println("WebSocket server is running on :8032/ws")
	if err := http.ListenAndServe(":8032", nil); err != nil {
		fmt.Println(err)
	}
}
