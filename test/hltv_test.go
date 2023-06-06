package hltv_test

import (
	"testing"

	hltv "github.com/stef16robbe/hltv-api-go/pkg"
	"github.com/stretchr/testify/assert"
)

func TestCreateChrome(t *testing.T) {
	assert.NotPanics(t, func() {
		_, cancel := hltv.InitChrome()

		defer cancel()
	})

	// url := "https://www.hltv.org/stats/teams/maps/6667/faze?startDate=2023-02-28&endDate=2023-05-30&matchType=Lan&rankingFilter=Top20"
	// ctx, cancel := hltv.InitChrome()
	// defer cancel()
	// ms, err := hltv.GetTeamMapStats(ctx, url)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(ms)
}
