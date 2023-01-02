package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const PREFIX = "https://collections.peabody.yale.edu"

var MANIFESTS sync.Map
var URLS sync.Map
var manifestRe = regexp.MustCompile(`href=".*?manifest.*?"`)
var recordRe = regexp.MustCompile(`href="/search/Record.*?"`)

func handleRecord(url string) {
	response, _ := http.Get(url)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	hrefs := manifestRe.FindAll(body, -1)
	for _, href := range hrefs {
		link := string(href)
		if strings.Contains(link, "manifest=https://manifests.collections.yale.edu") {
			_, handled := MANIFESTS.Load(link)
			if !handled {
				MANIFESTS.Store(link, true)
				manifest := strings.Split(link, "manifest=")[1]
				manifest = strings.Split(manifest, "&canvas")[0]
				fmt.Println(manifest)
			}
		}
	}

}

func handlePage(pageUrl string, response *io.ReadCloser) {
	body, err := io.ReadAll(*response)
	if err != nil {
		return
	}

	hrefs := recordRe.FindAll(body, -1)
	for _, href := range hrefs {
		link := string(href)
		if !strings.Contains(link, "/Save") &&
			!strings.Contains(link, "/Export") {
			recordURL := PREFIX + strings.ReplaceAll(link, `href="`, "")
			recordURL = strings.Trim(recordURL, `"`)
			_, handled := URLS.Load(recordURL)
			if !handled {
				URLS.Store(recordURL, true)
				handleRecord(recordURL)
			}
		}
	}
	// logging to stderr to be able to restart from specific page 
	io.WriteString(os.Stderr, pageUrl+"\n")

}

func main() {
	for i := 548; i < 414174; i++ {
		// for i := 1; i < 414174; i++ {
		pageUrl := PREFIX + "/search/Search/Results?type=AllFields&page=" + strconv.Itoa(i)
		response, _ := http.Get(pageUrl)
		go handlePage(pageUrl, &response.Body)
	}
}
