package routes

import (
	"net/http"

	"github.com/unpWn4bl3/myChitchat/handlers"
)

// 存储路由信息
// 将Method-Pattern映射到HandlerFunc
type WebRoute struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type WebRoutes []WebRoute

var webRoutes = WebRoutes{
	WebRoute{
		"home",
		"GET",
		"/",
		handlers.Index,
	},
	WebRoute{
		"signup",
		"GET",
		"/signup",
		handlers.Signup,
	},
	WebRoute{
		"signupAccount",
		"POST",
		"/signup_account",
		handlers.SignupAccount,
	},
	WebRoute{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},
	WebRoute{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},
	WebRoute{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
	WebRoute{
		"newThread",
		"GET",
		"/thread/new",
		handlers.NewThread,
	},
	WebRoute{
		"createThread",
		"POST",
		"/thread/create",
		handlers.CreateThread,
	},
	WebRoute{
		"readThread",
		"GET",
		"/thread/read",
		handlers.ReadThread,
	},
	WebRoute{
		"postThread",
		"POST",
		"/thread/post",
		handlers.PostThread,
	},
	WebRoute{
		"error",
		"GET",
		"/err",
		handlers.Err,
	},
}
