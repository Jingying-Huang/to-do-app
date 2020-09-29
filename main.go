package main

import (
	"net/http"
	"github.com/Jingying-Huang/to-do-app/routes"
	"github.com/Jingying-Huang/to-do-app/models"
	"github.com/Jingying-Huang/to-do-app/utils"
)

func main() {
	models.Init()
	utils.LoadTemplates("templates/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
