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
	if r.Method == "POST" {
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
