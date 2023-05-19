package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/NazymAbdrakhman/go_app/login"
)

type Book struct {
	Title   string
	Image   string
	Stars   int
	Price   string
	BuyLink string
}

type PageData struct {
	Books []Book
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("html/registration.html")
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
		confirm := r.FormValue("confirm")

		// Check if passwords match
		if password != confirm {
			tmpl, err := template.ParseFiles("html/registration-error.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, "Passwords do not match")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		// Open the users file in append mode
		file, err := os.OpenFile("users.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Write the user's credentials to the users file
		userData := fmt.Sprintf("%s,%s\n", username, password)
		_, err = file.WriteString(userData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect the user to the login page
		http.Redirect(w, r, "/authorize", http.StatusFound)
	}
}
func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("html/cat-book.html")
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
func BuybookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("html/buy-book.html")
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
	http.HandleFunc("/cat-book", CategoriesHandler)
	http.HandleFunc("/main", homeHandler)
	http.HandleFunc("/register", registrationHandler)
	http.HandleFunc("/authorize", login.LoginHandler)
	http.HandleFunc("/buy-book", BuybookHandler)
	log.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	books := []Book{
		{
			Title:   "Complete Tales and Poems",
			Image:   "https://i.pinimg.com/564x/bc/61/f3/bc61f3b4dc997c4fd4d41d10d91cf9f1.jpg",
			Stars:   4,
			Price:   "3250 tg",
			BuyLink: "#",
		},
		{
			Title:   "Heartless Wood",
			Image:   "https://i.pinimg.com/564x/ab/d0/45/abd045a943dc9d8a85099f4a139dfc07.jpg",
			Stars:   3,
			Price:   "6500 tg",
			BuyLink: "#",
		},
	}

	data := PageData{
		Books: books,
	}

	renderHTML(w, "html/authorization.html", data)
}

func renderHTML(w http.ResponseWriter, templateFile string, data interface{}) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
