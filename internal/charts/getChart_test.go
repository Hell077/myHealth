package charts

import (
	"bytes"
	"errors"
	"image/png"
	"testing"
)

var (
	ErrInvalidTypeDaily   = errors.New("invalid type for daily chart (expected int)")
	ErrInvalidTypeMonthly = errors.New("invalid type for monthly chart (expected uint16)")
)

func TestDrawChart(t *testing.T) {
	dailyValues := []float64{1.2, 2.5, 3.7, 4.8}
	monthValues := []float64{10.5, 12.3, 8.6, 15.9}

	// 🟢 Тест 1: Daily Chart
	buf, err := DrawChart("daily", dailyValues, 15)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if buf == nil || buf.Len() == 0 {
		t.Fatal("buffer is empty, expected PNG data")
	}
	_, err = png.Decode(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatal("buffer does not contain valid PNG image")
	}

	// 🟢 Тест 2: Monthly Chart
	buf, err = DrawChart("monthly", monthValues, uint16(2025))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if buf == nil || buf.Len() == 0 {
		t.Fatal("buffer is empty, expected PNG data")
	}
	_, err = png.Decode(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatal("buffer does not contain valid PNG image")
	}

	// 🔴 Тест 3: Ошибка при неправильном типе времени
	_, err = DrawChart("daily", dailyValues, uint16(2025)) // daily ожидает int
	if err == nil {
		t.Fatal("expected an error for invalid type, but got nil")
	}
	if err == ErrInvalidTypeDaily {
		return
	}

	_, err = DrawChart("monthly", monthValues, 15) // monthly ожидает uint16
	if err == nil {
		t.Fatal("expected an error for invalid type, but got nil")
	}
	if err == ErrInvalidTypeMonthly {
		return
	}

	// 🔴 Тест 4: Ошибка при неизвестном типе графика
	_, err = DrawChart("weekly", dailyValues, 15)
	if err == nil {
		t.Fatal("expected an error for unknown chart type, but got nil")
	}
}
