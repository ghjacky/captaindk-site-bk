package blog

import (
	"captaindk.site/model"
	"captaindk.site/utils"
)

func AddTag(tag *model.STag) (err error) {
	return utils.WrapError(tag.Add(), "无法添加标签")
}

func FetchTags(tags *model.STags) (err error) {
	return utils.WrapError(tags.Fetch(), "无法获取标签列表")
}
