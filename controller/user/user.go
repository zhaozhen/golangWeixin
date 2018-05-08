package user

import (
	"strconv"
	"time"
	"fmt"
	"net/http"
	"math/rand"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"golangWeixin/common"
	"golangWeixin/utils"
	"golangWeixin/config"
	"golangWeixin/model"
)


// 登录类型
type UsernameLogin struct {
	SigninInput string `json:"username" binding:"required,min=4,max=20"`
	Password    string `json:"password" binding:"required,min=6,max=20"`
	//LuosimaoRes string `json:"luosimaoRes"`
}
// 注册类型
type UserReqData struct {
	Name          string `json:"username" binding:"required,min=4,max=20"`
	Password      string `json:"password" binding:"required,min=6,max=20"`
	PasswordAgain string `json:"checkPassword" binding:"required,min=6,max=20"`
}

type UserVo struct {
	Name  string `json:"username" `
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

const (
	activeDuration = 24 * 60 * 60
	resetDuration  = 24 * 60 * 60
)

// Signin 用户登录
func Signin(c *gin.Context) {
	SendErrJSON := common.SendErrJSON

	//var emailLogin EmailLogin
	var usernameLogin UsernameLogin
	var signinInput string
	var password string
	var luosimaoRes string
	var sql string

	//if c.Query("loginType") == "email" {
	//	if err := c.ShouldBindWith(&emailLogin, binding.JSON); err != nil {
	//		fmt.Println(err.Error())
	//		SendErrJSON("邮箱或密码错误", c)
	//		return
	//	}
	//	signinInput = emailLogin.SigninInput
	//	password = emailLogin.Password
	//	luosimaoRes = emailLogin.LuosimaoRes
	//	sql = "email = ?"
	//} else
	//if c.Query("loginType") == "username" {
	if err := c.ShouldBindWith(&usernameLogin, binding.JSON); err != nil {
		fmt.Println(err.Error())
		SendErrJSON("用户名或密码错误", c)
		return
	}
	signinInput = usernameLogin.SigninInput
	password = usernameLogin.Password
	//luosimaoRes = usernameLogin.LuosimaoRes
	sql = "name = ?"
	//}

	verifyErr := utils.LuosimaoVerify(config.ServerConfig.LuosimaoVerifyURL, config.ServerConfig.LuosimaoAPIKey, luosimaoRes)

	if verifyErr != nil {
		SendErrJSON(verifyErr.Error(), c)
		return
	}

	var user model.User
	if err := common.DB.Where(sql, signinInput).First(&user).Error; err != nil {
		SendErrJSON("账号不存在", c)
		return
	}

	if user.CheckPassword(password) {
		if user.Status == model.UserStatusInActive {
			encodedEmail := base64.StdEncoding.EncodeToString([]byte(user.Email))
			c.JSON(200, gin.H{
				"errNo": common.ErrorCode.InActive,
				"msg":   "账号未激活",
				"data": gin.H{
					"email": encodedEmail,
				},
			})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": user.ID,
		})
		tokenString, err := token.SignedString([]byte(config.ServerConfig.TokenSecret))
		if err != nil {
			fmt.Println(err.Error())
			SendErrJSON("内部错误", c)
			return
		}

		if err := model.UserToRedis(user); err != nil {
			SendErrJSON("内部错误.", c)
			return
		}

		c.SetCookie("token", tokenString, config.ServerConfig.TokenMaxAge, "/", "", true, true)

		c.JSON(http.StatusOK, gin.H{
			"errNo": common.ErrorCode.SUCCESS,
			"msg":   "success",
			"data": gin.H{
				"token": tokenString,
				"user":  user,
			},
		})
		return
	}
	SendErrJSON("账号或密码错误", c)
}

// Register 用户注册
func Register(c *gin.Context) {
	SendErrJSON := common.SendErrJSON

	var userData UserReqData
	if err := c.ShouldBindWith(&userData, binding.JSON); err != nil {
		fmt.Println(err)
		SendErrJSON("参数无效", c)
		return
	}

	userData.Name = utils.AvoidXSS(userData.Name)
	userData.Name = strings.TrimSpace(userData.Name)
	userData.PasswordAgain = strings.TrimSpace(userData.PasswordAgain)
	userData.Password = strings.TrimSpace(userData.Password)
	if userData.Password!=userData.PasswordAgain {
		SendErrJSON("两次密码不一致", c)
		return
	}

	if strings.Index(userData.Name, "@") != -1 {
		SendErrJSON("用户名中不能含有@字符", c)
		return
	}

	var user model.User
	if err := common.DB.Where("name = ?",userData.Name).Find(&user).Error; err == nil {
		if user.Name == userData.Name {
			SendErrJSON("用户名 "+user.Name+" 已被注册", c)
			return
		}
	}

	var newUser model.User
	newUser.Name = userData.Name
	newUser.Email = "stronger@qq.com"
	newUser.Pass = newUser.EncryptPassword(userData.Password, newUser.Salt())
	newUser.Role = model.UserRoleNormal
	newUser.Status = model.UserStatusInActive
	newUser.Sex = model.UserSexMale
	newUser.AvatarURL = "/images/avatar/" + strconv.Itoa(rand.Intn(2)) + ".png"

	if err := common.DB.Create(&newUser).Error; err != nil {
		SendErrJSON("error", c)
		return
	}

	curTime := time.Now().Unix()
	activeUser := fmt.Sprintf("%s%d", model.ActiveTime, newUser.ID)

	RedisConn := common.RedisPool.Get()
	defer RedisConn.Close()

	if _, err := RedisConn.Do("SET", activeUser, curTime, "EX", activeDuration); err != nil {
		fmt.Println("redis set failed:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": common.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  newUser,
	})
}

// Signout 退出登录
func Signout(c *gin.Context) {
	userInter, exists := c.Get("user")
	var user model.User
	if exists {
		user = userInter.(model.User)

		RedisConn := common.RedisPool.Get()
		defer RedisConn.Close()

		if _, err := RedisConn.Do("DEL", fmt.Sprintf("%s%d", model.LoginUser, user.ID)); err != nil {
			fmt.Println("redis delelte failed:", err)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errNo": common.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  gin.H{},
	})
}

func Users(c *gin.Context) {
	//处理参数
	queryString, exists := c.GetQuery("name")
	pageQ := c.DefaultQuery("page", "1")
	limitQ := c.DefaultQuery("limit", "20")

	page,_:=strconv.Atoi(pageQ)
	limit,_:=strconv.Atoi(limitQ)

	// 初始化参数
	users := make([]model.User, 0)
	//错误处理对象
	SendErrJSON := common.SendErrJSON
	if exists {
		sql := "name like ?"
		if err := common.DB.Where(sql, "%"+queryString+"%").Offset((page-1)*limit).Limit(limit).Find(&users).Error; err != nil {
			SendErrJSON("查找用户出错", c)
			return
		}

	}else {
		if err := common.DB.Offset((page-1)*limit).Limit(limit).Find(&users).Error; err != nil {
			SendErrJSON("查找全部用户出错", c)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"errNo": common.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  gin.H{
			"users":users,
			"total":len(users),
		},
	})

}