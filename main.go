package main

import (
	"applicant/database"
	"applicant/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := database.ConnectDatabase()

	routes.Api(r, db)

	r.Run(":9888")
}
