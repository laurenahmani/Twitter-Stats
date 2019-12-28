package main

import (
	"html/template"
	"log"
	"net/http"
)

// UserVars is the user Info
type UserVars struct {
	Username  string
	Followers int
	Following int
}

func main() {
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":8080", nil)
}

// HomePage renders our html
func HomePage(w http.ResponseWriter, r *http.Request) {
	vars := UserVars{
		Username:  "laurenahmani",
		Followers: 1,
		Following: 1,
	}
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Print("template parsing err: ", err)
	}
	err = t.Execute(w, vars)
	if err != nil {
		log.Print("template executing error:", err)
	}
}
