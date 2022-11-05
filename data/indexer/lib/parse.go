package lib

import (
	"encoding/json"
	"strings"
)

func parseV2(manifestData []byte) (ManifestInfo, error) {

	var info ManifestInfo
	var manifest ManifestV2
	err := json.Unmarshal(manifestData, &manifest)
	if err != nil {
		return info, err
	}
	info.Identifier = manifest.ID
	if len(manifest.Sequences) == 0 {
		return info, err
	}
	if len(manifest.Sequences[0].Canvases) == 0 {
		return info, err
	}
	if len(manifest.Sequences[0].Canvases[0].Images) == 0 {
		return info, err
	}
	if manifest.Sequences[0].Canvases[0].Thumbnail.ID != "" {
		info.Thumbnail = manifest.Sequences[0].Canvases[0].Thumbnail.ID
	} else {
		info.Thumbnail = manifest.Sequences[0].Canvases[0].Images[0].Resource.ID
	}
	info.Label, err = InterfaceToString(manifest.Label)
	if err != nil {
		return info, err
	}
	info.Summary, err = InterfaceToString(manifest.Label)
	if err != nil {
		return info, err
	}
	return info, nil
}

func parseV3(manifestData []byte) (ManifestInfo, error) {

	var info ManifestInfo
	var manifest ManifestV3
	err := json.Unmarshal(manifestData, &manifest)
	if err != nil {
		return info, err
	}
	info.Identifier = manifest.ID
	if len(manifest.Items) == 0 {
		return info, err
	}
	info.Thumbnail = manifest.Items[0].Items[0].Items[0].Body.ID
	info.Label, err = InterfaceToString(manifest.Label)
	if err != nil {
		return info, err
	}
	info.Summary, err = InterfaceToString(manifest.Summary)
	if err != nil {
		return info, err
	}
	return info, nil
}

func ParseManifest(manifestData []byte) (ManifestInfo, error) {

	var info ManifestInfo
	var manifest Manifest
	err := json.Unmarshal(manifestData, &manifest)
	if err != nil {
		return info, err
	}
	version := ""
	if manifest.Context != "" {
		version = strings.Split(manifest.Context, "/")[5]
	} else {
		version = "2"
	}
	if version == "3" {
		info, err = parseV3(manifestData)
		if err != nil {
			return info, err
		}
	} else {
		info, err = parseV2(manifestData)
		if err != nil {
			return info, err
		}
	}

	return info, nil
}
