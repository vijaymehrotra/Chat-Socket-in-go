package ws

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if _, ok := h.Rooms[client.RoomID]; ok {
				r := h.Rooms[client.RoomID]
				if _, ok := r.Clients[client.ID]; !ok {
					r.Clients[client.ID] = client
				}
			}
		case client := <-h.Unregister:
			if _, ok := h.Rooms[client.RoomID]; ok {
				r := h.Rooms[client.RoomID]
				if _, ok := r.Clients[client.ID]; ok {
					if len(h.Rooms[client.RoomID].Clients) != 0 {
						message := &Message{
							Content: client.Username + " has left the room",
							Username: client.Username,
							RoomID: client.RoomID,
						}
						h.Broadcast <- message
					}

					delete(r.Clients, client.ID)
					close(client.Message)
				}
			}
		case message := <-h.Broadcast:
			if _, ok := h.Rooms[message.RoomID]; ok {
				for _, client := range h.Rooms[message.RoomID].Clients {
					select {
					case client.Message <- message:
					default:
						close(client.Message)
						delete(h.Rooms[message.RoomID].Clients, client.ID)
					}
				} 
			}
		}
	}
}
