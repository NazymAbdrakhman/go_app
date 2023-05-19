package login

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "html/authorization.html")
	} else if r.Method == "POST" {
		// handle login logic here
	}
}
