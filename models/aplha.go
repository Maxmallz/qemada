package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// AlphaMonthly defines the main struct.
type AlphaMonthly struct {
	MetaData struct {
		Information   string    `json:"1. Information"`
		Symbol        string    `json:"2. Symbol"`
		LastRefreshed time.Time `json:"3. Last Refreshed"`
		TimeZone      string    `json:"4. Time Zone"`
	} `json:"Meta Data"`
	MonthlyTimeSeries map[string]struct {
		Open   float64 `json:"1. open"`
		High   float64 `json:"2. high"`
		Low    float64 `json:"3. low"`
		Close  float64 `json:"4. close"`
		Volume float64 `json:"5. volume"`
	} `json:"Monthly Time Series"`
}

// UnmarshalJSON custom implementation for AlphaMonthly
func (a *AlphaMonthly) UnmarshalJSON(data []byte) error {
	// Create an alias to avoid recursion
	type Alias AlphaMonthly
	aux := &struct {
		MetaData struct {
			Information   string `json:"1. Information"`
			Symbol        string `json:"2. Symbol"`
			LastRefreshed string `json:"3. Last Refreshed"`
			TimeZone      string `json:"4. Time Zone"`
		} `json:"Meta Data"`
		MonthlyTimeSeries map[string]map[string]string `json:"Monthly Time Series"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", aux.MetaData.LastRefreshed)
	if err != nil {
		return fmt.Errorf("failed to parse LastRefreshed: %v", err)
	}
	a.MetaData.LastRefreshed = t

	a.MonthlyTimeSeries = make(map[string]struct {
		Open   float64 `json:"1. open"`
		High   float64 `json:"2. high"`
		Low    float64 `json:"3. low"`
		Close  float64 `json:"4. close"`
		Volume float64 `json:"5. volume"`
	})

	for date, values := range aux.MonthlyTimeSeries {
		open, _ := strconv.ParseFloat(values["1. open"], 64)
		high, _ := strconv.ParseFloat(values["2. high"], 64)
		low, _ := strconv.ParseFloat(values["3. low"], 64)
		closeValue, _ := strconv.ParseFloat(values["4. close"], 64)
		volume, _ := strconv.ParseFloat(values["5. volume"], 64)

		a.MonthlyTimeSeries[date] = struct {
			Open   float64 `json:"1. open"`
			High   float64 `json:"2. high"`
			Low    float64 `json:"3. low"`
			Close  float64 `json:"4. close"`
			Volume float64 `json:"5. volume"`
		}{
			Open:   open,
			High:   high,
			Low:    low,
			Close:  closeValue,
			Volume: volume,
		}
	}

	return nil
}
