package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleDefault)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./frontend/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./frontend/js"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./frontend/images"))))
	http.ListenAndServe(":8080", nil)
}

//HandleDefault .
func HandleDefault(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./frontend/index.html")
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	t.Execute(w, nil)
}
