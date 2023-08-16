package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/unpWn4bl3/myChitchat/models"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "All posts")
}

func PostThread(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	err = r.ParseForm()
	if err != nil {
		warning("Cannot parse form: " + err.Error())
		return
	}
	user, err := sess.User()
	if err != nil {
		warning("Cannot get user: " + err.Error())
		return
	}
	body := r.PostFormValue("body")
	uuid := r.PostFormValue("uuid")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		warning("Cannot find thread: " + err.Error())
	}
	if _, err := user.CreatePost(thread, body); err != nil {
		warning("Cannot create post: " + err.Error())
	}
	url := fmt.Sprint("/thread/read?uuid=", uuid)
	http.Redirect(w, r, url, 302)
}
