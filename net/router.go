package net

import (
	"github.com/farrout/goaway/middleware"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"net/http"
)

var log = logging.MustGetLogger("net")

func NewRouter() *mux.Router {
	// Create new Multiplexing router
	r := mux.NewRouter()

	// Attach middleware
	r.Use(middleware.LogRequest)

	// -- Define routing

	// Home page
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	//Serve assets
	r.PathPrefix("/web/static/").Handler(http.StripPrefix("/web/static/", http.FileServer(http.Dir("./web/static/"))))

	return r
}


//HomeHandler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/views/index.html")
}

//NotFoundHandler
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Warning("Not found:", r.URL)
}
