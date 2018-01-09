package service

import (

	"NastyAir/Blog/entity"
	"time"
	"github.com/dgrijalva/jwt-go"
	"log"
	"errors"
	"NastyAir/Blog/dao"
	"NastyAir/Blog/common"
)

const (
	SecretKey = "This is a key for NastyAir"
)

func Login(account, password string) common.RestMsg {

	var user entity.UserCredentials
	user = dao.UserFindByAccount(account, password);
	msg := new(common.RestMsg)
	//var data []byte
	if user.Password == password {
		token, _ := generateToken()
		msg.Data = make(map[string]string)
		msg.Data["token"] = token
		msg.Code = common.SUCCESS
		msg.Msg = "auth success"
	} else {
		msg.Code = common.FAIL
		msg.Msg = "auth fail"
	}
	return *msg
}
func generateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		fatal(err)
		return "", errors.New("Error while signing the token")
	}
	return tokenString, nil
}
func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
