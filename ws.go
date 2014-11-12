package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"os"
)

type Handler func(url.Values) map[string]interface{}
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
			fmt.Fprint(w, parse(route.Handler(r.Form)))
			return
		}).Methods(route.Method)
	}

	http.Handle("/", rtr)
	http.ListenAndServe(getPort(), nil)
}

func parse(m map[string]interface{}) (s string) {
	b, err := json.Marshal(m)
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