package service

import (
	"NastyAir/Blog/entity"
	"NastyAir/Blog/common"
	"NastyAir/Blog/dao"
	"encoding/json"
)

func BlogCreate(blog entity.Blog) (common.RestMsg, error) {

	_, err := dao.BlogInsert(blog)
	msg := new(common.RestMsg)
	//var data []byte
	if err != nil {
		msg.Code = common.FAIL
		msg.Msg = "create blog fail"
		msg.Data = make(map[string]string)
		msg.Data["err"] = err.Error()

	} else {
		//msg.Data = make(map[string]string)
		msg.Code = common.SUCCESS
		msg.Msg = "create blog success"
	}
	return *msg, err
}

func BlogList(offset, length int) (common.RestMsg, error) {

	blogs, err := dao.BlogSelectByPage(offset, length)
	msg := new(common.RestMsg)
	if err != nil {
		msg.Code = common.FAIL
		msg.Msg = "create blog fail"
		msg.Data = make(map[string]string)
		msg.Data["err"] = err.Error()

	} else {
		//msg.Data = make(map[string]string)
		msg.Code = common.SUCCESS
		msg.Msg = "create blog success"
		msg.Data = make(map[string]string)
		str, _ := json.Marshal(blogs)
		msg.Data["blog"] = string(str)
	}
	return *msg, err
}
