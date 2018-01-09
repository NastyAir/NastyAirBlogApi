package service

import (
	"../dao"
	"../common"
	"NastyAir/Blog/entity"
)

func Login(account, password string) common.RestMsg {

	var user entity.UserCredentials
	user = dao.UserFindByAccount(account, password);
	msg := new(common.RestMsg)
	//var data []byte
	if user.Password == password {
		msg.Code = common.SUCCESS
		msg.Msg = "auth success"
	} else {
		msg.Code = common.FAIL
		msg.Msg = "auth fail"
	}
	return *msg
}
