package metric

import (
	"time"

	"github.com/google/uuid"
)

type Metric struct {
	Id         uuid.UUID
	StartedAt  time.Time
	FinishedAt time.Time
	Latency    time.Duration
	Watcher    Watcher
	Routine    Routine
	RoutineID  uuid.UUID
	Metadata   map[string]any
	Indicators []Indicator
	Histograms []Histogram
	Logs       []Log
}

type Watcher struct {
	Id    uuid.UUID
	Name  string
	RunAt time.Time
}

type Routine struct {
	Id           uuid.UUID
	WatcherID    uuid.UUID
	Name         string
	Desc         string
	Path         string
	RunAt        time.Time
	Interval     time.Duration
	LoadInterval time.Duration
	Start        int64
	End          int64
}

type Indicator struct {
	Key       string
	Value     any
	CreatedAt time.Time
}

type Histogram struct {
	Key    string
	Values []HistogramValue
}

type HistogramValue struct {
	Value     any
	CreatedAt time.Time
}

type Log struct {
	Level     string
	Message   string
	Timestamp time.Time
}
