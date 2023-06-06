package test_data

import hltv "github.com/stef16robbe/hltv-api-go/pkg"

func TeamMapStatsData() []hltv.MapsStats {
	return []hltv.MapsStats{
		{
			Name:   "Inferno",
			Wins:   5,
			Draws:  0,
			Losses: 5,
		},
		{
			Name:   "Overpass",
			Wins:   5,
			Draws:  0,
			Losses: 5,
		},
		{
			Name:   "Mirage",
			Wins:   5,
			Draws:  0,
			Losses: 4,
		},
		{
			Name:   "Nuke",
			Wins:   4,
			Draws:  0,
			Losses: 5,
		},
		{
			Name:   "Ancient",
			Wins:   5,
			Draws:  0,
			Losses: 2,
		},
		{
			Name:   "Anubis",
			Wins:   3,
			Draws:  0,
			Losses: 4,
		},
	}
}
