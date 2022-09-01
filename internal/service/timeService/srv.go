package timeService

import (
	"time"
)

type TimeSrv interface {
	Now() time.Time
}

type timesrv struct {
}

func NewTimesrv() *timesrv {
	return &timesrv{}
}

func (t *timesrv) Now() time.Time {
	return time.Now()
}
