package repository

import (
	"applicant/app/models/applicant/domain"
)

type ApplicantRepository interface {
	FindAll() []domain.Applicant
	FindById(id int) domain.Applicant
	Save(applicant domain.Applicant) (*domain.Applicant, error)
	Update(applicant domain.Applicant) (*domain.Applicant, error)
	Delete(applicant domain.Applicant) (*domain.Applicant, error)
}
