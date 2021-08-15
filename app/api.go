package app

type api struct {
	World world `json:"world"`
	// Information information `json: "information"`
}

type world struct {
	PlayersOnline []player `json:"players_online"`
}

type player struct {
	Name     string `json:"name"`
	Level    int    `json:"level"`
	Vocation string `json:"vocation"`
}
