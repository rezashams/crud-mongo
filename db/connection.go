/**
 * Copyright (c) 2019 TRIALBLAZE PTY. LTD. All rights reserved.
 *
 * Created by Reza Shams (reza.shams@trialblaze.com).
 * User: reza
 * Date: 6/30/19
 * Time: 11:39 AM
 *
 * Description: connection
 *
 */
package db

import (
	"context"
	"fmt"
	"github.com/rezashams/crud-mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

}


func getPerson() string {
	filter := bson.D{}
	result := model.Car{}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("poc-db").Collection("people")
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return ""
}

func InsertCar(car model.Car)  {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("poc-db").Collection("car")
	insertResult, err := collection.InsertOne(context.TODO(), car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
}


func GetCars() [] model.Car {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	var results []model.Car
	collection := client.Database("poc-db").Collection("car")
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem model.Car

		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results
}




func GetCar() model.Car {
	filter := bson.D{}
	var  car model.Car
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("poc-db").Collection("car")
	err = collection.FindOne(context.Background(), filter).Decode(&car)
	if err != nil {
		log.Fatal(err)
	}
	return car
}
