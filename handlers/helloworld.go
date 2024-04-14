package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandlers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World") // mendapatkan inputan string untuk dikembalikan ke response writer sebagai response
}