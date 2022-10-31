package timeService

import (
	"errors"
	"time"
)

type TimeSrv interface {
	Now() time.Time
	IsExpired(time2 time.Time) error
}

type timesrv struct {
}

func (t *timesrv) IsExpired(time2 time.Time) error {
	today := time.Now()
	res := today.Before(time2)
	if res != true {
		return errors.New("expired already")
	} else {
		return nil
	}
}

func (t *timesrv) Now() time.Time {
	return time.Now()
}

func NewTimesrv() TimeSrv {
	return &timesrv{}
}
