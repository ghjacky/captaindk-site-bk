package model

import (
	"captaindk.site/common"
	uuid "github.com/satori/go.uuid"
)

type SFile struct {
	ID      uuid.UUID `json:"id" gorm:"primary_key"`
	Name    string    `json:"name"`
	Dir     string    `json:"dir"`
	Content []byte    `json:"content" gorm:"-"`
}

func (*SFile) TableName() string {
	return "x_file"
}

func (f *SFile) Add() (err error) {
	return common.Mysql.Debug().Create(f).Error
}

func (f *SFile) Get() (err error) {
	return common.Mysql.First(f).Error
}

func (f *SFile) Delete() (err error) {
	return common.Mysql.Unscoped().Delete(f).Error
}
