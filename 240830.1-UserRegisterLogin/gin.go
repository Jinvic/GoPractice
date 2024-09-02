// gin.go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jakehl/goid"
)

func main() {
	route := gin.Default()

	// 渲染HTML
	route.LoadHTMLGlob("./templates/*")

	// 路由重定向
	route.GET("/", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/index"
		route.HandleContext(ctx)
		// ctx.Redirect(http.StatusSeeOther, "/index")
	})

	// 首页
	route.GET("/index", func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("session_id")
		// 没有会话，进入首页
		fmt.Println(cookie)
		if err != nil || cookie == nil {
			ctx.HTML(http.StatusOK, "index.html", nil)
		}
		seessionID := cookie.Value
		userID := getUserID(seessionID)
		if userID > 0 {
			ctx.Redirect(http.StatusSeeOther, "/profile")
		} else {
			ctx.HTML(http.StatusOK, "index.html", nil)
		}

	})

	// 登录
	route.GET("/login_page", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	route.POST("/login", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		user, errmsg := login(username, password)

		if errmsg != "" {
			errorPage(ctx, errmsg)
			return
		}

		// 如果已有会话，登出
		cookie, err := ctx.Request.Cookie("session_id")
		if err == nil && cookie != nil {
			seessionID := cookie.Value
			userID := getUserID(seessionID)
			if userID > 0 {
				deleteSessionID(seessionID)
			}
		}

		// 记录会话
		sessionID := generateSessionID()
		setSessionID(user.ID, sessionID)
		http.SetCookie(ctx.Writer, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Now().Add(time.Hour),
		})

		// 进入主页
		// ctx.HTML(http.StatusOK, "profile.html", gin.H{
		// 	"username": user.Username,
		// 	"email":    user.Email,
		// })
		ctx.Redirect(http.StatusSeeOther, "/profile")
	})

	// 注册
	route.GET("/register_page", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "register.html", nil)
	})
	route.POST("/register", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		email := ctx.PostForm("email")

		var errmsg string

		if username == "" || password == "" {
			errmsg = "用户名和密码不能为空"
			errorPage(ctx, errmsg)
			return
		}

		user := User{
			Username: username,
			Password: password,
			Email:    email,
		}
		// fmt.Println(user)
		errmsg = register(&user)

		if errmsg != "" {
			errorPage(ctx, errmsg)
			return
		}

		// 进入登录界面
		ctx.Request.URL.Path = "/login"
		route.HandleContext(ctx)
	})

	// 个人主页
	route.GET("/profile", func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("session_id")
		// 没有会话，回到首页
		if err != nil || cookie == nil {
			ctx.Request.URL.Path = "/index"
			route.HandleContext(ctx)
		}

		seessionID := cookie.Value
		userID := getUserID(seessionID)
		user := getUser(userID)

		// 进入主页
		ctx.HTML(http.StatusOK, "profile.html", gin.H{
			"username": user.Username,
			"email":    user.Email,
		})
	})

	// 登出
	route.Any("/logout", func(ctx *gin.Context) {
		cookie, _ := ctx.Request.Cookie("session_id")
		seessionID := cookie.Value
		deleteSessionID(seessionID)

		ctx.Redirect(http.StatusSeeOther, "/index")
	})

	route.Run("localhost:8080")
}

func errorPage(ctx *gin.Context, errmsg string) {
	ctx.HTML(http.StatusOK, "error.html", gin.H{
		"errmsg": errmsg,
	})
}

func generateSessionID() (uuid string) {
	uuid = goid.NewV4UUID().String()
	return
}
