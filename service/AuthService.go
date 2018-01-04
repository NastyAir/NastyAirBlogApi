package service

import (
	"../dao"
	"../common"
)

func Login(account, password string) common.RestMsg {

	var isAuth bool
	isAuth = dao.UserFindByAccount(account, password);
	msg := new(common.RestMsg)
	//var data []byte
	if isAuth {
		msg.Code = common.SUCCESS
		//data, _ = json.Marshal(msg)
	} else {
		msg.Code = common.FAIL
	}
	return *msg
}
