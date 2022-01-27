# speedtest

Minimal library for testing an internet connection using Ookla's (https://www.speedtest.net) and Netflix's (https://fast.com/.)

## Installation

```
go get -u github.com/a-s8h/speedtest
```

## Usage

```Go
import "github.com/a-s8h/speedtest/pkg/speedtest"

func main() {
	// You can choose between ookla and netflix, netflixHeadlessChrome test providers. 
	// Result contains the Mbps for both download and upload
	tr, err := speedtest.Test(speedtest.ookla)
	fmt.Printf("speedtest.net download: %f Mbps, upload: %f Mbps", r.Download, r.Upload)
}
```

## Dependencies

* showwin/speedtest-go, uses for speedtest.net
* gesquive/fast-cli, used to get api url for fast.com provider
* chromedp/chromedp, used to scrape https://fast.com, google chrome required