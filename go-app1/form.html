package main

import (
	"net/http"
	"text/template"
)

type Contact struct {
	Name            string
	Username        string
	Email           string
	Password        string
	ConfirmPassword string
}

func main() {

	tmpl := template.Must(template.ParseFiles("form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := Contact{
			Name:            r.FormValue("name"),
			Username:        r.FormValue("username"),
			Email:           r.FormValue("email"),
			Password:        r.FormValue("password"),
			ConfirmPassword: r.FormValue("confirm"),
		}

		_ = details

		tmpl.Execute(w, struct{ success bool }{true})
	})
	http.ListenAndServe(":8080", nil)
}
