package domain

import "time"

type Rate struct {
	Date     time.Time `json:"date"`
	ValueUSD float64   `json:"valueusd"`
}
