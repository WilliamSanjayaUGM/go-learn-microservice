package models

import (
	"go-learn-api/constant"
	"go-learn-api/utils"
	"log"
)

type MstUser struct {
	Id      int64
	UserId  string
	Name    string
	Email   string
	Address string
}

func GetUserLogin(userId string, email string) *MstUser {
	transaction, errData := utils.DB.Database.Begin()
	if errData != nil {
		log.Panic(errData)
	}

	res, errData := transaction.Query(`SELECT * FROM `+constant.SCHEMA_VERSION+`fn_get_user_login_tst($1,$2);`, userId, email)
	if errData != nil {
		log.Panic(errData)
	}

	var user MstUser
	log.Print("Print Reach Here")
	for res.Next() {
		err := res.Scan(&user.Id, &user.UserId, &user.Name, &user.Email, &user.Address)
		if err != nil {
			log.Panic(err)
		}
	}

	return &user
}
