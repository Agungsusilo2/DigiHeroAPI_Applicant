package applicant

type UpdateApplicantInput struct {
	ApplicantsName       string `json:"applicant_name" validate:"required"`
	EventName            string `json:"event_name" validate:"required"`
	Date                 string `json:"date" validate:"required"`
	EventVenues          string `json:"event_venues" validate:"required"`
	RequirementMaterials string `json:"requirement_materials" validate:"required"`
}
