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
	ID              uint             `json:"id"`
	Key             string           `json:"key"`
	MsgType         int              `json:"msg_type"` // text,image,voice, --video,music,news
	Value           string           `json:"value"`
	KeyReplyNewsVos []KeyReplyNewsVo `json:"keyReplyNewsVos"`
	KeyReplyMusicVo                  `json:"KeyReplyMusicVo"`
	KeyReplyVedioVo                  `json:"KeyReplyVedioVo"`
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
	pageQ := c.DefaultQuery("page", "0")
	limitQ := c.DefaultQuery("limit", "20")

	page, _ := strconv.Atoi(pageQ)
	limit, _ := strconv.Atoi(limitQ)

	reply, err := model.FindAllKeysReplyPage(queryString, page, limit)
	if err != nil {
		common.SendErrJSON("查找全部用户出错", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": common.ErrorCode.SUCCESS,
		"msg":   "success",
		"data": gin.H{
			"keyplies": reply,
			"total":len(reply),
		},
	})

}

//该接口新增关键字
func KeyReplyAdd(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	var keyReplyVo KeyReplyVo

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
		break
	case model.KeywordsReplyMsgVideo:
		//1,保存关键字
		key, err := keyReplyAddfunc(&keyReplyVo)
		if err != nil {
			common.SendErrJSON("保存失败")
		}
		var keyReplayVideo model.KeywordsReplyVideoSub
		data, err := msgpack.Marshal(&keyReplyVo.KeyReplyMusicVo)
		if err == nil {
			fmt.Println("err:", "序列化出错")
		}
		err = msgpack.Unmarshal(data, &keyReplayVideo)
		fmt.Println("err:", "解序列化出错")
		if err != nil {
			panic(err)
		}
		//2：填充数据
		keyReplayVideo.CreatedAt = time.Now()
		keyReplayVideo.CreatedPerson = model.SystemUser
		keyReplayVideo.Status = model.StatusNormal
		keyReplayVideo.ReplyId = key.ID
		//3：新增数据
		keyReplayVideo.Insert()
		break

		// 2：video
	case model.KeywordsReplyMsgMusic:

		key, err := keyReplyAddfunc(&keyReplyVo)
		if err != nil {
			common.SendErrJSON("保存失败")
		}
		var keyReplyMusic model.KeywordsReplyMusicSub
		data, err := msgpack.Marshal(&keyReplyVo.KeyReplyMusicVo)
		if err == nil {
			fmt.Println("err:", "序列化出错")
		}
		err = msgpack.Unmarshal(data, &keyReplyMusic)
		fmt.Println("err:", "解序列化出错")
		if err != nil {
			panic(err)
		}
		//2：填充数据
		keyReplyMusic.CreatedAt = time.Now()
		keyReplyMusic.CreatedPerson = model.SystemUser
		keyReplyMusic.Status = model.StatusNormal
		keyReplyMusic.ReplyId = key.ID
		//3：新增数据
		keyReplyMusic.Insert()
		break //可以添加

	case model.KeywordsReplyMsgNews:
		// 4：news

		key, err := keyReplyAddfunc(&keyReplyVo)
		if err != nil {
			common.SendErrJSON("保存失败")
		}
		var keyReplayVideos []model.KeywordsReplyNewsSub
		data, err := msgpack.Marshal(&keyReplyVo.KeyReplyNewsVos)
		if err == nil {
			fmt.Println("err:", "序列化出错")
		}
		err = msgpack.Unmarshal(data, &keyReplayVideos)
		fmt.Println("err:", "解序列化出错")
		if err != nil {
			panic(err)
		}
		//2：填充数据
		for _, value := range keyReplayVideos {
			value.CreatedAt = time.Now()
			value.CreatedPerson = model.SystemUser
			value.Status = model.StatusNormal
			value.ReplyId = key.ID
			//3：新增数据
			value.Insert()
		}

	default:
		SendErrJSON("你有毒？", c)
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": common.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  "success",
	})

}

