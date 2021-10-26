package handle

import (
	"email-center/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetDirtyWords ... 按页获取
func GetDirtyWords(c *gin.Context) {
	pageID := c.GetInt("page")
	entry := c.GetInt("entry")
	if entry > 100 || entry <= 1 {
		entry = 20
	}
	items, err := model.DirtyModel.GetItemsByPage(pageID, entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	cnt, err := model.DirtyModel.GetItemsCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": gin.H{"items:": items, "amount": cnt}})
}

//SearchDirtyWords ...
func SearchDirtyWords(c *gin.Context) {
	keywords, _ := c.GetQuery("idx")
	if keywords == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "bad request"})
	}
	items, err := model.DirtyModel.SearchItemsByIdx(keywords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
}

//DelDirtyWords ... 获取异常的分行组织机构
func DelDirtyWords(c *gin.Context) {
	id := c.GetInt("id")
	fmt.Println(id)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ""})
}
