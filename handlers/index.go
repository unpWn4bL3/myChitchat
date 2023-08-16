package handlers

import (
	"net/http"

	"github.com/unpWn4bl3/myChitchat/models"
)

// GET /index
func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := models.Threads()
	if err != nil {
		warning(err.Error())
		return
	}
	_, err = session(w, r)
	if err != nil {
		generateHTML(w, threads, "layout", "navbar", "index")
		return
	}
	generateHTML(w, threads, "layout", "auth.navbar", "index")
}

// GET /err?msg={msg}
func Err(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, params.Get("msg"), "layout", "navbar", "error")
		return
	}
	generateHTML(w, params.Get("msg"), "layout", "auth.navbar", "error")
}
