package models

import (
	"go-learn-api/utils"
)

func ConnectDatabase() {
	utils.DatabaseConfig().Database.SetMaxOpenConns(10)
}
