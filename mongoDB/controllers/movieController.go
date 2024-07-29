package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AVVKavvk/mongoWithGO/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const collectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const colName = "watchList"
var collection *mongo.Collection

func init() {

	clientOption := options.Client().ApplyURI(collectionString);
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("mongoDB connectiob success")
	collection=client.Database(dbName).Collection(colName);

	fmt.Println("collection instance is ready")
}


func insertOneMovie (movie model.Movies){

	res,err:=collection.InsertOne(context.Background(),movie);

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("inserted one movie with _id",res.InsertedID)
}

func updateOneMovie( movieId string){
	id,err:=primitive.ObjectIDFromHex(movieId);

	if err!=nil {
		log.Fatal(err);
	}

	filter:=bson.M{"_id":id};
	update:=bson.M{"$set":bson.M{"watched":true}};

	res,err:=collection.UpdateOne(context.Background(),filter,update);

	if err!=nil {
		log.Fatal(err);
	}
	fmt.Println("updated successfully",res.ModifiedCount)
	
}
func deleteOneMovie(movieId string){
	id,err:=primitive.ObjectIDFromHex(movieId);
	if err!=nil{
		log.Fatal(err);
	}
	filter:=bson.M{"_id":id};

	res,err:=collection.DeleteOne(context.Background(),filter);

	if err!=nil {
		log.Fatal(err);
	}
	fmt.Println("moive deleted by count",res.DeletedCount);

}

func deleteAllMovie() int64{
	filter:=bson.M{};

	res,err:=collection.DeleteMany(context.Background(),filter);

	if err!=nil {
		log.Fatal(err);
	}
	fmt.Println("moives deleted count is ",res.DeletedCount);
	return res.DeletedCount;

}

func getAllMovies() []primitive.M {
	cursor,err:=collection.Find(context.Background(),bson.M{});
	if err!=nil {
		log.Fatal(err);
	}

	var movies []primitive.M;

	for cursor.Next(context.Background()){
		var movie bson.M;
		err:=cursor.Decode(&movie);
		if err!=nil{
			log.Fatal(err);
		}
		movies=append(movies, primitive.M(movie))
	}
	defer cursor.Close(context.Background());
	return movies;
}

func GetAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json");
	allMovies:=getAllMovies();
	json.NewEncoder(w).Encode(allMovies);
	return ;
}
func CreateMovie(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","application/json");
	w.Header().Set("Allow-Control-Allow-Methods","POST");

	var movie model.Movies;
	err:=json.NewDecoder(r.Body).Decode(&movie);
	if err!=nil {
		log.Fatal(err);
	}
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie);
}
func MarkAsWatched(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json");
	w.Header().Set("Allow-Control-Allow-Methods","PUT");

	params:=mux.Vars(r);
	_id:=params["id"]
	updateOneMovie(_id);
	json.NewEncoder(w).Encode(_id);
	return ;
}

func DeleteOneMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json");
	w.Header().Set("Allow-Control-Allow-Methods","DELETE");

	params:=mux.Vars(r);
	_id:=params["id"]
	deleteOneMovie(_id);
	return ;
}
func DeleteAllMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json");
	w.Header().Set("Allow-Control-Allow-Methods","DELETE");
	count :=deleteAllMovie();
	json.NewEncoder(w).Encode(count);
	return ;
}