package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"
)

const PREFIX = "https://www.qdl.qa/en/search/site?page="
const SUFFIX = "&f%5B0%5D=document_source%3Aarchive_source"

var MANIFESTS sync.Map
var URLS sync.Map
var recordRe = regexp.MustCompile(`https://www.qdl.qa/en/archive/.*?"`)

func handlePage(pageUrl string, response *io.ReadCloser) {
	time.Sleep(1)
	body, err := io.ReadAll(*response)
	if err != nil {
		return
	}

	fmt.Println(string(body))
	hrefs := recordRe.FindAll(body, -1)
	for _, href := range hrefs {
		link := string(href)
		fmt.Println(link)
	}
	// logging to stderr to be able to restart from specific page (549)
	// io.WriteString(os.Stderr, pageUrl+"\n")

}

func main() {
	for i := 0; i < 2244; i++ {
		pageUrl := PREFIX + strconv.Itoa(i) + SUFFIX
		response, _ := http.Get(pageUrl)
		handlePage(pageUrl, &response.Body)
	}
}
