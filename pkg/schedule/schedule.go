package schedule

import (
	"encoding/json"
	"errors"
	"time"

	interval "github.com/go-follow/time-interval"
)

// Schedule controls period-based intervals
type Schedule struct {
	weekdays interval.SpanMany
	weekends interval.SpanMany
	empty    bool
}

func mkTime(h, m int) time.Time {
	return time.Date(1970, 1, 1, h, m, 0, 0, time.UTC)
}

func New(r Representation) (*Schedule, error) {
	if !r.IsValid() {
		return nil, errors.New("invalid schedule settings")
	}

	s := Schedule{
		weekdays: interval.NewMany(),
		weekends: interval.NewMany(),
	}

	for _, i := range r.Intervals {
		t1 := mkTime(i.Begin.Hours, i.Begin.Minutes)
		t2 := mkTime(i.End.Hours, i.End.Minutes+1)
		span, err := interval.New(t1, t2)
		if err != nil {
			return nil, err
		}

		if s.set(i.IsHoliday).IsIntersection(span) {
			return nil, errors.New("schedule interval has intersections")
		}

		if err = s.set(i.IsHoliday).Add(t1, t2); err != nil {
			return nil, err
		}
	}

	s.empty = len(r.Intervals) == 0

	return &s, nil
}

func Parse(schedule string) (*Schedule, error) {
	var r Representation
	if err := json.Unmarshal([]byte(schedule), &r); err != nil {
		return nil, err
	}
	return New(r)
}

func (s *Schedule) set(isHoliday bool) *interval.SpanMany {
	if isHoliday {
		return &s.weekends
	} else {
		return &s.weekdays
	}
}

func (s *Schedule) IsActive(t time.Time) bool {
	normalizedTime := mkTime(t.Hour(), t.Minute())
	isHoliday := t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
	span, err := interval.New(normalizedTime, normalizedTime.Add(time.Millisecond))
	if err != nil {
		panic(err)
	}
	return s.set(isHoliday).IsContains(span)
}

func (s *Schedule) IsActiveNow() bool {
	return s.IsActive(time.Now())
}

func (s *Schedule) Empty() bool {
	return s.empty
}
