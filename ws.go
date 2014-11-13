package ws

import (
	"encoding/json"
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
			r.ParseMultipartForm(0)
			w.Header().Set("Content-Type", "application/json")
			rep, _ := json.Marshal(route.Handler(r.Form))
			//validate error here
			w.Write(rep)
		}).Methods(route.Method)
	}

	http.Handle("/", rtr)
	http.ListenAndServe(getPort(), nil)
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8181"
	}
	return ":" + port
}
