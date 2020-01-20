package model

import "captaindk.site/common"

type SCategory struct {
	Name     string      `json:"name" gorm:"primary_key"`
	Parent   string      `json:"parent"`
	Children SCategories `json:"children"`
}

type SCategories []*SCategory

func (*SCategory) TableName() string {
	return "blog_category"
}

func (c *SCategory) Add() (err error) {
	return common.Mysql.Debug().Create(c).Error
}

func (c *SCategory) Delete() (err error) {
	return common.Mysql.Debug().Delete(c).Error
}

func (c *SCategory) Update() (err error) {
	return common.Mysql.Debug().Save(c).Error
}

func (c *SCategory) Fetch() (err error) {

}
