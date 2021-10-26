package router

import (
	"email-center/handle"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/email/trend/proportion", handle.GetPartition)    //获取一次邮件比例
	r.GET("/email/trend/sendertop", handle.GetTopSender)     //获取头部异常发信者名单
	r.GET("/email/trend/departtop", handle.GetTopDepartment) //获取头部异常组织机构名单
	r.GET("/email/trend/period", handle.GetPeriod)           //获取异常邮件趋势图
	r.GET("/email/trend/hours", handle.GetHours)             //获取异常邮件分时趋势图
}
