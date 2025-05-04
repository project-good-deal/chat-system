package util

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Create creates a new chat message in the database
func GetConnection() (database *mongo.Database, ctx context.Context, cancel context.CancelFunc) {	
	user, pass, host, port, authDB := getDbEnv()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s", user, pass, host, port, authDB)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
  client, error := mongo.Connect(options.Client().ApplyURI(uri))
	
	if error != nil {
		log.Printf("Error connecting to MongoDB: %v\n", error)
		panic(error)
	}

	return client.Database("chat"), ctx, cancel
}

func getDbEnv() (string, string, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	user := os.Getenv("MONGO_USER")
	pass := os.Getenv("MONGO_PASS")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	authDB := os.Getenv("MONGO_AUTH_DB")
	return user, pass, host, port, authDB
}