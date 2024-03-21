package model

import (
	"time"
)

func (UnorBkn) TableName() string {
	return "master_unor_bkn"
}

type SearchUnorBkn struct {
	Id          string `json:"id,omitempty"`
	NamaUnor    string `gorm:"column:NamaUnor" json:"NamaUnor,omitempty"`
	NamaJabatan string `gorm:"column:NamaJabatan" json:"NamaJabatan,omitempty"`
	CepatKode   string `gorm:"column:CepatKode" json:"CepatKode,omitempty"`
}

type DeleteUnorBkn struct {
	Id string `json:"id" validate:"required"`
}

type UnorBkn struct {
	Id            string     `gorm:"primaryKey;autoIncrement:false" json:"Id" validate:"required"`
	InstansiId    string     `gorm:"column:InstansiId" json:"InstansiId" validate:"required"`
	DiatasanId    string     `gorm:"column:DiatasanId" json:"DiatasanId" validate:"required"`
	NamaUnor      string     `gorm:"column:NamaUnor" json:"NamaUnor" validate:"required"`
	EselonId      string     `gorm:"column:EselonId" json:"EselonId"`
	NamaJabatan   string     `gorm:"column:NamaJabatan" json:"NamaJabatan"`
	CepatKode     string     `gorm:"column:CepatKode" json:"CepatKode"`
	IndukUnorId   string     `gorm:"column:IndukUnorId" json:"IndukUnorId"`
	PemimpinPnsId string     `gorm:"column:PemimpinPnsId" json:"PemimpinPnsId"`
	JenisUnorId   string     `gorm:"column:JenisUnorId" json:"JenisUnorId"`
	CreatedBy     string     `gorm:"<-:create" json:"created_by"`
	CreatedAt     *time.Time `gorm:"<-:create" json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedBy     string     `gorm:"<-:update" json:"updated_by"`
	UpdatedAt     *time.Time `gorm:"<-:update" json:"updated_at"` // Automatically managed by GORM for update time
}
