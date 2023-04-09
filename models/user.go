package models

import (
	"DTS/Chapter-3/chapter3-challenge2/helpers"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required, minstringlength(6)~Password minimum length of 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL" json:"products"`
	RoleID   uint      `gorm:"not null" json:"role_id" form:"role_id" valid:"required~Role is required"`
}

func (u *User) TableName() string {
	return "tb_users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPassword(u.Password)

	return
}
