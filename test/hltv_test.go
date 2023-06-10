package hltv_test

import (
	"testing"

	hltv "github.com/stef16robbe/hltv-api-go/pkg"
	td "github.com/stef16robbe/hltv-api-go/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert.NotPanics(t, func() {
		_, cancel := hltv.Init()

		defer cancel()
	})
}

func TestGetTeamMapStats(t *testing.T) {
	msTest := td.TeamMapStatsData()

	// not using chromedp to scrape with these requests,
	// so we don't care about the context
	ms, err := hltv.GetTeamMapStats(mockGetPage, nil, "get_team_map_stats.html")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, msTest, ms)
}

// TODO: FIXME
func TestGetTeamVetoStats(t *testing.T) {
	vstatTest := td.VetoStatsData()

	vs, err := hltv.GetTeamRecentMatchLinks(mockGetPage, nil, "get_veto_stats.html")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, vstatTest, vs)
}
