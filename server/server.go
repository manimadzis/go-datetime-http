package server

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func SetHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/date", http.StatusFound)
	})

	yearHandlerFunc := getHandlerFuncWithTZ(Year, YearLoc)
	monthHandlerFunc := getHandlerFuncWithTZ(Month, MonthLoc)
	dayHandlerFunc := getHandlerFuncWithTZ(Day, DayLoc)

	dateHandlerFunc := getHandlerFuncWithTZ(Date, DayLoc)
	timeHandlerFunc := getHandlerFuncWithTZ(Time, TimeLoc)
	weekdayHandlerFunc := getHandlerFuncWithTZ(Weekday, WeekdayLoc)

	http.HandleFunc("/date", dateHandlerFunc)
	http.HandleFunc("/time", timeHandlerFunc)
	http.HandleFunc("/weekday", weekdayHandlerFunc)

	http.HandleFunc("/day", dayHandlerFunc)
	http.HandleFunc("/month", monthHandlerFunc)
	http.HandleFunc("/year", yearHandlerFunc)
}

func getHandlerFuncWithTZ(get func() string, getLoc func(*time.Location) string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Some internal error"))
			log.Println("getYear: %w", err)
			return
		}

		tz_str := r.Form.Get("tz")

		var data string
		if tz_str == "" {
			data = get()
		} else {
			var (
				tz  int
				err error
			)

			if tz, err = strconv.Atoi(tz_str); err != nil {
				w.WriteHeader(400)
				w.Write([]byte("Failed to read timezone"))
				return
			}

			if tz > 12 || tz < -12 {
				w.WriteHeader(400)
				w.Write([]byte("Failed to read timezone"))
				return
			}

			loc := time.FixedZone("", tz*3600)

			data = getLoc(loc)
		}

		w.Write([]byte(data))
	}
}
