package keyrepaly

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"golangWeixin/model"
	"golangWeixin/common"
)

func KeyRpeyls(c *gin.Context) {

	// 初始化参数
	//keyReplys := make([]model.KeywordsReply, 0)

	//reply := make([]model.KeywordsReply, 0)
	var reply []model.KeywordsReply;

	if err :=common.DB.Where("status = ?" ,model.StatusNormal).Find(&reply).Error;err != nil{
		common.SendErrJSON("查找全部用户出错", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"errNo": common.ErrorCode.SUCCESS,
		"msg":   "success",
		"data": gin.H{
			"keyplies": reply,
			"total": len(reply),
		},
	})

}