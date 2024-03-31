package entity

// GetResponse contains types with json tags like in elasticsearch struct
type GetResponse struct {
	Index   string `json:"_index"`
	ID      string `json:"_id"`
	Version string `json:"_version"`
	Source  *Test  `json:"_source"` // replace
}

type SearchResponse struct {
	Hits struct {
		Total struct {
			Value int64 `json:"value"`
		} `json:"total"`
		Hits []*struct {
			Source *Test `json:"_source"` // replace
		} `json:"hits"`
	} `json:"hits"`
}
