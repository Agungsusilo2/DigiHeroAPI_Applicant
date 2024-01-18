package controllers

import (
	as "applicant/app/models/applicant/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ApplicantControllerImp struct {
	applicantService as.ApplicantService
	ctx              *gin.Context
}

func NewApplicantController(applicantService as.ApplicantService, ctx *gin.Context) ApplicantControllerImp {
	return ApplicantControllerImp{applicantService, ctx}
}

func (ac *ApplicantControllerImp) Index(ctx *gin.Context) {
	data := ac.applicantService.GetAll()

	ctx.JSON(http.StatusOK, data)
}

func (ac *ApplicantControllerImp) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := ac.applicantService.GetByID(id)

	ctx.JSON(http.StatusOK, data)
}

func (ac *ApplicantControllerImp) Create(ctx *gin.Context) {
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

func (ac *ApplicantControllerImp) Delete(ctx *gin.Context) {
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

func (ac *ApplicantControllerImp) Update(ctx *gin.Context) {
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
