package blog

import "captaindk.site/utils"

// cannotFindArticleByIDError 无法找到相应ID的文章错误
func cannotFindArticleByIDError(e error) (err error) {
	return utils.WrapError(e, "无法找到相应ID的文章")
}
