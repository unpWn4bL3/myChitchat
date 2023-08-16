package handlers

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	io.WriteString(w, "Return user info with id = "+id)
}
