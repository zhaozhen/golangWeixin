package weixin

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"golangWeixin/utils"
	"golangWeixin/constant"
	"github.com/gin-gonic/gin/binding"
)

type WeixinParam struct {
	Signature string `form:"signature" json:"signature" binding:"exists"`
	Timestamp string `form:"timestamp" json:"timestamp" binding:"exists"`
	Nonce     string `form:"nonce" json:"nonce" binding:"exists"`
	Echostr   string `form:"echostr" json:"echostr" binding:"exists"`
}

func WeixinAction(c *gin.Context)  {
	var weixinParam WeixinParam
	if err := c.ShouldBindWith(&weixinParam,binding.Form); err != nil {
		fmt.Println(err.Error())
		fmt.Println("解析参数出错！")
		return
	}
	s:=utils.SignatureMethod(constant.WEIXIN_TOKEN,weixinParam.Timestamp,weixinParam.Nonce)
	if weixinParam.Signature==s {
		c.Writer.WriteString(weixinParam.Echostr)
	}
}