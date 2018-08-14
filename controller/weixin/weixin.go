package weixin

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"golangWeixin/utils"
	"golangWeixin/constant"
	"github.com/gin-gonic/gin/binding"
	"encoding/xml"
	"time"
	"golangWeixin/model"
)

type WeixinParam struct {
	Signature string `form:"signature" json:"signature" binding:"exists"`
	Timestamp string `form:"timestamp" json:"timestamp" binding:"exists"`
	Nonce     string `form:"nonce" json:"nonce" binding:"exists"`
	Echostr   string `form:"echostr" json:"echostr" binding:"exists"`
}
type WeixinPostParam struct {
	ToUserName   string `binding:"exists"`
	FromUserName string `binding:"exists"`
	Content      string `binding:"exists"`
}

type WeixinTextData struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}


func WeixinAction(c *gin.Context) {

	if c.Request.Method == "GET" {
		var weixinParam WeixinParam
		if err := c.ShouldBindWith(&weixinParam, binding.Form); err != nil {
			fmt.Println(err.Error())
			fmt.Println("解析参数出错！")
			return
		}
		s := utils.SignatureMethod(constant.WEIXIN_TOKEN, weixinParam.Timestamp, weixinParam.Nonce)
		if weixinParam.Signature == s {
			c.Writer.WriteString(weixinParam.Echostr)
		}
	} else {
		var weixinPostParam WeixinPostParam
		if err := c.ShouldBindBodyWith(&weixinPostParam, binding.XML); err != nil {
			fmt.Println(err.Error())
			fmt.Println("解析参数出错！")
			return
		}
		var toUserName=weixinPostParam.FromUserName
		var fromUserName=weixinPostParam.ToUserName

		var key model.KeywordsReply
		if err := model.FindOneByFiled("keywords_reply", "key_word", weixinPostParam.Content, &key); err != nil {
			return
		}

		switch key.MsgType {
		case model.KeywordsReplyMsgText:
			var data = WeixinTextData{
				ToUserName:   fromUserName,
				FromUserName: toUserName,
				Content:      key.Value,
				MsgType:      "text",
				CreateTime:   time.Now().Unix(),
			}
			d, _ := xml.Marshal(&data)
			c.Writer.WriteString(string(d))
			break
		case model.KeywordsReplyMsgImage:
			break
		case model.KeywordsReplyMsgVoice:
			break
		case model.KeywordsReplyMsgVideo:
			break
		case model.KeywordsReplyMsgMusic:
			break
		case model.KeywordsReplyMsgNews:
			break
		default:
			var data = WeixinTextData{
				ToUserName:   fromUserName,
				FromUserName: toUserName,
				Content:      "what?",
				MsgType:      "text",
				CreateTime:   time.Now().Unix(),
			}
			d, _ := xml.Marshal(&data)
			c.Writer.WriteString(string(d))
			break
		}


	}

}
