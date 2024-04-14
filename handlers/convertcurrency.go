package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

const exchangeRate = 15000.0

func ConverDolarToRupiah(w http.ResponseWriter, r *http.Request) {

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
