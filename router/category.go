package router

import (
	"captaindk.site/model"
	"captaindk.site/module/blog"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func categoryRouter() {
	category := Router.Group(fmt.Sprintf("%s/categories", apiPrefix))
	{
		category.GET("", fetchCategories)
		category.POST("", addCategory)
		category.PUT("", updateCategory)
		category.DELETE("/:id", deleteCategory)
	}
}

func deleteCategory(ctx *gin.Context) {
	cid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(200, newResponse(1000401, fmt.Sprintf("参数错误：%s", err.Error()), nil))
		return
	}
	category := &model.SCategory{}
	category.ID = uint(cid)
	if err := blog.DeleteCategory(category); err != nil {
		ctx.JSON(200, newResponse(1000402, err.Error(), nil))
		return
	}
	ctx.JSON(200, newResponse(100000, "分类删除成功", category))
	return
}

func updateCategory(ctx *gin.Context) {
	category := &model.SCategory{}
	if err := ctx.BindJSON(category); err != nil {
		ctx.JSON(200, newResponse(1000301, fmt.Sprintf("参数错误：%s", err.Error()), nil))
		return
	}
	if err := blog.UpdateCategory(category); err != nil {
		ctx.JSON(200, newResponse(1000302, err.Error(), nil))
		return
	}
	ctx.JSON(200, newResponse(100000, "分类名称更新成功", category))
	return
}

func fetchCategories(ctx *gin.Context) {
	tree, err := strconv.ParseBool(ctx.Query("tree"))
	if err != nil {
		ctx.JSON(200, newResponse(1000101, fmt.Sprintf("参数错误：%s", err.Error()), nil))
		return
	}
	var categories model.SCategories
	if tree {
		if err := blog.FetchCategoriesInTree(&categories); err != nil {
			ctx.JSON(200, newResponse(1000102, err.Error(), nil))
			return
		}
	} else {
		if err := blog.FetchCategories(&categories); err != nil {
			ctx.JSON(200, newResponse(1000102, err.Error(), nil))
			return
		}
	}
	ctx.JSON(200, newResponse(100000, "成功获取分类", categories))
	return
}

func addCategory(ctx *gin.Context) {
	var category model.SCategory
	if err := ctx.BindJSON(&category); err != nil {
		ctx.JSON(200, newResponse(1000101, fmt.Sprintf("参数错误"), nil))
		return
	}
	if err := blog.AddCategory(&category); err != nil {
		ctx.JSON(200, newResponse(1000102, err.Error(), nil))
		return
	}
	ctx.JSON(200, newResponse(100000, "成功添加分类", category))
	return
}
