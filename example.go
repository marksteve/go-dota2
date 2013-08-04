package main

import (
	"fmt"
	"github.com/marksteve/go-dota2/dota2"
)

func main() {
	cli := dota2.NewDota2Client("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	for _, l := range cli.GetLeagueListing() {
		fmt.Printf(
			"\nName: %s\nLeagueId: %d\nDescription: %s\nTournamentUrl: %s\nItemDef: %d\n",
			l.Name, l.LeagueId, l.Description, l.TournamentUrl, l.ItemDef,
		)
	}
}
