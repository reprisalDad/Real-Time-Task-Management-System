package main

import (
    "os"

    "github.com/gin-gonic/gin"
    "taskmanager/db"
    "taskmanager/routes"
)

func main() {
    // Load MongoDB URI from env (or use default)
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb://localhost:27017"
    }
    db.Connect(mongoURI)

    router := gin.Default()
    routes.SetupRoutes(router)
    router.Run(":8000")
}
