package service

import (
	"applicant/app/models/applicant/domain"
	"github.com/gin-gonic/gin"
)

type ApplicantService interface {
	GetAll() []domain.Applicant
	GetByID(id int) domain.Applicant
	Create(ctx *gin.Context) (*domain.Applicant, error)
	Update(ctx *gin.Context) (*domain.Applicant, error)
	Delete(ctx *gin.Context) (*domain.Applicant, error)
}
