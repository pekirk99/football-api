package nfl

type Team struct {
	Name string `json:"name"`
	Offense []Player `json:"offense"`
	Defense []Player `json:"defense"`
}