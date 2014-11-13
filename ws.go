package ws

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"os"
)

type Handler func(url.Values) (int, interface{})
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

			if r.FormValue("access_token") != os.Getenv("ACCESS_TOKEN") {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				authErr, _ := json.Marshal(map[string]string{"error": "Incorrect access token provided"})
				w.Write(authErr)
			} else {
				status, response := route.Handler(r.Form)
				jsonResponse, _ := json.Marshal(response)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(status)
				w.Write(jsonResponse)
			}

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
