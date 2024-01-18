package service

import (
	"applicant/app/models/applicant/domain"
	"applicant/app/models/applicant/dto"
	"applicant/app/models/applicant/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type ApplicantServiceImpl struct {
	applicantRepository repository.ApplicantRepository
}

func NewApplicantService(applicantRepository repository.ApplicantRepository) ApplicantService {
	return &ApplicantServiceImpl{applicantRepository}
}

func (as *ApplicantServiceImpl) GetAll() []domain.Applicant {
	return as.applicantRepository.FindAll()
}

func (as *ApplicantServiceImpl) GetByID(id int) domain.Applicant {
	return as.applicantRepository.FindById(id)
}

func (as *ApplicantServiceImpl) Create(ctx *gin.Context) (*domain.Applicant, error) {
	var input dto.CreateApplicantInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	a := domain.Applicant{
		IdIndentity:          input.IdIndentity,
		EventName:            input.EventName,
		Date:                 input.Date,
		EventVenues:          input.EventVenues,
		RequirementMaterials: input.RequirementMaterials,
		PlaceOfExecution:     input.PlaceOfExecution,
	}

	result, err := as.applicantRepository.Save(a)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (as *ApplicantServiceImpl) Update(ctx *gin.Context) (*domain.Applicant, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.UpdateApplicantInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	a := domain.Applicant{
		ID:                   int32(id),
		IdIndentity:          input.IdIndentity,
		EventName:            input.EventName,
		Date:                 input.Date,
		EventVenues:          input.EventVenues,
		RequirementMaterials: input.RequirementMaterials,
		PlaceOfExecution:     input.PlaceOfExecution,
	}

	result, err := as.applicantRepository.Update(a)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (as *ApplicantServiceImpl) Delete(ctx *gin.Context) (*domain.Applicant, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	a := domain.Applicant{
		ID: int32(id),
	}

	result, err := as.applicantRepository.Delete(a)

	if err != nil {
		return nil, err
	}

	return result, nil
}
