package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type test struct {
	Classnum int    `json:"classnum" bson:"classnum"`
	Username string `json:"username" bson:"username"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	return nil, err
	// }

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	return nil, err
	// }

	coll := client.Database("DBTEST").Collection("mongodb")

	//입력할 데이터
	bae := test{1107, "Bae"}

	coll.InsertOne(context.TODO(), bae)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("insert 성공")
	// return coll, nil
}
