package monthFile

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/takutotsuchie/study/types"
)

var monthFile = "/Users/t.tsuchie/Documents/all-practice/study/study/json/month.json"

func Read() []types.MonthData {
	f, err := os.Open(monthFile)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	var ls []types.MonthData
	err = json.Unmarshal(data, &ls)
	if err != nil {
		log.Fatal(err)
	}
	return ls
}

func Write(ls []types.MonthData) {
	f, err := os.Create(monthFile)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(ls)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(data)
}

// 今月に足す関数。
func AddThisMonth(d time.Duration) {
	ls := Read()
	num := len(ls) - 1
	ls[num].Total += d
	Write(ls)
}

// 月が変わる時
func ChangeMonth() {
	now := time.Now()
	ls := Read()
	num := len(ls)
	mon := ls[num-1].Month
	if int(now.Month()) != mon {
		nextmon := mon + 1
		if nextmon == 13 {
			nextmon = 1
			ls = append(ls, types.MonthData{Total: 0, Month: nextmon})
			Write(ls)
		}
	}
}
