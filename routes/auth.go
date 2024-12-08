package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"ginSample/database"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key")

// RegisterAuthRoutes godoc
// @Summary      Register authentication routes
// @Description  Sets up authentication endpoints
func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.POST("/register", Register)
	auth.POST("/login", Login)
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register godoc
// @Summary      Register a new user
// @Description  Create a new user account with username and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        registerInput  body      RegisterInput  true  "Register Input"
// @Success      200            {object}  map[string]string
// @Failure      400            {object}  map[string]string
// @Failure      500            {object}  map[string]string
// @Router       /auth/register [post]
func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	_, err = database.DB.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", input.Username, string(hashedPassword))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary      Authenticate user
// @Description  Authenticate a user with username and password, and return a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        loginInput  body      LoginInput  true  "Login Input"
// @Success      200         {object}  map[string]string
// @Failure      400         {object}  map[string]string
// @Failure      401         {object}  map[string]string
// @Failure      500         {object}  map[string]string
// @Router       /auth/login [post]
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedPassword string
	err := database.DB.QueryRow("SELECT password_hash FROM users WHERE username=$1", input.Username).Scan(&storedPassword)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": input.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
