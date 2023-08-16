package handlers

import (
	"net/http"

	"github.com/unpWn4bl3/myChitchat/models"
)

// GET /login
func Login(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("auth.layout", "navbar", "login")
	t.Execute(w, nil)
}

// GET /signup
// 注册页面
func Signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "auth.layout", "navbar", "signup")
}

// POST /signup_account
// 注册接口
func SignupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		warning("Cannot parse form: ", r.Body)
		return
	}
	user := models.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		warning("Cannot create user: " + err.Error())
		return
	}
	http.Redirect(w, r, "/login", 302)
}

// POST /authenticate
// 验证用户的接口
func Authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		warning("Cannot find user with email: " + r.PostFormValue("email"))
		w.Write([]byte("Cannot find user with email: " + r.PostFormValue("email")))
		return
	}
	if user.Password != models.Encrypt(r.PostFormValue("password")) {
		http.Redirect(w, r, "/login", 302)
		return
	}
	session, err := user.CreateSession()
	if err != nil {
		warning("Cannot create session: " + err.Error())
		return
	}
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", 302)
}

// GET /logout
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning("Failed to get cookie")
		session := models.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)
}
