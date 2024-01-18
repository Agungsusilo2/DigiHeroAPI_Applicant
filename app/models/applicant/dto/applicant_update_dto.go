package dto

type UpdateApplicantInput struct {
	EventName            string `json:"event_name" validate:"required"`
	Date                 string `json:"date" validate:"required"`
	EventVenues          string `json:"event_venues" validate:"required"`
	RequirementMaterials string `json:"requirement_materials" validate:"required"`
	IdIndentity          string `json:"id_indentity" validate:"required"`
	PlaceOfExecution     string `json:"place_of_execution" validate:"required"`
}
