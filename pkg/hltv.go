package hltv

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	page "github.com/stef16robbe/hltv-api-go/internal/page"

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
func InitChrome() (context.Context, context.CancelFunc) {
	ctx, cancel, err := cu.New(cu.NewConfig(
		cu.WithTimeout(10 * time.Second),
		// cu.WithHeadless(),
	))
	if err != nil {
		panic(err)
	}

	return ctx, cancel
}

func GetTeamMapStats(ctx context.Context, url string) ([]MapsStats, error) {
	var body string
	var ms []MapsStats

	if err := page.GetPage(ctx, url, &body); err != nil {
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
