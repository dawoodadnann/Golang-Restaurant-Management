// database/db.go
package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
	err := godotenv.Load() // Load environment variables from .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI") // Mongo URI from .env
	if mongoURI == "" {
		log.Fatal("Mongo URI is empty!")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	err = client.Ping(nil, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	Client = client
	log.Println("Connected to MongoDB!")
}
