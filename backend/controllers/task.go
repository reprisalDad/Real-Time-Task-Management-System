package controllers

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"

    "taskmanager/db"
    "taskmanager/models"
    "taskmanager/services"
)

var taskCollection = db.GetCollection("taskdb", "tasks")

// CreateTaskHandler creates a new task.
func CreateTaskHandler(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    task.ID = primitive.NewObjectID()
    task.CreatedAt = time.Now()
    task.UpdatedAt = time.Now()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    _, err := taskCollection.InsertOne(ctx, task)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task created"})
}

// GetTasksHandler returns all tasks.
func GetTasksHandler(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    cursor, err := taskCollection.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    var tasks []models.Task
    if err := cursor.All(ctx, &tasks); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

// UpdateTaskHandler updates a task.
func UpdateTaskHandler(c *gin.Context) {
    taskID := c.Param("id")
    objID, err := primitive.ObjectIDFromHex(taskID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    var updateData models.Task
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    updateData.UpdatedAt = time.Now()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    update := bson.M{"$set": updateData}
    _, err = taskCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

// DeleteTaskHandler deletes a task.
func DeleteTaskHandler(c *gin.Context) {
    taskID := c.Param("id")
    objID, err := primitive.ObjectIDFromHex(taskID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    _, err = taskCollection.DeleteOne(ctx, bson.M{"_id": objID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

// AISuggestionsHandler returns AI-powered task suggestions.
func AISuggestionsHandler(c *gin.Context) {
    var input struct {
        Description string `json:"description" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    suggestion, err := services.GetTaskSuggestions(input.Description)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"suggestion": suggestion})
}
