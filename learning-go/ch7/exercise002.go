package main

import "fmt"

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResults(team1 string, score1 int, team2 string, score2 int) {
	// Ensure they key exists in the map.
	if _, ok := l.Teams[team1]; !ok {
		return
	}

	// Ensure they key exists in the map.
	if _, ok := l.Teams[team2]; !ok {
		return
	}

	if score1 > score2 {
		l.Wins[team1]++
	} else {
		l.Wins[team2]++
	}

}

func main() {

	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}

	l.MatchResults("Italy", 50, "France", 70)
	l.MatchResults("India", 85, "Niger", 70)
	l.MatchResults("Italy", 60, "India", 70)
	l.MatchResults("France", 100, "Nigeria", 70)
	l.MatchResults("Italy", 65, "Nigeria", 70)
	l.MatchResults("France", 95, "India", 70)

	fmt.Println(l.Wins)
}
