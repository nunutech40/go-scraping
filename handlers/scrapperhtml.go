package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HtmlScrapper(w http.ResponseWriter, r *http.Request) {
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
