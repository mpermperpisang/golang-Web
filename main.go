package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

/*Details : berisikan keterangan tanggal dan waktu*/
type Details struct {
	Day  string
	Date string
	Time string
}

func main() {
	currentTime := time.Now()
	day := currentTime.Format("Monday")
	date := currentTime.Format("02 January 2006")
	time := currentTime.Format("15:04:05")
	details := Details{day, date, time}
	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "index.html", details); err != nil {
			http.Error(w, err.Error(), http.StatusNoContent)
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
