package gigasecond

import "time"

const (
	testVersion = 4
	gigaSecond  = 1000000000 * time.Second
)

func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
