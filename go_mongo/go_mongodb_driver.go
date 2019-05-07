package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	_"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Trainer struct {
	Name string `json:"name" bson:"name"`
	Age int `json:"age" bson:"age"`
	City string `json:"city" bson:"city"`
	Created time.Time `json:"created" bson:"created"`
}

func main(){
	
	// MongoDB 연결 하기 
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	
	if err != nil {
		log.Fatal(err)
	}

	
	// 연결 체크 

	fmt.Println("Connected to MongoDB!")

	
	//Collection 가져오기
	collection := client.Database("test").Collection("trainers")

	
	//입력할 데이터 
	ash := Trainer{"Ash", 10, "Pallet Town",time.Now()}
	misty := Trainer{"Misty", 10, "Cerulean City",time.Now()}
	brock := Trainer{"Brock", 15, "Pewter City",time.Now()}

	//1개의 Document 추가하기
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
	log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)


	trainers := []interface{}{misty, brock}

	// 여러개의 Document 추가하기 
	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
	log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)


	
	// 업데이트
	filter := bson.D{{"name", "Ash"}}



	if err != nil {
	log.Fatal(err)
	}
	ash.Age = 120
	update := bson.D{
		{"$set", ash}}
	updateOption := options.Update()
	updateOption.SetUpsert(true)
	
	
	updateResult, err := collection.UpdateOne(context.TODO(), filter,update,updateOption )
	if err != nil {
		log.Fatal(err)
	}


	
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// create a value into which the result can be decoded
	var result Trainer

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)


	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*Trainer

	cur, err := collection.Find(context.TODO(),bson.D{{}}   ,findOptions)
	
	if err != nil {
		log.Fatal(err)
	}
	log.Println("completed find")

	for cur.Next(context.TODO()){

		var elem Trainer
		err := cur.Decode(&elem)
		fmt.Printf(" document: %+v\n", elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results,&elem)
		
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		
	}


	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	cur.Close(context.TODO())

	

}


