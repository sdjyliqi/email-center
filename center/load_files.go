package center

import (
	"email-center/utils"
	"fmt"
	"github.com/golang/glog"
)

//定义邮件合法性标记
type LegalTag int8

const (
	UnknownTag LegalTag = 0 //无状态
	ValidTag   LegalTag = 1 //合法
	InvalidTag LegalTag = 2 //非法
)

type EmailFile struct {
	path string   //设置邮件分类的目录
	name string   //设置邮件分类的名称
	tag  LegalTag //按照目录导入的时候，可能已经对内容进行了合法性分类
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

//SetDataPath ...设置数据文件配置信息
func (l *LoadEmailBody) SetDataPath(item []EmailFile) error {
	l.emailItems = item
	return nil
}

func (l *LoadEmailBody) ReadEmailData() error {
	for _, v := range l.emailItems {
		files, err := utils.GetFileNames(v.path, "")
		if err != nil {
			glog.Errorf("Call utils.GetFileNames failed,err:%+v", err)
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
			fmt.Println("=====LIQI=====", info)
			return nil
			//err = model.BodyModel.NoExistedInsertItem(info)
			//if err != nil {
			//	glog.Errorf("Call model.NoExistedInsertItem failed,err:%+v", err)
			//}
		}
	}
	return nil
}
