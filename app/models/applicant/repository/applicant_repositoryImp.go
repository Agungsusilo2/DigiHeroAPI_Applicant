package repository

import (
	"applicant/app/models/applicant/domain"
	"gorm.io/gorm"
)

type ApplicantRepositoryImpl struct {
	db *gorm.DB
}

func NewApplicantRepository(db *gorm.DB) ApplicantRepository {
	return &ApplicantRepositoryImpl{db}
}

func (ar *ApplicantRepositoryImpl) FindAll() []domain.Applicant {
	var applicants []domain.Applicant

	_ = ar.db.Find(&applicants)

	return applicants
}

func (ar *ApplicantRepositoryImpl) FindById(id int) domain.Applicant {
	var applicant domain.Applicant

	_ = ar.db.First(&applicant, id)

	return applicant
}

func (ar *ApplicantRepositoryImpl) Save(applicant domain.Applicant) (*domain.Applicant, error) {
	result := ar.db.Create(&applicant)

	if result.Error != nil {
		return nil, result.Error
	}

	return &applicant, nil
}

func (ar *ApplicantRepositoryImpl) Update(applicant domain.Applicant) (*domain.Applicant, error) {
	result := ar.db.Save(&applicant)

	if result.Error != nil {
		return nil, result.Error
	}

	return &applicant, nil
}

func (ar *ApplicantRepositoryImpl) Delete(applicant domain.Applicant) (*domain.Applicant, error) {
	result := ar.db.Delete(&applicant)

	if result.Error != nil {
		return nil, result.Error
	}

	return &applicant, nil
}
