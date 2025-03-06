package charts

import (
	"bytes"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"iter"
	"sync"
	"time"
)

var CurrentMonth = int8(time.Now().Month())
var CurrentYear = int16(time.Now().Year())

var DailyChartPool = sync.Pool{
	New: func() any {
		return &DailyChart{}
	},
}

var Buffer = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

var MonthChartPool = sync.Pool{
	New: func() any {
		return &MonthChart{}
	},
}

type DailyChart struct {
	value []float64
	hours []time.Time
}

type MonthChart struct {
	value []float64
	days  []int
	year  uint16
}

func newDailyChart(id uuid.UUID, value []float64) (*DailyChart, error) {
	chAny := DailyChartPool.Get()
	ch, ok := chAny.(*DailyChart)
	if !ok {
		return nil, errors.New("failed to get DailyChart from pool")
	}
	ch.hours = getHourlyTimeRange()
	ch.value = value
	return ch, nil
}

func newMonthlyChart(id uuid.UUID, value []float64, month, year uint8) (*MonthChart, error) {
	chAny := MonthChartPool.Get()
	ch, ok := chAny.(*MonthChart)
	if !ok {
		return nil, errors.New("failed to get DailyChart from pool")
	}
	var curMonth []int
	for num := range getVal(daysInMonth(CurrentMonth, CurrentYear)) {
		curMonth = append(curMonth, num)
	}
	ch.days = curMonth
	ch.value = value
	return ch, nil
}

func

func GetMonthChart[T any](id uuid.UUID,db *sql.DB, fn func(*sql.Rows,int64,time.Time)(T,error)) ([]T,error) {
	obj, err := newMonthlyChart(id)
}

func getVal(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}
