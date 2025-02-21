package routes

import (
    "github.com/gin-gonic/gin"
    "taskmanager/controllers"
    "taskmanager/middleware"
    "taskmanager/ws"
)

func SetupRoutes(router *gin.Engine) {
    // Public endpoints
    router.POST("/signup", controllers.SignupHandler)
    router.POST("/login", controllers.LoginHandler)

    // Protected endpoints (JWT required)
    auth := router.Group("/")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.GET("/tasks", controllers.GetTasksHandler)
        auth.POST("/tasks", controllers.CreateTaskHandler)
        auth.PUT("/tasks/:id", controllers.UpdateTaskHandler)
        auth.DELETE("/tasks/:id", controllers.DeleteTaskHandler)
        auth.POST("/tasks/suggestions", controllers.AISuggestionsHandler)
    }

    // WebSocket endpoint (real-time updates)
    router.GET("/ws/tasks", ws.HandleWebSocket)
}
