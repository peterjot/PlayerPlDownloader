package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_getUrlByVideoId1(t *testing.T) {
	type VideoId string
	type TestCase struct {
		name    string
		videoId VideoId
		want    string
	}

	tests := []TestCase{
		{"Same cyferki", VideoId("3213123"), "http://player.pl/api/?platform=ConnectedTV&terminal=Panasonic&format=json&authKey=064fda5ab26dc1dd936f5c6e84b7d3c2&v=3.1&m=getItem&id=3213123"},
		{"Z przecinkami", VideoId("asd,321,312,1,23"), "http://player.pl/api/?platform=ConnectedTV&terminal=Panasonic&format=json&authKey=064fda5ab26dc1dd936f5c6e84b7d3c2&v=3.1&m=getItem&id=asd,321,312,1,23"},
		{"Literki", VideoId("ddddd"), "http://player.pl/api/?platform=ConnectedTV&terminal=Panasonic&format=json&authKey=064fda5ab26dc1dd936f5c6e84b7d3c2&v=3.1&m=getItem&id=ddddd"},
		{"Znaki specjalne", VideoId("#123#213#3!@#@#@#%213123"), "http://player.pl/api/?platform=ConnectedTV&terminal=Panasonic&format=json&authKey=064fda5ab26dc1dd936f5c6e84b7d3c2&v=3.1&m=getItem&id=#123#213#3!@#@#@#%213123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareApiUrl(string(tt.videoId)); got != tt.want {
				t.Errorf("prepareApiUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_downloadByVideoId(t *testing.T) {
	videoId := "157813"
	targetDir := anyTempDirPath()

	savedFilePath := DownloadByVideoId(videoId, targetDir)

	if !strings.HasPrefix(savedFilePath, targetDir) {
		t.Error("File should be in target dir")
	}

	if !fileExists(savedFilePath) {
		t.Error("File should exists")
	}
	_ = os.Remove(savedFilePath)
	_ = os.Remove(targetDir)
}

func anyTempDirPath() string {
	targetDir, _ := ioutil.TempDir(os.TempDir(), "pvd_testing")
	return targetDir
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
