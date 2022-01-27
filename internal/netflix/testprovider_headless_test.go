package netflix

import "testing"

func TestFastHeadlessMeasurement(t *testing.T) {
	d, u, err := MeasureHeadless()
	if err != nil {
		t.Fatal("error:", err)
	}
	t.Logf("download: %f Mbps, upload: %f Mbps", d, u)
}

func BenchmarkFastHeadlessMeasurement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = MeasureHeadless()
	}
}
