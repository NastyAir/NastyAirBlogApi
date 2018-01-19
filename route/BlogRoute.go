package route

import (
	"net/http"
	"NastyAir/Blog/entity"
	"encoding/json"
	"NastyAir/Blog/Utils"
	"NastyAir/Blog/service"
	"strconv"
)

func Blog(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var offset, length []string

	if r.Method == "GET" {
		offset = r.Form["offset"]
		length = r.Form["length"]
		var offsetNum int
		var lengthNum int
		if len(offset) > 0 {
			offsetNum, _ = strconv.Atoi(offset[0])
		} else {
			offsetNum = 0
		}
		if len(length) > 0 {
			lengthNum, _ = strconv.Atoi(length[0])
		} else {
			lengthNum = 10
		}
		msg, _ :=service.BlogList(offsetNum, lengthNum)
		Utils.JsonResponse(msg, w)
	} else if r.Method == "POST" {
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
