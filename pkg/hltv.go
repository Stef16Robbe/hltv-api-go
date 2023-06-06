package hltv

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	cu "github.com/Davincible/chromedp-undetected"
)

// TODO: allow config to be passed
func Init() (context.Context, context.CancelFunc) {
	ctx, cancel, err := cu.New(cu.NewConfig(
		cu.WithTimeout(10 * time.Second),
		// cu.WithHeadless(),
	))
	if err != nil {
		panic(err)
	}

	return ctx, cancel
}

// https://www.hltv.org/stats/teams/maps/6667/faze?startDate=2023-02-28&endDate=2023-05-30&matchType=Lan&rankingFilter=Top20
func GetTeamMapStats(getPage func(context.Context, string, *string) error, ctx context.Context, url string) ([]MapsStat, error) {
	var body string
	var ms []MapsStat

	if err := getPage(ctx, url, &body); err != nil {
		return nil, err
	}

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(body))
	doc.Find("div.col").Each(func(_ int, sel *goquery.Selection) {
		stat := MapsStat{}
		if n := sel.Find("div.map-pool-map-name").Text(); !strings.Contains(n, "%") && n != "" {
			m, err := ParseMap(n)
			if err != nil {
				panic("incorrect map")
			}

			stat.Map = m
		}
		if s := sel.Find("div.stats-row:first-child").Find("span:last-of-type").Text(); s != "" {
			fmt.Sscanf(s, "%d / %d / %d", &stat.Wins, &stat.Draws, &stat.Losses)
		}

		if stat != (MapsStat{}) {
			ms = append(ms, stat)
		}
	})

	return ms, nil
}
