package main

import (
	"fmt"

	hltv "github.com/stef16robbe/hltv-api-go/pkg"
)

// This exists to mimic how an end-user would use this module
func main() {
	// url := "https://www.hltv.org/stats/teams/maps/6667/faze?startDate=2023-02-28&endDate=2023-05-30&matchType=Lan&rankingFilter=Top20"
	// url := "https://www.hltv.org/results?startDate=2023-03-10&endDate=2023-06-10&team=6667&matchType=Lan"
	url := "https://www.hltv.org/matches/2364750/vitality-vs-faze-blast-premier-spring-final-2023"

	ctx, cancel := hltv.Init()
	defer cancel()

	// TODO: we have match links, now follow them and return vetoes...
	stat, err := hltv.GetMatchVetoes(hltv.GetPage, ctx, url)
	if err != nil {
		msg := fmt.Sprintf("err retrieving stats: %v", err)
		panic(msg)
	}

	fmt.Println(stat)
}
