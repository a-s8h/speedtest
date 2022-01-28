package speedtest

type TestProvider int

const (
	OOKLA TestProvider = iota
	Netflix
)

type TestResult struct {
	Download uint
	Upload   uint
}

// Test runs speed test using test provider, returns the Mbps for both download and upload
func Test(provider TestProvider) (TestResult, error) {
	return TestResult{}, nil
}
