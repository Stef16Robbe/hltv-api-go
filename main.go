package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

type MapsStats struct {
	name   string
	wins   uint16
	losses uint16
}

// https://www.hltv.org/stats/teams/maps/6667/faze?startDate=2023-02-28&endDate=2023-05-30&matchType=Lan&rankingFilter=Top20
func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://www.hltv.org/stats/teams/maps/6667/faze?startDate=2023-02-28&endDate=2023-05-30&matchType=Lan&rankingFilter=Top20",
				g.Opt.ParseFunc)
		},
		ParseFunc: testing,
	}).Start()
}

func testing(g *geziyor.Geziyor, r *client.Response) {
	// var mapsStats []MapsStats

	p := r.HTMLDoc.Find("div.col")
	p.Each(func(_ int, sel *goquery.Selection) {
		// TODO:
		// get properly rid of first "Map highlight map-pool-map-name"
		if c := sel.Find("div.map-pool-map-name").Text(); !strings.Contains(c, "%") && c != "" {
			fmt.Println(c)
			fmt.Println("=====")
		}
	})
}
