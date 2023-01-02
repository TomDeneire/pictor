package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
)

const PREFIX = "https://hdl.loc.gov/loc.wdl/wdl."

var MANIFESTS sync.Map
var URLS sync.Map
var manifestRe = regexp.MustCompile(`https://www.loc.gov/item/.*?/manifest.json`)

func handlePage(pageUrl string, response *io.ReadCloser) {
	body, err := io.ReadAll(*response)
	if err != nil {
		return
	}

	hrefs := manifestRe.FindAll(body, -1)
	for _, href := range hrefs {
		link := string(href)
		fmt.Println(link)
	}
	// logging to stderr to be able to restart from specific page (549)
	io.WriteString(os.Stderr, pageUrl+"\n")

}

func main() {
    // get number of pages from downloaded dataset, see https://www.loc.gov/item/2020446966/
	for i := 1; i < 28005; i++ {
		pageUrl := PREFIX + strconv.Itoa(i)
		response, _ := http.Get(pageUrl)
		go handlePage(pageUrl, &response.Body)
	}
}
