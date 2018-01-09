package route

import (
	"net/http"
	"NastyAir/Blog/Utils"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//t, _ := template.ParseFiles("index.html")
		//t.Execute(w, nil)
		Utils.PageResponse("index.html", w)
	}
}
