package schedule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	r := Representation{
		Intervals: []Interval{
			{
				Begin:     TimePoint{10, 00},
				End:       TimePoint{12, 71},
				IsHoliday: false,
			},
		},
	}
	_, err := New(r)
	assert.Error(t, err)

	r = Representation{
		Intervals: []Interval{
			{
				Begin:     TimePoint{10, 00},
				End:       TimePoint{24, 00},
				IsHoliday: false,
			},
		},
	}
	_, err = New(r)
	assert.Error(t, err)

	r = Representation{
		Intervals: []Interval{
			{
				Begin:     TimePoint{12, 00},
				End:       TimePoint{10, 00},
				IsHoliday: false,
			},
		},
	}
	_, err = New(r)
	assert.Error(t, err)

	r = Representation{
		Intervals: []Interval{
			{
				Begin:     TimePoint{10, 00},
				End:       TimePoint{23, 00},
				IsHoliday: false,
			},
			{
				Begin:     TimePoint{11, 00},
				End:       TimePoint{17, 00},
				IsHoliday: false,
			},
		},
	}
	_, err = New(r)
	assert.Error(t, err)
}

func TestSchedule_IsActive(t *testing.T) {
	r := Representation{
		Intervals: []Interval{
			{
				Begin:     TimePoint{10, 00},
				End:       TimePoint{11, 00},
				IsHoliday: false,
			},
			{
				Begin:     TimePoint{15, 00},
				End:       TimePoint{16, 00},
				IsHoliday: false,
			},
		},
	}
	s, err := New(r)
	assert.NoError(t, err)

	now := mkTime(11, 30)
	assert.False(t, s.IsActive(now))

	now = mkTime(10, 30)
	assert.True(t, s.IsActive(now))

	now = mkTime(10, 59)
	assert.True(t, s.IsActive(now))
}

func TestSchedule_Empty(t *testing.T) {
	r := Representation{}
	s, err := New(r)
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.True(t, s.Empty())

	r = Representation{
		Intervals: []Interval{
			{
				Begin:     TimePoint{10, 00},
				End:       TimePoint{11, 00},
				IsHoliday: true,
			},
		},
	}
	s, err = New(r)
	assert.NoError(t, err)
	assert.False(t, s.Empty())
}
