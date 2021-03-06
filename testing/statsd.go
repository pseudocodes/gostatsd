package statsd

import (
	"strconv"
	"time"
)

// Satisfies the StatsReporter interface to make testing easier.
type MockStatsdClient struct {
	Counts  map[string]string
	Gauges  map[string]string
	Timings map[string]string
}

func (c *MockStatsdClient) Flush() error {
	return nil
}

func (c *MockStatsdClient) Count(bucket string, value, sampleRate float64) {
	valueString := strconv.FormatFloat(value, 'f', -1, 64)
	c.Counts[bucket] = valueString
}

func (c *MockStatsdClient) Gauge(bucket string, value float64) {
	valueString := strconv.FormatFloat(value, 'f', -1, 64)
	c.Gauges[bucket] = valueString
}

func (c *MockStatsdClient) Timing(bucket string, value float64) {
	valueString := strconv.FormatFloat(value, 'f', -1, 64)
	c.Timings[bucket] = valueString
}

func (c *MockStatsdClient) TimingDuration(bucket string, value time.Duration) {
	c.Timing(bucket, float64(value)/float64(time.Millisecond))
}

func (c *MockStatsdClient) CountUnique(bucket, value string) {
}
