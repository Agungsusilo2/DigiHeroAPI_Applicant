package applicant

import (
	"gorm.io/gorm"
)

type ApplicantRepository interface {
	FindAll() []Applicant
	FindOne(id int) Applicant
	Save(applicant Applicant) (*Applicant, error)
	Update(applicant Applicant) (*Applicant, error)
	Delete(applicant Applicant) (*Applicant, error)
}

type ApplicantRepositoryImpl struct {
	db *gorm.DB
}

func NewApplicantRepository(db *gorm.DB) ApplicantRepository {
	return &ApplicantRepositoryImpl{db}
}

func (ar *ApplicantRepositoryImpl) FindAll() []Applicant {
	var applicants []Applicant

	_ = ar.db.Find(&applicants)

	return applicants
}

func (ar *ApplicantRepositoryImpl) FindOne(id int) Applicant {
	var applicant Applicant

	_ = ar.db.First(&applicant, id)

	return applicant
}

func (ar *ApplicantRepositoryImpl) Save(applicant Applicant) (*Applicant, error) {
	result := ar.db.Create(&applicant)

	if result.Error != nil {
		return nil, result.Error
	}

	return &applicant, nil
}

func (ar *ApplicantRepositoryImpl) Update(applicant Applicant) (*Applicant, error) {
	result := ar.db.Save(&applicant)

	if result.Error != nil {
		return nil, result.Error
	}

	return &applicant, nil
}

func (ar *ApplicantRepositoryImpl) Delete(applicant Applicant) (*Applicant, error) {
	result := ar.db.Delete(&applicant)

	if result.Error != nil {
		return nil, result.Error
	}

	return &applicant, nil
}
