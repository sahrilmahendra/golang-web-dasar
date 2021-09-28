package handler

import (
	"golang-web/entity"
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
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(rw, "Website sedang mengalami Error", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"contentHome": "Hello World, saya sekarang sedang belajar Golang Web",
	// 	"title":       "Home Page",
	// }

	// data := entity.Product{Id: 1, Name: "Mouse", Price: 90000, Stock: 5}

	data := []entity.Product{
		{Id: 1, Name: "Mouse", Price: 87000, Stock: 7},
		{Id: 2, Name: "Keyboard", Price: 243000, Stock: 3},
		{Id: 3, Name: "Flashdisk", Price: 60000, Stock: 12},
	}

	err = tmpl.Execute(rw, data)
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
	data := map[string]interface{}{
		"content": idNumber,
		"title":   "Halaman Product",
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(rw, "Website sedang mengalami Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(rw, data)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Website sedang mengalami Error", http.StatusInternalServerError)
		return
	}
}

func PostGet(rw http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		rw.Write([]byte("Method Get"))
	case "POST":
		rw.Write([]byte("Method Post"))
	default:
		http.Error(rw, "Error terjadi", http.StatusBadRequest)
	}
}

func Form(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(rw, "sedang terjadi Error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(rw, nil)
		if err != nil {
			log.Println(err)
			http.Error(rw, "sedang terjadi Error", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(rw, "sedang terjadi Error", http.StatusBadRequest)
}
