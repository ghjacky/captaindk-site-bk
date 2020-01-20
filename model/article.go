package model

import (
	"captaindk.site/common"
	"fmt"
	"strings"
)

type SArticle struct {
	SBaseModel
	Author     string      `json:"author"`     // 作者
	Title      string      `json:"title"`      // 标题
	Content    string      `json:"content"`    // 内容
	Categories SCategories `json:"categories"` // 分类
	Tags       STags       `json:"tags"`       // 标签
	Count      uint        `json:"count"`      // 阅读次数
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

func (a *SArticle) Get() (err error) {
	return common.Mysql.First(a).Error
}

type SArticleQuery struct {
	Category  string
	Tags      []string
	Author    string
	Year      string
	Month     string
	OrderBy   string
	OrderType string
	Page      uint
	Limit     uint
}

// 支持分页，支持多维度条件查询，支持排序
func (as *SArticles) Fetch(query SArticleQuery) (err error) {
	var orderClause = fmt.Sprintf("%s %s", query.OrderBy, query.OrderType)
	var whereClause []string
	if len(query.Category) > 0 {
		whereClause = append(whereClause, fmt.Sprintf("category = %s", query.Category))
	}
	if len(query.Tags) > 0 {
		var _wc []string
		for _, tag := range query.Tags {
			if len(tag) > 0 {
				_wc = append(_wc, fmt.Sprintf("tag = %s", tag))
			}
		}
		whereClause = append(whereClause, strings.Join(_wc, " and "))
	}
	if len(query.Author) > 0 {
		whereClause = append(whereClause, fmt.Sprintf("author = %s", query.Author))
	}
	if len(query.Year) > 0 {
		whereClause = append(whereClause, fmt.Sprintf("year = %s", query.Year))
	}
	if len(query.Month) > 0 {
		whereClause = append(whereClause, fmt.Sprintf("month = %s", query.Month))
	}
	return common.Mysql.Where(whereClause).Order(orderClause).Offset((query.Page - 1) * query.Limit).Limit(query.Limit).Find(as).Error
}
