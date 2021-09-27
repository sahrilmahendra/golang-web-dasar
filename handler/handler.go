package handler

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(rw, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		log.Println(err)
		http.Error(rw, "Website sedang mengalami Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(rw, nil)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Website sedang mengalami Error", http.StatusInternalServerError)
		return
	}
}
func HelloHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello World, Saya sedang belajar Golang Web"))
}
func SahrilHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Moch. Syahryil Mahendra"))
}

func ProductHandler(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil || idNumber < 1 {
		http.NotFound(rw, r)
		return
	}
	fmt.Fprintf(rw, "Product Page : %d", idNumber)
}
