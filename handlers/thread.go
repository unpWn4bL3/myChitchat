package handlers

import (
	"net/http"

	"github.com/unpWn4bl3/myChitchat/models"
)

// GET /threads/new
func NewThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	generateHTML(w, nil, "layout", "auth.navbar", "new.thread")
}

// POST /thread/create
func CreateThread(w http.ResponseWriter, r *http.Request) {
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
		warning("Cannot get user in session: " + err.Error())
		return
	}
	topic := r.PostFormValue("topic")
	if _, err := user.CreateThread(topic); err != nil {
		warning("Cannot create thread: " + err.Error())
		return
	}
	http.Redirect(w, r, "/", 302)
}

// GET /thread/read?uuid={uuid}
func ReadThread(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uuid := params.Get("uuid")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		warning("Cannot read thread: " + err.Error())
		return
	}
	_, err = session(w, r)
	if err != nil { //这里的&thread要和(thread *Thread) User()保持一致否则不会被调用
		generateHTML(w, &thread, "layout", "navbar", "thread")
	} else {
		generateHTML(w, &thread, "layout", "auth.navbar", "auth.thread")
	}
}
