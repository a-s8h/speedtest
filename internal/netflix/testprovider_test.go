package netflix

import (
	"testing"
)

func TestFastMeasurement(t *testing.T) {
	d, u, err := Measure()
	if err != nil {
		t.Fatal("error:", err)
	}
	t.Logf("download: %f Mbps, upload: %f Mbps", d, u)
}

func BenchmarkFastMeasurement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = Measure()
	}
}
