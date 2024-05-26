package metric

import (
	"time"

	"github.com/google/uuid"
)

type Metric struct {
	Id         uuid.UUID      `json:"id"`
	StartedAt  time.Time      `json:"started_at"`
	FinishedAt time.Time      `json:"finished_at"`
	Latency    time.Duration  `json:"latency"`
	Watcher    Watcher        `json:"watcher"`
	Routine    Routine        `json:"routine"`
	Metadata   map[string]any `json:"metadata"`
	Indicators []Indicator    `json:"indicators"`
	Histograms []Histogram    `json:"histograms"`
	Logs       []Log          `json:"logs"`
}

type Watcher struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	RunAt time.Time `json:"run_at"`
}

type Routine struct {
	Id           uuid.UUID     `json:"routine_id"`
	WatcherID    uuid.UUID     `json:"-"`
	Name         string        `json:"name"`
	Desc         string        `json:"desc"`
	Path         string        `json:"path"`
	RunAt        time.Time     `json:"run_at"`
	Interval     time.Duration `json:"interval"`
	LoadInterval time.Duration `json:"load_interval"`
	Start        int64         `json:"start"`
	End          int64         `json:"end"`
}

type Indicator struct {
	Key       string    `json:"key"`
	Value     any       `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

type Histogram struct {
	Key    string           `json:"key"`
	Values []HistogramValue `json:"values"`
}

type HistogramValue struct {
	Value     any       `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

type Log struct {
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type SetupIn struct {
	Watcher Watcher `json:"watcher"`
	Routine Routine `json:"routine"`
}
