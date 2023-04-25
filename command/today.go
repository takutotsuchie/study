package command

import (
	"fmt"
	"time"

	todayFile "github.com/takutotsuchie/study/today-file"
)

func Today() {
	t := todayFile.Read()
	ans := t.TodayTotal
	if !t.StartTime.IsZero() {
		ans -= time.Until(t.StartTime)
	}
	fmt.Println("今日の勉強時間", ans.Truncate(time.Minute))
}
