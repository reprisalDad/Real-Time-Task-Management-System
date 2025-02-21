package db

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect(uri string) *mongo.Client {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }
    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal(err)
    }
    log.Println("Connected to MongoDB")
    Client = client
    return client
}

func GetCollection(database, collection string) *mongo.Collection {
    return Client.Database(database).Collection(collection)
}
