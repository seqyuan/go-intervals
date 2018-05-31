package intervalset

import (
	"fmt"
)

type Span struct {
	Id  string
	Min int
	Max int
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
	return fmt.Sprintf("%s [%d, %d)", s.Id, s.Min, s.Max)
}

func (s *Span) Id_name() string {
	return s.Id
}

func (s *Span) Equal(t *Span) bool {
	return s.Min == t.Min && s.Max == t.Max
}

// Intersect returns the intersection of an interval with another
// interval. The function may panic if the other interval is incompatible.

func (s *Span) Intersect(tInt Interval) Interval {
	t := cast(tInt)
	result := &Span{
		t.Id,
		max(s.Min, t.Min),
		min(s.Max, t.Max),
	}
	if result.Min < result.Max {
		return result
	}
	return zero()
}

// Before returns true if the interval is completely before another interval.
func (s *Span) Before(tInt Interval) bool {
	t := cast(tInt)
	return s.Max <= t.Min
}

// IsZero returns true for the zero value of an interval.
func (s *Span) IsZero() bool {
	return s.Min == 0 && s.Max == 0
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
	maybeZero := func(Id string, Min, Max int) *Span {
		if Min == Max {
			return zero()
		}
		return &Span{Id, Min, Max}
	}
	return maybeZero(intersection.Id, s.Min, intersection.Min), maybeZero(intersection.Id, intersection.Max, s.Max)

}

// Adjoin returns the union of two intervals, if the intervals are exactly
// adjacent, or the zero interval if they are not.
func (s *Span) Adjoin(tInt Interval) Interval {
	t := cast(tInt)
	if s.Max == t.Min {
		return &Span{t.Id, s.Min, t.Max}
	}
	if t.Max == s.Min {
		return &Span{t.Id, t.Min, s.Max}
	}
	return zero()
}

// Encompass returns an interval that covers the exact extents of two
// intervals.
func (s *Span) Encompass(tInt Interval) Interval {
	t := cast(tInt)
	return &Span{t.Id, min(s.Min, t.Min), max(s.Max, t.Max)}
}

func (s *Span) ID() string {
	return s.Id
}

func (s *Span) MIN() int {
	return s.Min
}

func (s *Span) MAX() int {
	return s.Max
}
