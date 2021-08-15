//go:generate goversioninfo -icon=./assets/red-skull-icon.ico -manifest=./resources/goversioninfo.exe.manifest
package main

import (
	"fmt"

	"github.com/brunoquindeler/check-hunteds-on/app"
)

func main() {

	// Servidor que deseja checar.
	var server string
	fmt.Print("Server: ")
	fmt.Scan(&server)
	// Tempo em segundos de checagem dos players
	var timeCheck = 3
	// Url da API
	var url = fmt.Sprintf("https://api.tibiadata.com/v2/world/%s.json", server)

	app.CheckPlayers(server, url, timeCheck)
}
