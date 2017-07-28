package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// struct containing all the properties
// that an article should have
type Article struct {
	Id      int    `json:"Id"`
	Title   string `json: "Title"`
	Desc    string `json: "desc"`
	Content string `json:"content"`
}

// array of articles
type Articles []Article

// returns all articles as a Json
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description 1", Content: "Article Content 1"},
		Article{Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage")
	fmt.Println("Endpoint Hit: homePage")
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	// var1 := vars["var1"]
	// var2 := vars["var2"]

	// fmt.Println("Var 1: " + var1)
	// fmt.Println("Var 2: " + var2)
	fmt.Fprintf(w, "Key: "+key)
}

// routing
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	// myRouter.HandleFunc("/article/{key}/{var1}/{var2}/", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
