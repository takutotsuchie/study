package main

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

func readMonth() StudyMonthList {
	f, _ := os.Open(studyFile)
	data, _ := io.ReadAll(f)
	var ls StudyMonthList
	_ = json.Unmarshal(data, &ls)
	return ls
}

func writeMonth(s StudyMonthList) {
	f, _ := os.Create(studyFile)
	data, _ := json.Marshal(s)
	f.Write(data)
}

// 今月に足す関数。
func addThisMonth(d time.Duration) {
	ls := readMonth()
	num := len(ls.List) - 1
	ls.List[num].Total += d
	writeMonth(ls)
}

// 月が変わる時
func changeMonth() {
	ls := readMonth()
	num := len(ls.List)
	mon := ls.List[num-1].Month
	nextmon := mon + 1
	if nextmon == 13 {
		nextmon = 1
	}
	ls.List = append(ls.List, StudyMonth{Total: 0, Month: nextmon})
	writeMonth(ls)
}
