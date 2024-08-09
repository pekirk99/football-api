package main

import (
	"net/http"

	"football-api/internal/nfl"
)

func main() {
	api := nfl.NewAPI()

	http.HandleFunc("/teams", api.GetTeams)
	http.HandleFunc("/teams/", api.GetTeam)
	http.ListenAndServe(":8080", nil)

}