package main

import "sync"

type MonitorHub struct{}

var monitorhub sync.Map

func NewMonitor() *Counters {
	return &Counters{
		m: make(map[string]int),
	}
}
