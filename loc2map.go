package loc2map

import (
	"context"
	"errors"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"io/ioutil"
	"log"
	"math"
	"net/url"

	"github.com/chromedp/chromedp"
)

// Convert accepts a lat, lng, and filepath (and name ending with .png)
// and saves a map in the specified file
func Convert(lat, lng float64, filePath string) error {
	if filePath == "" {
		return errors.New("filepath is required")
	}

	// Format the URL with the lat, lng, and api key
	mapUrl, err := getURL(lat, lng)
	if err != nil {
		return err
	}
	// Get the screenshot image data
	buf, err := getMapScreen(mapUrl)
	if err != nil {
		return err
	}

	// Write the image data to file
	if err := ioutil.WriteFile(filePath, buf, 0644); err != nil {
		log.Fatal(err)
	}

	return nil
}

// getMapScreen accepts a mapUrl and returns image data as []byte & nil, or nil & an error
func getMapScreen(mapUrl string) ([]byte, error) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// capture screenshot of an element
	var buf []byte
	if err := chromedp.Run(ctx, fullScreenshot(mapUrl,90, &buf)); err != nil {
		return nil, err
	}
	return buf, nil
}

// getURL creates a google maps url with latitude, longitude, and google maps api key
func getURL(lat, lng float64) (string, error) {
	mapUrlTemplate := "https://www.google.com/maps/@?api=1&map_action=map&center=%v,%v&zoom=11"
	u, err := url.Parse(fmt.Sprintf(mapUrlTemplate, lat, lng))
	if err != nil {
		return "", err
	}
	fmt.Printf("got url: %s", u.String())
	return u.String(), nil
}

// Note: copied from: https://github.com/chromedp/examples/blob/master/screenshot/main.go
// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Liberally copied from puppeteer's source.
//
// Note: this will override the viewport emulation settings.
func fullScreenshot(urlstr string, quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// get layout metrics
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			// capture screenshot
			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
