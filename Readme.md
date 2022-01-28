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
	// You can choose between OOKLA and NETFLIX test providers. 
	// Result contains the Mbps for both download and upload
	tr, err := speedtest.Test(speedtest.OOKLA)
	fmt.Printf("speedtest.net download: %d Mbps, upload: %d Mbps", r.Download, r.Upload)
}
```