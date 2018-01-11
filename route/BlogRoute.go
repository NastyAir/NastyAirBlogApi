package route

import (
	"net/http"
	"NastyAir/Blog/entity"
	"encoding/json"
	"NastyAir/Blog/Utils"
	"NastyAir/Blog/service"
)

func Blog(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var offset, length []string

	if r.Method == "GET" {
		offset = r.Form["offset"]
		length = r.Form["length"]
		service.
	} else    if r.Method == "POST" {
		var blog entity.Blog
		err := json.NewDecoder(r.Body).Decode(&blog)
		if err != nil {
			Utils.CustomHttpCodeResponse("Json 解析错误", http.StatusBadRequest, w)
			return
		}
		data, _ := service.BlogCreate(blog)
		Utils.JsonResponse(data, w)
	}
}
