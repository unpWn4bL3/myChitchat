package routes

import "github.com/gorilla/mux"

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true)
	for _, route := range webRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}
	return
}
