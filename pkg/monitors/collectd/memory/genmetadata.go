// Code generated by monitor-code-gen. DO NOT EDIT.

package memory

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "collectd/memory"

var groupSet = map[string]bool{}

const (
	memoryBuffered   = "memory.buffered"
	memoryCached     = "memory.cached"
	memoryFree       = "memory.free"
	memorySlabRecl   = "memory.slab_recl"
	memorySlabUnrecl = "memory.slab_unrecl"
	memoryUsed       = "memory.used"
)

var metricSet = map[string]monitors.MetricInfo{
	memoryBuffered:   {Type: datapoint.Gauge},
	memoryCached:     {Type: datapoint.Gauge},
	memoryFree:       {Type: datapoint.Gauge},
	memorySlabRecl:   {Type: datapoint.Gauge},
	memorySlabUnrecl: {Type: datapoint.Gauge},
	memoryUsed:       {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	memoryBuffered:   true,
	memoryCached:     true,
	memoryFree:       true,
	memorySlabRecl:   true,
	memorySlabUnrecl: true,
	memoryUsed:       true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:       "collectd/memory",
	DefaultMetrics:    defaultMetrics,
	Metrics:           metricSet,
	MetricsExhaustive: false,
	Groups:            groupSet,
	GroupMetricsMap:   groupMetricsMap,
	SendAll:           false,
}
