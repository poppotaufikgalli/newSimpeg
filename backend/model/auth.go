package model

import (
	"time"
)

func (UsersToken) TableName() string {
	return "users_token"
}

type UsersToken struct {
	Nip       string     `json:"nip"`
	Nama      string     `json:"nama"`
	IdOpd     string     `json:"id_opd"`
	ParentOpd string     `json:"parent_opd"`
	TokenId   string     `json:"token_id"`
	Gid       int        `json:"gid"`
	LoginAt   *time.Time `json:"login_at"`
	LogoutAt  *time.Time `json:"logout_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type SIAPResponse struct {
	Message     string     `json:"message"`
	Success     bool       `json:"success"`
	LoginStatus bool       `json:"loginStatus"`
	Status      *int       `json:"status"`
	Datauser    UsersToken `json:"datauser"`
}
