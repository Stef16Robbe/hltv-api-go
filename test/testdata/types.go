package test_data

import hltv "github.com/stef16robbe/hltv-api-go/pkg"

func TeamMapStatsData() []hltv.MapsStat {
	return []hltv.MapsStat{
		{
			Map:    hltv.Inferno,
			Wins:   5,
			Draws:  0,
			Losses: 5,
		},
		{
			Map:    hltv.Overpass,
			Wins:   5,
			Draws:  0,
			Losses: 5,
		},
		{
			Map:    hltv.Mirage,
			Wins:   5,
			Draws:  0,
			Losses: 4,
		},
		{
			Map:    hltv.Nuke,
			Wins:   4,
			Draws:  0,
			Losses: 5,
		},
		{
			Map:    hltv.Ancient,
			Wins:   5,
			Draws:  0,
			Losses: 2,
		},
		{
			Map:    hltv.Anubis,
			Wins:   3,
			Draws:  0,
			Losses: 4,
		},
	}
}
