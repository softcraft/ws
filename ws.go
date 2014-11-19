package ws

import (
	"encoding/json"
	"net/http"
	"os"
)

type handler func(w http.ResponseWriter, r *http.Request)

func Start(path string, handler handler) {
	var port = ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8181"
	}

	http.HandleFunc(path, Authorize(handler))
	http.ListenAndServe(port, nil)
}

func Success(w http.ResponseWriter, status int, response map[string]interface{}) {
	jsonResponse, _ := json.Marshal(response)

	w.WriteHeader(status)
	w.Write(jsonResponse)
}

func Error(w http.ResponseWriter, status int, err string) {
	response := map[string]string{"error": err}
	jsonResponse, _ := json.Marshal(response)

	w.WriteHeader(status)
	w.Write(jsonResponse)
}

func Authorize(f handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("access_token") != os.Getenv("ACCESS_TOKEN") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			authErr, _ := json.Marshal(map[string]string{"error": "Incorrect access token provided"})
			w.Write(authErr)
			return
		}
		f(w, r)
	}
}
