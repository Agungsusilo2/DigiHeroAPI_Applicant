package applicant

import (
	as "applicant/app/models/applicant"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ApplicantController struct {
	applicantService as.ApplicantService
	ctx              *gin.Context
}

func NewApplicantController(applicantService as.ApplicantService, ctx *gin.Context) ApplicantController {
	return ApplicantController{applicantService, ctx}
}

func (ac *ApplicantController) Index(ctx *gin.Context) {
	data := ac.applicantService.GetAll()

	ctx.JSON(http.StatusOK, data)
}

func (ac *ApplicantController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := ac.applicantService.GetByID(id)

	ctx.JSON(http.StatusOK, data)
}

func (ac *ApplicantController) Create(ctx *gin.Context) {
	data, err := ac.applicantService.Create(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (ac *ApplicantController) Delete(ctx *gin.Context) {
	data, err := ac.applicantService.Delete(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (ac *ApplicantController) Update(ctx *gin.Context) {
	data, err := ac.applicantService.Update(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, data)
}
