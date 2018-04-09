package utils


import (
	"sort"
	"crypto/sha1"
	"io"
	"fmt"
	"golangDemo/common"
	"github.com/garyburd/redigo/redis"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint   `json:"expires_in"`
}

func SignatureMethod(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetAccessToken() string {
	accessToken, err := redis.String(common.C.Do("GET", "access_token"))
	if err != nil {
		fmt.Println("redis get failed----------->_<----------:", err)
	}
	if len(accessToken)<=0{
		//没有accessToken，从新获取token，并设置到redis里面做个缓存
		url:=fmt.Sprintf(common.WEIXIN_HTTP_TOKEN,common.WEIXIN_APPID,common.WEIXIN_APP_SECRET)

		resp,err:=http.Get(url)
		if err !=nil {
			fmt.Println("get token fail----------->_<----------:", err)
		}
		var token Token
		data, err := ioutil.ReadAll(resp.Body)
		if err == nil && data != nil {
			err = json.Unmarshal(data, &token)
		}
		//设置token过期时间ex为秒，px为毫秒
		common.C.Do("SET","access_token",token.AccessToken,"EX",token.ExpiresIn)
		//返回结果设置值
		accessToken=token.AccessToken

		defer resp.Body.Close()
	}
	return accessToken
}