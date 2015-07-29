package cron
import "time"

type TimeRangeSchedule struct {
	originalSchedule Schedule
	startTime        time.Time
	endTime          time.Time
}

func (schedule TimeRangeSchedule) Next(t time.Time) time.Time {
	schedule.adjustTime(t)
	if t.After(schedule.startTime) && t.Before(schedule.endTime) {
		return schedule.originalSchedule.Next(t)
	}else if t.Before(schedule.startTime) {
		return schedule.originalSchedule.Next(schedule.startTime)
	}else {
		return schedule.originalSchedule.Next(schedule.startTime.AddDate(0, 0, 1))
	}
}

func (trs TimeRangeSchedule)adjustTime(t time.Time) {
	yearDay := t.YearDay()
	if trs.startTime.YearDay()!=yearDay {
		trs.startTime = time.Date(t.Year(), t.Month(), t.Day(), trs.startTime.Hour(),
			trs.startTime.Minute(), 0, 0, time.Local)
	}
	if trs.endTime.YearDay()!=yearDay {
		trs.endTime = time.Date(t.Year(), t.Month(), t.Day(), trs.endTime.Hour(),
			trs.endTime.Minute(), 0, 0, time.Local)
	}
}