package routine

import (
	"time"

	"github.com/google/uuid"
)

type StateRoutine string

const (
	StateRunning = "running"
	StateStoped  = "stoped"
)

type RoutineRes struct {
	Identifier string       `json:"identifier"`
	Name       string       `json:"name"`
	Desc       string       `json:"desc"`
	State      StateRoutine `json:"state"`
	CreatedAt  string       `json:"created_at"`
	LastRun    string       `json:"last_run"`
	Interval   string       `json:"interval"`
	Hours      string       `json:"hours"`
	Path       string       `json:"path"`
}

type IndicatorsRes struct {
	TotalRequests   Indicator `json:"total_requests"`
	ErrorRequests   Indicator `json:"error_requests"`
	AverageLatency  Indicator `json:"average_latency"`
	ErrorPercentage Indicator `json:"error_percentage"`
	Timestamp       time.Time `json:"timestamp"`
}

type Indicator struct {
	Name  string      `json:"name"`
	Desc  string      `json:"desc"`
	Value interface{} `json:"value"`
}

type StatusLog string

const (
	StatusLogSuccess = "success"
	StatusLogError   = "error"
)

type LogsRes struct {
	Identifier  uuid.UUID `json:"identifier"`
	Latency     string    `json:"latency"`
	Initialized string    `json:"initialized"`
	Terminated  string    `json:"terminated"`
	Status      StatusLog `json:"status"`
	CreatedAt   string    `json:"created_at"`
}
