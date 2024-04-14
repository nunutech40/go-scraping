package main

import (
	"github.com/nunutech40/go-scraping/handlers"
	"log"
	"net/http"
)

func main() {
	// route dan handler
	http.HandleFunc("/", handlers.HelloHandlers) // jadi saat client nge request localhost:8080/ -> "/", akan mentriger fungsi hello handler yang akan me return response "hello world"
	http.HandleFunc("/converdolartorp", handlers.ConverDolarToRupiah)
	http.HandleFunc("/scraphtml", handlers.HtmlScrapper)

	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil) //open dan listen TCP pada port 8080
	if err != nil {
		log.Fatal("Listen and serve: ", err)
	}
}
