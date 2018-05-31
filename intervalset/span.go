package intervalset

import (
	"fmt"
)

type Span struct {
	Id  string
	min int
	max int
}

func cast(i Interval) *Span {
	x, ok := i.(*Span)
	if !ok {
		panic(fmt.Errorf("interval must be an Span: %v", i))
	}
	return x
}

// zero returns the zero value for Span.
func zero() *Span {
	return &Span{}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s *Span) String() string {
	return fmt.Sprintf("%s [%d, %d)", s.Id, s.min, s.max)
}

func (s *Span) Id_name() string {
	return s.Id
}

func (s *Span) Equal(t *Span) bool {
	return s.min == t.min && s.max == t.max
}

// Intersect returns the intersection of an interval with another
// interval. The function may panic if the other interval is incompatible.

func (s *Span) Intersect(tInt Interval) Interval {
	t := cast(tInt)
	result := &Span{
		t.Id,
		max(s.min, t.min),
		min(s.max, t.max),
	}
	if result.min < result.max {
		return result
	}
	return zero()
}

// Before returns true if the interval is completely before another interval.
func (s *Span) Before(tInt Interval) bool {
	t := cast(tInt)
	return s.max <= t.min
}

// IsZero returns true for the zero value of an interval.
func (s *Span) IsZero() bool {
	return s.min == 0 && s.max == 0
}

// Bisect returns two intervals, one on either lower side of x and one on the
// upper side of x, corresponding to the subtraction of x from the original
// interval. The returned intervals are always within the range of the
// original interval.
func (s *Span) Bisect(tInt Interval) (Interval, Interval) {
	intersection := cast(s.Intersect(tInt))
	if intersection.IsZero() {
		if s.Before(tInt) {
			return s, zero()
		}
		return zero(), s
	}
	maybeZero := func(Id string, min, max int) *Span {
		if min == max {
			return zero()
		}
		return &Span{Id, min, max}
	}
	return maybeZero(intersection.Id, s.min, intersection.min), maybeZero(intersection.Id, intersection.max, s.max)

}

// Adjoin returns the union of two intervals, if the intervals are exactly
// adjacent, or the zero interval if they are not.
func (s *Span) Adjoin(tInt Interval) Interval {
	t := cast(tInt)
	if s.max == t.min {
		return &Span{t.Id, s.min, t.max}
	}
	if t.max == s.min {
		return &Span{t.Id, t.min, s.max}
	}
	return zero()
}

// Encompass returns an interval that covers the exact extents of two
// intervals.
func (s *Span) Encompass(tInt Interval) Interval {
	t := cast(tInt)
	return &Span{t.Id, min(s.min, t.min), max(s.max, t.max)}
}

func (s *Span) ID() string {
	return s.Id
}

func (s *Span) Min() int {
	return s.min
}

func (s *Span) Max() int {
	return s.max
}
