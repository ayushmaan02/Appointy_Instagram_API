package main

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	Id              string `json:"id"`
	Caption         string `json:"caption"`
	imageURL        string `json:"imageURL"`
	PostedTimeStamp string `json:"timeStamp"`
}

type Users []User
type Posts []Post

var gid = "ohh_hmmmm" //Global Id

//Pagination Function
func Pagination(r *http.Request, FindOptions *options.FindOptions) (int64, int64) {
	if r.URL.Query().Get("page") != "" && r.URL.Query().Get("limit") != "" {
		page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
		limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
		if page == 1 {
			FindOptions.SetSkip(0)
			FindOptions.SetLimit(limit)
			return page, limit
		}

		FindOptions.SetSkip((page - 1) * limit)
		FindOptions.SetLimit(limit)
		return page, limit

	}
	FindOptions.SetSkip(0)
	FindOptions.SetLimit(0)
	return 0, 0
}

//1.Function to create a user
func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST worked")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.jxmqb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	options := options.Find()
	Pagination(r, options)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	instagramDB := client.Database("instagramDB")
	usersCollection := instagramDB.Collection("users")

	pass := "A@1243"
	uEnc := b64.URLEncoding.EncodeToString([]byte(pass)) //Encoded the password using URL encoding for more security

	userResult, err := usersCollection.InsertOne(ctx, bson.D{
		{"id", "ohh_hmmmm"},
		{"Name", "Ayushmaan"},
		{"Email", "asr00@gmail.com"},
		{"Pass", uEnc},
	})

	// To decode we can use uDec, _ := b64.URLEncoding.DecodeString(uEnc)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userResult.InsertedID)
}

//2.GET User by ID
func getUsingId(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.jxmqb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	options := options.Find()
	Pagination(r, options)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	instagramDB := client.Database("instagramDB")
	usersCollection := instagramDB.Collection("users")

	//Code to retrieve by FILTER

	filterCursor, err := usersCollection.Find(ctx, bson.M{"id": gid})
	if err != nil {
		log.Fatal(err)
	}
	var usersFiltered []bson.M
	if err = filterCursor.All(ctx, &usersFiltered); err != nil {
		log.Fatal(err)
	}

	var id = ""
	var Name = ""
	var Email = ""
	var Pass = ""
	for _, usersFiltereds := range usersFiltered {
		id = usersFiltereds["id"].(string)
		Name = usersFiltereds["Name"].(string)
		Email = usersFiltereds["Email"].(string)
		Pass = usersFiltereds["Pass"].(string)
	}

	Users := []User{
		User{Id: id, Name: Name, Email: Email, Password: Pass},
	}
	fmt.Println("Endpoint sucessfully triggered")
	json.NewEncoder(w).Encode(Users)

}

//3.Function to create a post
func createPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST Worked")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.jxmqb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	options := options.Find()
	Pagination(r, options)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	instagramDB := client.Database("instagramDB")
	postsCollection := instagramDB.Collection("posts")

	currentTime := time.Now() //Time Stamp added
	postResult, err := postsCollection.InsertOne(ctx, bson.D{
		{"id", "ohh_hmmmm"},
		{"caption", "Pic 1 hmm"},
		{"imageURL", "https://go.dev/blog/go-brand/logos.jpg"}, //Golang Logo
		{"Posted TimeStamp", currentTime.Format("02-January-2000 15:04:05")},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(postResult.InsertedID)
}

//4.GET post by ID
func getPostUsingId(w http.ResponseWriter, r *http.Request) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.jxmqb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	options := options.Find()
	Pagination(r, options)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	instagramDB := client.Database("instagramDB")
	postsCollection := instagramDB.Collection("posts")

	//Code to retrieve by FILTER

	filterCursor, err := postsCollection.Find(ctx, bson.M{"id": gid})
	if err != nil {
		log.Fatal(err)
	}
	var postsFiltered []bson.M

	if err = filterCursor.All(ctx, &postsFiltered); err != nil {
		log.Fatal(err)
	}

	var id = ""
	var caption = ""
	var ImageURL = ""
	var postedTimeStamp = ""
	for _, postsFiltereds := range postsFiltered {
		id = postsFiltereds["id"].(string)
		caption = postsFiltereds["caption"].(string)
		ImageURL = postsFiltereds["imageURL"].(string)
		postedTimeStamp = postsFiltereds["Posted TimeStamp"].(string)
	}

	Posts := []Post{
		Post{Id: id, Caption: caption, imageURL: ImageURL, PostedTimeStamp: postedTimeStamp},
	}
	fmt.Println("Endpoint sucessfully triggered")
	json.NewEncoder(w).Encode(Posts)

}

//5.GET all post by particular user id
func getAllPostUsingId(w http.ResponseWriter, r *http.Request) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.jxmqb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	options := options.Find()
	Pagination(r, options)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	instagramDB := client.Database("instagramDB")
	postsCollection := instagramDB.Collection("posts")

	//Code retrieve by FILTER

	filterCursor, err := postsCollection.Find(ctx, bson.M{"id": gid})
	if err != nil {
		log.Fatal(err)
	}
	var postsFiltered []bson.M
	if err = filterCursor.All(ctx, &postsFiltered); err != nil {
		log.Fatal(err)
	}

	var id = ""
	var caption = ""
	var ImageURL = ""
	var postedTimeStamp = ""
	for _, postsFiltereds := range postsFiltered {
		id = postsFiltereds["id"].(string)
		caption = postsFiltereds["caption"].(string)
		ImageURL = postsFiltereds["imageURL"].(string)
		postedTimeStamp = postsFiltereds["Posted TimeStamp"].(string)
		post_all_1 := Posts{
			Post{Id: id, Caption: caption, imageURL: ImageURL, PostedTimeStamp: postedTimeStamp},
		}
		json.NewEncoder(w).Encode(post_all_1)
	}

	fmt.Println("Endpoint sucessfully triggered")

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "By Ayushmaan: Apponity Tech Round Instagram Backend API")
}

//Request Handler for all requests
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	getIdPath := "/users/" + gid
	getPostPathById := "/posts/" + gid
	getAllPostPathById := "/posts/users/" + gid
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", createUser).Methods("POST")                 // POST request to create a new user http://localhost:8000/users
	myRouter.HandleFunc(getIdPath, getUsingId).Methods("GET")                 // GET request to display user details using gid http://localhost:8000/users/ohh_hmmmm
	myRouter.HandleFunc("/posts", createPost).Methods("POST")                 // POST request to create a new post for a user http://localhost:8000/posts
	myRouter.HandleFunc(getPostPathById, getPostUsingId).Methods("GET")       // GET request to display user's post details using gid http://localhost:8000/posts/ohh_hmmmm
	myRouter.HandleFunc(getAllPostPathById, getAllPostUsingId).Methods("GET") // GET request to display particular user's post details using gid http://localhost:8000/posts/users/ohh_hmmmm

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	handleRequests()
}
