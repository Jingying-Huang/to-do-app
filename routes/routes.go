package routes

import (
	"net/http"

	"github.com/Jingying-Huang/to-do-app/models"
	"github.com/Jingying-Huang/to-do-app/utils"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	return r
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	utils.ExecuteTemplate(w, "index.html", tasks)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	task := r.PostForm.Get("to-do")
	err := models.PostTask(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	http.Redirect(w, r, "/", 302)
}
