package model

import (
	"captaindk.site/common"
	"fmt"
	"regexp"
)

type SCategory struct {
	SBaseModel
	Name     string      `json:"name" gorm:"unique_index"`
	Parent   string      `json:"parent"`
	Children SCategories `json:"children"`
	Articles SArticles   `json:"articles" gorm:"many2many:articles_categories"`
	Author   string      `json:"author" gorm:"unique_index:na1"`
}

type SCategories []*SCategory

const (
	DefaultCategoryLengthLimit = 20
)

func (*SCategory) TableName() string {
	return "blog_category"
}

// 在create和save之前检测category的名称是否非法
func (c *SCategory) BeforeCreate() error {
	if c.Valide() {
		return nil
	} else {
		return fmt.Errorf("name of category is invalide")
	}
}
func (c *SCategory) BeforeSave() error {
	if c.Valide() {
		return nil
	} else {
		return fmt.Errorf("name of category is invalide")
	}
}

func (c *SCategory) Get() (err error) {
	return common.Mysql.Where("author = ?", c.Author).Where("name = ?", c.Name).First(c).Error
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

func (cs *SCategories) Fetch() (err error) {
	return common.Mysql.Find(cs).Error
}

func (cs *SCategories) FetchTree() (err error) {
	if len(*cs) == 0 {
		if e := common.Mysql.Where("parent=''").Find(cs).Error; e != nil {
			return e
		}
	}
	for _, c := range *cs {
		if e := common.Mysql.Where("parent=?", c.Name).Find(&c.Children).Error; e != nil {
			return e
		} else if len(c.Children) > 0 {
			return (&c.Children).FetchTree()
		}
	}
	return
}

func (c *SCategory) Valide() bool {
	if len(c.Name) > DefaultCategoryLengthLimit {
		return false
	}
	return regexp.MustCompile("^[^~!@#$%^&*(){}:\"'<>?,./`=]+$").MatchString(c.Name)
}
