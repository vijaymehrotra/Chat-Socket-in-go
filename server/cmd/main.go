package main

import (
	"github.com/vijaymehrotra/go-next-ts_chat/db"
	"github.com/vijaymehrotra/go-next-ts_chat/routes"
	"github.com/vijaymehrotra/go-next-ts_chat/ws"
)

func main() {
	db.InitilizeDB()

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()
	
	routes.SetupRoutes(wsHandler)
}
