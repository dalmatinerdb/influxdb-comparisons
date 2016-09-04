package main

import "time"

// DdbDevopsGroupby produces Ddb-specific queries for the devops groupby case.
type DdbDevopsGroupby struct {
	DdbDevops
}

func NewDdbDevopsGroupBy(dbConfig DatabaseConfig, start, end time.Time) QueryGenerator {
	underlying := newDdbDevopsCommon(dbConfig, start, end).(*DdbDevops)
	return &DdbDevopsGroupby{
		DdbDevops: *underlying,
	}

}

func (d *DdbDevopsGroupby) Dispatch(i, scaleVar int) Query {
	q := NewHTTPQuery() // from pool
	d.MeanCPUUsageDayByHourAllHostsGroupbyHost(q, scaleVar)
	return q
}
