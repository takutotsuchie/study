package main

import (
	"encoding/json"
	"io"
	"os"
)

func readToday() Today {
	f, _ := os.Open(todayFile)
	data, _ := io.ReadAll(f)
	var today Today
	_ = json.Unmarshal(data, &today)
	return today
}

func writeToday(t Today) {
	f, _ := os.Create(todayFile)
	data, _ := json.Marshal(t)
	f.Write(data)
}
