package main

import "time"

type Today struct {
	TodayTotal time.Duration `json:"todayTotal"`
	StartTime  time.Time     `json:"startTime"`
}

type StudyMonth struct {
	Month int           `json:"month"`
	Total time.Duration `json:"total"`
}

type StudyMonthList struct {
	List []StudyMonth `json:"list"`
}
