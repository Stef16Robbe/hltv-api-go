package hltv

import (
	"context"
	"fmt"
	"regexp"
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
func GetTeamMapStats(getPage func(context.Context, *string, string) error, ctx context.Context, url string) ([]MapsStat, error) {
	var body string
	var ms []MapsStat

	if err := getPage(ctx, &body, url); err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, err
	}

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

func GetTeamRecentMatchLinks(getPage func(context.Context, *string, string) error, ctx context.Context, url string, limit int) ([]string, error) {
	var body string
	// TODO: max size of `limit`
	matchLinks := []string{}

	if err := getPage(ctx, &body, url); err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	doc.Find("a.a-reset").Each(func(_ int, sel *goquery.Selection) {
		ml, ok := sel.Attr("href")
		if ok && strings.Contains(ml, "matches") {
			matchLinks = append(matchLinks, ml)
		}
	})

	return matchLinks, nil
}

// TODO: I want this in internal, how to without cyclical import?
func parseVeto(strVeto string) VetoStat {
	vs := VetoStat{}
	re := regexp.MustCompile(`\S+`)

	submatchall := re.FindAllString(strVeto, -1)
	// TODO: handle error cases
	// TODO: handle last veto: in HLTV it will say "map X was left over", which effectively means that the team that vetoed last picked the map that was left over...
	vs.Team = submatchall[1]
	vs.PickBan, _ = VetoString(submatchall[2])
	vs.Map, _ = MapString(submatchall[3])
	return vs
}

func GetMatchVetoes(getPage func(context.Context, *string, string) error, ctx context.Context, url string) ([]VetoStat, error) {
	var body string
	matchVetoes := []VetoStat{}

	if err := getPage(ctx, &body, url); err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	// not sure why but 3rd child works...
	doc.Find("div.standard-box.veto-box:nth-child(3) div.padding").Each(func(_ int, sel *goquery.Selection) {
		// HLTV has shitty HTML in this section, all of the map vetoes
		// are seen as a single string, so we have to split on "\n" and remove
		// any empty lines:
		vetoes := strings.Split(sel.Text(), "\n")
		for _, v := range vetoes {
			tmp := strings.Join(strings.Fields(v), " ")
			if tmp != "" {
				matchVetoes = append(matchVetoes, parseVeto(tmp))
			}
		}
	})

	return matchVetoes, nil
}
