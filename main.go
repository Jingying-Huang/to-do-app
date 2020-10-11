package main

import (
	"github.com/Jingying-Huang/to-do-app/models"
	"github.com/Jingying-Huang/to-do-app/routes"
	"github.com/Jingying-Huang/to-do-app/utils"
)

func main() {
	models.Init()
	utils.LoadTemplates("templates/*.html")
	routes.NewRouter()
}
