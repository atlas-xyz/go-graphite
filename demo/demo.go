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

// the Graphite endpoint connection is automatically resolved
// from the GRAPHITE_HOST and GRAPHITE_PORT environment variables
func DefaultHostFromEnv() {
	meter := graphite.GetOrCreateMeter("prefix")
	log.Printf("Created meter...: %#v", meter)
	timer := graphite.NewTimer("prefix", "metric.name", time.Second)
	log.Printf("Created timer...: %#v", timer)
}

// Using a custom host
func UseCustomHost() {
	meter := graphite.GetOrCreateMeterWithHost("prefix", "localhost", 2003)
	log.Printf("Created meter...: %#v", meter)
	timer := graphite.NewTimerWithHost("prefix", "metric.name", time.Second, "localhost", 2003)
	log.Printf("Created timer...: %#v", timer)
}

func UsingUnderlying() {
	Graphite, _ := graphite.NewGraphite("localhost", 2003)
	log.Printf("Created Graphite...: %#v", Graphite)
}

func main() {
	UsingRawGraphite()
	UsingTimer()
	UsingMark()

	DefaultHostFromEnv()
	UseCustomHost()
	UsingUnderlying()
}
