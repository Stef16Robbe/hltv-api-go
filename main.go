package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"

	cu "github.com/Davincible/chromedp-undetected"
)

type MapsStats struct {
	name   string
	wins   int
	draws  int
	losses int
}

// TODO:
// allow config to be passed
func createChrome() (context.Context, context.CancelFunc) {
	ctx, cancel, err := cu.New(cu.NewConfig(
		cu.WithTimeout(10 * time.Second),
		// cu.WithHeadless(),
	))
	if err != nil {
		panic(err)
	}

	return ctx, cancel
}

func getPage(ctx context.Context, url string, body *string) error {
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.OuterHTML("html", body, chromedp.ByQuery),
	); err != nil {
		return err
	}

	return nil
}

func getTeamMapStats(ctx context.Context, url string) ([]MapsStats, error) {
	var body string
	var ms []MapsStats

	if err := getPage(ctx, url, &body); err != nil {
		return nil, err
	}

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(body))
	doc.Find("div.col").Each(func(_ int, sel *goquery.Selection) {
		stat := MapsStats{}
		if n := sel.Find("div.map-pool-map-name").Text(); !strings.Contains(n, "%") && n != "" {
			stat.name = n
		}
		if s := sel.Find("div.stats-row:first-child").Find("span:last-of-type").Text(); s != "" {
			fmt.Sscanf(s, "%d / %d / %d", &stat.wins, &stat.draws, &stat.losses)
		}

		if stat != (MapsStats{}) {
			ms = append(ms, stat)
		}
	})

	return ms, nil
}

func main() {
	url := "https://www.hltv.org/stats/teams/maps/6667/faze?startDate=2023-02-28&endDate=2023-05-30&matchType=Lan&rankingFilter=Top20"

	ctx, cancel := createChrome()
	defer cancel()

	ms, err := getTeamMapStats(ctx, url)
	if err != nil {
		panic(err)
	}

	fmt.Println(ms)
}
