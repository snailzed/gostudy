package account

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gostudy/answer/single-arch/session"
)

//account中间件初始化
func Init() {
	session.Init()
}

func ProcessRequest(c *gin.Context) {
	var userSession session.Session
	//延迟设置
	defer func() {
		if userSession == nil {
			userSession, _ = session.CreateSession()
		}
		c.Set(UserSessionName, userSession)
	}()
	//获取cookie
	cookie, err := c.Request.Cookie(CookieSessionId)
	if err != nil {
		c.Set(UserId, int64(0))
		c.Set(UserLoginStatus, false)
		return
	}

	//获取cookie对应的session
	sessionId := cookie.Value
	if len(sessionId) == 0 {
		c.Set(UserId, int64(0))
		c.Set(UserLoginStatus, false)
		return
	}
	userSession, err = session.Get(cookie.Value)
	if err != nil {
		c.Set(UserId, int64(0))
		c.Set(UserLoginStatus, false)
		return
	}
	//获取用户id
	userId, err := userSession.Get(UserId)
	if err != nil {
		c.Set(UserId, int64(0))
		c.Set(UserLoginStatus, false)
		return
	}
	//设置到请求上下文
	intUserId, ok := userId.(int64)
	if !ok {
		c.Set(UserId, int64(0))
		c.Set(UserLoginStatus, false)
		return
	}

	//设置用户id和登录状态
	c.Set(UserId, intUserId)
	c.Set(UserLoginStatus, true)
}

//获取用户Id
func GetUserId(ctx *gin.Context) (userId int64, err error) {
	tempUserId, exists := ctx.Get(UserId)
	if !exists {
		err = errors.New("The user id does not exists.")
		return
	}
	userId, ok := tempUserId.(int64)
	if !ok {
		err = errors.New("Convert user id to int64 failed.")
		return
	}
	return
}

//判断是否登录
func IsLogin(ctx *gin.Context) bool {
	tempLogin, exists := ctx.Get(UserLoginStatus)
	if !exists {
		return false
	}
	ret, ok := tempLogin.(bool)
	if !ok {
		return false
	}
	return ret
}

//处理响应，设置cookie
func ProcessResponse(ctx *gin.Context) {
	tempUserSession, exists := ctx.Get(UserSessionName)
	if !exists {
		return
	}
	userSession, ok := tempUserSession.(session.Session)
	if !ok {
		return
	}
	//userSession 设置cookie
	ctx.SetCookie(CookieSessionId, userSession.Id(), CookieMaxAge, "/", "localhost", false, true)

	//cookie := &http.Cookie{
	//	Name:     CookieSessionId,
	//	Value:    userSession.Id(),
	//	Expires:  time.Now().Add(1 * 3600 * time.Second),
	//	Path:     "/",
	//	Secure:   false,
	//	HttpOnly: true,
	//}
	//http.SetCookie(ctx.Writer, cookie)
}
