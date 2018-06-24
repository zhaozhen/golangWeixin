package model


// redis相关常量, 为了防止从redis中存取数据时key混乱了，在此集中定义常量来作为各key的名字
const (
	// ActiveTime 生成激活账号的链接
	ActiveTime = "activeTime"

	// ResetTime 生成重置密码的链接
	ResetTime = "resetTime"

	// LoginUser 用户信息
	LoginUser = "loginUser"

	// ArticleMinuteLimit 用户每分钟最多能发表的文章数
	ArticleMinuteLimit = "articleMinuteLimit"

	// ArticleDayLimit 用户每天最多能发表的文章数
	ArticleDayLimit = "articleDayLimit"

	// CommentMinuteLimit 用户每分钟最多能发表的评论数
	CommentMinuteLimit = "commentMinuteLimit"

	// CommentDayLimit 用户每天最多能发表的评论数
	CommentDayLimit = "commentDayLimit"
)
