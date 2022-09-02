package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

var db *sql.DB

func Loaddatabase() (string, error) {
	file, err := ini.Load("myconfig.ini")
	dbtype := file.Section("DataBase").Key("typeBase").Value()
	switch dbtype {
	case "sqlite3":
		FileExistVer("database.db")
		ConectionDBSqlite()
		fmt.Println("Banco sqlite")

	case "mysql":
		fmt.Println("Banco mysql")

	default:
		fmt.Println(db)
		fmt.Println("Banco de dados não suportado")
	}
	return (file.Section("Server").Key("typeBase").Value()), err

}

func FileExistVer(path string) bool {
	_, e := os.Stat(path)
	if e != nil {
		if os.IsNotExist(e) {
			return false
		}
	}
	return true
}

func ConectionDBSqlite() {

	database, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println("Error sqlite")
	}
	db = database
}
