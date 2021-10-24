package handle

import (
	"email-center/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetPartition ...
func GetPartition(c *gin.Context) {
	result := map[string]int{}
	items, err := model.SituationForm{}.GetAllItems()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	for _, v := range items {
		result[v.Name] = v.Amount
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": result})
}

//GetTopSender ...
func GetTopSender(c *gin.Context) {
	idx := "sender"
	items, err := model.SituationTop{}.GetAllItems(idx)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
}

//GetTopDepartment ... 获取异常的分行组织机构
func GetTopDepartment(c *gin.Context) {
	idx := "department"
	items, err := model.SituationTop{}.GetAllItems(idx)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
}

//GetPeriod ... 获取异常的分行组织机构
func GetPeriod(c *gin.Context) {
	result := map[string][]*model.SituationPeriod{}
	items, err := model.SituationPeriod{}.GetAllItems(30)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	for _, v := range items {
		periodItems, ok := result[v.Name]
		if ok {
			periodItems = append(periodItems, v)
			result[v.Name] = periodItems
		} else {
			result[v.Name] = []*model.SituationPeriod{v}
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": result})
}
