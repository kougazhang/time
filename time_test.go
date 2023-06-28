package time

import (
	"fmt"
	"testing"
	"time"
)

func TestIsOnFiveMinutes(t *testing.T) {
	start, _ := time.ParseInLocation("2006-01-02T15:04:05", "2023-06-27T00:00:00", time.Local)
	for i := 0; i < 3600; i++ {
		if IsOnFiveMinutes(start) {
			fmt.Println(start)
		}
		start = start.Add(time.Second)
	}
}

func TestOnFiveMinutes(t *testing.T) {
	start, _ := time.ParseInLocation("2006-01-02T15:04:05", "2023-06-27T00:00:00", time.Local)
	for i := 0; i < 3600; i++ {
		fmt.Println(start, OnFiveMinutes(start))
		start = start.Add(time.Second)
	}
}

func TestOnMinute(t *testing.T) {
	start, _ := time.ParseInLocation("2006-01-02T15:04:05", "2023-06-27T00:00:00", time.Local)
	for i := 0; i < 3600; i++ {
		fmt.Println(start, OnMinute(start))
		start = start.Add(time.Second)
	}
}

func TestUnitDigitOfMinute(t *testing.T) {
	for i := 0; i < 60; i++ {
		fmt.Println(i, unitDigit(i))
	}
}
