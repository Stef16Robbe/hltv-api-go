package main

import (
	"fmt"

	hltv "github.com/stef16robbe/hltv-api-go/pkg"
)

// This exists to mimic how an end-user would use this module
func main() {
	// https://www.hltv.org/stats/teams/maps/6667/faze?startDate=2023-02-28&endDate=2023-05-30&matchType=Lan&rankingFilter=Top20
	// https://www.hltv.org/results?startDate=2023-03-10&endDate=2023-06-10&team=6667&matchType=Lan
	// https://www.hltv.org/matches/2364750/vitality-vs-faze-blast-premier-spring-final-2023

	ctx, cancel := hltv.Init()
	defer cancel()

	links, err := hltv.GetTeamRecentMatchLinks(hltv.GetPage, ctx, "https://www.hltv.org/results?startDate=2023-03-10&endDate=2023-06-10&team=6667&matchType=Lan", 3)

	if err != nil {
		msg := fmt.Sprintf("err retrieving stats: %v", err)
		panic(msg)
	}
	fmt.Println("found these links:\n", links)

	var allVetoes []hltv.VetoStat
	for _, l := range links {
		vetoes, err := hltv.GetMatchVetoes(hltv.GetPage, ctx, l)
		allVetoes = append(allVetoes, vetoes...)

		if err != nil {
			msg := fmt.Sprintf("err retrieving stats: %v", err)
			panic(msg)
		}
	}

	fmt.Println("And here are the vetoes from FaZe:\n", allVetoes)
}
