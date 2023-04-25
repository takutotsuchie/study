package command

import (
	"fmt"
	"time"

	monthFile "github.com/takutotsuchie/study/month-file"
	todayFile "github.com/takutotsuchie/study/today-file"
)

//Endは2つのファイルに書き込む。
// 処理を分けよう

func End() {
	t := todayFile.Read()
	// startしてない時
	if t.StartTime.IsZero() {
		fmt.Println("あなたはstartを記録していないため、記録が失敗しました")
		return
	}

	fmt.Println("記録を終了します")
	// 指定された時間が過去なので、time.Untilは負の数を返す。
	duration := -time.Until(t.StartTime)
	fmt.Println("増加した時間", duration)
	t.TodayTotal += duration
	fmt.Println("today total: ", t.TodayTotal.Truncate(time.Minute))
	monthFile.AddThisMonth(-time.Until(t.StartTime))
	t.StartTime = time.Time{}
	todayFile.Write(t)
}
