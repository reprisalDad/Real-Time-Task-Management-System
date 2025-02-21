package controllers

import (
    "context"
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "golang.org/x/crypto/bcrypt"

    "taskmanager/db"
    "taskmanager/middleware"
    "taskmanager/models"
)

var validate = validator.New()
var userCollection = db.GetCollection("taskdb", "users")

// HashPassword hashes a plaintext password.
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// VerifyPassword compares a hashed password with its plaintext equivalent.
func VerifyPassword(hashedPassword, password string) (bool, string) {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        return false, "Invalid password"
    }
    return true, ""
}

// SignupHandler creates a new user.
func SignupHandler(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := validate.Struct(user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Ensure email is lowercase
    user.Email = strings.ToLower(user.Email)

    count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if count > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
        return
    }
    hashedPwd, err := HashPassword(user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
        return
    }
    user.Password = hashedPwd
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    user.ID = primitive.NewObjectID()

    _, err = userCollection.InsertOne(ctx, user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// LoginHandler authenticates a user and returns a JWT.
func LoginHandler(c *gin.Context) {
    var input struct {
        Email    string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    var user models.User
    err := userCollection.FindOne(ctx, bson.M{"email": strings.ToLower(input.Email)}).Decode(&user)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }
    valid, errMsg := VerifyPassword(user.Password, input.Password)
    if !valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
        return
    }
    token, err := middleware.GenerateToken(user.ID.Hex())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": token})
}
