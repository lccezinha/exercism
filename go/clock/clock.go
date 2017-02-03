package clock

import "fmt"

const testVersion = 4
const minutesInDay = 1440
const minutesInHour = 60

type Clock struct {
	hour, minute, minutes int
}

func New(hour, minute int) Clock {
	minutes := toMinute(hour, minute)
	clock := Clock{minutes: minutes}.prepare()

	return clock
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
	clock.minutes = (clock.minutes + minutes) % 1440
	return clock.prepare()
}

func (clock Clock) prepare() Clock {
	if clock.minutes < 0 {
		clock.minutes += minutesInDay
	}
	clock.minute = clock.minutes % minutesInHour
	clock.hour = clock.minutes / minutesInHour

	return clock
}

func toMinute(hour, minute int) int {
	return ((hour * minutesInHour) + minute) % minutesInDay
}
