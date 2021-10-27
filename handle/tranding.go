package handle

import (
	"email-center/model"
	"email-center/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

//GetPartition ...
func GetPartition(c *gin.Context) {
	result := map[string]int{}
	items, err := model.SituationFormModel.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	for _, v := range items {
		result[v.Name] = v.Amount
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": result})
}

//GetHours ...
func GetHours(c *gin.Context) {
	items, err := model.SituationHoursModel.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
}

//GetTopSender ...
func GetTopSender(c *gin.Context) {
	idx := "sender"
	items, err := model.SituationTopModel.GetAllItems(idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
}

//GetTopDepartment ... 获取异常的分行组织机构
func GetTopDepartment(c *gin.Context) {
	idx := "department"
	items, err := model.SituationTop{}.GetAllItems(idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
}

type periodItem struct {
	EventDay string `json:"event_day"`
	Amount   int    `json:"合计"`
	Bill     int    `json:"发票"`
	Adv      int    `json:"广告"`
	Dirty    int    `json:"色情"`
}

type periodItemList []*periodItem

func (I periodItemList) Len() int {
	return len(I)
}
func (I periodItemList) Less(i, j int) bool {
	return I[i].EventDay < I[j].EventDay
}
func (I periodItemList) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}

//GetPeriod ... 获取异常的分行组织机构
func GetPeriod(c *gin.Context) {
	result := map[string]map[string]int{}
	var resultsWeb periodItemList
	items, err := model.SituationPeriod{}.GetAllItems(30)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	for _, v := range items {
		dayIdx := v.Event.Format(utils.DayCommonFormat)
		node, ok := result[dayIdx]
		if ok {
			node[v.Name] = v.Amount
			node["amount"] = node["amount"] + v.Amount
			result[dayIdx] = node
			continue
		}
		node = map[string]int{
			v.Name:   v.Amount,
			"amount": v.Amount,
		}
		result[dayIdx] = node
	}
	//做一下转换
	for k, v := range result {
		node := &periodItem{
			EventDay: k,
			Amount:   v["amount"],
			Bill:     v["发票"],
			Adv:      v["广告"],
			Dirty:    v["色情"],
		}
		resultsWeb = append(resultsWeb, node)
	}
	sort.Sort(resultsWeb)

	for k, v := range resultsWeb {
		fmt.Println(k, *v)
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": resultsWeb})
}

//GetEmailAmount ... 获取异常的分行组织机构
func GetEmailAmount(c *gin.Context) {
	items, err := model.SituationAmountModel.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
}
