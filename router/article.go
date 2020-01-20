package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// article相关子路由
func articleRouter() {
	article := Router.Group(fmt.Sprintf("%s/articles", apiPrefix))
	{
		article.GET("", fetchArticles)
		article.POST("", addArticle)
	}
}

/*
	article相关api对应的handler
*/

// fetchArticles 获取
func fetchArticles(ctx *gin.Context) {

}

// addArticle 新增文章
func addArticle(ctx *gin.Context) {

}
