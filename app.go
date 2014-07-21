package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var (
	symbol                                                   *walk.ComboBox
	units, openPrice, stopPrice, closePrice                  *walk.LineEdit
	pipToUSDLabel, profitTargetLabel, stopLossLabel, RRLabel *walk.Label
)

func main() {
	MainWindow{
		Title:   "rock.fx",
		MinSize: Size{355, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{},
				Children: []Widget{
					// Market
					Label{
						Row:    1,
						Column: 0,
						Text:   "Market",
					},
					ComboBox{
						AssignTo: &symbol,
						Row:      1,
						Column:   1,
						Value:    "EUR/USD",
						Model:    []string{"EUR/USD", "USD/JPY", "GBP/USD"},
						OnCurrentIndexChanged: GetResult,
					},

					// Units
					Label{
						Row:    2,
						Column: 0,
						Text:   "Units",
					},
					LineEdit{
						AssignTo:      &units,
						Row:           2,
						Column:        1,
						Text:          "10000",
						OnTextChanged: GetResult,
					},

					// Open
					Label{
						Row:    3,
						Column: 0,
						Text:   "Open",
					},
					LineEdit{
						AssignTo:      &openPrice,
						Row:           3,
						Column:        1,
						OnTextChanged: GetResult,
					},

					// Close
					Label{
						Row:    4,
						Column: 0,
						Text:   "Target",
					},
					LineEdit{
						AssignTo:      &closePrice,
						Row:           4,
						Column:        1,
						OnTextChanged: GetResult,
					},

					// Stop
					Label{
						Row:    5,
						Column: 0,
						Text:   "Stop",
					},
					LineEdit{
						AssignTo:      &stopPrice,
						Row:           5,
						Column:        1,
						OnTextChanged: GetResult,
					},

					// Final info
					Label{
						AssignTo:   &pipToUSDLabel,
						Row:        6,
						Column:     0,
						ColumnSpan: 2,
					},

					Label{
						AssignTo:   &profitTargetLabel,
						Row:        7,
						Column:     0,
						ColumnSpan: 2,
					},

					Label{
						AssignTo:   &stopLossLabel,
						Row:        8,
						Column:     0,
						ColumnSpan: 2,
					},

					Label{
						AssignTo:   &RRLabel,
						Row:        9,
						Column:     0,
						ColumnSpan: 2,
					},
				},
			},
		},
	}.Run()
}

func WriteLabels() {
	// Pip to USD
	pipToUSDLabel.SetText("1 PIP: " + f2str(pipValue) + " USD")

	// Stop loss
	sl := pipValue * (stopValue - openValue) * unitValue
	stopLossLabel.SetText("Loss: " + f2str(sl) + " USD")

	// Profit target
	pt := pipValue * (closeValue - openValue) * unitValue
	profitTargetLabel.SetText("Profit: " + f2str(pt) + " USD")

	// Risk to reward
	rr := -(pt / sl)
	RRLabel.SetText("R/R: " + f2str(rr) + ":1")
}
