package main

import (
	"golang-web/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// aboutHandler := func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Write([]byte("About Page"))
	// }
	// mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/sahril", handler.SahrilHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	// mux.HandleFunc("/profile", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Write([]byte("Moch. Syahryil Mahendra | Web Developer"))
	// })

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
