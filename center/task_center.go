package center

import "email-center/utils"

//TaskCenter() ...创建任务中心
func TaskCenter() {
	task := CreateLoadEmailBody()
	homes := []EmailFile{{
		path: "D:\\gowork\\src\\email-center\\data\\发票类\\异常",
		name: "发票类",
		tag:  utils.InvalidTag,
	}}

	task.SetDataPath(homes)
	task.ReadEmailData()
	//task.ExtractEmailData()
}
