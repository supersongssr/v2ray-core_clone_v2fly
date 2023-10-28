//go:build !confonly
// +build !confonly

package stats

import (
	"strings" // song
	"sync/atomic"
)

// Counter is an implementation of stats.Counter.
type Counter struct {
	value int64
	ips   []string
}

// Value implements stats.Counter.
func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

// Set implements stats.Counter.
func (c *Counter) Set(newValue int64) int64 {
	return atomic.SwapInt64(&c.value, newValue)
}

// Add implements stats.Counter.
func (c *Counter) Add(delta int64) int64 {
	return atomic.AddInt64(&c.value, delta)
}

// Add IP
func (c *Counter) AddIP(ip string) {
	for _, v := range c.ips {
		if v == ip {
			return
		}
	}
	c.ips = append(c.ips, ip)
}

// Get IP
func (c *Counter) GetIP() (int64, string) {
	return int64(len(c.ips)), strings.Join(c.ips, ",")
}

// ResetIP
func (c *Counter) ResetIP() {
	c.ips = []string{}
}
