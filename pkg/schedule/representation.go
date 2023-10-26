package schedule

import "encoding/json"

// Representation represents schedule for recording cameras or other repeatable actions
type Representation struct {
	Intervals []Interval
}

type Interval struct {
	Begin     TimePoint
	End       TimePoint
	IsHoliday bool
}

func (i Interval) IsValid() bool {
	return i.Begin.IsValid() && i.End.IsValid()
}

type TimePoint struct {
	Hours   int
	Minutes int
}

func (tp TimePoint) IsValid() bool {
	return tp.Hours >= 0 && tp.Hours < 24 && tp.Minutes >= 0 && tp.Minutes <= 59
}

func (r Representation) String() string {
	raw, err := json.Marshal(&r)
	if err != nil {
		panic(err)
	}
	return string(raw)
}

func (r Representation) IsValid() bool {
	for _, i := range r.Intervals {
		if !i.IsValid() {
			return false
		}
	}
	return true
}
