package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type person struct {
	name          string
	lastname      string
	musicalGenres []string
}

func isUserRegistered(collection *mongo.Collection, name string, lastname string) bool {
	filter := bson.M{"lastname": lastname, "name": name}
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found %d documents with filter %v\n", count, filter)
	return count > 0
}

func main() {
	// Set the URI of your MongoDB Atlas cluster
	uri := "mongodb+srv://SdinarNetlabs:123dinar@cluster0.2ss0vbw.mongodb.net/?retryWrites=true&w=majority"

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB Atlas
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB Atlas!")

	// Create a new database
	database := client.Database("myDatabase")

	// Create a new collection in the database
	collection := database.Collection("myCollection")

	var p person

	fmt.Print("Ingrese su nombre: ")
	fmt.Scan(&p.name)

	fmt.Print("Ingrese su apellido: ")
	fmt.Scan(&p.lastname)

	if isUserRegistered(collection, p.name, p.lastname) {
		fmt.Println("Ya tienes una cuenta registrada.")
		return
	}

	fmt.Print("Ingrese sus géneros musicales favoritos: ")
	var genres string
	fmt.Scan(&genres)
	p.musicalGenres = strings.Split(genres, ",")

	fmt.Println("Usuario registrado?", isUserRegistered(collection, p.name, p.lastname))

	result, err := collection.InsertOne(context.Background(), p)
	if err != nil {
		panic(err)
	}

	fmt.Println("Inserted document with ID:", result.InsertedID)
}
