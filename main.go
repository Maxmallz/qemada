package main

import (
	"encoding/json"
	"fmt"
	"os"
	"qemada/models"
	"strconv"
	"strings"
)

const (
	GOOGL_MONTHLY = "/Users/maxmallz/dev/repo/qemada/googl-monthly.json"
)

var qubits []uint = []uint{1, 2, 3, 4, 5, 7, 40} // max qubit cannot exceed the difference in sample_years

type SampleInput struct {
	AverageHigh float64
	AverageLow  float64
}

type SampleRange [2]uint

func main() {
	b, _ := os.ReadFile(GOOGL_MONTHLY)
	var data models.AlphaMonthly
	err := json.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}

	monthlyInfo := make(map[string]SampleInput)

	i := 0
	for date, data := range data.MonthlyTimeSeries {
		multiplier := 1000.0

		realHigh := data.Volume / float64(data.High*multiplier)
		realLow := data.Volume / float64(data.Low*multiplier)

		if i == 0 {
			monthlyInfo[date] = SampleInput{
				AverageHigh: realHigh,
				AverageLow:  realLow,
			}
			i++
			continue
		}

		monthlyInfo[date] = SampleInput{
			AverageHigh: realHigh,
			AverageLow:  realLow,
		}
		i++
	}

	for date, data := range data.MonthlyTimeSeries {
		year, _ := strconv.Atoi((strings.Split(date, "-"))[0])

		for _, qubit := range qubits {
			if year == 2021 {

				fmt.Println(year, date, data.Close, qubit)
			}
		}
	}

	fmt.Println(monthlyInfo)
}
