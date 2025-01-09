package pkg

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// generate response
func generateResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	r := Response{data}
	res, err := json.Marshal(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func filterData(start, end, text string) ([]Data, error) {
	var result []Data

	if len(data) == 0 {
		return data, errors.New("no logs found")
	}

	for _, val := range data {
		timeFlag, textFlag := false, false

		startTime, _ := strconv.Atoi(start)
		endTime, _ := strconv.Atoi(end)
		if start != "" && end != "" {
			if val.Time >= int64(startTime) && val.Time <= int64(endTime) {
				timeFlag = true
			}
		} else if start != "" {
			if val.Time >= int64(startTime) {
				timeFlag = true
			}
		} else if end != "" {
			if val.Time <= int64(endTime) {
				timeFlag = true
			}
		} else {
			timeFlag = true
		}

		if text != "" {
			if strings.Contains(val.Log.Body, text) {
				textFlag = true
			}
		} else {
			textFlag = true
		}

		if timeFlag && textFlag {
			result = append(result, val)
		}

	}

	if len(result) == 0 {
		return result, errors.New("no match found for criteria")
	}

	return result, nil
}
