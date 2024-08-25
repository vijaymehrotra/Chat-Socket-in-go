package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/vijaymehrotra/go-next-ts_chat/db"
	"github.com/vijaymehrotra/go-next-ts_chat/models"
)

type formData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var data formData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid form data"})
		return
	}
	if data.Username == "" || data.Email == "" || data.Password == "" {
		c.JSON(400, gin.H{"error": "Empty fields not allowed"})
		return
	}

	var user models.User
	user.Username = data.Username
	user.Email = data.Email
	user.Password = HashPassword(data.Password)


	err = db.DB.Create(&user).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "Error registering user" , "error": err})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully", "user": user})
}

func Login(c *gin.Context) {
	var data formData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid form data"})
		return
	}
	var user models.User
	fmt.Println(data)
	err = db.DB.Where("email = ?", data.Email).Or("username = ?" , data.Username).First(&user).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	if !CheckPasswordHash(data.Password, user.Password) {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}

	token, err := GenerateToken(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error generating token Or Token Expired"})
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	c.JSON(200, gin.H{"message": "User logged in successfully", "user": user , "token": token})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	db.DB.Delete(&user)
	c.JSON(200, gin.H{"message": "User deleted successfully","user": user})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	err := db.DB.Find(&users).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(200, gin.H{"users": users})
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "",-1 , "","" , false, true)
	c.JSON(200, gin.H{"message": "User logged out successfully"})
}

func User(c *gin.Context) {
    fmt.Println("Request to Get User")

    // Get the "jwt" cookie by name
    cookie, err := c.Cookie("jwt")
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "No JWT token provided Or No user Logged in"})
        return
    }

    // Parse the JWT token
    token, err := jwt.ParseWithClaims(cookie, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })
    if err != nil || !token.Valid {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing token or token is invalid"})
        return
    }

    // Extract and process claims as before
    claims, ok := token.Claims.(*customClaims)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse claims"})
        return
    }

	id := claims.UserID

    // Fetch the user from the database
    user := models.User{ID: id}
    if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "This was the user logged in", "user": user})
}
