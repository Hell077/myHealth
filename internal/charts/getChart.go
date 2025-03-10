package charts

import (
	"bytes"
	"errors"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

func DrawChart[T int | uint16](chartType string, values []float64, timeValue T) (*bytes.Buffer, error) {
	var data interface{}

	switch chartType {
	case "daily":
		day, ok := any(timeValue).(int)
		if !ok {
			return nil, errors.New("invalid type for daily chart (expected int)")
		}
		dc := DailyChartPool.Get().(*DailyChart)
		dc.value = values
		dc.hours = len(values)
		dc.day = day
		data = dc

	case "monthly":
		year, ok := any(timeValue).(uint16)
		if !ok {
			return nil, errors.New("invalid type for monthly chart (expected uint16)")
		}
		mc := MonthChartPool.Get().(*MonthChart)
		mc.value = values
		mc.days = len(values)
		mc.year = year
		data = mc

	default:
		return nil, fmt.Errorf("unknown chart type: %s", chartType)
	}

	// Создание графика
	p := plot.New()
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Value"

	var points plotter.XYs
	switch v := data.(type) {
	case *DailyChart:
		p.Title.Text = fmt.Sprintf("Daily Chart (Day %d)", v.day)
		points = make(plotter.XYs, len(v.value))
		for i := range v.value {
			points[i].X = float64(i)
			points[i].Y = v.value[i]
		}

	case *MonthChart:
		p.Title.Text = fmt.Sprintf("Monthly Chart (%d)", v.year)
		points = make(plotter.XYs, len(v.value))
		for i := range v.value {
			points[i].X = float64(i + 1)
			points[i].Y = v.value[i]
		}
	}

	line, err := plotter.NewLine(points)
	if err != nil {
		return nil, err
	}
	line.Color = color.RGBA{R: 0, G: 0, B: 0, A: 100}
	p.Add(line)

	buf := BufferPool.Get().(*bytes.Buffer)
	buf.Reset()

	w, err := p.WriterTo(6*vg.Inch, 4*vg.Inch, "png")
	if err != nil {
		return nil, err
	}
	if _, err := w.WriteTo(buf); err != nil {
		return nil, err
	}

	switch chartType {
	case "daily":
		DailyChartPool.Put(data)
	case "monthly":
		MonthChartPool.Put(data)
	}

	return buf, nil
}
