package controllers

import (
	"fmt"

	"github.com/go-ini/ini"
	"golang.org/x/text/date"
)

// Loadconfigs struct que recebe as config. do myconfig.ini
type Loadconfigs struct {
	Logs struct {
		Logpath string
	}
	ConfigDB struct {
		User string
		Pass string
		Name string
	}
	ConfiServer struct {
		Port string
		Url  string
	}
}

// LoadConfigDataBase carrega as configurações do banco de dados e disponibiliza na struct
func LoadConfigDataBase() Loadconfigs {
	file, err := ini.Load("myconfig.ini")
	if err != nil {
		fmt.Println(date.Brasilia, "Não foi possivel carregar o arquivo de configurações do banco de dados: ", err)
	}
	db := Loadconfigs{}

	db.ConfigDB.User = file.Section("DataBase").Key("user").Value()
	db.ConfigDB.Pass = file.Section("DataBase").Key("pass").Value()
	db.ConfigDB.Name = file.Section("DataBase").Key("ipserver").Value()

	return db
}

// LoadConfigServer carrega as configurações do Servidor e disponibiliza na struct
func LoadConfigServer() Loadconfigs {
	file, err := ini.Load("myconfig.ini")
	if err != nil {
		fmt.Println(date.Brasilia, "Não foi possivel carregar o arquivo de configurações do Servidor: ", err)
	}
	scfg := Loadconfigs{}

	scfg.ConfiServer.Port = file.Section("Server").Key("port").Value()
	scfg.ConfiServer.Url = file.Section("Server").Key("url").Value()

	return scfg
}
