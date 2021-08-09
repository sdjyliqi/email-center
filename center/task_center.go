package center

//TaskCenter() ...创建任务中心
func TaskCenter() {
	task := CreateLoadEmailBody()
	task.ExtractEmailData()
}
