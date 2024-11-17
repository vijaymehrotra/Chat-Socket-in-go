# WebSocket Chat Application in Go
![go-websocket](https://github.com/user-attachments/assets/15c735d2-83e7-4f1a-b32c-3c097071418d)
## Overview

This project is a real-time chat application built with Go using the Gin framework. It features user authentication via JWT tokens, WebSockets for real-time communication, and PostgreSQL for database management. Users can create and join chat rooms, where they can interact with multiple participants in real-time.

![image](https://github.com/user-attachments/assets/e43736ed-f5f6-4402-a163-ce4f9d91bc19)

![image](https://github.com/user-attachments/assets/71b23058-3612-4332-bb2a-7e48d6995c0f)

![image](https://github.com/user-attachments/assets/7f315ce3-a58a-435f-862c-bc8027d84ad8)

## Features

  User Registration & Login: Secure authentication using JWT tokens.
  
  Real-Time Messaging: WebSockets allow real-time chat between users.
  
  Room Management: Users can create and join chat rooms.
  
  Session Management: Cookie-based sessions to track user activity.
  
  PostgreSQL Integration: Persistent storage for users and chat room data.
    

## Tech Stack

  Go: Backend programming language.
  
  Gin: Web framework for building the API.
  
  WebSockets: Real-time communication between clients.

  JWT: Token-based authentication.
  
  PostgreSQL: Database to store user and room data.
  
  Cookie Sessions: Managing user sessions through cookies.

  ![image](https://github.com/user-attachments/assets/477f28d0-609e-434a-b6db-6e0d5a70dd86)


## Project Setup
### Prerequisites

  Go: Version 1.16 or higher
  
  PostgreSQL: Installed and running locally or on a cloud platform
  
  Git: For version control and repository management
    

## Installation

  Clone the repository:
    
    git clone https://github.com/your-username/Chat-Socket-in-go.git
    
    cd websocket-chat-app

  Install dependencies:
   
      go mod tidy

  Configure environment variables in a .env file:

    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_username
    DB_PASSWORD=your_password
    DB_NAME=chat_app_db
    JWT_SECRET=your_jwt_secret

  Set up the PostgreSQL database:

    psql -U postgres -c "CREATE DATABASE chat_app_db;"

  Run the application:

    go run main.go

  Here's a template for your GitHub README file for the WebSocket-based chat application project in Go:
WebSocket Chat Application in Go
Overview

This project is a real-time chat application built with Go using the Gin framework. It features user authentication via JWT tokens, WebSockets for real-time communication, and PostgreSQL for database management. Users can create and join chat rooms, where they can interact with multiple participants in real-time.
Features

    User Registration & Login: Secure authentication using JWT tokens.
    Real-Time Messaging: WebSockets allow real-time chat between users.
    Room Management: Users can create and join chat rooms.
    Session Management: Cookie-based sessions to track user activity.
    PostgreSQL Integration: Persistent storage for users and chat room data.

Tech Stack

    Go: Backend programming language.
    Gin: Web framework for building the API.
    WebSockets: Real-time communication between clients.
    JWT: Token-based authentication.
    PostgreSQL: Database to store user and room data.
    Cookie Sessions: Managing user sessions through cookies.

Project Setup
Prerequisites

    Go: Version 1.16 or higher
    PostgreSQL: Installed and running locally or on a cloud platform
    Git: For version control and repository management

Installation

    Clone the repository:

    git clone https://github.com/your-username/websocket-chat-app.git
    cd websocket-chat-app

## install dependencies:

    go mod tidy

## Configure environment variables in a .env file:

    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_username
    DB_PASSWORD=your_password
    DB_NAME=chat_app_db
    JWT_SECRET=your_jwt_secret

## Set up the PostgreSQL database:


    psql -U postgres -c "CREATE DATABASE chat_app_db;"

Run the application:


    go run main.go

# API Endpoints
## Authentication

    auth.POST("/register", controllers.Register)
  	auth.POST("/login", controllers.Login)
  	auth.GET("/users", controllers.GetUsers)          
  	auth.GET("/user", controllers.User)               
  	auth.DELETE("/delete/:id", controllers.DeleteUser)
  	auth.POST("/logout", controllers.Logout)

## Chat Rooms

    ws.POST("/create-room", wsHandler.CreateRoomHandler)
  	ws.GET("/join-room/:room_id", wsHandler.JoinRoomHandler) // websocker request
  	ws.GET("/get-rooms", wsHandler.GetRooms)
  	ws.GET("/get-clients/:room_id", wsHandler.GetClients)

## WebSockets

    Join Chat Room via WebSocket: ws://localhost:8081/ws/join-room/{roomId}?user_id={userId}&username={username}

## Project Structure

    ├── controller
    │   └── chat_controller.go
    ├── middleware
    │   └── auth_middleware.go
    ├── models
    │   └── user.go
    │   └── room.go
    ├── websocket
    │   └── websocket_handler.go
    ├── main.go
    └── .env

## Future Improvements

  Implementing typing indicators and online/offline status.
  Adding message persistence in PostgreSQL.
  Enhancing UI with a frontend framework such as React or Vue.js.
    
