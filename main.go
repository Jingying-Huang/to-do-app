package main

import (
	"html/template"
	"net/http"

	"./models"
	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	models.Init()
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
	tasks, err := models.GetTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	templates.ExecuteTemplate(w, "index.html", tasks)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	task := r.PostForm.Get("task")
	err := models.PostTask(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	http.Redirect(w, r, "/", 302)
}
