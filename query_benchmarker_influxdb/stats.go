package main

import "fmt"
import "sort"


// Stat represents one statistical measurement.
type Stat struct {
	Label []byte
	Value float64
}

// Init safely initializes a stat while minimizing heap allocations.
func (s *Stat) Init(label []byte, value float64) {
	s.Label = s.Label[:0] // clear
	s.Label = append(s.Label, label...)
	s.Value = value
}

// StatGroup collects simple streaming statistics.
type StatGroup struct {
	Min  float64
	Max  float64
	Mean float64
	Sum  float64
	Items []float64
	Count int64
}

// Push updates a StatGroup with a new value.
func (s *StatGroup) Push(n float64) {
	if s.Count == 0 {
		s.Min = n
		s.Max = n
		s.Mean = n
		s.Count = 1
		s.Sum = n
		s.Items = append(s.Items, n)
		return
	}

	s.Items = append(s.Items, n)

	sort.Float64s(s.Items)

	if n < s.Min {
		s.Min = n
	}
	if n > s.Max {
		s.Max = n
	}
	s.Sum += n

	// constant-space mean update:
	sum := s.Mean*float64(s.Count) + n
	s.Mean = sum / float64(s.Count+1)

	s.Count++
}
func (s *StatGroup) Percentile(P float64) float64 {
	var Len int
	Len = len(s.Items)
	return s.Items[int(float64(Len) * P)]
}

// String makes a simple description of a StatGroup.
func (s *StatGroup) String() string {
	return fmt.Sprintf("Min: %f/%f, max: %f, mean: %f, count: %d, sum: %f", s.Min, s.Items[0], s.Max, s.Mean, s.Count, s.Sum)
}
