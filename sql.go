package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func getSqlSettings(key string) string {
	err := godotenv.Load()
	if err != nil {
		return fmt.Sprintln("error: ", err)
	}
	return os.Getenv(key)
}

func dbConnection() {
	connectionParams := getSqlSettings("User") + ":" + getSqlSettings("Password") + "@" + "tcp(" + getSqlSettings("dbIp") + ":" + getSqlSettings("dbPort") + ")/" + getSqlSettings("dbName")
	db, err := sql.Open("mysql", connectionParams)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to db")

}
