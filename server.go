package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type message struct {
	Text string `json:"text"`
}

type vars struct {
	Var1 string `json:"var1"`
	Var2 string `json:"var2"`
	Var3 string `json:"var3"`
}

type Request struct {
	Date   string          `json:"date"`
	Link   string          `json:"link"`
	Method string          `json:"method"`
	Body   json.RawMessage `json:"body"`
}

func insertReq(link string, method string, body string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println(err)
	}

	collection := client.Database("local").Collection("requests")

	res, err := collection.InsertOne(context.Background(), bson.M{"date": time.DateTime, "link": link, "method": method, "body": body})

	if err != nil {
		log.Println(err)
	}

	id := res.InsertedID
	fmt.Println(id)

	if err := client.Disconnect(context.TODO()); err != nil {
		log.Println(err)
	}

}

func handleRequests() {
	http.HandleFunc("/getMessage", getMessage)
	http.HandleFunc("/postMessage", postMessage)
	http.HandleFunc("/postVar", postVar)
	http.HandleFunc("/", allReq)

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}

func allReq(w http.ResponseWriter, r *http.Request) {
	insertReq(r.RequestURI, r.Method, "")
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	insertReq(r.RequestURI, r.Method, "")
	fmt.Fprint(w, "Test GET method")
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	var m message
	json.NewDecoder(r.Body).Decode(&m)
	fmt.Println("The user sent a message: " + m.Text)
	fmt.Fprintf(w, "Your message: \"%s\" was successfully received", m.Text)
	insertReq(r.RequestURI, r.Method, m.Text)
}

func postVar(w http.ResponseWriter, r *http.Request) {
	var v vars
	json.NewDecoder(r.Body).Decode(&v)
	fmt.Printf("User submitted 3 variables: \n 1 var: %s \n 2 var: %s \n 3 var: %s \n", v.Var1, v.Var2, v.Var3)
	fmt.Fprint(w, v.Var1+v.Var2+v.Var3)
	insertReq(r.RequestURI, r.Method, v.Var1+v.Var2+v.Var3)
}

func insertData() {
}
