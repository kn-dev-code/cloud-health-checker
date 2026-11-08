package main

import (
	"fmt"
	"net/http"
	"time"
)


type Endpoint struct {
	URL string
	Status string
	Latency time.Duration
}


func (e *Endpoint) CheckHealth() {
	start := time.Now()

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(e.URL)
	e.Latency = time.Since(start)

	if err != nil {
		e.Status = "DOWN (Error)"
		return 
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		e.Status = "UP"
	} else {
		e.Status = fmt.Sprintf("DOWN (Status: %d)", resp.StatusCode)
	}
}

func main() {
	endpoints := []Endpoint{
		{URL: "https://httpbin.org/status/200"},
		{URL: "https://httpbin.org/delay/1"},
		{URL: "https://this-domain-definitely-does-not-exist-12345.com"},
	}

	fmt.Println("Starting cloud health checks...")

	for i := range endpoints {
		endpoints[i].CheckHealth()
		fmt.Printf("Target: %s | Status: %s | Latency: %v\n", 
			endpoints[i].URL, endpoints[i].Status, endpoints[i].Latency)
	}
}