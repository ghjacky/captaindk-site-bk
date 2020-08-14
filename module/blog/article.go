package blog

import (
	"captaindk.site/model"
	"captaindk.site/utils"
	"github.com/jinzhu/gorm"
)

// 获取文章列表 (可根据category、tag、author、月份、年份等来获取对应文章列表)
func FetchArticles(articles *model.SArticles, query model.SArticleQuery) (err error) {
	return utils.WrapError(articles.Fetch(query), "获取文章失败")
}

// 新增文章
func AddArticle(article *model.SArticle) (err error) {
	// 首先判断所包含的tag是否已经存在，如果不存在需要先创建改tag
	for _, tag := range article.Tags {
		if e := tag.Get(); gorm.IsRecordNotFoundError(e) {
			// 如果不存在，则创建
			_ = AddTag(tag)
		} else if e != nil {
			return utils.WrapError(e, "数据库错误")
		}
	}
	// 检测分类是否存在，不存在则返回错误
	for _, category := range article.Categories {
		category.Author = article.Author
		if e := category.Get(); e != nil {
			return utils.WrapError(e, "分类不存在")
		}
	}
	return utils.WrapError(article.Add(), "文章添加失败")
}

// 更新文章标题、内容
func UpdateArticle(article *model.SArticle) (err error) {
	// 首先获取老的文章信息，更新标题和内容字段
	var oldArticle = &model.SArticle{}
	oldArticle.ID = article.ID
	if e := oldArticle.Get(); e != nil {
		return cannotFindArticleByIDError(e)
	} else {
		oldArticle.Title = article.Title
		oldArticle.Content = article.Content
	}
	// 然后更新
	return utils.WrapError(oldArticle.Update(), "文章更新失败")
}

func AddReadCountOfAritcle(article *model.SArticle) (err error) {
	var oldArticle = &model.SArticle{}
	oldArticle.ID = article.ID
	if e := oldArticle.Get(); e != nil {
		return cannotFindArticleByIDError(e)
	} else {
		oldArticle.Count += 1
	}
	return utils.WrapError(oldArticle.UpdateCount(), "阅读次数增加失败")
}

// 更新tags（删除老的tags，重新设置tags）
func ChangeArticleTags(article *model.SArticle) (err error) {
	// 首先获取老的文章信息，重设tags
	var oldArticle = &model.SArticle{}
	oldArticle.ID = article.ID
	if e := oldArticle.Get(); e != nil {
		return cannotFindArticleByIDError(e)
	} else {
		oldArticle.Tags = article.Tags
	}
	// 然后更新文章tags信息
	return utils.WrapError(oldArticle.UpdateTags(), "无法更新文章标签信息")
}

// 更新分类
func ChangeArticleCategories(article *model.SArticle) (err error) {
	// 首先获取老的文章信息，重新分类
	var oldArticle = &model.SArticle{}
	oldArticle.ID = article.ID
	if e := oldArticle.Get(); e != nil {
		return cannotFindArticleByIDError(e)
	} else {
		oldArticle.Categories = article.Categories
	}
	// 然后更新分类
	return utils.WrapError(oldArticle.UpdateCategories(), "无法更新文章分类信息")
}