func keyReplyAddfunc(key *KeyReplyVo) (*model.KeywordsReply, error) {
	var keyReply model.KeywordsReply

	data, err := msgpack.Marshal(&key)
	if err == nil {
		fmt.Println("err:", "序列化出错")
		return nil, err
	}
	err = msgpack.Unmarshal(data, &keyReply)
	fmt.Println("err:", "解序列化出错")
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

func keyReplyUpdatefunc(key *KeyReplyVo) (*model.KeywordsReply, error) {
	var keyReply model.KeywordsReply
	data, err := msgpack.Marshal(&key)
	if err == nil {
		fmt.Println("err:", "序列化出错")
		return nil, err
	}
	err = msgpack.Unmarshal(data, &keyReply)
	fmt.Println("err:", "解序列化出错")
	if err != nil {
		panic(err)
	}
	//2：填充数据
	keyReply.UpdatedAt = time.Now()
	keyReply.UpdatedPerson = model.SystemUser
	keyReply.Status = model.StatusNormal
	//3：更新数据
	keyReply.Update()
	return &keyReply, nil
}

func keyReplyDeletefunc(key *KeyReplyVo) (*model.KeywordsReply, error) {
	var keyReply model.KeywordsReply
	data, err := msgpack.Marshal(&key)
	if err == nil {
		fmt.Println("err:", "序列化出错")
		return nil, err
	}
	err = msgpack.Unmarshal(data, &keyReply)
	fmt.Println("err:", "解序列化出错")
	if err != nil {
		panic(err)
	}
	//2：填充数据
	delteDate := time.Now()
	keyReply.DeletedAt = &delteDate
	keyReply.DeletedPerson = model.SystemUser
	keyReply.Status = model.StatusNormal
	//3：更新数据
	keyReply.Update()
	return &keyReply, nil
}

//根据子表进行更新
func KeyReplyUpdate(c *gin.Context) {
	var keyReplyVo KeyReplyVo

	if err := c.ShouldBindWith(&keyReplyVo, binding.JSON); err != nil {
		fmt.Println(err.Error())
		common.SendErrJSON("解析参数有误", c)
		return
	}
	switch keyReplyVo.MsgType {
	case model.KeywordsReplyMsgText:
	case model.KeywordsReplyMsgImage:
	case model.KeywordsReplyMsgVoice:
		keyReplyUpdatefunc(&keyReplyVo)
		break
	case model.KeywordsReplyMsgVideo:
		_, err := keyReplyUpdatefunc(&keyReplyVo)
		if err != nil {
			common.SendErrJSON("保存失败")
		}
		var keyReplyVideo model.KeywordsReplyVideoSub
		data, err := msgpack.Marshal(&keyReplyVo.KeyReplyVedioVo)
		if err == nil {
			fmt.Println("err:", "序列化出错")
		}
		err = msgpack.Unmarshal(data, &keyReplyVideo)
		fmt.Println("err:", "解序列化出错")
		if err != nil {
			panic(err)
		}
		keyReplyVideo.Update()
		break
	case model.KeywordsReplyMsgMusic:

		_, err := keyReplyUpdatefunc(&keyReplyVo)
		if err != nil {
			common.SendErrJSON("保存失败")
		}
		var keyReplyMusic model.KeywordsReplyMusicSub
		data, err := msgpack.Marshal(&keyReplyVo.KeyReplyNewsVos)
		if err == nil {
			fmt.Println("err:", "序列化出错")
		}
		err = msgpack.Unmarshal(data, &keyReplyMusic)
		fmt.Println("err:", "解序列化出错")
		if err != nil {
			panic(err)
		}

		keyReplyMusic.Update()

		break
	case model.KeywordsReplyMsgNews:
		_, err := keyReplyUpdatefunc(&keyReplyVo)
		if err != nil {
			common.SendErrJSON("保存失败")
		}
		var keyReplyNewsSubs []model.KeywordsReplyNewsSub
		data, err := msgpack.Marshal(&keyReplyVo.KeyReplyNewsVos)
		if err == nil {
			fmt.Println("err:", "序列化出错")
		}
		err = msgpack.Unmarshal(data, &keyReplyNewsSubs)
		fmt.Println("err:", "解序列化出错")
		if err != nil {
			panic(err)
		}

		for _, value := range keyReplyNewsSubs {
			value.Update()
		}

		break
	default:
		fmt.Print("------更新失败了------")
		break
	}

}

//得到对应的子表记录
func getKeyReplySubsfunc(c *gin.Context) {

	replyId, exists := c.GetQuery("reply_id")
	keyType, _ := c.GetQuery("type")

	//转换类型
	keyTypeInt, _ := strconv.Atoi(keyType)

	//var keyWord model.KeywordsReply
	if exists {
		switch keyTypeInt {
		case model.KeywordsReplyMsgVideo:
			//初始化参数
			var keySub model.KeywordsReplyVideoSub
			subKeyModel, err := model.FindKeywordsReplyVideoSubByReplyId(true, replyId)

			//序列化
			data, err := msgpack.Marshal(&subKeyModel)
			if err == nil {
				fmt.Println("err:", "序列化出错")
			}
			err = msgpack.Unmarshal(data, &keySub)
			fmt.Println("err:", "解序列化出错")
			if err != nil {
				panic(err)
			}

			//返回参数
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"errNo": common.ErrorCode.SUCCESS,
					"msg":   "success",
					"data": gin.H{
						"keySub": keySub,
					},
				})
			}
			break
			//1,保存关键字
		case model.KeywordsReplyMsgMusic:
			//初始化
			var keySub model.KeywordsReplyMusicSub
			subKeyModel, err := model.FindKeywordsReplyMusicSubByReplyId(true, replyId)

			//序列化
			data, err := msgpack.Marshal(&subKeyModel)
			if err == nil {
				fmt.Println("err:", "序列化出错")
			}
			err = msgpack.Unmarshal(data, &keySub)
			fmt.Println("err:", "解序列化出错")
			if err != nil {
				panic(err)
			}

			//返回参数
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"errNo": common.ErrorCode.SUCCESS,
					"msg":   "success",
					"data": gin.H{
						"keySub": keySub,
					},
				})
			}
			break
		case model.KeywordsReplyMsgNews:
			var keySubs []model.KeywordsReplyNewsSub
			subKeyModel, err := model.FindKeywordsReplyNewsSubByReplyId(true, replyId)
			//序列化
			data, err := msgpack.Marshal(&subKeyModel)
			if err == nil {
				fmt.Println("err:", "序列化出错")
			}
			err = msgpack.Unmarshal(data, &keySubs)
			fmt.Println("err:", "解序列化出错")
			if err != nil {
				panic(err)
			}
			//返回参数
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"errNo": common.ErrorCode.SUCCESS,
					"msg":   "success",
					"data": gin.H{
						"keySub": keySubs,
					},
				})
			}
			break
		default:
			common.SendErrJSON("子表回复类型错误", c)
		}
	}

}
