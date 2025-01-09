package pkg

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

func GetData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	q := r.URL.Query()

	start := q.Get("start")
	end := q.Get("end")
	text := q.Get("text")

	filtered, err := filterData(start, end, text)

	if err != nil {
		generateResponse(w, err.Error())
		return
	}
	generateResponse(w, filtered)
}

func StoreData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// parse and store data in data variable

	var raw []string

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&raw)
	if err != nil {
		generateResponse(w, err.Error())
		return
	}

	for _, val := range raw {
		valArr := strings.Split(val, " ")
		if len(valArr) < 4 {
			generateResponse(w, "incorrect data format")
			return
		}
		t, err := time.Parse(time.RFC3339, valArr[0])
		if err != nil {
			generateResponse(w, "error parsing time")
			return
		}

		// Convert the time to Unix timestamp (seconds since epoch)
		unixTimestamp := t.Unix()

		severity := valArr[1]
		valArr[2] = strings.ReplaceAll(valArr[2], "[", "")
		valArr[2] = strings.ReplaceAll(valArr[2], "]", "")
		service := valArr[2]
		body := strings.Join(valArr[3:], " ")

		tempData := Data{Time: unixTimestamp, Log: Info{
			Body: body, Service: service, Severity: severity,
		}}

		data = append(data, tempData)
	}

	// convert data into struct
	// [  time severity [server] body
	//    ”2023-10-11T10:31:00Z INFO [apache] Received GET request from 192.168.0.1 for /index.html”,
	//    ”2023-10-11T10:32:15Z INFO [apache] Request from 10.0.0.2 failed with status code 404 for /page-not-found.html”,
	//    “2023-10-11T11:33:30Z WARN [nginx] Received POST request from 192.168.0.3 for /submit-form”
	// ] 1685426738

	generateResponse(w, "successfully added")
}

// reset data variable to empty
func RefreshData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data = []Data{}

	generateResponse(w, "data refreshed")
}
