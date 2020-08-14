package model

import (
	"captaindk.site/common"
	"fmt"
	"regexp"
)

type STag struct {
	Name     string    `json:"name" gorm:"primary_key"`
	Articles SArticles `json:"articles" gorm:"blog_articles_tags"`
}

type STags []*STag

const (
	DefaultTagLengthLimit = 16
)

func (*STag) TableName() string {
	return "blog_tag"
}

// 在create和save之前检测tag的名称是否非法
func (t *STag) BeforeCreate() error {
	if t.Valide() {
		return nil
	} else {
		return fmt.Errorf("name of tag is invalide")
	}
}
func (t *STag) BeforeSave() error {
	if t.Valide() {
		return nil
	} else {
		return fmt.Errorf("name of tag is invalide")
	}
}

func (t *STag) Add() error {
	return common.Mysql.Debug().Create(t).Error
}

func (t *STag) Delete() error {
	return common.Mysql.Debug().Delete(t).Error
}

func (t *STag) Get() error {
	return common.Mysql.First(t).Error
}

func (ts *STags) Fetch() error {
	return common.Mysql.Find(ts).Error
}

func (t *STag) Valide() bool {
	if len(t.Name) > DefaultTagLengthLimit {
		return false
	}
	return regexp.MustCompile("^[^~!@#$%^&*(){}:\"'<>?,./`=]+$").MatchString(t.Name)
}
