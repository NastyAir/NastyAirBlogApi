package Utils

import (
	"net/http"
	"encoding/json"
	"html/template"
	"fmt"
)

func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func PageResponse(pageUrl string, w http.ResponseWriter, ) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, _ := template.ParseFiles(pageUrl)
	t.Execute(w, nil)
}
func CustomHttpCodeResponse(msg string,httpCode int, w http.ResponseWriter, ) {
	w.WriteHeader(httpCode)
	fmt.Fprintln(w, msg)
}
