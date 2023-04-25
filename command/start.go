package command

import (
	"fmt"
	"time"

	todayFile "github.com/takutotsuchie/study/today-file"
)

func Start() {
	t := todayFile.Read()
	if !t.StartTime.IsZero() {
		fmt.Println("あなたはすでにstartを記録しています")
		return
	}
	fmt.Println("記録を開始します")
	t.StartTime = time.Now()
	todayFile.Write(t)
}
