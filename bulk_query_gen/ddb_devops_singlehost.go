package main

import "time"

// DdbDevopsSingleHost produces Ddb-specific queries for the devops single-host case.
type DdbDevopsSingleHost struct {
	DdbDevops
}

func NewDdbDevopsSingleHost(dbConfig DatabaseConfig, start, end time.Time) QueryGenerator {
	underlying := newDdbDevopsCommon(dbConfig, start, end).(*DdbDevops)
	return &DdbDevopsSingleHost{
		DdbDevops: *underlying,
	}
}

func (d *DdbDevopsSingleHost) Dispatch(i, scaleVar int) Query {
	q := NewHTTPQuery() // from pool
	d.MaxCPUUsageHourByMinuteOneHost(q, scaleVar)
	return q
}
