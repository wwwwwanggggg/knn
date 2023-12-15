package model

import (
	"time"
)

type Show struct {
	BaseModel
	Title        string    `json:"title" gorm:"not null"`
	Starttime    time.Time `json:"starttime" time_fotmat:"2006-01-02T15:04:05" gorm:"not null"`
	Endtime      time.Time `json:"endtime" time_format:"2006-01-02T15:04:05" gorm:"not null"`
	Location     string    `json:"location" gorm:"not null"`
	MaxCapacity  int       `json:"maxcapacity" gorm:"not null"`
	ShowContent  string    `json:"showcontent"`
	Promo        string    `json:"promo"`
	Performers   []Star    `json:"performers" gorm:"many2many:show_star"`
	Sold         int       `json:"-" gorm:"not null"`
	CurrCapacity int       `json:"currcapacity" gorm:"not null"`
	Users        []User    `json:"-" gorm:"many2many:show_user"`
}
