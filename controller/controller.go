package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shamsfiroz/mongoApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "your-mongourl"
const dbName = "youtube-clone"
const colName = "watchlist"

// Most important method here:
var collection *mongo.Collection

//connect with mongodb

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Connented to server successfully!")
	collection = client.Database(dbName).Collection(colName)

	//if collection ready for me
	fmt.Println("Collection instance is ready!")

}

//Mongodb helper

func insertOneMovie(movie model.Netflix) {
	result, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie Added successfully", result.InsertedID)

}

// Update one
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie update successfully", result.ModifiedCount)

}

// delete one
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie deleted successfully", result.DeletedCount)

}

//Delete many or all movie

func deleteAllMovie() {
	filter := bson.M{}
	result, err := collection.DeleteMany(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie deleted All successfully", result.DeletedCount)

}

// get All movie
func getAllMovie() []primitive.M {
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies

}

// get one movie by id

// func getMovieById(movieId string)  {
// 	id, err := primitive.ObjectIDFromHex(movieId)
// 	if err != nil {
// 		log.Fatal("Invalid ID format: ", err)
// 	}

// 	filter := bson.M{"_id": id}
// 	var result

// 	err = collection.FindOne(context.Background(), filter).Decode(&result)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			fmt.Println("No document found with the given ID")
// 			return nil
// 		}
// 		log.Fatal("Error fetching document: ", err)
// 	}

// 	fmt.Println("Movie fetched successfully")
// 	return &result
// }

func GetAllMoviesList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovie()
	json.NewEncoder(w).Encode(allMovies)
}
