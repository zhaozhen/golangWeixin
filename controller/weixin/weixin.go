package weixin

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"golangWeixin/common"
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

type WeixinBaseData struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
}

type WeixinImgData struct {
	WeixinBaseData
	Image Image `xml:"Image"`
}

type Image struct {
	MediaId string `xml:"MediaId"`
}

type WeixinMusicData struct {
	WeixinBaseData
	Music Music `xml:"Music"`
}

type Music struct {
	Title        string `xml:"Title"`
	Description  string `xml:"Description"`
	MusicUrl     string `xml:"MusicUrl"`
	HQMusicUrl   string `xml:"HQMusicUrl"`
	ThumbMediaId string `xml:"ThumbMediaId"`
}

type WeixinNewsData struct {
	WeixinBaseData
	ArticleCount string `xml:"ArticleCount"`
	Articles     Item   `xml:"Articles"`
}

type Item struct {
	XMLName    xml.Name `xml:"Articles"`
	News []News `xml:"item"`
}
type News struct {
	XMLName    xml.Name `xml:"item"`
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl      string `xml:"PicUrl"`
	Url         string `xml:"Url"`
}

//6590986134459777000
func WeixinAction(c *gin.Context) {

	if c.Request.Method == "GET" {
		var weixinParam WeixinParam
		if err := c.ShouldBindWith(&weixinParam, binding.Form); err != nil {
			fmt.Println(err.Error())
			fmt.Println("解析参数出错！")
			return
		}
		s := common.SignatureMethod(constant.WEIXIN_TOKEN, weixinParam.Timestamp, weixinParam.Nonce)
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
			key.MsgType = 9
		}

		switch key.MsgType {
		case model.KeywordsReplyMsgText:
			var data = WeixinTextData{
				ToUserName:   toUserName,
				FromUserName: fromUserName,
				Content:      key.Value,
				MsgType:      "text",
				CreateTime:   time.Now().Unix(),
			}
			d, _ := xml.Marshal(&data)
			fmt.Printf(string(d))
			c.Writer.WriteString(string(d))
			break
		case model.KeywordsReplyMsgImage:
			var data = WeixinImgData{}
			data.ToUserName = toUserName
			data.FromUserName = fromUserName
			data.MsgType = "image"
			data.CreateTime = time.Now().Unix()
			data.Image.MediaId = "100000001"
			d, _ := xml.Marshal(&data)
			fmt.Printf(string(d))
			c.Writer.WriteString(string(d))
			break
		case model.KeywordsReplyMsgVoice:
			break
		case model.KeywordsReplyMsgVideo:
			break
		case model.KeywordsReplyMsgMusic:
			break
		case model.KeywordsReplyMsgNews:
			subs, err := model.FindKeywordsReplyNewsSubByReplyId(true, key.ID)
			if err != nil {
				break
			}
			var data = WeixinNewsData{}
			data.ToUserName = toUserName
			data.FromUserName = fromUserName
			data.MsgType = "news"
			data.CreateTime = time.Now().Unix()
			data.ArticleCount = "3"

			var items Item
			for _, value := range *subs {
				var new News
				new.Title = value.Title
				new.Description = value.Description
				new.PicUrl = value.PicUrl
				new.Url = value.Url
				items.News=append(items.News,new)
			}
			data.Articles=items
			d, _ := xml.Marshal(&data)
			fmt.Printf(string(d))
			c.Writer.WriteString(string(d))
			break
		default:
			var data = WeixinTextData{
				ToUserName:   toUserName,
				FromUserName: fromUserName,
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
