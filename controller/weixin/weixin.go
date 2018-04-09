package weixin

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"golangWeixin/utils"
	"golangWeixin/constant"
)

type WeixinParam struct {
	Signature string `json:"signature" binding:"required"`
	Timestamp string `json:"timestamp" binding:"required"`
	Nonce     string `json:"nonce" binding:"required"`
	Echostr   string `json:"echostr" binding:"required"`
}

func WeixinAction(c *gin.Context)  {
	var weixinParam WeixinParam
	if err := c.ShouldBindWith(&weixinParam, binding.Query); err != nil {
		fmt.Println(err.Error())
		fmt.Println("解析参数出错！")
		return
	}
	s:=utils.SignatureMethod(constant.WEIXIN_TOKEN,weixinParam.Timestamp,weixinParam.Nonce)
	if weixinParam.Signature==s {
		c.Writer.WriteString(weixinParam.Echostr)
	}
}