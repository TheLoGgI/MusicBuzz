package ws

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type ClientInfo struct {
	UserId   uuid.UUID `json:"userId"`
	Name     string    `json:"name"`
	GroupID  string    `json:"groupId"`
	IsLeader bool      `json:"isLeader"`
}

type ClientWS struct {
	Connection *websocket.Conn `json:"-"`
	ClientInfo
}

func (client *ClientWS) toJson() string {
	clientJson, _ := json.Marshal(client)
	return string(clientJson)
}

type ServerWS struct {
	groups map[string][]*ClientWS
}

// Get total connections in the server
func (server *ServerWS) GetTotalConnections() int {
	fmt.Println("Groups", server.groups)
	var totalCount = 0
	for index := range server.groups {
		// fmt.Println(index, group)
		for index, group := range server.groups[index] {
			fmt.Println(index, group)
			totalCount++
		}
	}
	return totalCount
}

func (server *ServerWS) MarshalGroup(groupId string) []byte {
	var group = server.groups[groupId]
	marshaledGroup, _ := json.Marshal(group)
	return marshaledGroup
}

func (server *ServerWS) removeClientFromGroup(groupId string, clientId uuid.UUID) []*ClientWS {
	var foundIndex = -1
	var slice = server.groups[groupId]
	for i, client := range slice {
		if client.UserId == clientId {
			foundIndex = i
			break
		}
	}

	return append(slice[:foundIndex], slice[foundIndex+1:]...)
}

func (server *ServerWS) GetGroupConnectionsJSON(groupId string) string {
	connectionsMap := make(map[uuid.UUID]ClientWS)
	for _, client := range server.groups[groupId] {
		connectionsMap[client.UserId] = *client
	}
	connections, _ := json.Marshal(connectionsMap)

	return string(connections)
}

func CreateServer() *ServerWS {
	return &ServerWS{
		// conns:    make(map[*websocket.Conn]ClientWS),
		groups: make(map[string][]*ClientWS),
	}
}

func (server *ServerWS) AddWebSocketHandler(ws *websocket.Conn) {
	var queryParms = ws.Request().URL.Query()
	var name = queryParms.Get("name")
	var groupID = queryParms.Get("groupId")

	var client = ClientWS{
		Connection: ws,
		ClientInfo: ClientInfo{
			UserId:   uuid.New(), // TODO: add verication to users e.g autentication
			Name:     name,
			GroupID:  groupID,
			IsLeader: false,
		},
	}

	if client.GroupID == "" {
		groupID = uuid.New().String()
		client.GroupID = groupID
		client.IsLeader = true
	}

	if name == "" {
		fmt.Println("Name is missing")

		ws.Close()
		return
	}

	// fmt.Println("New Connection", client.Name, client.UserId, client.GroupID)

	server.groups[groupID] = append(server.groups[groupID], &client)

	fmt.Println("Number of connections in group", len(server.groups[groupID]))
	fmt.Println("Total Connections", server.GetTotalConnections())

	if len(server.groups[groupID]) > 1 {
		var leader = server.groups[groupID][0]
		websocket.Message.Send(leader.Connection, "MEMBER_ADDED")

		// Send the group members to the new member
		err := websocket.Message.Send(ws, "CONNECTION!")
		// websocket.Message.Send(ws, server.GetGroupConnectionsJSON(groupID))
		fmt.Println("Error", err)
	}

	// Keep websocket connection alive
	for {
		var msg string
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			fmt.Println("Error", err)
			server.groups[groupID] = server.removeClientFromGroup(groupID, client.UserId)
			return
		}
	}

	// TODO: add a mutex
}

type GroupHandlerBody struct {
	GroupID string `json:"groupId"`
	Msg     string `json:"message"`
}

type GetGroupHandlerBody struct {
	GroupID string `json:"groupId" form:"groupId"`
}

func (server *ServerWS) SendGroupMessageHandler(w http.ResponseWriter, r *http.Request) {

	var body GroupHandlerBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if body.GroupID == "" {
		http.Error(w, "Group ID is missing", http.StatusBadRequest)
		return
	}

	if len(server.groups[body.GroupID]) == 0 {
		http.Error(w, "No connections", http.StatusNoContent)
		return
	}

	fmt.Println("Sending message to all in group", len(server.groups[body.GroupID]))
	for _, client := range server.groups[body.GroupID] {
		websocket.Message.Send(client.Connection, body.Msg)
	}

}

func (server *ServerWS) GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Group Members")
	var body GetGroupHandlerBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(body)
	fmt.Println("Group ID", body.GroupID, server.groups)
	if body.GroupID == "" {
		http.Error(w, "Group ID is missing", http.StatusBadRequest)
		return
	}

	var marshaledGroup = server.MarshalGroup(body.GroupID)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaledGroup)
}
