package hltv_test

import (
	"testing"

	hltv "github.com/stef16robbe/hltv-api-go/pkg"
	td "github.com/stef16robbe/hltv-api-go/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCreateChrome(t *testing.T) {
	assert.NotPanics(t, func() {
		_, cancel := hltv.InitChrome()

		defer cancel()
	})
}

func TestGetTeamMapStats(t *testing.T) {
	msTest := td.GetTestGetTeamMapStatsData()
	url := "https://www.hltv.org/stats/teams/maps/6667/faze?startDate=2023-02-28&endDate=2023-05-30&matchType=Lan&rankingFilter=Top20"

	ms, err := hltv.GetTeamMapStats(mockGetPage, nil, url)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, msTest, ms)
}
