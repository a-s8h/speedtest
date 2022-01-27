package speedtest

import (
	nf "github.com/a-s8h/speedtest/internal/netflix"
	speedtest "github.com/a-s8h/speedtest/internal/ookla"
)

// TestProvider type, should be used to choose speed test provider
type TestProvider int

const (
	// test provider, speedtest.net
	ookla TestProvider = iota
	// netflix test provider, fast.com
	netflix
	// netflixHeadlessChrome test provider, fast.com, uses headless chrome
	netflixHeadlessChrome
)

// TestResult contains download and upload speed in Mbps
type TestResult struct {
	Download float64
	Upload   float64
}

// Test runs speed test using test provider, returns the Mbps for both download and upload
func Test(provider TestProvider) (TestResult, error) {
	switch provider {
	case ookla:
		d, u, err := speedtest.Measure()
		if err != nil {
			return TestResult{}, err
		}
		return TestResult{Download: d, Upload: u}, nil
	case netflix:
		d, u, err := nf.Measure()
		if err != nil {
			return TestResult{}, err
		}
		return TestResult{Download: d, Upload: u}, nil
	case netflixHeadlessChrome:
		d, u, err := nf.Measure()
		if err != nil {
			return TestResult{}, err
		}
		return TestResult{Download: d, Upload: u}, nil
	}
	return TestResult{}, nil
}
