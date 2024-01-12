package applicant

type Applicant struct {
	ID                   int64  `json:"id" gorm:"primaryKey;auto_increment:true;index"`
	ApplicantsName       string `json:"applicant_name" gorm:"type:varchar(150)"`
	EventName            string `json:"event_name" gorm:"type:varchar(250)"`
	Date                 string `json:"date" gorm:"type:varchar(150)"`
	EventVenues          string `json:"event_venues" gorm:"type:varchar(250)"`
	RequirementMaterials string `json:"requirement_materials" gorm:"type:varchar(250)"`
}
