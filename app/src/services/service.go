package services

import (
	"github.com/scarrionv/simple-api/app/src/model"
	"github.com/sirupsen/logrus"
)

func GetUsers() model.User {

	logrus.Info("[getUser] : Getting user")
	return model.User{Name: "Santiago", Surname: "Carrion Vivanco", Mail: "carrionvs@gmail.com"}

}
