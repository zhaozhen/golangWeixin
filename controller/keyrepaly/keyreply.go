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
	"golangWeixin/constant"
)

type KeyReplyVo struct {
	ID              int              `json:"id"`
	Key             string           `json:"key"`
	MsgType         int              `json:"msg_type"` // text,image,voice, --video,music,news
	Value           string           `json:"value"`
	KeyReplyNewsVos []KeyReplyNewsVo `json:"key_reply_news_vos"`
	KeyReplyMusicVo KeyReplyMusicVo  `json:"Key_reply_music_vo"`
	KeyReplyVedioVo KeyReplyVedioVo  `json:"Key_reply_vedio_vo"`
	//HttpMethod      string           `json:"method"`
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

	// entity_model
	var keyReplys []KeyReplyVo

	for _, entity := range reply {
		var keyReply KeyReplyVo
		data, err := msgpack.Marshal(&entity)
		if err != nil {
			common.SendErrJSON("序列化出错", c)
			return
		}
		err = msgpack.Unmarshal(data, &keyReply)
		if err != nil {
			common.SendErrJSON("解序列化出错", c)
			return
		}
		//序列化之后清除掉空值
		//&keyReply.KeyReplyMusicVo=nil
		//&keyReply.KeyReplyVedioVo=nil

		//添加相应的子类

		switch entity.MsgType {
		case model.KeywordsReplyMsgVideo:
			var keyReplyVideo KeyReplyVedioVo
			//查找数据
			video, err := model.FindKeywordsReplyVideoSubByReplyId(true, entity.ID)
			if err != nil {
				common.SendErrJSON("查找子类video出错", c)
			}

			//序列化成data
			data, err := msgpack.Marshal(&video)
			if err != nil {
				common.SendErrJSON("序列化出错", c)
			}
			err = msgpack.Unmarshal(data, &keyReplyVideo)
			if err != nil {
				common.SendErrJSON("解序列化出错", c)
			}
			//复值
			keyReply.KeyReplyVedioVo = keyReplyVideo

			//添加相应的切片
			keyReplys = append(keyReplys, keyReply)
			break
		case model.KeywordsReplyMsgMusic:

			var keyReplyMusic KeyReplyMusicVo
			//查找数据
			music, err := model.FindKeywordsReplyMusicSubByReplyId(true, entity.ID)
			if err != nil {
				common.SendErrJSON("查找子类video出错", c)
			}

			//序列化成data
			data, err := msgpack.Marshal(&music)
			if err != nil {
				common.SendErrJSON("序列化出错", c)
			}
			err = msgpack.Unmarshal(data, &keyReplyMusic)
			if err != nil {
				common.SendErrJSON("解序列化出错", c)
			}
			//复值
			keyReply.KeyReplyMusicVo = keyReplyMusic

			//添加相应的切片
			keyReplys = append(keyReplys, keyReply)
			break
		case model.KeywordsReplyMsgNews:

			var keyReplyNews []KeyReplyNewsVo
			//查找数据
			news, err := model.FindKeywordsReplyNewsSubByReplyId(true, entity.ID)
			if err != nil {
				common.SendErrJSON("查找子类video出错", c)
			}

			//序列化成data
			data, err := msgpack.Marshal(&news)
			if err != nil {
				common.SendErrJSON("序列化出错", c)
			}
			err = msgpack.Unmarshal(data, &keyReplyNews)
			if err != nil {
				common.SendErrJSON("解序列化出错", c)
			}
			//复值
			keyReply.KeyReplyNewsVos = keyReplyNews
			//添加相应的切片
			keyReplys = append(keyReplys, keyReply)
			break
		}
		//添加相应的切片
		keyReplys = append(keyReplys, keyReply)

	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": common.ErrorCode.SUCCESS,
		"msg":   "success",
		"data": gin.H{
			"keyplies": keyReplys,
			"total":    len(keyReplys),
		},
	})

}

func Test(c *gin.Context){
	//queryString, _ := c.GetQuery("name")
	//var table model.KeywordsReply
	//model.FindOne(&table,true,"1")
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"errNo": common.ErrorCode.SUCCESS,
	//	"msg":   "success",
	//	"data": table,
	//})
	c.JSON(http.StatusOK, gin.H{
		"errNo": common.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  "success",
	})
}


