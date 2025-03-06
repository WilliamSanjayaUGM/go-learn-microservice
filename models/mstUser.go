package models

import (
	"go-learn-api/constant"
	"go-learn-api/utils"
	"log"
)

type MstUser struct {
	Id              int64
	UserId          string
	FullName        string
	Email           string
	PhoneNo         int64
	NationalityCode string
	NationalityName string
	Address         string
}

func GetUserLogin(email string) *MstUser {
	transaction, errData := utils.DB.Database.Begin()
	if errData != nil {
		log.Panic(errData)
	}

	res, errData := transaction.Query(`SELECT * FROM `+constant.SCHEMA_VERSION+`fn_get_user_login($1);`, email)
	if errData != nil {
		log.Panic(errData)
	}

	var user MstUser
	log.Print("Print Reach Here")
	for res.Next() {
		err := res.Scan(&user.Id, &user.UserId, &user.FullName, &user.Email, &user.PhoneNo, &user.NationalityCode, &user.NationalityName, &user.Address)
		if err != nil {
			log.Panic(err)
		}
	}

	return &user
}
