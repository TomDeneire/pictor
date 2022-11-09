package lib

type ManifestInfo struct {
	Label      string
	Summary    string
	Thumbnail  string
	Identifier string
}

type Manifest struct {
	Context string `json:"@context"`
}

type ManifestV2 struct {
	Context   string      `json:"@context"`
	ID        string      `json:"@id"`
	Label     interface{} `json:"label"`
	Sequences []struct {
		ID       string `json:"@id"`
		Type     string `json:"@type"`
		Canvases []struct {
			ID     string `json:"@id"`
			Type   string `json:"@type"`
			Height int64  `json:"height"`
			Images []struct {
				ID         string `json:"@id"`
				Type       string `json:"@type"`
				Motivation string `json:"motivation"`
				On         string `json:"on"`
				Resource   struct {
					ID      string `json:"@id"`
					Type    string `json:"@type"`
					Format  string `json:"format"`
					Height  int64  `json:"height"`
					Service struct {
						Context string `json:"@context"`
						ID      string `json:"@id"`
						Profile string `json:"profile"`
					} `json:"service"`
					Width int64 `json:"width"`
				} `json:"resource"`
			} `json:"images"`
			Thumbnail struct {
				ID     string `json:"@id"`
				Type   string `json:"@type"`
				Format string `json:"format"`
				Height int64  `json:"height"`
				Width  int64  `json:"width"`
			} `json:"thumbnail"`
			Label string `json:"label"`
			Width int64  `json:"width"`
		} `json:"canvases"`
		Label interface{} `json:"label"`
	} `json:"sequences"`
}

type ManifestV3 struct {
	Context string `json:"@context"`
	ID      string `json:"id"`
	Items   []struct {
		Height int64  `json:"height"`
		ID     string `json:"id"`
		Items  []struct {
			ID    string `json:"id"`
			Items []struct {
				Body struct {
					Format string `json:"format"`
					ID     string `json:"id"`
					Type   string `json:"type"`
				} `json:"body"`
				ID         string `json:"id"`
				Motivation string `json:"motivation"`
				Target     string `json:"target"`
				Type       string `json:"type"`
			} `json:"items"`
			Type string `json:"type"`
		} `json:"items"`
		Label struct {
			En []string `json:"en"`
			Fr []string `json:"fr"`
			Nl []string `json:"nl"`
		} `json:"label"`
		Thumbnail []struct {
			Format string `json:"format"`
			Height int64  `json:"height"`
			ID     string `json:"id"`
			Type   string `json:"type"`
			Width  int64  `json:"width"`
		} `json:"thumbnail"`
		Type  string `json:"type"`
		Width int64  `json:"width"`
	} `json:"items"`
	Label    interface{} `json:"label"`
	Metadata []struct {
		Label struct {
			En []string `json:"en"`
			Nl []string `json:"nl"`
		} `json:"label"`
		Value struct {
			En []string `json:"en"`
			Nl []string `json:"nl"`
		} `json:"value"`
	} `json:"metadata"`
	Summary interface{} `json:"summary"`
}
