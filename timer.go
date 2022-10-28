package graphite

import (
	"strconv"
	"time"
)

type Timer struct {
	Key        string
	Meter      Meter
	Resolution time.Duration
	TimeStart  int64
	TimeStop   int64
	Delta      int64
}

// Constructs a new Timer
func NewTimer(prefix string, key string, resolution time.Duration) Timer {
	m := GetOrCreateMeter(prefix)
	return Timer{
		Meter:      m,
		Resolution: resolution,
	}
}

// Constructs a new Timer With Host
func NewTimerWithHost(prefix string, key string, resolution time.Duration, host string, port int) Timer {
	m := GetOrCreateMeterWithHost(prefix, host, port)
	return Timer{
		Meter:      m,
		Resolution: resolution,
	}
}

// Sets the Timer's StartTime property
func (t Timer) Start(key string) Timer {
	t.Key = key
	t.TimeStart = t.Now()
	return t
}

// Sets the Timer's StopTime and Delta properties
// and sends the Delta to Graphite
func (t Timer) Stop() Timer {
	t.TimeStop = t.Now()
	t.Delta = t.TimeStop - t.TimeStart
	delta := strconv.FormatInt(t.Delta, 10)
	t.Meter.Mark(t.Key, delta)
	return t
}

// Resets the Timer's properties (TimeStart, TimeStop, Delta)
func (t Timer) Reset() Timer {
	t.TimeStart = t.Now()
	t.TimeStop = 0
	t.Delta = 0
	return t
}

// Returns the current time for a timestamp
// Resolution granularity
func (t Timer) Now() int64 {
	now := time.Now()
	var result int64
	switch t.Resolution {
	case time.Second:
		result = now.Unix()
	case time.Millisecond:
		result = now.UnixMilli()
	case time.Microsecond:
		result = now.UnixMicro()
	default:
		result = now.Unix()
	}
	return result
}
