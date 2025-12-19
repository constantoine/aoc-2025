package main

type Range struct {
	Min, Max int64
}

func (r Range) Contains(x int64) bool {
	return x >= r.Min && x <= r.Max
}

// Intersect expects r to start before rhs.
func (r Range) Intersect(rhs Range) bool {
	return (r.Max <= rhs.Max && r.Max >= rhs.Min) || // r has intersection with rhs
		(r.Min >= rhs.Min && r.Max <= rhs.Max) || // r is subset of rhs
		(r.Min <= rhs.Min && r.Max >= rhs.Min) // r is superset of rhs
}

func (r Range) Size() int64 {
	return r.Max - r.Min + 1
}
