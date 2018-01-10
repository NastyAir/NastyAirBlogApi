package service

import (
	"NastyAir/Blog/entity"
	"NastyAir/Blog/common"
	"NastyAir/Blog/dao"
)

func BlogCreate(blog entity.Blog) (common.RestMsg, error) {

	_, err := dao.BlogInsert(blog)
	msg := new(common.RestMsg)
	//var data []byte
	if err != nil {
		msg.Code = common.FAIL
		msg.Msg = "create blog fail"
	} else {
		msg.Data = make(map[string]string)
		msg.Code = common.SUCCESS
		msg.Msg = "create blog success"
	}
	return *msg, err
}
