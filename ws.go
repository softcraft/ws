package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"os"
)

type Response map[string]interface{}
type Handler func(url.Values) Response
type Route struct {
	Method  string
	Path    string
	Handler Handler
}

func StartServer(routes []Route) {
	rtr := mux.NewRouter()

	for _, route := range routes {
		rtr.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			r.ParseMultipartForm(264)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, route.Handler(r.Form))
			return
		}).Methods(route.Method)
	}

	http.Handle("/", rtr)
	http.ListenAndServe(getPort(), nil)
}

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8181"
	}
	return ":" + port
}
