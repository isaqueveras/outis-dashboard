package metric

import (
	"time"

	"github.com/google/uuid"
)

type Metric struct {
	Watcher Watcher
}

type Watcher struct {
	Id        uuid.UUID
	Name      string
	StartedAt time.Time
}
