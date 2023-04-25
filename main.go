package main

import (
	"os"

	"github.com/takutotsuchie/study/command"
	monthFile "github.com/takutotsuchie/study/month-file"
	todayFile "github.com/takutotsuchie/study/today-file"
)

func main() {
	// 引数が設定されていないときはhelpを出力。
	if len(os.Args) == 1 {
		command.Help()
		return
	}
	monthFile.ChangeMonth()
	todayFile.ChangeDay()
	switch os.Args[1] {
	case "start":
		command.Start()
		return
	case "end":
		command.End()
		return
	case "today":
		// 今日の勉強時間を出力
		command.Today()
		return
	case "ls":
		command.Ls()
		return
	default:
		command.Help()
		return
	}
}
