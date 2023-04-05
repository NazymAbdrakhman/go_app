package main

import (
	"html/template"
	"log"
	"net/http"
)

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("registration.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		// Get the form data from the request
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Save the user in the database or perform any other necessary actions

		// Render a success page
		tmpl, err := template.ParseFiles("registration-success.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func authorizationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("registration.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		// Get the form data from the request
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Check the user's credentials against the database or other authentication service
		if username == "admin" && password == "password123" {
			// Authenticx`x	ation successful, redirect the user to the dashboard
			http.Redirect(w, r, "/dashboard", http.StatusFound)
		} else {
			// Authentication failed, render an error page
			tmpl, err := template.ParseFiles("authorization-error.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func MainBookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("main-book.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("cat.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/cat", CategoriesHandler)
	http.HandleFunc("/register", registrationHandler)
	http.HandleFunc("/authorize", authorizationHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/main-book", MainBookHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
