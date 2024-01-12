package applicant

import (
	dto "applicant/app/models/applicant/dto"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApplicantService interface {
	GetAll() []Applicant
	GetByID(id int) Applicant
	Create(ctx *gin.Context) (*Applicant, error)
	Update(ctx *gin.Context) (*Applicant, error)
	Delete(ctx *gin.Context) (*Applicant, error)
}

type ApplicantServiceImpl struct {
	applicantRepository ApplicantRepository
}

func NewApplicantService(applicantRepository ApplicantRepository) ApplicantService {
	return &ApplicantServiceImpl{applicantRepository}
}

func (as *ApplicantServiceImpl) GetAll() []Applicant {
	return as.applicantRepository.FindAll()
}

func (as *ApplicantServiceImpl) GetByID(id int) Applicant {
	return as.applicantRepository.FindOne(id)
}

func (as *ApplicantServiceImpl) Create(ctx *gin.Context) (*Applicant, error) {
	var input dto.CreateApplicantInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	applicant := Applicant{
		ApplicantsName:       input.ApplicantsName,
		EventName:            input.EventName,
		Date:                 input.Date,
		EventVenues:          input.EventVenues,
		RequirementMaterials: input.RequirementMaterials,
	}

	result, err := as.applicantRepository.Save(applicant)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (as *ApplicantServiceImpl) Update(ctx *gin.Context) (*Applicant, error) {
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

	applicant := Applicant{
		ID:                   int64(id),
		ApplicantsName:       input.ApplicantsName,
		EventName:            input.EventName,
		Date:                 input.Date,
		EventVenues:          input.EventVenues,
		RequirementMaterials: input.RequirementMaterials,
	}

	result, err := as.applicantRepository.Update(applicant)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (as *ApplicantServiceImpl) Delete(ctx *gin.Context) (*Applicant, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	applicant := Applicant{
		ID: int64(id),
	}

	result, err := as.applicantRepository.Delete(applicant)

	if err != nil {
		return nil, err
	}

	return result, nil
}
