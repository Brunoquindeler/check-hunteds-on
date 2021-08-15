package main

import (
	"fmt"
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

	checkPlayers(server, url, timeCheck)
}
