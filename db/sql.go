package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/nicus101/scaling-giggle/config"
)

var dbConn *sql.DB

func InitConnection() error {
	dbCfg := config.Get().Db

	connectionParams := mysql.Config{
		User:                 dbCfg.User,
		Passwd:               dbCfg.Pass,
		Net:                  "tcp",
		Addr:                 dbCfg.Addr,
		DBName:               dbCfg.Name,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", connectionParams.FormatDSN())
	if err != nil {
		return err
	}

	log.Println("Connected to db")
	dbConn = db
	return nil
}

func GetConnection() *sql.DB {
	return dbConn
}

func Close() {
	dbConn.Close()
}
