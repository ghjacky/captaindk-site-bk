package router

import (
	"captaindk.site/common"
	"captaindk.site/model"
	"captaindk.site/module/file"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
)

func FileRouter() {
	fr := Router.Group(fmt.Sprintf("%s/files", apiPrefix))
	{
		fr.POST("/upload", uploadFile)
		fr.GET("/download/:id", downloadFile)
		fr.DELETE("/remove/:id", removeFile)
	}
	static := Router.Group(fmt.Sprintf("%s/static", apiPrefix))
	{
		static.Static("", common.Config.FileStore)
	}
}

func removeFile(ctx *gin.Context) {
	fid, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(200, newResponse(1000301, fmt.Sprintf("参数错误：%s", err.Error()), nil))
		return
	}
	f := &model.SFile{ID: fid}
	if err := file.GetFile(f); err != nil {
		ctx.JSON(200, newResponse(1000302, err.Error(), nil))
		return
	}
	if err := file.RemoveFile(f); err != nil {
		ctx.JSON(200, newResponse(1000302, err.Error(), nil))
		return
	}
	ctx.JSON(200, newResponse(100000, "成功删除文件", f))
	return
}

func downloadFile(ctx *gin.Context) {
	fid, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(200, newResponse(1000201, fmt.Sprintf("参数错误：%s", err.Error()), nil))
		return
	}
	f := &model.SFile{ID: fid}
	if err := file.GetFile(f); err != nil {
		ctx.JSON(200, newResponse(1000202, err.Error(), nil))
		return
	}
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", f.Name))
	ctx.Header("Content-Type", "application/text/plain")
	ctx.Header("Accept-Length", fmt.Sprintf("%d", len(f.Content)))
	_, err = ctx.Writer.Write([]byte(f.Content))
	if err != nil {
		ctx.JSON(200, newResponse(1000202, fmt.Sprintf("下载失败：%s", err.Error()), nil))
		return
	}
	ctx.JSON(200, newResponse(100000, "文件下载成功", nil))
	return
}

func uploadFile(ctx *gin.Context) {
	f, h, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(200, newResponse(1000101, fmt.Sprintf("参数错误：%s", err.Error()), nil))
		return
	}
	fileContent, err := ioutil.ReadAll(f)

	if err != nil {
		ctx.JSON(200, newResponse(1000102, fmt.Sprintf("文件内容读取错误：%s", err.Error()), nil))
		return
	}
	fid, err := file.StoreFile(h.Filename, fileContent, common.Config.FileStore)
	if err != nil {
		ctx.JSON(200, newResponse(1000102, err.Error(), nil))
		return
	}
	ctx.JSON(200, newResponse(100000, "上传文件成功", fid))
	return
}
