package model

import (
	"time"
)

func (Singkronisasi) TableName() string {
	return "master_singkronisasi"
}

type SearchSingkronisasi struct {
	Id   string `json:"id,omitempty"`
	Host string `json:"host,omitempty"`
}

type DeleteSingkronisasi struct {
	Id string `json:"id" validate:"required"`
}

type UpdateTokenSingkronisasi struct {
	Type     string `json:"type" validate:"required"`
	Ckey     string `json:"ckey,omitempty"`
	Csecret  string `json:"csecret,omitempty"`
	ClientId string `json:"client_id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type ResultTokenSso struct {
	AccessToken  string  `json:"access_token"`
	ExpiresIn    float64 `json:"expires_in"`
	RefreshToken string  `json:"refresh_token"`
}

type Singkronisasi struct {
	Id                     string     `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Host                   *string    `json:"host"`
	Uri                    *string    `json:"uri"`
	Ckey                   *string    `json:"ckey"`
	Username               *string    `json:"username"`
	Csecret                *string    `json:"csecret"`
	Password               *string    `json:"password"`
	ClientId               *string    `json:"client_id"`
	GrantType              *string    `json:"grant_type"`
	TokenSso               *string    `json:"-"`
	TokenSsoExpired        *time.Time `json:"token_sso_expired"`
	TokenApimanager        *string    `json:"-"`
	TokenApimanagerExpired *time.Time `json:"token_apimanager_expired"`
	Status                 *int       `json:"status"`
	CreatedBy              string     `gorm:"<-:create" json:"created_by"`
	CreatedAt              *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy              string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt              *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
