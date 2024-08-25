package routes

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vijaymehrotra/go-next-ts_chat/controllers"
	"github.com/vijaymehrotra/go-next-ts_chat/ws"
)

func SetupRoutes(wsHandler *ws.Handler) {
	router := gin.Default()
	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Create an authentication group
	auth := router.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)
	auth.GET("/users", controllers.GetUsers)          
	auth.GET("/user", controllers.User)               
	auth.DELETE("/delete/:id", controllers.DeleteUser)
	auth.POST("/logout", controllers.Logout)

	// Create a chat group
	ws := router.Group("/ws")
	ws.POST("/create-room", wsHandler.CreateRoomHandler)
	ws.GET("/join-room/:room_id", wsHandler.JoinRoomHandler) // websocker request
	ws.GET("/get-rooms", wsHandler.GetRooms)
	ws.GET("/get-clients/:room_id", wsHandler.GetClients)



	err := router.Run(":8081")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server running on port 8081")
}