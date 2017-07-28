package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// struct containing all the properties
// that an article should have
type Article struct {
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

// routing
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
