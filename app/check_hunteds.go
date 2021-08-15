package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func CheckPlayers(server, url string, timeCheck int) {
	for {

		var h = Hunteds{}

		h.NewHunteds()

		var huntedsOn []string

		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(body)
		}

		content := api{}

		err = json.Unmarshal(body, &content)
		if err != nil {
			log.Fatal(err)
		}

		for _, player := range content.World.PlayersOnline {

			for _, hunted := range h.hunteds {
				if player.Name == hunted {
					huntedsOn = append(huntedsOn, player.Name)
					fmt.Printf("Hunted ON: %v\n", player.Name)
				}
			}
		}

		msg := fmt.Sprintf("Total de Hunteds ON: %v/%v\n", len(huntedsOn), len(h.hunteds))
		fmt.Printf(msg)

		// Notificações para Windows 10
		// go get github.com/go-toast/toast
		//
		// notification := toast.Notification{
		// 	AppID:   "Hunteds Check",
		// 	Title:   "Hunteds ON",
		// 	Message: msg,
		// 	//Icon:    "/alerta.png", // This file must exist (remove this line if it doesn't)
		// 	Actions: []toast.Action{
		// 		{"protocol", "Ok", ""},
		// 	},
		// }
		// errnotf := notification.Push()
		// if errnotf != nil {
		// 	log.Fatalln(errnotf)
		// }

		time.Sleep(time.Second * time.Duration(timeCheck))
		CallClear()
	}
}
