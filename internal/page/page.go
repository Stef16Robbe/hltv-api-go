package page

import (
	"context"

	"github.com/chromedp/chromedp"
)

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
