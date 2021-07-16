package center

//TaskCenter() ...创建任务中心
func TaskCenter() {
	task := CreateLoadEmailBody()
	homes := []EmailFile{{
		path: "D:\\gowork\\src\\email-center\\data\\金融诈骗",
		name: "金融",
	}}
	task.SetDataPath(homes)
	task.ReadEmailData()
}
