package hltv

import (
	"context"

	"github.com/chromedp/chromedp"
)

// TODO: I want this in internal/, but can't rn because how the tests are run...
func GetPage(ctx context.Context, url string, body *string) error {
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.OuterHTML("html", body, chromedp.ByQuery),
	); err != nil {
		return err
	}

	return nil
}
