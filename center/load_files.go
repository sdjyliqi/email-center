package center

import (
	"email-center/model"
	"email-center/utils"
	"github.com/golang/glog"
)

type EmailFile struct {
	path string //设置邮件分类的目录
	name string //设置邮件分类的名称
}

type LoadEmailBody struct {
	emailItems []EmailFile
}

//CreateLoadEmailBody 创建结构体
func CreateLoadEmailBody() *LoadEmailBody {
	return &LoadEmailBody{
		emailItems: []EmailFile{},
	}
}

//
func (l *LoadEmailBody) SetDataPath(item []EmailFile) error {
	l.emailItems = item
	return nil
}

func (l *LoadEmailBody) ReadEmailData() error {
	for _, v := range l.emailItems {
		files, err := utils.GetFileNames(v.path, "")
		if err != nil {
			glog.Error("Call utils.GetFileNames failed,err:%+v", err)
			return err
		}
		for _, vv := range files {
			info, err := utils.PickupEmail(vv)
			if err != nil {
				glog.Errorf("Call utils.PickupEmail failed,err:%+v", err)
				return err
			}
			info.Category = v.name
			info.FileName = vv
			err = model.BodyModel.NoExistedInsertItem(info)
			if err != nil {
				glog.Errorf("Call model.NoExistedInsertItem failed,err:%+v", err)
			}
		}
	}
	return nil
}
