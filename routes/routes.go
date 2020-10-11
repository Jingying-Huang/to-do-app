package routes

import (
	"database/sql"
	"net/http"

	"github.com/Jingying-Huang/to-do-app/models"
	"github.com/Jingying-Huang/to-do-app/utils"
)

func NewRouter() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dashboard", dashboardIndex)
	http.HandleFunc("/dashboard/show", tasksShow)
	http.HandleFunc("/dashboard/create", tasksCreateForm)
	http.HandleFunc("/dashboard/create/process", tasksCreateProcess)
	// http.HandleFunc("/dashboard/update", tasksUpdateForm)
	// http.HandleFunc("/dashboard/update/process", tasksUpdateProcess)
	// http.HandleFunc("/dashboard/delete/process", tasksDeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func dashboardIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := models.Db.Query("SELECT * FROM tasks")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	tks := make([]models.Task, 0)
	for rows.Next() {
		tk := models.Task{}
		err := rows.Scan(&tk.ID, &tk.Description, &tk.Deadline, &tk.Priority)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		tks = append(tks, tk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	utils.ExecuteTemplate(w, "dashboard.html", tks)
}

func tasksShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := models.Db.QueryRow("SELECT * FROM books WHERE id = $1", id)

	tk := models.Task{}
	err := row.Scan(&tk.ID, &tk.Description, &tk.Deadline, &tk.Priority)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	utils.ExecuteTemplate(w, "show.html", tk)
}

func tasksCreateForm(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "create.html", nil)
}

func tasksCreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// get form values
	tk := models.Task{}
	tk.ID = r.FormValue("id")
	tk.Description = r.FormValue("description")
	tk.Deadline = r.FormValue("deadline")
	p := r.FormValue("priority")

	// validate form values
	if tk.ID == "" || tk.Description == "" || tk.Deadline == "" || p == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// insert values
	models.Db.Exec("INSERT INTO tasks (id, description, deadline, priority) VALUES ($1, $2, $3, $4)", tk.ID, tk.Description, tk.Deadline, tk.Priority)

	// confirm insertion
	utils.ExecuteTemplate(w, "created.html", tk)
}
