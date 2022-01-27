package speedtest

import "testing"

func TestOokla(t *testing.T) {
	r, err := Test(ookla)
	if err != nil {
		t.Fatal("error:", err)
	}
	t.Logf("download: %f Mbps, upload: %f Mbps", r.Download, r.Upload)
}

func TestNetflix(t *testing.T) {
	r, err := Test(netflix)
	if err != nil {
		t.Fatal("error:", err)
	}
	t.Logf("download: %f Mbps, upload: %f Mbps", r.Download, r.Upload)
}

func TestNetflixHeadless(t *testing.T) {
	r, err := Test(netflixHeadlessChrome)
	if err != nil {
		t.Fatal("error:", err)
	}
	t.Logf("download: %f Mbps, upload: %f Mbps", r.Download, r.Upload)
}
