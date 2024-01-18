package domain

import "time"

type Applicant struct {
	ID                   int32     `json:"id" gorm:"primaryKey;auto_increment:true;index"`
	IdIndentity          string    `json:"id_identity" gorm:"type:varchar(150);index;"`
	EventName            string    `json:"event_name" gorm:"type:varchar(250)"`
	Date                 string    `json:"date" gorm:"type:varchar(150)"`
	EventVenues          string    `json:"event_venues" gorm:"type:varchar(250)"`
	RequirementMaterials string    `json:"requirement_materials" gorm:"type:varchar(250)"`
	PlaceOfExecution     string    `json:"place_of_execution" gorm:"type:varchar(250)"`
	CreateAt             time.Time `json:"create_at" gorm:"type:varchar(250)"`
}
