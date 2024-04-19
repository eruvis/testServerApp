package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

type Vars struct {
	Var1 string `json:"var1"`
	Var2 string `json:"var2"`
	Var3 string `json:"var3"`
}

func handleRequests() {
	http.HandleFunc("/getMessage", getMessage)
	http.HandleFunc("/postMessage", postMessage)
	http.HandleFunc("/postVar", postVar)

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test GET method")
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	var m Message
	json.NewDecoder(r.Body).Decode(&m)
	fmt.Println("The user sent a message: " + m.Text)
	fmt.Fprintf(w, "Your message: \"%s\" was successfully received", m.Text)
}

func postVar(w http.ResponseWriter, r *http.Request) {
	var v Vars
	json.NewDecoder(r.Body).Decode(&v)
	fmt.Printf("User submitted 3 variables: \n 1 var: %s \n 2 var: %s \n 3 var: %s \n", v.Var1, v.Var2, v.Var3)
	fmt.Fprint(w, v.Var1+v.Var2+v.Var3)
}
