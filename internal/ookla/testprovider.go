package ookla

import (
	"github.com/showwin/speedtest-go/speedtest"
)

// Measure calculates download and upload speed using speedtest.net
func Measure() (download, upload float64, err error) {
	user, _ := speedtest.FetchUserInfo()
	serverList, err := speedtest.FetchServerList(user)
	if err != nil {
		return 0, 0, err
	}
	testServer := serverList.Servers[0]

	err = testServer.DownloadTest(false)
	if err != nil {
		return 0, 0, err
	}

	err = testServer.UploadTest(false)
	if err != nil {
		return 0, 0, err
	}

	return testServer.DLSpeed, testServer.ULSpeed, nil
}
