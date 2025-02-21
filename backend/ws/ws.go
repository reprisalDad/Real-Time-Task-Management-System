package ws

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan interface{})

// HandleWebSocket upgrades the HTTP connection and registers the WebSocket.
func HandleWebSocket(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    clients[conn] = true

    go func(conn *websocket.Conn) {
        defer conn.Close()
        for {
            if _, _, err := conn.ReadMessage(); err != nil {
                delete(clients, conn)
                break
            }
        }
    }(conn)

    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    for {
        select {
        case msg := <-broadcast:
            if err := conn.WriteJSON(msg); err != nil {
                delete(clients, conn)
                return
            }
        case <-ticker.C:
            if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
                delete(clients, conn)
                return
            }
        }
    }
}

// BroadcastTaskUpdate sends a task update to all connected WebSocket clients.
func BroadcastTaskUpdate(update interface{}) {
    broadcast <- update
}
