package main

import (
	"github.com/lxn/walk"
	"strconv"
	"strings"
)

var (
	openValue, stopValue, closeValue, pipValue, unitValue, stopLoss, profitTarget float64
)

func GetResult() {
	openValue = str2f(openPrice)
	stopValue = str2f(stopPrice)
	closeValue = str2f(closePrice)
	unitValue = str2f(units)

	if strings.Contains(symbol.Text(), "/USD") {
		pipValue = 1
	} else {
		pipValue = (0.01 / openValue) * 10000
	}

	WriteLabels()
}

func f2str(f float64) string {
	return strconv.FormatFloat(f, 'f', 4, 64)
}

func str2f(t *walk.LineEdit) float64 {
	s, _ := strconv.ParseFloat(t.Text(), 64)
	return s
}
