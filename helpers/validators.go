package helpers

import (
	"fmt"
	"net/url"
)

func IsValidURL(cekUrlString string) bool {
	_, err := url.ParseRequestURI(cekUrlString) // return to uri atau to err ()
	return err == nil                           // jika hasilnya uri, berarti err nya nil, jika err == nil, isValid == true
}
