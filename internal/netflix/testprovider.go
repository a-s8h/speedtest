package netflix

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gesquive/fast-cli/fast"
	"golang.org/x/sync/errgroup"
)

const (
	workersCount  = 8
	payloadSizeMB = 25.0 // download payload is by default 25MB, make upload 25MB also
)

func measureNetworkSpeed(operation func() error) (float64, error) {
	eg := errgroup.Group{}
	sTime := time.Now()
	for i := 0; i < workersCount; i++ {
		eg.Go(func() error {
			return operation()
		})
	}
	if err := eg.Wait(); err != nil {
		return 0, err
	}
	fTime := time.Now()
	return payloadSizeMB * 8 * float64(workersCount) / fTime.Sub(sTime).Seconds(), nil
}

func downloadOperation(client http.Client, url string) error {
	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	ioutil.ReadAll(res.Body)
	return nil
}

func uploadOperation(client http.Client, uri string, uploadPayload string) error {
	v := url.Values{}
	//10b * x MB / 10 = x MB
	v.Add("content", uploadPayload)
	resp, err := client.PostForm(uri, v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)
	return nil
}

// Measure calculates download and upload speed using fast.com
func Measure() (download, upload float64, err error) {
	urls := fast.GetDlUrls(1)
	if len(urls) == 0 {
		return 0, 0, errors.New("server urls are not available")
	}
	client := http.Client{}
	uploadPayload := strings.Repeat("0123456789", payloadSizeMB*1024*1024/10)
	download, err = measureNetworkSpeed(func() error {
		return downloadOperation(client, urls[0])
	})
	upload, err = measureNetworkSpeed(func() error {
		return uploadOperation(client, urls[0], uploadPayload)
	})
	if err != nil {
		return 0, 0, err
	}
	return download, upload, nil
}
