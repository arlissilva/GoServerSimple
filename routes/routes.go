package routes

import (
	control "main/controllers"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", control.RequestController)
	http.HandleFunc("/post", control.RequestController)

	cfgS := new(control.Loadconfigs)

	cfgS.ConfiServer = control.LoadConfigServer().ConfiServer

	URI := cfgS.ConfiServer.Url + ":" + cfgS.ConfiServer.Port
	http.ListenAndServe(URI, nil)
}
