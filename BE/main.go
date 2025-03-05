package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
)

var (
	upgrader     = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }}
	clients      = make(map[*websocket.Conn]bool)
	broadcast    = make(chan Message, 100)
	mutex        = &sync.Mutex{}
	messagesFile = "messages.json"
)

type Message struct {
	ID        int    `json:"id"`
	UserID    string `json:"userId"`
	Timestamp int64  `json:"timestamp"`
	ServerID  int    `json:"serverId"`
	Content   string `json:"content"`
}

type Payload struct {
	Type     string `json:"type"`
	ServerID int    `json:"serverId"`
	Content  string `json:"content"`
}

type Channel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	ServerID int       `json:"serverId"`
	Name     string    `json:"name"`
	Channels []Channel `json:"channels"`
}

// Fetches the server list
func getServers() ([]Server, error) {
	// Open the servers.json file
	file, err := os.Open("servers.json")
	if err != nil {
		return nil, fmt.Errorf("error opening servers.json: %v", err)
	}
	defer file.Close()

	// Create a struct to hold the servers from the JSON
	var data struct {
		Servers []Server `json:"servers"`
	}

	// Decode the JSON data from the file
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding JSON from servers.json: %v", err)
	}

	// Return the list of servers
	return data.Servers, nil
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			removeClient(conn)
			return
		}

		var payload Payload
		if err := json.Unmarshal(data, &payload); err != nil {
			fmt.Println("JSON unmarshal error:", err)
			continue
		}

		if payload.Type == "get_messages" {
			sendPreviousMessages(conn, payload.ServerID)
			continue
		}

		newMessage := Message{
			ID:        rand.Int(),
			UserID:    fmt.Sprintf("%p", conn),
			Timestamp: time.Now().Unix(),
			ServerID:  payload.ServerID,
			Content:   payload.Content,
		}

		broadcast <- newMessage
	}
}

func handleMessages() {
	for msg := range broadcast {
		saveMessage(msg)
		sendToAllClients(msg)
	}
}

func sendPreviousMessages(conn *websocket.Conn, serverID int) {
	messages, err := loadMessages()
	if err != nil {
		fmt.Println("Error loading messages:", err)
		return
	}

	for _, msg := range messages {
		if msg.ServerID == serverID {
			if err := sendMessage(conn, msg); err != nil {
				return
			}
		}
	}
}

func sendToAllClients(msg Message) {
	mutex.Lock()
	defer mutex.Unlock()

	for client := range clients {
		if err := sendMessage(client, msg); err != nil {
			removeClient(client)
		}
	}
}

func sendMessage(conn *websocket.Conn, msg Message) error {
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return err
	}
	return conn.WriteMessage(websocket.TextMessage, msgJSON)
}

func removeClient(conn *websocket.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	conn.Close()
	delete(clients, conn)
}

func saveMessage(msg Message) {
	mutex.Lock()
	defer mutex.Unlock()

	messages, err := loadMessages()
	if err != nil {
		fmt.Println("Error loading messages for saving:", err)
		return
	}

	messages = append(messages, msg)

	data, err := json.Marshal(messages)
	if err != nil {
		fmt.Println("JSON marshal error while saving:", err)
		return
	}

	if err := os.WriteFile(messagesFile, data, 0644); err != nil {
		fmt.Println("File write error:", err)
	}
}

func loadMessages() ([]Message, error) {
	data, err := os.ReadFile(messagesFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Message{}, nil
		}
		return nil, err
	}

	var messages []Message
	if err := json.Unmarshal(data, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

// New HTTP handler to get the list of servers
func getServersHandler(w http.ResponseWriter, r *http.Request) {
	// Allow cross-origin requests (CORS)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Respond with the server list as JSON
	servers, err := getServers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching servers: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(servers); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding server data: %v", err), http.StatusInternalServerError)
	}
}

func main() {
	if _, err := loadMessages(); err != nil {
		fmt.Println("Error loading initial messages:", err)
	}

	// WebSocket handler
	http.HandleFunc("/ws", handleWebSocket)

	// New endpoint for getting the list of servers
	http.HandleFunc("/get_servers", getServersHandler)

	// Apply CORS to all routes using the gorilla handler
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins, or specify the frontend URL
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
	)(http.DefaultServeMux)

	go handleMessages()

	fmt.Println("WebSocket server is running on :8032/ws")
	fmt.Println("HTTP server is running on :8032/get_servers")
	if err := http.ListenAndServe(":8032", corsHandler); err != nil {
		fmt.Println("Server error:", err)
	}
}
