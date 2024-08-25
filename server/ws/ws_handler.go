package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	Hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{
		Hub: hub,
	}
}

type CreateRoomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoomHandler(c *gin.Context) {
	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	h.Hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room created successfully", "room": req})
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Change this when Implementing Frontend
	},
}

func (h *Handler) JoinRoomHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	roomID := c.Param("room_id")
	clientID := c.Query("user_id")
	username := c.Query("username")

	log.Printf("Client %s joined room %s", clientID, roomID)

	client := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		RoomID:   roomID,
		ID:       clientID,
		Username: username,
	}

	message := &Message{
		Content:  username + " has joined the room",
		Username: username,

		RoomID:   roomID,
	}

	h.Hub.Register <- client

	h.Hub.Broadcast <- message

	go client.ReadMessage(h.Hub)
	go client.WriteMessage()
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {
	var rooms []RoomRes
	for _, r := range h.Hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID: r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{"rooms": rooms})
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
func (h *Handler) GetClients(c *gin.Context) {
	var clients []ClientRes
	roomID := c.Param("room_id")

	if _, ok := h.Hub.Rooms[roomID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found"})
		return
	}

	for _, client := range h.Hub.Rooms[roomID].Clients {
		clients = append(clients, ClientRes{
			ID: client.ID,
			Username: client.Username,
		})
	}
	c.JSON(http.StatusOK, clients)
}
