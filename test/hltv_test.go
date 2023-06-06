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

	ms, err := hltv.GetTeamMapStats(mockGetPage, nil, "")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, msTest, ms)
}
