package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var Numerals = map[int]string{
	10: "X",
	9:  "IX",
	8:  "VIII",
	7:  "VII",
	6:  "VI",
	5:  "V",
	4:  "IV",
	3:  "III",
	2:  "II",
	1:  "I",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Bulunamadı"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(Numerals[number]))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Hatalı istek"))
		}
	})

	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
