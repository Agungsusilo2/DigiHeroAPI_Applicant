package controllers

import "github.com/gin-gonic/gin"

type ApplicantController interface {
	Index(ctx *gin.Context)

	GetByID(ctx *gin.Context)

	Create(ctx *gin.Context)

	Delete(ctx *gin.Context)

	Update(ctx *gin.Context)
}
