package charts

import (
	"bytes"
	"sync"
	"time"
)

var CurrentMonth = int(time.Now().Month())
var CurrentYear = uint16(time.Now().Year())

type DailyChart struct {
	value []float64
	hours int
	day   int
}

type MonthChart struct {
	value []float64
	days  int
	year  uint16
}

var DailyChartPool = sync.Pool{
	New: func() any {
		return &DailyChart{}
	},
}

var MonthChartPool = sync.Pool{
	New: func() any {
		return &MonthChart{}
	},
}

var BufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}
