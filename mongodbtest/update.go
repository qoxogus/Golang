package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type updateuserinfo struct {
	Classnum int    `json:"classnum" bson:"classnum"`
	Username string `json:"username" bson:"username"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	fmt.Println("mongoDB Connect")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err) //에러 문자열이 출력되고 종료
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("DBTEST").Collection("mongodb")

	//mongoDB update
	//update

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongoDB Disconnect")
}
