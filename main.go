package main

import (
	"github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
)

func handleThis(res http.ResponseWriter, req *http.Request) {
	templatePage, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	name := req.FormValue("name")
	age := req.FormValue("age")
	sex := req.FormValue("sex")
	loc := req.FormValue("location")

	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			// Secure: true,
			Name:     "session-fino",
			Value:    id.String() + "," + name + "," + age + "," + sex + "," + location,
			HttpOnly: true,
		}
	}

	http.SetCookie(res, cookie)
	templatePage.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", handleThis)
	http.ListenAndServe(":8080", nil)
}
