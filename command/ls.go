package command

import (
	"fmt"
	"time"

	monthFile "github.com/takutotsuchie/study/month-file"
)

func Ls() {
	ls := monthFile.Read()
	fmt.Println("歴代の勉強時間")
	for _, l := range ls {
		fmt.Println("月", l.Month, "勉強時間", l.Total.Truncate(time.Minute))
	}
}
