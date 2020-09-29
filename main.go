package main

import (
	"context"
	"html/template"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var client *redis.Client
var templates *template.Template

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	tasks, err := client.LRange(ctx, "tasks", 0, 10).Result()
	if err != nil {
		return
	}
	templates.ExecuteTemplate(w, "index.html", tasks)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ctx := context.TODO()
	task := r.PostForm.Get("task")
	client.LPush(ctx, "tasks", task)
	http.Redirect(w, r, "/", 302)
}
