package todayFile

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/takutotsuchie/study/types"
)

var todayFile = "/Users/t.tsuchie/Documents/all-practice/study/study/json/today.json"

func Read() types.Today {
	f, err := os.Open(todayFile)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	var today types.Today
	err = json.Unmarshal(data, &today)
	if err != nil {
		log.Fatal(err)
	}
	return today
}

func Write(t types.Today) {
	f, err := os.Create(todayFile)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(data)
}

// 基本は呼ばれない
func ChangeDay() {
	t := Read()
	if t.StartTime.IsZero() {
		return
	}
	now := time.Now().Day()
	if t.StartTime.Day() != now {
		t.StartTime = t.StartTime.AddDate(0, 0, 1)
	}
	Write(t)
}
