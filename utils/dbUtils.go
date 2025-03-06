package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Database *sql.DB
}

var DB Database

func DatabaseConfig() Database {
	//Kalau Postgres, tambahkan ssid/sslmode nya
	database, err := sql.Open(ReadConfig().DbType, ReadConfig().DbType+"://"+ReadConfig().Username+":"+ReadConfig().Password+"@"+ReadConfig().Host+":"+ReadConfig().Port+"/"+ReadConfig().DbName+"?"+ReadConfig().SslMode)
	if err != nil {
		log.Panic(err)
	}
	log.Print("Database Connected Successfully")

	DB.Database = database
	return Database{Database: database}
}
