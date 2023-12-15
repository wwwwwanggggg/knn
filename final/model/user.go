package model

type User struct {
	BaseModel
	Name     string `json:"name" gorm:"not null,unique"`
	Password string `json:"password" gorm:"not null"`
	Show     []Show `json:"shows" gorm:"many2many:show_user"`
}
