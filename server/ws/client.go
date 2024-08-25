package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Message chan *Message
	ID string `json:"id"`
	Username string `json:"username"`
	RoomID string `json:"room_id"`
}

type Message struct {
	Content string `json:"content"`
	Username string `json:"username"`
	RoomID string `json:"room_id"`
}

func (c *Client) WriteMessage() {
	defer c.Conn.Close()

	for message := range c.Message {
		err := c.Conn.WriteJSON(message)
		if err != nil {
			break
		}  
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func () {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_,m,err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		
		msg := Message{
			Content: string(m),
			Username: c.Username,
			RoomID: c.RoomID,
		}

		hub.Broadcast <- &msg
	}
}