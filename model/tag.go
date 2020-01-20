package model

type STag struct {
	SBaseModel
	Name string `json:"name" gorm:"unique"`
}

type STags []*STag
