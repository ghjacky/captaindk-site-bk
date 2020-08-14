package blog

import (
	"captaindk.site/model"
	"captaindk.site/utils"
)

func AddCategory(category *model.SCategory) (err error) {
	return utils.WrapError(category.Add(), "无法创建分类")
}

func FetchCategoriesInTree(categories *model.SCategories) (err error) {
	return utils.WrapError(categories.FetchTree(), "获取分类失败")
}

func FetchCategories(categories *model.SCategories) (err error) {
	return utils.WrapError(categories.Fetch(), "获取分类失败")
}

func UpdateCategory(category *model.SCategory) (err error) {
	return utils.WrapError(category.Update(), "分类名称更新失败")
}

func DeleteCategory(category *model.SCategory) (err error) {
	return utils.WrapError(category.Delete(), "分类删除失败")
}
