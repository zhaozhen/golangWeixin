package keyrepaly

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"golangWeixin/model"
	"golangWeixin/common"
	"strconv"
	"github.com/gin-gonic/gin/binding"
	"fmt"
	"github.com/vmihailenco/msgpack"
	"time"
)

type KeyReplyVo struct {
	ID      uint   `json:"id"`
	Key     string `json:"key"`
	MsgType int    `json:"msg_type"` // text,image,voice, --video,music,news
	Value   string `json:"value"`
	KeyReplyNewsVos []KeyReplyNewsVo
	KeyReplyMusicVo
	KeyReplyVedioVo
}

//music
type KeyReplyMusicVo struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	MusicUrl     string `json:"music_url"`
	HqMusicUrl   string `json:"hq_music_url"`
	ThumbMediaId string `json:"thumb_media_id"`
	ReplyId      string `json:"reply_id"`
}

//news
type KeyReplyNewsVo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PicUrl      string `json:"pic_url"`
	Url         string `json:"url"`
	ReplyId     string `json:"reply_id"`
}

//vedio
type KeyReplyVedioVo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	MediaId     string `json:"media_id"`
	ReplyId     string `json:"reply_id"`
}





//根据关键字和page，limit，进行分页模糊查找
func KeyRpeyls(c *gin.Context) {

	queryString, _ := c.GetQuery("name")
	pageQ := c.DefaultQuery("page", "1")
	limitQ := c.DefaultQuery("limit", "20")

	page, _ := strconv.Atoi(pageQ)
	limit, _ := strconv.Atoi(limitQ)

	reply, err := model.FindAllKeysReplyPage(queryString, page, limit);
	if err != nil {
		common.SendErrJSON("查找全部用户出错", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": common.ErrorCode.SUCCESS,
		"msg":   "success",
		"data": gin.H{
			"keyplies": reply,
			"total":    len(reply),
		},
	})

}

//该接口新增关键字
func KeyReplyAdd(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	var keyReplyVo KeyReplyVo;

	if err := c.ShouldBindWith(&keyReplyVo, binding.JSON); err != nil {
		fmt.Println(err.Error())
		SendErrJSON("解析参数有误", c)
		return
	}
	//新增有4类

	switch keyReplyVo.MsgType {
	case model.KeywordsReplyMsgText:
	case model.KeywordsReplyMsgImage:
	case model.KeywordsReplyMsgVoice:
		// 1：text,image,voice,
		//var keyReply model.KeywordsReply
		//data, err := msgpack.Marshal(&keyReplyVo)
		//if err == nil {
		//	fmt.Println("err:", "序列化成为model类型出错")
		//}
		//err = msgpack.Unmarshal(data, &keyReply)
		//if err != nil {
		//	panic(err)
		//}
		////2：填充数据
		//keyReply.CreatedAt=time.Now()
		//keyReply.CreatedPerson=model.SystemUser
		//keyReply.Status=model.StatusNormal
		////3：新增数据
		//keyReply.Insert()
		keyReplyAddfunc(&keyReplyVo)

	case model.KeywordsReplyMsgVideo:
		//1,保存关机那字
		_,err:=keyReplyAddfunc(&keyReplyVo)
		if err!=nil {
			common.SendErrJSON("保存失败")
		}



		// 2：video
	case model.KeywordsReplyMsgMusic:
		// 3：music
		break //可以添加
	case model.KeywordsReplyMsgNews:
		// 4：news
	default:
		SendErrJSON("你有毒？", c)
	}

}

func keyReplyAddfunc(key *KeyReplyVo) (*model.KeywordsReply, error) {
	var keyReply model.KeywordsReply

	data, err := msgpack.Marshal(&key)
	if err == nil {
		fmt.Println("err:", "序列化出错")
		return nil, err
	}
	err = msgpack.Unmarshal(data, &keyReply)
	fmt.Println("err:", "结序列化出错")
	if err != nil {
		panic(err)
	}
	//2：填充数据
	keyReply.CreatedAt = time.Now()
	keyReply.CreatedPerson = model.SystemUser
	keyReply.Status = model.StatusNormal
	//3：新增数据
	keyReply.Insert()

	return &keyReply, nil
}

