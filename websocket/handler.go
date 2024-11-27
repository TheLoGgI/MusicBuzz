package ws

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type ClientWS struct {
	Ip       *websocket.Conn
	Name     string
	GroupID  string
	IsLeader bool
}

type ServerWS struct {
	conns    map[*websocket.Conn]ClientWS
	groupIds map[string][]*websocket.Conn
}

func CreateServer() *ServerWS {
	return &ServerWS{
		conns:    make(map[*websocket.Conn]ClientWS),
		groupIds: make(map[string][]*websocket.Conn),
	}
}

func (server *ServerWS) AddWebSocketHandler(ws *websocket.Conn) {
	var queryParms = ws.Request().URL.Query()
	var name = queryParms.Get("name")
	var groupID = queryParms.Get("groupId")
	var client = ClientWS{
		Ip:   ws,
		Name: name,
	}

	if groupID == "" {
		groupID = uuid.New().String()
		client.GroupID = groupID
		client.IsLeader = true
		server.groupIds[groupID] = []*websocket.Conn{client.Ip}
	} else {
		fmt.Println("Have Group ID", groupID)
		server.groupIds[groupID] = append(server.groupIds[groupID], client.Ip)
	}

	if name == "" {
		fmt.Println("Name is missing")

		ws.Close()
		return
	}

	fmt.Println("New Connection", ws.RemoteAddr())
	server.conns[ws] = client

	fmt.Println("Number of connections", len(server.conns))
	fmt.Println("Connections", server.conns)
	for {

		var msg string
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			fmt.Println("Error", err)
			delete(server.conns, ws)
			return
		}

		for conn := range server.conns {
			if conn != ws {
				websocket.Message.Send(conn, msg)
			}
		}
	}

	// Todo add a mutex

	// ws.Write([]byte("Hello, World!"))
	// ws.WriteClose(200)
}

type GroupHandlerBody struct {
	GroupID string `json:"groupId"`
	Msg     string `json:"message"`
}

func (server *ServerWS) SendGroupMessageHandler(w http.ResponseWriter, r *http.Request) {

	var body GroupHandlerBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(body)
	fmt.Println("Group ID", body.GroupID, server.groupIds)
	if body.GroupID == "" {
		http.Error(w, "Group ID is missing", http.StatusBadRequest)
		return
	}

	if len(server.groupIds[body.GroupID]) == 0 {
		http.Error(w, "No connections", http.StatusNoContent)
		return
	}

	fmt.Println("Sending message to all in group", len(server.groupIds[body.GroupID]))
	for _, conn := range server.groupIds[body.GroupID] {
		websocket.Message.Send(conn, body.Msg)
	}

}
