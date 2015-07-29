package cron

import (
	"testing"
	"time"
)

func TestTimeRangeNext(t *testing.T) {
	tests := []struct {
		time     string
		delay    time.Duration
		expected string
	}{
		//超过当日区间
		{"Mon Jul 9 14:45 2012", 5*time.Minute, "Mon Jul 10 10:05 2012"},
		//还没有到区间
		{"Mon Jul 9 09:10 2012", 5*time.Minute, "Mon Jul 9 10:05 2012"},
		//区间内部
		{"Mon Jul 9 11:10 2012", 5*time.Minute, "Mon Jul 9 11:15 2012"},
	}
	startTime := getTime("Mon Jul 9 10:00 2012")
	endTime := getTime("Mon Jul 9 12:00 2012")
	for _, c := range tests {
		schedule := Every(c.delay)
		timeRangeSchedule := &TimeRangeSchedule{schedule, startTime, endTime}
		actual := timeRangeSchedule.Next(getTime(c.time))
		expected := getTime(c.expected)
		if actual != expected {
			t.Errorf("%s, \"%s\": (expected) %v != %v (actual)", c.time, c.delay, expected, actual)
		}
	}


}
