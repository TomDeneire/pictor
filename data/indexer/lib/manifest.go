package lib

import (
	"fmt"
	"io"
	"net/http"
)

func ManifestWrapper(manifests []string) error {

	HandleManifest := func(i int) (interface{}, error) {

		url := manifests[i]
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		manifest, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		// handle manifest
		info, err := ParseManifest(manifest)
		if err != nil {
			return nil, err
		}

		// log manifest metadata
		fmt.Println(url + " tdn:T " + info.Thumbnail)
		fmt.Println(url + " tdn:L " + info.Label)
		fmt.Println(url + " tdn:S " + info.Summary)

		// log index info
		words, err := SplitIntoIndexWords(info.Label + " " + info.Summary)
		if err != nil {
			panic(err)
		}
		for _, word := range words {
			fmt.Println(url + " tdn:W " + word)
		}

		return nil, nil
	}

	NMap(len(manifests), 0, HandleManifest)

	return nil
}
