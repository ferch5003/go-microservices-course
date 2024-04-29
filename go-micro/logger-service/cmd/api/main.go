package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log-service/internal/data"
)

const (
	_webPort  = "80"
	_rpcPort  = "5001"
	_mongoURL = "mongodb://mongo:27017"
	_gRPCPort = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	// connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panicln(err)
	}
	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	// go app.serve()
	log.Println("Starting Service on Port", _webPort)
	srv := &http.Server{
		Addr:    ":" + _webPort,
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panicln(err)
	}
}

//func (app *Config) serve() {
//	srv := &http.Server{
//		Addr:    ":" + _webPort,
//		Handler: app.routes(),
//	}
//
//	if err := srv.ListenAndServe(); err != nil {
//		log.Panicln(err)
//	}
//}

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(_mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting", err)
		return nil, err
	}

	log.Println("Connected to Mongo!")

	return c, nil
}
