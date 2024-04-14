package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const exchangeRate = 15000.0

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World") // mendapatkan inputan string untuk dikembalikan ke response writer sebagai response
}

func converDolarToRupiah(w http.ResponseWriter, r *http.Request) {

	// menerima paramter rupiah dari query string
	rupiahStr := r.URL.Query().Get("rupiah")
	if rupiahStr == "" {
		http.Error(w, "Anda perlu memasukan bilangan dalam rupiah", http.StatusBadRequest)
		return
	}

	// mengkonversi rupiah  dari string ke float
	rupiah, err := strconv.ParseFloat(rupiahStr, 64)
	if err != nil {
		http.Error(w, "Invalid Rupiah Ammount", http.StatusBadRequest)
		return
	}

	// mengkonversi rupiah ke dolar
	dollars := rupiah / exchangeRate

	// tulis ke response dan return ke client
	fmt.Fprintf(w, "%.2f USD", dollars)
}

func main() {
	// route dan handler
	http.HandleFunc("/", helloHandler) // jadi saat client nge request localhost:8080/ -> "/", akan mentriger fungsi hello handler yang akan me return response "hello world"
	http.HandleFunc("/converdolartorp", converDolarToRupiah)

	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil) //open dan listen TCP pada port 8080
	if err != nil {
		log.Fatal("Listen and serve: ", err)
	}
}