//该接口新增关键字
func KeyReplyAddAndUpdate(c *gin.Context) {
	SendErrJSON := common.SendErrJSON
	var keyReplyVo KeyReplyVo

	if err := c.ShouldBindWith(&keyReplyVo, binding.JSON); err != nil {
		fmt.Println(err.Error())
		SendErrJSON("解析参数有误", c)
		return
	}
	//新增有4类
	//添加主表
	var key *model.KeywordsReply

	data, err := msgpack.Marshal(&keyReplyVo)
	if err != nil {
		SendErrJSON("序列化出错", c)
	}
	err = msgpack.Unmarshal(data, &key)
	if err != nil {
		SendErrJSON("解序列化出错", c)
	}

	if keyReplyVo.ID > 0 {
		err := key.Update(model.SystemUser)
		if err != nil {
			common.SendErrJSON("更新失败")
		}
	} else if keyReplyVo.ID == 0 {
		err := key.Insert(model.SystemUser)
		if err != nil {
			common.SendErrJSON("保存失败")
		}
	} else if keyReplyVo.ID == -1 {
		err := key.Delete(model.SystemUser)
		if err != nil {
			common.SendErrJSON("删除失败")
		}
	}

	//得到原始的值
	oldKey, err := model.FindKeyWordReplyByOne(true, keyReplyVo.ID)
	if err != nil {
		common.SendErrJSON("查找失败")
	}

	//说明更换类型了,删除对应的子表记录
	if keyReplyVo.MsgType != oldKey.MsgType {
		switch keyReplyVo.MsgType {
		case model.KeywordsReplyMsgVideo:
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

			keyReplayVideo.Delete(model.SystemUser)
		case model.KeywordsReplyMsgMusic:
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

			keyReplyMusic.Delete(model.SystemUser)
		case model.KeywordsReplyMsgNews:
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
			for _, value := range keyReplayVideos {
				value.Delete(model.SystemUser)
			}

		}

		//子表为新增
		keyReplyVo.ID = constant.Zero

	}

	//更新新增对应的子表
	switch keyReplyVo.MsgType {
	case model.KeywordsReplyMsgVideo:
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

		if keyReplyVo.ID > constant.Zero {
			err := keyReplayVideo.Update(model.SystemUser)
			if err != nil {
				common.SendErrJSON("更新Video失败")
			}
		} else if keyReplyVo.ID == constant.Zero {
			keyReplayVideo.ID = key.ID
			err := keyReplayVideo.Insert(model.SystemUser)
			if err != nil {
				common.SendErrJSON("新增Video失败")
			}
		} else if keyReplyVo.ID < constant.NegativeOne {
			err := keyReplayVideo.Delete(model.SystemUser)
			if err != nil {
				common.SendErrJSON("删除Video失败")
			}
		}
		break
	case model.KeywordsReplyMsgMusic:

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

		if keyReplyVo.ID > constant.Zero {
			err := keyReplyMusic.Update(model.SystemUser)
			if err != nil {
				common.SendErrJSON("更新Video失败")
			}
		} else if keyReplyVo.ID == constant.Zero {
			keyReplyMusic.ID = key.ID
			err := keyReplyMusic.Insert(model.SystemUser)
			if err != nil {
				common.SendErrJSON("新增Video失败")
			}
		} else if keyReplyVo.ID < constant.NegativeOne {
			err := keyReplyMusic.Delete(model.SystemUser)
			if err != nil {
				common.SendErrJSON("删除Video失败")
			}
		}
		break
	case model.KeywordsReplyMsgNews:

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
			if keyReplyVo.ID > constant.Zero {
				err := value.Update(model.SystemUser)
				if err != nil {
					common.SendErrJSON("更新Video失败")
				}
			} else if keyReplyVo.ID == constant.Zero {
				value.ID = key.ID
				err := value.Insert(model.SystemUser)
				if err != nil {
					common.SendErrJSON("新增Video失败")
				}
			} else if keyReplyVo.ID < constant.NegativeOne {
				err := value.Delete(model.SystemUser)
				if err != nil {
					common.SendErrJSON("删除Video失败")
				}
			}
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
