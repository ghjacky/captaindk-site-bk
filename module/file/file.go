package file

import (
	"captaindk.site/model"
	"captaindk.site/utils"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"os"
	"path"
)

func StoreFile(name string, content []byte, dir string) (id uuid.UUID, err error) {
	f := &model.SFile{ID: uuid.NewV4(), Name: name, Dir: dir}
	if e := f.Add(); e != nil {
		return f.ID, utils.WrapError(e, "无法保存文件")
	}
	return f.ID, utils.WrapError(ioutil.WriteFile(path.Join(dir, f.ID.String()), content, 0644), "无法保存文件")
}

func GetFile(file *model.SFile) (err error) {
	if e := file.Get(); e != nil {
		return utils.WrapError(e, "无法获取文件信息")
	}
	f, e := os.Open(path.Join(file.Dir, file.ID.String()))
	if e != nil {
		return utils.WrapError(e, "无法读取文件内容")
	}
	content, e := ioutil.ReadAll(f)
	if e != nil {
		return utils.WrapError(e, "无法读取文件内容")
	}
	file.Content = content
	return nil
}

func RemoveFile(file *model.SFile) (err error) {
	// 首先删除库中记录
	if e := file.Delete(); e != nil {
		return utils.WrapError(e, "文件删除失败")
	}
	// 删除存储的文件
	return utils.WrapError(os.Remove(path.Join(file.Dir, file.ID.String())), "文件删除失败")
}
