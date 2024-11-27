package main

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"

// 	"github.com/gorilla/websocket"
// )

// // https://medium.com/@dogabudak/websockets-with-go-with-tests-88f47837a4e5

// func TestWebsocket() {
// 	// Create a test server and client
// 	server := httptest.NewServer(http.HandlerFunc(HandleWebSocket))
// 	defer server.Close()

// 	url := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
// 	client, _, err := websocket.DefaultDialer.Dial(url, nil)
// 	if err != nil {
// 		t.Fatalf("could not connect to WebSocket server: %v", err)
// 	}
// 	defer client.Close()

// 	// Test sending and receiving messages
// 	message := []byte("hello, world")
// 	err = client.WriteMessage(websocket.TextMessage, message)
// 	if err != nil {
// 		t.Fatalf("could not write message to WebSocket server: %v", err)
// 	}

// 	_, response, err := client.ReadMessage()
// 	if err != nil {
// 		t.Fatalf("could not read message from WebSocket server: %v", err)
// 	}

// 	if string(response) != string(message) {
// 		t.Errorf("expected response '%s', but got '%s'", string(message), string(response))
// 	}

// }
