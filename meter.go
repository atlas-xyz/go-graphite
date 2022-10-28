package graphite

type Meter struct {
	Key   string
	Value string
	g     *Graphite
}

// Sends a value to Meter with key
// ex: Mark("mystat")
func (c Meter) Mark(key, value string) Meter {
	c.Key = key
	c.Value = value
	c.g.SimpleSend(key, value)
	return c
}

// Pushes a new value to a previously
// recorded Mark
func (c Meter) Update(value string) Meter {
	c.Value = value
	c.g.SimpleSend(c.Key, c.Value)
	return c
}

// Creates a Meter with Metric prefix
// (ex: "analyticsengine")
// which will resolve to the key "analyticsengine.mystat"
// for Mark("mystat")
func GetOrCreateMeter(prefix string) Meter {
	return GetOrCreateMeterWithHost(prefix, host, int(port))
}

// Creates a Meter with Metric prefix
// (ex: "analyticsengine")
// which will resolve to the key "analyticsengine.mystat"
// for Mark("mystat")
func GetOrCreateMeterWithHost(prefix string, host string, port int) Meter {
	g, _ := NewGraphiteWithMetricPrefix(host, int(port), prefix)
	return Meter{
		g: g,
	}
}
