package main

import (
	"fmt"
	"github.com/nunutech40/go-scraping/handlers"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const exchangeRate = 15000.0

func converDolarToRupiah(w http.ResponseWriter, r *http.Request) {

	// menerima paramter rupiah dari query string
	rupiahStr := r.URL.Query().Get("rupiah") // menggunakan query untuk menerima inputan dari client alamatdomain/convert?rupiah=30000
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

func htmlScrapper(w http.ResponseWriter, r *http.Request) {
	// get url from client
	urlToScrape := r.URL.Query().Get("url")
	if urlToScrape == "" {
		http.Error(w, "Please provide a URL to scrape", http.StatusBadRequest)
		return
	}

	// cek apakah query adalah URL ?
	cekStrUri := isValidURL(urlToScrape)
	if cekStrUri == false {
		http.Error(w, "Mohon masukan URL yang benar", http.StatusBadRequest)
		return
	}

	// lakukan http get request ke URL yang diberikan
	resp, err := http.Get(urlToScrape)
	if err != nil {
		http.Error(w, "Failed to retrieve website", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close() // setelah mendapatkan response dari connect dan get request ke url, lalu tutup

	// baca isi dari body response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read the website content", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "This is scrapped result:\n\n%s", body) // return result dengan header, This is scrapped result, lalu enter dan return di response of scrapped content
}

func isValidURL(cekUrlString string) bool {
	_, err := url.ParseRequestURI(cekUrlString) // return to uri atau to err ()
	return err == nil                           // jika hasilnya uri, berarti err nya nil, jika err == nil, isValid == true
}

func main() {
	// route dan handler
	http.HandleFunc("/", handlers.HelloHandlers) // jadi saat client nge request localhost:8080/ -> "/", akan mentriger fungsi hello handler yang akan me return response "hello world"
	http.HandleFunc("/converdolartorp", converDolarToRupiah)
	http.HandleFunc("/scraphtml", htmlScrapper)

	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil) //open dan listen TCP pada port 8080
	if err != nil {
		log.Fatal("Listen and serve: ", err)
	}
}
