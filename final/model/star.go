package model

type Star struct {
	BaseModel
	Name  string `json:"name" gorm:"not null,unique"`
	Intro string `json:"intro"`
	Show  []Show `json:"-" gorm:"many2many:show_star"`
}
