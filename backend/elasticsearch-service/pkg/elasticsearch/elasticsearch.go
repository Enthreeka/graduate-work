package elasticsearch

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"math"
	"time"
)

type Elastic struct {
	*elasticsearch.Client
}

func New() (*Elastic, error) {
	cfg := elasticsearch.Config{
		Addresses:     []string{"http://localhost:9200"},
		RetryOnStatus: []int{429, 502, 503, 504},
		RetryBackoff: func(attempt int) time.Duration {
			d := time.Duration(math.Exp2(float64(attempt))) * time.Second
			fmt.Printf("Attempt: %d | Sleeping for %s...\n", attempt, d)
			return d
		},
	}

	client, err := elasticsearch.NewClient(cfg)

	e := &Elastic{
		client,
	}

	return e, err
}
