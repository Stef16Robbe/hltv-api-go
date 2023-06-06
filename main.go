package main

import (
	"fmt"

	hltv "github.com/stef16robbe/hltv-api-go/pkg"
)

// This exists to mimic how an end-user would use this module
func main() {
	url := "https://www.hltv.org/stats/teams/maps/6667/faze?startDate=2023-02-28&endDate=2023-05-30&matchType=Lan&rankingFilter=Top20"

	ctx, cancel := hltv.Init()
	defer cancel()

	ms, err := hltv.GetTeamMapStats(hltv.GetPage, ctx, url)
	if err != nil {
		panic("err retrieving mapstats")
	}

	fmt.Println(ms)
}
