package main

import (
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

func main() {
	// url := "https://nowsecure.nl"
	// chromedp.WaitVisible(`//div[@class="hystericalbg"]`),

	url := "https://www.hltv.org/stats/teams/maps/6667/faze?startDate=2023-02-28&endDate=2023-05-30&matchType=Lan&rankingFilter=Top20"

	// New creates a new context for use with chromedp. With this context
	// you can use chromedp as you normally would.
	ctx, cancel, err := cu.New(cu.NewConfig(
		// Remove this if you want to see a browser window.
		// cu.WithHeadless(),

		// If the webelement is not found within 10 seconds, timeout.
		cu.WithTimeout(10 * time.Second),
	))
	if err != nil {
		panic(err)
	}
	defer cancel()

	var body string

	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.OuterHTML("html", &body, chromedp.ByQuery),
	); err != nil {
		panic(err)
	}

	var ms []MapsStats

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(body))
	doc.Find("div.col").Each(func(_ int, sel *goquery.Selection) {
		stat := MapsStats{}
		if n := sel.Find("div.map-pool-map-name").Text(); !strings.Contains(n, "%") && n != "" {
			stat.name = n
		}
		if s := sel.Find("div.stats-row:first-child").Find("span:last-of-type").Text(); s != "" {
			var wins, draws, losses int

			fmt.Sscanf(s, "%d / %d / %d", &wins, &draws, &losses)

			stat.wins = wins
			stat.draws = draws
			stat.losses = losses
		}

		if stat != (MapsStats{}) {
			ms = append(ms, stat)
		}
	})

	fmt.Println(ms)
}
