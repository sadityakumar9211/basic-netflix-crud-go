package mymongo

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sadityakumar9211/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongoDB
func init() {
	// loading environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}


	connectionString := os.Getenv("MONGODB_ATLAS_URI")
	const dbName = "netflix"
	const collectionName = "watchlist"

	// client options
	clientOptions := options.Client().ApplyURI(string(connectionString))

	// connect to mongoDB
	// context has informatino regarding the connections with th mongoDB API
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection Successful!")

	collection = client.Database(dbName).Collection(collectionName)
}

func GetCollection() *mongo.Collection {
	return collection
}

func InsertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted in DB")
	fmt.Printf("%v", inserted)
}

// update 1 record
func UpdateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId) //converts to objectId from hexadecimal string
	if err != nil {
		log.Fatal(err)
	}

	// 2 parts - 1. filter 2. update
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	updateCount, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modify Count: ", updateCount)
}

// deletion of the record
func DeleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The delete count is: ", deleteCount)
}

func DeleteAllMovie() uint64 {
	// filter := bson.D{{}}    // selecting everything
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The delete count is: ", deleteResult.DeletedCount)
	return uint64(deleteResult.DeletedCount)
}

/*
	Type conversion happens when we assign the value of one data type to another.
	Statically typed languages like C/C++, Java, provide the support for Implicit
	Type Conversion but Golang is different, as it doesn't support the Automatic
	Type Conversion or Implicit Type Conversion even if the data types are compatible.
*/

func GetAllMovies() []primitive.M {

	// reading the values from mongoDB gives you a cursor and you have to loop through the cursor to get all the data
	// you need to close the connection to the cursor, after all the reading is done.
	curr, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer curr.Close(context.Background())

	var movies []primitive.M

	for curr.Next(context.Background()) {
		var movie bson.M
		err := curr.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}
