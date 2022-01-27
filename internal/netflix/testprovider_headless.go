package netflix

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

// MeasureHeadless calculates download and upload speed using fast.com
func MeasureHeadless() (download, upload float64, err error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	var d, u string
	err = chromedp.Run(ctx,
		chromedp.Navigate(`https://fast.com`),
		chromedp.WaitVisible(`#show-more-details-link`),
		chromedp.Click(`#show-more-details-link`, chromedp.NodeVisible),
		chromedp.WaitVisible(`#upload-value.succeeded`),
		chromedp.Text(`#speed-value`, &d),
		chromedp.Text(`#upload-value`, &u),
	)
	if err != nil {
		return 0, 0, err
	}

	download, err = strconv.ParseFloat(d, 64)
	if err != nil {
		return 0, 0, err
	}
	upload, err = strconv.ParseFloat(u, 64)
	if err != nil {
		return 0, 0, err
	}
	return download, upload, nil
}
