package main

import (
	"net/http"
	"encoding/json"
	// "io/ioutil"
	"log"
	"fmt"
)

type Article struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

type Articles []Article

var articles = Articles {
	Article{Title: "judul 1",Content: "Content 1"},
	Article{Title: "judul 2",Content: "Content 2"},
	Article{Title: "judul 3",Content: "Content 3"},
}

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/about", withLogging(getAbout, "get"))
	http.HandleFunc("/articles", withLogging(getArticles, "get"))
	http.HandleFunc("/post-article", withLogging(postArticle, "post"))
	log.Println("application started at port :8000")
  log.Fatal(http.ListenAndServe(":8888", nil))
}

func getHome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello World, you're in home!"))
}

func getAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, this is about Docker!")
}

func getArticles(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func postArticle(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "POST" {
		if r.Header.Get("Content-type") == "application/json" {
			var a Article
			err := json.NewDecoder(r.Body).Decode(&a)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			articles = append(articles, a)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(articles)
		} else {
			http.Error(w, "Content-Type not support, gunakan 'application/json'", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}
}

func withLogging(next http.HandlerFunc, method string) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		log.Printf("%s dari %s", method, r.RemoteAddr)
		next.ServeHTTP(w,r)
	}
}