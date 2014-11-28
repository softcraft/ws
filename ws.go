package ws

import (
	"encoding/json"
	"net/http"
	"os"
)

type handler func(w http.ResponseWriter, r *http.Request)

func Start(path string, port string, handler handler) {
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.Header().Set("Allow", "GET,POST,PUT,DELETE,OPTIONS")
			return
		}

		if r.FormValue("access_token") != os.Getenv("ACCESS_TOKEN") {
			w.WriteHeader(http.StatusUnauthorized)
			authErr, _ := json.Marshal(map[string]string{"error": "Incorrect access token provided"})
			w.Write(authErr)
			return
		}
		f(w, r)
	}
}
