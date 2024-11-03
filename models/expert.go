package models

import "github.com/jinzhu/gorm"

type Expert struct {
	gorm.Model
	Name         string   `json:"name" binding:"required"`
	Email        string   `json:"email" gorm:"unique" binding:"required,email"`
	Password     string   `json:"password" binding:"required"`
	Bio          string   `json:"bio"`
	Expertise    string `json:"expertise" binding:"required"`
	Rating       float32  `json:"rating" gorm:"default:0"`
	Appointments int      `json:"appointments" gorm:"default:0"`
	Verified     bool     `json:"verified" gorm:"default:false"`
}
