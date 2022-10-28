# Go-Graphite

A simple Graphite interface.

[![Go Reference](https://pkg.go.dev/badge/github.com/srslafazan/go-graphite.svg)](https://pkg.go.dev/github.com/srslafazan/go-graphite)

> Note: Supports TCP, only.
> Underlying interfaces support both TCP and UDP.

## TL;DR

```golang
go get "github.com/srslafazan/go-graphite"
```

## Requirements

- Go (1.19)
- Docker

## Usage

```golang
package main

import (
	"log"
	"time"

	"fmt"

	"github.com/srslafazan/go-graphite"
)

func UsingRawGraphite() {
  // Using raw graphite
	Graphite, err := graphite.NewGraphite("localhost", 2003)
	log.Printf("Loaded Graphite connection: %#v", Graphite)
	if err != nil {
		fmt.Println(err)
	}
	Graphite.SimpleSend("stats.graphite_loaded", "1")
}

func UsingTimer() {
  // Creating a timer
	timer := graphite.NewTimer("prefix", "metric.name", time.Second)
	log.Printf("Created timer...: %#v", timer)
	clock := timer.Start("mytimestat")
	log.Printf("Created clock...: %#v", clock)
	time.Sleep(125 * time.Millisecond)
	clock = clock.Stop()
	log.Printf("Stopped clock...: %#v", clock)
}

func UsingMark() {
  // Marking a single Metric
	meter := graphite.GetOrCreateMeter("prefix")
	log.Printf("Created meter...: %#v", meter)
	mark := meter.Mark("metric.name", "1")
	log.Printf("Created mark...: %#v", mark)
	mark = mark.Update("2")
	log.Printf("Updated mark...: %#v", mark)
	mark = mark.Update("3")
	log.Printf("Updated mark...: %#v", mark)
}

func main() {
	UsingRawGraphite()
  UsingTimer()
	UsingMark()
}
```

## Connecting to Graphite

```golang
// the Graphite endpoint connection is automatically resolved
// from the GRAPHITE_HOST and GRAPHITE_PORT environment variables
func DefaultHostFromEnv() {
	meter := graphite.GetOrCreateMeter("prefix")
	timer := graphite.NewTimer("prefix", "metric.name", time.Second)
}

// Using a custom host
func UseCustomHost() {
	meter := graphite.GetOrCreateMeterWithHost("prefix", "localhost", 2003)
	timer := graphite.NewTimerWithHost("prefix", "metric.name", time.Second, "localhost", 2003)
}

func UsingUnderlying() {
  Graphite, err := graphite.NewGraphite("localhost", 2003)
}
```

## Graphite

To run graphite locally (with Docker):

```bash
docker compose up
```

## Demo

```bash
docker compose up -d

cd demo \
  && go run .
```

## Test

```bash
go test
```
