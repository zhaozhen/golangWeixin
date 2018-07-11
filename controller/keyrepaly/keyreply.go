package keyrepaly

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"golangWeixin/model"
	"golangWeixin/common"
	"strconv"
	"github.com/gin-gonic/gin/binding"
	"fmt"
	"github.com/jinzhu/gorm"
	"golangWeixin/utils"
	"golangWeixin/constant"
)

type KeyReplyVo struct {
	ID              int              `json:"id"`
	Key             string           `json:"key"`
	MsgType         int              `json:"msg_type"` // text,image,voice, --video,music,news
	Value           string           `json:"value"`
	KeyReplyNewsVos []KeyReplyNewsVo `json:"key_reply_news_vos,omitempty"`
	KeyReplyMusicVo KeyReplyMusicVo  `json:"key_reply_music_vo,omitempty"`
	KeyReplyVedioVo KeyReplyVedioVo  `json:"key_reply_vedio_vo,omitempty"`
	Status          int              `json:"method"`
}

//music
type KeyReplyMusicVo struct {
	ID           int    `json:"id"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	MusicUrl     string `json:"music_url,omitempty"`
	HqMusicUrl   string `json:"hq_music_url,omitempty"`
	ThumbMediaId string `json:"thumb_media_id,omitempty"`
	ReplyId      int    `json:"reply_id,omitempty"`
}

//news
type KeyReplyNewsVo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PicUrl      string `json:"pic_url"`
	Url         string `json:"url"`
	ReplyId     int    `json:"reply_id"`
}

//vedio
type KeyReplyVedioVo struct {
	ID          int    `json:"id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	MediaId     string `json:"media_id,omitempty"`
	ReplyId     int    `json:"reply_id,omitempty"`
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

	for _, entity := range *reply {
		var keyReply KeyReplyVo

		if err := utils.TransformToOther(&entity, &keyReply); err != nil {
			common.SendErrJSON("序列化出错", c)
			return
		}
		//添加相应的子类

		switch entity.MsgType {
		case model.KeywordsReplyMsgVideo:
			var keyReplyVideo KeyReplyVedioVo
			//查找数据
			video, err := model.FindKeywordsReplyVideoSubByReplyId(true, entity.ID)
			if err != nil {
				common.SendErrJSON("查找子类video出错", c)
				return
			}

			if err := utils.TransformToOther(&video, &keyReplyVideo); err != nil {
				common.SendErrJSON("序列化出错", c)
				return
			}
			//复值
			keyReply.KeyReplyVedioVo = keyReplyVideo
			break
		case model.KeywordsReplyMsgMusic:

			var keyReplyMusic KeyReplyMusicVo
			//查找数据
			music, err := model.FindKeywordsReplyMusicSubByReplyId(true, entity.ID)
			if err != nil {
				common.SendErrJSON("查找子类Music出错", c)
				return
			}

			if err := utils.TransformToOther(&music, &keyReplyMusic); err != nil {
				common.SendErrJSON("序列化出错", c)
				return
			}
			//复值
			keyReply.KeyReplyMusicVo = keyReplyMusic
			break
		case model.KeywordsReplyMsgNews:

			var keyReplyNews []KeyReplyNewsVo
			var keyReplyNewEntity []model.KeywordsReplyNewsSub
			//查找数据
			err := model.FindOneByFiledAll("keywords_reply_news_sub", "reply_id", entity.ID, &keyReplyNewEntity)
			if err != nil {
				common.SendErrJSON("查找子类New出错", c)
				return
			}
			if err := utils.TransformToOther(&keyReplyNewEntity, &keyReplyNews); err != nil {
				common.SendErrJSON("序列化出错", c)
				return
			}
			//复值
			keyReply.KeyReplyNewsVos = keyReplyNews
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

func Test(c *gin.Context) {
	//queryString, _ := c.GetQuery("name")
	//var table model.KeywordsReply
	//model.FindOne(&table,true,"1")
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"errNo": common.ErrorCode.SUCCESS,
	//	"msg":   "success",
	//	"data": table,
	//})

	var key model.KeywordsReply

	//核武器
	if err := model.FindOneByFiled("keywords_reply", "id", 15, &key); err != nil {
		return
	}

	fmt.Printf(key.Value)

	c.JSON(http.StatusOK, gin.H{
		"errNo": common.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  "success",
	})
}

//该接口新增关键字
func KeyReplyAddAndUpdate(c *gin.Context) {
	// 使用gorm的事务管理框架
	tx := common.DB.Begin()

	SendErrJSON := common.SendErrJSON
	var keyReplyVo KeyReplyVo

	if err := c.ShouldBindWith(&keyReplyVo, binding.JSON); err != nil {
		fmt.Println(err.Error())
		SendErrJSON("解析参数有误", c)
		return
	}
	// 新增更新的时候，重复的关键字判断
	if keyReplyVo.ID == 0 {
		if _, err := model.FindKeyWordReplyByKey(keyReplyVo.Key); err != gorm.ErrRecordNotFound {
			SendErrJSON("重复的关键字", c)
			return
		}
	}

	//新增有4类
	//添加主表
	var key *model.KeywordsReply

	if err := utils.TransformToOther(&keyReplyVo, &key); err != nil {
		common.SendErrJSON("序列化出错", c)
		return
	}

	if keyReplyVo.Status == -1 {
		err := key.Delete(tx, model.SystemUser)
		if err != nil {
			common.SendErrJSON("删除失败")
		}
	} else if keyReplyVo.ID > 0 {
		err := key.Update(tx, model.SystemUser)
		if err != nil {
			common.SendErrJSON("更新失败")
		}
	} else if keyReplyVo.ID == 0 {
		err := key.Insert(tx, model.SystemUser)
		if err != nil {
			common.SendErrJSON("保存失败")
		}
	}

	//得到原始的值
	oldKey, err := model.FindKeyWordReplyByOne(true, keyReplyVo.ID)
	if err != nil && err == gorm.ErrRecordNotFound {
		common.SendErrJSON("查找失败")
		return
	}

	//说明更换类型了,删除对应的子表记录,否则说明没有更换类型
	if keyReplyVo.MsgType != oldKey.MsgType {
		switch oldKey.MsgType {
		case model.KeywordsReplyMsgVideo:
			var keyReplayVideo model.KeywordsReplyVideoSub
			keyReplayVideo.ReplyId = key.ID
			keyReplayVideo.Delete(tx, model.SystemUser)
			break
		case model.KeywordsReplyMsgMusic:
			var keyReplyMusic model.KeywordsReplyMusicSub
			keyReplyMusic.ReplyId = key.ID
			keyReplyMusic.Delete(tx, model.SystemUser)
			break
		case model.KeywordsReplyMsgNews:
			var keyReplayVideos []model.KeywordsReplyNewsSub
			for _, value := range keyReplayVideos {
				value.ReplyId = key.ID
				value.Delete(tx, model.SystemUser)
			}
			break
		}

		//子表为新增
		keyReplyVo.ID = constant.Zero

	}

	//更新新增对应的子表
	switch keyReplyVo.MsgType {
	case model.KeywordsReplyMsgVideo:
		var keyReplayVideo model.KeywordsReplyVideoSub
		if err := utils.TransformToOther(&keyReplyVo.KeyReplyMusicVo, &keyReplayVideo); err != nil {
			common.SendErrJSON("序列化出错", c)
			return
		}

		if keyReplyVo.Status == constant.NegativeOne {
			err := keyReplayVideo.Delete(tx, model.SystemUser)
			if err != nil {
				common.SendErrJSON("删除Video失败")
			}
		} else if keyReplyVo.ID > constant.Zero {
			err := keyReplayVideo.Update(tx, model.SystemUser)
			if err != nil {
				common.SendErrJSON("更新Video失败")
			}
		} else if keyReplyVo.ID == constant.Zero {
			keyReplayVideo.ReplyId = key.ID
			err := keyReplayVideo.Insert(tx, model.SystemUser)
			if err != nil {
				common.SendErrJSON("新增Video失败")
			}
		}
		break
	case model.KeywordsReplyMsgMusic:

		var keyReplyMusic model.KeywordsReplyMusicSub

		if err := utils.TransformToOther(&keyReplyVo.KeyReplyMusicVo, &keyReplyMusic); err != nil {
			common.SendErrJSON("序列化出错", c)
			return
		}

		if keyReplyVo.Status == constant.NegativeOne {
			err := keyReplyMusic.Delete(tx, model.SystemUser)
			if err != nil {
				common.SendErrJSON("删除Video失败")
			}
		} else if keyReplyVo.ID > constant.Zero {
			err := keyReplyMusic.Update(tx, model.SystemUser)
			if err != nil {
				common.SendErrJSON("更新Video失败")
			}
		} else if keyReplyVo.ID == constant.Zero {
			keyReplyMusic.ReplyId = key.ID
			err := keyReplyMusic.Insert(tx, model.SystemUser)
			if err != nil {
				common.SendErrJSON("新增Video失败")
			}
		}
		break
	case model.KeywordsReplyMsgNews:

		var keyReplayVideos []model.KeywordsReplyNewsSub

		if err := utils.TransformToOther(&keyReplyVo.KeyReplyNewsVos, &keyReplayVideos); err != nil {
			common.SendErrJSON("序列化出错", c)
			return
		}

		//2：填充数据
		for _, value := range keyReplayVideos {

			if keyReplyVo.ID < constant.NegativeOne {
				err := value.Delete(tx, model.SystemUser)
				if err != nil {
					common.SendErrJSON("删除Video失败")
				}
			} else if keyReplyVo.ID > constant.Zero {
				err := value.Update(tx, model.SystemUser)
				if err != nil {
					common.SendErrJSON("更新Video失败")
				}
			} else if keyReplyVo.ID == constant.Zero {
				value.ReplyId = key.ID
				if err := value.Insert(tx, model.SystemUser); err != nil {
					common.SendErrJSON("新增Video失败")
				}
			}
		}
		break
	}

	//最后提交数据
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"errNo": common.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  "success",
	})

}
