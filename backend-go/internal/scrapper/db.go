package scrapper

import (
	"context"
	"errors"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func ConnectDB() bool {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	dbPassword := os.Getenv("MONGODB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("MONGODB_PASSWORD is not set")
		return false
	}
	opts := options.Client().ApplyURI("mongodb+srv://user:" + dbPassword + "@refugenavigator.vgcqecj.mongodb.net/?retryWrites=true&w=majority&appName=refugenavigator").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	localClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	// Send a ping to confirm a successful connection
	if err := localClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	client = localClient
	return true
}

func CloseDB() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
func getCollection() *mongo.Collection {
	if client == nil {
		fmt.Println("Client is nil")
	}
	return client.Database("refugenavigator").Collection("features")
}

func (feature Feature) GetCommentsSummaryFromDb() (string, bool) {
	var dbFeature Feature
	err := getCollection().FindOne(context.TODO(), bson.M{"id": feature.Id}).Decode(&dbFeature)
	summaryDbIsValid := true
	if errors.Is(err, mongo.ErrNoDocuments) {
		fmt.Printf("Feature %v not found in db\n", feature.Id)
		_, err = getCollection().InsertOne(context.TODO(), feature)
		if err != nil {
			panic(err)
		}
		return "", false
	}
	if dbFeature.CommentData.Prompt != feature.CommentData.Prompt {
		fmt.Printf("Prompt is not the same for feature %v (db: %v, current: %v)\n", feature.Id, dbFeature.CommentData.Prompt, feature.CommentData.Prompt)
		summaryDbIsValid = false
	}
	if len(dbFeature.CommentData.Comments) != len(feature.CommentData.Comments) {
		fmt.Printf("Comments are not the same for feature %v (db: %v, current: %v)\n", feature.Id, len(dbFeature.CommentData.Comments), len(feature.CommentData.Comments))
		summaryDbIsValid = false
	}
	for i := range dbFeature.CommentData.Comments {
		if dbFeature.CommentData.Comments[i].Content != feature.CommentData.Comments[i].Content {
			fmt.Printf("Comment content is not the same for feature %v (db: %v, current: %v)\n", feature.Id, dbFeature.CommentData.Comments[i].Content, feature.CommentData.Comments[i].Content)
			summaryDbIsValid = false
		}
	}
	if summaryDbIsValid {
		return dbFeature.CommentData.Summary, true
	}
	return "", false
}

func (feature Feature) StoreCommentsSummary() {
	result, err := getCollection().UpdateOne(context.TODO(), bson.M{"id": feature.Id}, bson.M{"$set": bson.M{"commentdata": feature.CommentData}}, options.Update().SetUpsert(true))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated %v documents for feature %v\n", result.ModifiedCount, feature.Id)
}
