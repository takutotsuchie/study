package types

import "time"

type Today struct {
	TodayTotal time.Duration `json:"todayTotal"`
	StartTime  time.Time     `json:"startTime"`
}

type MonthData struct {
	Month int           `json:"month"`
	Total time.Duration `json:"total"`
}
