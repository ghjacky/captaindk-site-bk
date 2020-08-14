package model

import (
	"captaindk.site/common"
	"fmt"
	"strings"
)

type SArticle struct {
	SBaseModel
	Author     string      `json:"author"`                                               // 作者
	Title      string      `json:"title"`                                                // 标题
	Content    string      `json:"content"`                                              // 内容
	Categories SCategories `json:"categories" gorm:"many2many:blog_articles_categories"` // 分类
	Tags       STags       `json:"tags" gorm:"many2many:blog_articles_tags"`             // 标签
	Count      uint        `json:"count"`                                                // 阅读次数
}

type SArticles []*SArticle

func (*SArticle) TableName() string {
	return "blog_article"
}

func (a *SArticle) Add() (err error) {
	return common.Mysql.Debug().Create(a).Error
}

func (a *SArticle) Delete() (err error) {
	return common.Mysql.Debug().Delete(a).Error
}

func (a *SArticle) Update() (err error) {
	return common.Mysql.Debug().Save(a).Error
}

func (a *SArticle) UpdateCount() (err error) {
	return common.Mysql.Model(a).Update("count", a.Count).Error
}

func (a *SArticle) Get() (err error) {
	return common.Mysql.Preload("Categories").Preload("Tags").Preload("VotedDownByUsers").Preload("VotedUpByUsers").First(a).Error
}

// 更新文章tags
func (a *SArticle) UpdateTags() (err error) {
	return common.Mysql.Model(a).Association("Tags").Replace(a.Tags).Error
}

// 更新文章分类
func (a *SArticle) UpdateCategories() (err error) {
	return common.Mysql.Model(a).Association("Categories").Replace(a.Categories).Error
}

type SArticleQuery struct {
	Category  string   `form:"category"`
	Tags      []string `form:"tags"`
	Author    string   `form:"author"`
	Year      string   `form:"year"`
	Month     string   `form:"month"`
	OrderBy   string   `form:"order_by"`
	OrderType string   `form:"order_type"`
	Page      uint     `form:"page"`
	Limit     uint     `form:"limit"`
}

// 支持分页，支持多维度条件查询，支持排序
func (as *SArticles) Fetch(query SArticleQuery) (err error) {
	if len(query.OrderBy) == 0 {
		query.OrderBy = "id"
	}
	var orderClause = fmt.Sprintf("%s %s", query.OrderBy, query.OrderType)
	var whereClauseArr []string
	if len(query.Category) > 0 {
		whereClauseArr = append(whereClauseArr, fmt.Sprintf("category = %s", query.Category))
	}
	if len(query.Tags) > 0 {
		var _wc []string
		for _, tag := range query.Tags {
			if len(tag) > 0 {
				_wc = append(_wc, fmt.Sprintf("tag = %s", tag))
			}
		}
		whereClauseArr = append(whereClauseArr, strings.Join(_wc, " and "))
	}
	if len(query.Author) > 0 {
		whereClauseArr = append(whereClauseArr, fmt.Sprintf("author = %s", query.Author))
	}
	if len(query.Year) > 0 {
		whereClauseArr = append(whereClauseArr, fmt.Sprintf("year = %s", query.Year))
	}
	if len(query.Month) > 0 {
		whereClauseArr = append(whereClauseArr, fmt.Sprintf("month = %s", query.Month))
	}
	var whereClause = strings.Join(whereClauseArr, " and ")
	return common.Mysql.Where(whereClause).Preload("Categories").Preload("Tags").Order(orderClause).Offset((query.Page - 1) * query.Limit).Limit(query.Limit).Find(as).Error
}
