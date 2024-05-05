package elasticsearch

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"math"
	"net"
	"net/http"
	"time"
)

type Elastic struct {
	*elasticsearch.Client
}

func New(addresses []string) (*Elastic, error) {
	cfg := elasticsearch.Config{
		Addresses:     addresses,
		RetryOnStatus: []int{429, 502, 503, 504},
		RetryBackoff: func(attempt int) time.Duration {
			d := time.Duration(math.Exp2(float64(attempt))) * time.Second
			fmt.Printf("Attempt: %d | Sleeping for %s...\n", attempt, d)
			return d
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		},
		EnableMetrics: true,
	}

	client, err := elasticsearch.NewClient(cfg)

	e := &Elastic{
		client,
	}

	return e, err
}
