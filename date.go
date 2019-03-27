package utils

import "time"

func MilisToTime(milis int64) time.Time {
	return time.Unix(0, milis*int64(time.Millisecond))
}

func TimeToMilis(timeValue time.Time) int64 {
	return timeValue.UnixNano() / int64(time.Millisecond)
}
