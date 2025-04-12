package main

import (
	"fmt"

	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ibrahimbayburtlu/Building-Restful-Web-Services-With-Go/chapter-1/romanNumerals"
)

func main() {
	// http paketi, isteklerle ilgilenmek için metodlar içerir
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		// Eğer istek doğru sözdizimi ile GET ise
		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				// Eğer kaynak listede yoksa, Bulunamadı durumu gönder
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Bulunamadı"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(romanNumerals.Numerals[number]))
			}
		} else {
			// Diğer tüm istekler için, İstemcinin hatalı istek gönderdiğini söyle
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Hatalı istek"))
		}
	})

	// Bir sunucu oluştur ve 8000 portunda çalıştır
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Roma rakamları API sunucusu 8000 portunda çalışıyor...")
	s.ListenAndServe()
}
