package main

import (
	"html/template"
	"log"
	"net/http"

	"gitlab.com/evdokimovn/mosgor/inc"
)

func main() {
	in := inc.NewIncrementer(2)
	s := in.(inc.IncrementerServer)
	in.Start()
	http.Handle("/data/", s)
	http.Handle("/", MainPage(in.Size()))
	log.Println(http.ListenAndServe(":8002", nil))
}

func MainPage(i int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatal("Can't parse template")

		}

		// Most straight forward way
		a := make([]struct{}, i)
		t.Execute(w, a)
	}
}
