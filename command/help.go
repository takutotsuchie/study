package command

import "fmt"

func Help() {
	fmt.Println("help")
	fmt.Println("		start: 記録を開始")
	fmt.Println("		end: 記録を終了")
	fmt.Println("		ls: 歴代の月の記録を表示")
	fmt.Println("		today: 今日の勉強時間を表示")
}
