package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func SetHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/date", http.StatusFound)
	})
	http.HandleFunc("/date", getDate)
	http.HandleFunc("/time", getTime)
	http.HandleFunc("/weekday", getWeekday)

	http.HandleFunc("/day", getDay)
	http.HandleFunc("/month", getMonth)
	http.HandleFunc("/year", getYear)
}

func getDate(w http.ResponseWriter, r *http.Request) {
	year, month, day := time.Now().Date()

	r.ParseForm()
	v := r.Form.Get("var")
	fmt.Println(v)

	w.Write([]byte(fmt.Sprintf("%02d.%02d.%d", day, month, year)))
}

func getTime(w http.ResponseWriter, r *http.Request) {
	hour, min, sec := time.Now().Clock()
	w.Write([]byte(fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)))
}

func getWeekday(w http.ResponseWriter, r *http.Request) {
	cur_time := time.Now().Weekday().String()
	w.Write([]byte(cur_time))
}

func getDay(w http.ResponseWriter, r *http.Request) {
	cur_time := strconv.Itoa(time.Now().Day())
	w.Write([]byte(cur_time))
}

func getMonth(w http.ResponseWriter, r *http.Request) {
	cur_time := time.Now().Month().String()
	w.Write([]byte(cur_time))
}

func getYear(w http.ResponseWriter, r *http.Request) {
	cur_time := strconv.Itoa(time.Now().Year())
	w.Write([]byte(cur_time))
}
