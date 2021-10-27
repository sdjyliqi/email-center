package handle

import (
	"email-center/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"strconv"
)

//GetDirtyWords ... 按页获取
func GetDirtyWords(c *gin.Context) {
	pageID, entry := 0, 0
	var err error
	strPageID, ok := c.GetQuery("page")
	if ok {
		pageID, err = strconv.Atoi(strPageID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": err.Error(), "data": nil})
			return
		}
	}
	strEntry, ok := c.GetQuery("entry")
	if ok {
		entry, err = strconv.Atoi(strEntry)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": err.Error(), "data": nil})
			return
		}
	}
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
	strID, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "id invalid", "data": nil})
		return
	}
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	err = model.DirtyModel.DelItemByID(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ""})
}

//UpInsertDirtyWords ... 获取异常的分行组织机构
func UpInsertDirtyWords(c *gin.Context) {
	reqJson := &model.Dirty{}
	err := c.ShouldBindJSON(reqJson)
	if err != nil {
		glog.Errorf("The request %+v is invalid,please check.", c.Request)
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "bind json failed."})
		return
	}
	err = model.DirtyModel.UpdateItemByID(reqJson)
	if err != nil {
		glog.Errorf("The request %+v is invalid,please check.", c.Request)
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "bind json failed."})
		return
	}
}
