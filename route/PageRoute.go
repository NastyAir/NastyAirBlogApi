package route

import (
	"net/http"
	"NastyAir/Blog/Utils"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Utils.PageResponse("index.html", w)
	}
}
