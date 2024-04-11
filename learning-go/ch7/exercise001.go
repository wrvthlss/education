package main

import (
	"fmt"
	"sort"
)

type Class struct {
	Name string
}

type Duel struct {
	MatchUps map[string]Class
	Wins     map[string]int
}

func (d *Duel) DualResults(classOne, classTwo string, victoryOne bool) {
	//
	if _, ok := d.MatchUps[classOne]; !ok {
		return
	}
	if _, ok := d.MatchUps[classTwo]; !ok {
		return
	}

	if victoryOne {
		d.Wins[classOne]++
	} else {
		d.Wins[classTwo]++
	}
}

func Ranking(wins map[string]int) []string {
	// Creating a slice of keys from the map.
	keys := make([]string, 0, len(wins))

	for key := range wins {
		keys = append(keys, key)
	}

	// Sort the keys slice based on the victory counts in descending order
	sort.Slice(keys, func(i, j int) bool {
		return wins[keys[i]] > wins[keys[j]]
	})

	return keys

}

func main() {

	d := Duel{
		// Map class names to their corresponding `Class` struct representations.
		MatchUps: map[string]Class{
			"Templar":   {Name: "Templar"},
			"Assassin":  {Name: "Assassin"},
			"Cleric":    {Name: "Cleric"},
			"Gladiator": {Name: "Gladiator"},
			"Sorcerer":  {Name: "Sorcerer"},
		},
		Wins: make(map[string]int),
	}

	d.DualResults("Templar", "Assassin", true)
	d.DualResults("Cleric", "Assassin", false)
	d.DualResults("Sorcerer", "Assassin", false)
	d.DualResults("Gladiator", "Assassin", true)
	d.DualResults("Templar", "Gladiator", true)
	d.DualResults("Templar", "Cleric", false)
	d.DualResults("Templar", "Sorcerer", true)

	rankedClasses := Ranking(d.Wins)
	for i, className := range rankedClasses {
		fmt.Printf("%d. %s - %d wins\n", i+1, className, d.Wins[className])
	}

}
