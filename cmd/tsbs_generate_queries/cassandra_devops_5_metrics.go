package main

import (
	"time"

	"bitbucket.org/440-labs/influxdb-comparisons/query"
)

// CassandraDevops5Metrics contains info for Cassandra-devops test '5-metrics-*'
type CassandraDevops5Metrics struct {
	CassandraDevops
	hosts int
	hours int
}

// NewCassandraDevops5Metrics produces a new function that produces a new CassandraDevops5Metrics
func NewCassandraDevops5Metrics(hosts, hours int) QueryGeneratorMaker {
	return func(start, end time.Time) QueryGenerator {
		underlying := newCassandraDevopsCommon(start, end)
		return &CassandraDevops5Metrics{
			CassandraDevops: *underlying,
			hosts:           hosts,
			hours:           hours,
		}
	}
}

// Dispatch fills in the query.Query
func (d *CassandraDevops5Metrics) Dispatch(scaleVar int) query.Query {
	q := query.NewCassandra() // from pool
	d.MaxCPUMetricsByMinute(q, scaleVar, d.hosts, 5, time.Duration(int64(d.hours)*int64(time.Hour)))
	return q
}
