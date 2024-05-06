package model

import (
	"time"
)

func (Users) TableName() string {
	return "users"
}

type SearchUsers struct {
	Id  string `json:"id,omitempty"`
	Nip string `json:"nip,omitempty"`
}

type DeleteUsers struct {
	Id string `json:"id" validate:"required"`
}

type Users struct {
	Id              int        `gorm:"primaryKey" json:"id"`
	Name            string     `json:"name" validate:"required"`
	Email           string     `json:"email,omitempty" validate:"omitempty,email"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Password        string     `json:"password"`
	Nip             string     `json:"nip" validate:"required"`
	Gid             int        `json:"gid"`
	Stts            int        `json:"stts"`
	CreatedAt       *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedAt       *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
