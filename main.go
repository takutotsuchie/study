package main

import (
	"fmt"
	"os"
	"time"
)

var studyFile = "/Users/t.tsuchie/Documents/all-practice/study/study/json/month.json"
var todayFile = "/Users/t.tsuchie/Documents/all-practice/study/study/json/today.json"

func main() {
	// 引数が設定されていないときはhelpを出力。
	if len(os.Args) == 1 {
		help()
		return
	}
	var zero time.Time
	t := readToday()
	if !t.StartTime.IsZero() && t.StartTime.Month() != time.Now().Month() {
		fmt.Println("月が変わりました もう一度コマンドを入力してください")
		changeMonth()
		t.StartTime = zero
		return
	}
	switch os.Args[1] {
	case "start":
		if !t.StartTime.IsZero() {
			fmt.Println("あなたはすでにstartを記録しています")
			return
		}
		fmt.Println("記録を開始します")
		t.StartTime = time.Now()
		writeToday(t)
		return
	case "end":
		t := readToday()
		// startしてない時
		if t.StartTime.IsZero() {
			fmt.Println("あなたはstartを記録していないため、記録が失敗しました")
			return
		}
		fmt.Println("記録を終了します")
		t.TodayTotal -= time.Until(t.StartTime)
		fmt.Println("today total: ", t.TodayTotal.Truncate(time.Minute))
		addThisMonth(-time.Until(t.StartTime))
		t.StartTime = zero
		writeToday(t)

		return
	case "ls":
		ls := readMonth()
		fmt.Println("歴代の勉強時間")
		for _, l := range ls.List {
			fmt.Println("月", l.Month, "勉強時間", l.Total.Truncate(time.Minute))
		}
		return
	case "today":
		// 今日の勉強時間を出力
		t := readToday()
		ans := t.TodayTotal
		if !t.StartTime.IsZero() {
			ans -= time.Until(t.StartTime)
		}
		fmt.Println("今日の勉強時間", ans.Truncate(time.Minute))
		return
	default:
		help()
		return
	}
}
