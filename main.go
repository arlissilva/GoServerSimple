package main

import (
	"fmt"
	"main/database"
	"main/routes"
	"net/http"
	"time"
)

func Test(w http.ResponseWriter, r *http.Request) {

}

func main() {
	teste, err := database.Loaddatabase()
	if err != nil {
		fmt.Println("Deu ruim no arquivo", err)
	}

	routes.Routes()

	fmt.Println(time.Now().UTC().Format("02-01-2006 15:04:05"), "teste", teste)

}
