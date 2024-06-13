package main

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/lib"
	"net/http"
	"time"
)

type vegetaSettings struct {
	q        map[string]string
	query    []string
	url      string
	method   string
	metrics  *vegeta.Metrics
	attacker *vegeta.Attacker
	duration time.Duration
	rate     vegeta.ConstantPacer
}

func main() {
	metrics := &vegeta.Metrics{}
	v := vegetaSettings{
		url:      "http://localhost:8080/v1/api/movie/search",
		method:   http.MethodGet,
		duration: 2 * time.Minute,
		rate:     vegeta.Rate{Freq: 1000, Per: time.Second},
		attacker: vegeta.NewAttacker(),
		query:    []string{"harry", "iron", "harr", "what is it", "road"},
		metrics:  metrics,
	}
	Attacker(v)
	metrics.Close()

	MetricsInfo(metrics)
}

func Attacker(v vegetaSettings) {
	targeter := vegeta.NewStaticTargeter(
		vegeta.Target{
			Method: v.method,
			URL:    v.url + "?" + "query=" + v.query[0] + "&" + "cache=" + "true",
		},
		vegeta.Target{
			Method: v.method,
			URL:    v.url + "?" + "query=" + v.query[1],
		},
		vegeta.Target{
			Method: v.method,
			URL:    v.url + "?" + "query=" + v.query[2],
		},
		//vegeta.Target{
		//	Method: v.method,
		//	URL:    v.url + "?" + "query=" + v.query[3],
		//},
		//vegeta.Target{
		//	Method: v.method,
		//	URL:    v.url + "?" + "query=" + v.query[4] + "&" + "cache=" + "true",
		//},
	)

	for res := range v.attacker.Attack(targeter, v.rate, v.duration, "Big Bang!") {
		v.metrics.Add(res)
	}
}

func MetricsInfo(metrics *vegeta.Metrics) {
	fmt.Println("Request count:", metrics.Requests)
	fmt.Println("Time:", metrics.Duration)
	fmt.Println("Success rate:", metrics.Success, "%")
	fmt.Println("Errors:", metrics.Errors)
	fmt.Println("Average latency:", metrics.Latencies.Mean)
	fmt.Println("95th percentile:", metrics.Latencies.P95)
	fmt.Println("99th percentile:", metrics.Latencies.P99)
	fmt.Println("Mean latency:", metrics.Latencies.Mean)
	fmt.Println("Maximum latency:", metrics.Latencies.Max)
	fmt.Println("Throughput:", metrics.Throughput, "bytes/s")
}
