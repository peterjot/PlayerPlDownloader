package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	ApuUrl           = "http://player.pl/api/?platform=ConnectedTV&terminal=Panasonic&format=json&authKey=064fda5ab26dc1dd936f5c6e84b7d3c2&v=3.1&m=getItem&id=#videoId"
	DefaultTargetDir = "."
)

func Download(videoWebsiteUrl string, targetDir string) string {
	videoId := getVideoId(videoWebsiteUrl)
	fmt.Printf("Found video id:%s\n", videoId)
	return DownloadByVideoId(videoId, targetDir)
}

func DownloadByVideoId(videoId string, targetDir string) string {
	if videoId == "" {
		log.Fatal("Video id is empty")
	}
	videoSourceUrl := getVideoSourceUrl(videoId)
	return downloadFromSource(videoId, targetDir, videoSourceUrl)
}

func downloadFromSource(videoId string, targetDir string, videoSourceUrl string) string {
	fmt.Println("Fetching video...")
	resp, err := http.Get(videoSourceUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Creating file...")

	filePath := targetDir + "/" + videoId + ".mp4"
	out, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	fmt.Println("Reading to file...")
	if _, err = io.Copy(out, resp.Body); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Video saved to: " + filePath)
	return filePath
}

func getVideoSourceUrl(videoId string) string {
	videoApiUrl := prepareApiUrl(videoId)

	body := getResponseAsBytes(videoApiUrl)

	response := Response{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}

	return response.Item.Videos.Main.VideoContent[0].URL
}

func getResponseAsBytes(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func prepareApiUrl(videoId string) string {
	return strings.Replace(ApuUrl, "#videoId", videoId, 1)
}

func getVideoId(url string) string {
	id := serialIdParser(url)

	if id == "" {
		id = voidIdParser(url)
	}

	if id == "" {
		log.Fatal("Cannot find video id")
	}

	return id
}

func serialIdParser(url string) string {
	return getValue(url, "odcinki,(\\d+)/.*,(\\d+)", 2)
}

func voidIdParser(url string) string {
	return getValue(url, ",(\\d+)", 1)
}

func getValue(data string, expr string, position int) string {
	compile := regexp.MustCompile(expr)
	results := compile.FindStringSubmatch(data)

	if results == nil || len(results) < position-1 {
		return ""
	}

	return results[position]
}

func main() {
	urlParam := flag.String("urlParam", "", "Link to player.pl video")
	targetDirParam := flag.String("target", DefaultTargetDir, "Target directory")
	videoIdParam := flag.String("videoid", "", "You can put video id instead of url")
	flag.Parse()

	if *targetDirParam != DefaultTargetDir {
		fmt.Println("Provided target dir path: " + *targetDirParam)
	} else {
		fmt.Println("Target dir path has not been provided. Using default path: " + DefaultTargetDir)
	}

	if *videoIdParam != "" {
		fmt.Println("Video id provided. Downloading by id...")
		DownloadByVideoId(*videoIdParam, *targetDirParam)
	} else if *urlParam != "" {
		fmt.Println("Video id provided. Downloading by id...")
		Download(*urlParam, *targetDirParam)
	} else {
		log.Fatal("Please provide video urlParam or video id")
	}
}
