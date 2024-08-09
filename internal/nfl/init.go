package nfl

func initTeams() map[string]Team {
	return map[string]Team{
		"vikings": {
			Name: "Minnesota Vikings",
			Offense: []Player{
				{Name: "Kirk Cousins", Position: "QB"},
				{Name: "Dalvin Cook", Position: "RB"},
			},
			Defense: []Player{
				{Name: "Eric Kendricks", Position: "LB"},
				{Name: "Harrison Smith", Position: "S"},
			},
		},
		"bears": {
			Name: "Chicago Bears",
			Offense: []Player{
				{Name: "Justin Fields", Position: "QB"},
				{Name: "David Montgomery", Position: "RB"},
			},
			Defense: []Player{
				{Name: "Roquan Smith", Position: "LB"},
				{Name: "Eddie Jackson", Position: "S"},
			},
		},
	}
}