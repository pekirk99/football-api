package nfl

import (
	"encoding/json"
	"net/http"
)

type API struct {
	Teams map[string]Team
}

func NewAPI() *API {
	return &API{
		Teams: initTeams(),
	}
}

// GetTeams handles the /teams endpoint.
func (api *API) GetTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(api.Teams)
}

func (api *API) GetTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	teamName := r.URL.Path[len("/teams/"):]
	team, exists := api.Teams[teamName]
	if !exists {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(team)
}