package router

import (
	"captaindk.site/model"
	"captaindk.site/module/blog"
	"captaindk.site/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// article相关子路由
func articleRouter() {
	article := Router.Group(fmt.Sprintf("%s/articles", apiPrefix))
	{
		article.GET("", fetchArticles)
		article.POST("", addArticle)
		article.GET("/:id", getArticle)
		article.PUT("", updateArticle)
		article.DELETE("/:id", deleteArticle)
	}
}

/*
	article相关api对应的handler
*/

func getArticle(ctx *gin.Context) {
	aid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(200, newResponse(1000501, fmt.Sprintf("参数错误：%s", err.Error()), nil))
		return
	}
	article := &model.SArticle{}
	article.ID = uint(aid)
	if err := article.Get(); err != nil {
		ctx.JSON(200, newResponse(1000502, err.Error(), nil))
		return
	}
	ctx.JSON(200, newResponse(100000, "成功获取文章内容", article))
	return
}

func deleteArticle(ctx *gin.Context) {
	aid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(200, newResponse(1000401, fmt.Sprintf("参数错误：%s", err.Error()), nil))
		return
	}
	article := &model.SArticle{}
	article.ID = uint(aid)
	if err := article.Delete(); err != nil {
		ctx.JSON(200, newResponse(1000402, err.Error(), nil))
		return
	}
	ctx.JSON(200, newResponse(100000, "成功删除文章", article))
	return
}

// updateArticle 更新文章标题、内容、标签、分类等
func updateArticle(ctx *gin.Context) {
	var article = &model.SArticle{}
	if err := ctx.BindJSON(article); err != nil {
		ctx.JSON(200, newResponse(1000301, fmt.Sprintf("参数错误：%s", err.Error()), nil))
		return
	}
	var err error
	var fieldList = strings.Split(ctx.Query("field"), ",")
	if utils.StringInList("content", fieldList) {
		if err = blog.UpdateArticle(article); err != nil {
			ctx.JSON(200, newResponse(1000302, err.Error(), nil))
			return
		}
	}
	if utils.StringInList("tags", fieldList) {
		if err = blog.ChangeArticleTags(article); err != nil {
			ctx.JSON(200, newResponse(1000302, err.Error(), nil))
			return
		}
	}
	if utils.StringInList("categories", fieldList) {
		if err = blog.ChangeArticleCategories(article); err != nil {
			ctx.JSON(200, newResponse(1000302, err.Error(), nil))
			return
		}
	}
	if utils.StringInList("count", fieldList) {
		if err = blog.AddReadCountOfAritcle(article); err != nil {
			ctx.JSON(200, newResponse(1000302, err.Error(), nil))
			return
		}
	}
	ctx.JSON(200, newResponse(100000, "成功更新文章信息", article))
	return
}

// fetchArticles 获取
func fetchArticles(ctx *gin.Context) {
	var articles model.SArticles
	var query = model.SArticleQuery{}
	if err := ctx.BindQuery(&query); err != nil {
		ctx.JSON(200, newResponse(1000201, fmt.Sprintf("参数错误：%s", err.Error()), nil))
		return
	}
	if err := blog.FetchArticles(&articles, query); err != nil {
		ctx.JSON(200, newResponse(1000202, err.Error(), nil))
		return
	}
	ctx.JSON(200, newResponse(100000, "成功获取文章列表", articles))
	return
}

// addArticle 新增文章
func addArticle(ctx *gin.Context) {
	var article = model.SArticle{}
	if err := ctx.BindJSON(&article); err != nil {
		ctx.JSON(200, newResponse(1000101, fmt.Sprintf("参数错误：%s", err.Error()), nil))
		return
	}
	if err := blog.AddArticle(&article); err != nil {
		ctx.JSON(200, newResponse(1000102, err.Error(), nil))
		return
	}
	ctx.JSON(200, newResponse(100000, "文章添加成功", article))
	return
}
