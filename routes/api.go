package routes

import (
	ac "applicant/app/controllers"
	"applicant/app/models/applicant"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
)

func Api(router *gin.Engine, db *gorm.DB) {
	applicantRepository := applicant.NewApplicantRepository(db)
	applicantService := applicant.NewApplicantService(applicantRepository)
	applicantController := ac.NewApplicantController(applicantService, ctx)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/applicants", applicantController.Index)
		v1.GET("/applicants/:id", applicantController.GetByID)
		v1.POST("/applicants", applicantController.Create)
		v1.PATCH("/applicants/:id", applicantController.Update)
		v1.DELETE("/applicants/:id", applicantController.Delete)
	}
}
