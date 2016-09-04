package main

import "time"

// DdbDevops8Hosts produces Ddb-specific queries for the devops groupby case.
type DdbDevops8Hosts struct {
	DdbDevops
}

func NewDdbDevops8Hosts(dbConfig DatabaseConfig, start, end time.Time) QueryGenerator {
	underlying := newDdbDevopsCommon(dbConfig, start, end).(*DdbDevops)
	return &DdbDevops8Hosts{
		DdbDevops: *underlying,
	}
}

func (d *DdbDevops8Hosts) Dispatch(_, scaleVar int) Query {
	q := NewHTTPQuery() // from pool
	d.MaxCPUUsageHourByMinuteEightHosts(q, scaleVar)
	return q
}
