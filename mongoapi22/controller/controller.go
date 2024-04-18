package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Sunny955/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://sunny:1234@gocluster.tuaozoa.mongodb.net/?retryWrites=true&w=majority&appName=GoCluster"
const dbName = "netflix"
const colName = "watchlist"

// Most Important
var collection *mongo.Collection

// Connect with mongoDB
func init() {
	// client option
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection to mongoDB is successfull!")

	collection = client.Database(dbName).Collection(colName)

	// once collection reference/instance ready
	fmt.Println("Collection instance is ready!")
}

// MongoDB helpers - file

// insert 1 record
func insertOneMovie(movie model.Netflix) (*mongo.InsertOneResult, error) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		return &mongo.InsertOneResult{}, err
	}

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
	return inserted, nil
}

// update one record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated successfully!: ", result.ModifiedCount)
}

// delete one movie
func deleteOneMovie(movieId string) (int64, error) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return 0, err
	}

	// fmt.Println("Movie got deleted with delete count :", deleteCount)
	return deleteCount.DeletedCount, nil
}

// delete all movie
func deleteAllMovie() (int64, error) {
	// filter := bson.D{{}};

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		return 0, err
	}

	fmt.Println("Number of movies deleted: ", deleteResult.DeletedCount)

	return deleteResult.DeletedCount, nil
}

// get all movies from db

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M

		err := cur.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())

	return movies
}

// Actual controller - file (we will use this functions in routes file i.e they are CaptialCased)

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	// Write the message to the response writer
	w.WriteHeader(http.StatusOK)
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	if r.Body == nil {
		http.Error(w, "Empty body, please enter some data!", http.StatusBadRequest)
		return
	}

	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)

	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if movie.Movie == "" && !movie.Watched {
		http.Error(w, "Fields cannot be empty", http.StatusBadRequest)
		return
	}

	createdMovie, err := insertOneMovie(movie)

	if err != nil {
		log.Fatal("Unable to insert a new movie in the database:", err)
		http.Error(w, "Unable to insert a new movie in the database", http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Created the movie and it is : %s", createdMovie.InsertedID)

	// Write the message to the response writer
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func MarkedAsWatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])

	message := fmt.Sprintf("Updated the movie of id : %s", params["id"])

	// Write the message to the response writer
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteCount, err := deleteOneMovie(params["id"])

	if err != nil {
		log.Fatal(err)
	}
	message := fmt.Sprintf("Deleted the movie of id : %s with delete count as :%d", params["id"], deleteCount)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteCount, err := deleteAllMovie()

	if err != nil {
		log.Fatal(err)
	}
	// Write the message to the response writer
	w.WriteHeader(http.StatusOK)
	message := fmt.Sprintf("Deleted all the movies in databased, delete count is : %d", deleteCount)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
