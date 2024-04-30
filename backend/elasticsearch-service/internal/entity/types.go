package entity

import pb "github.com/Entreeka/proto-proxy/go"

// GetResponse contains types with json tags like in elasticsearch struct
type GetResponse struct {
	Index   string    `json:"_index"`
	ID      string    `json:"_id"`
	Version string    `json:"_version"`
	Source  *pb.Movie `json:"_source"` // replace
}

type SearchResponse struct {
	Hits struct {
		Total struct {
			Value int64 `json:"value"`
		} `json:"total"`
		Hits []*struct {
			Source *Movie `json:"_source"` // replace
		} `json:"hits"`
	} `json:"hits"`
}
