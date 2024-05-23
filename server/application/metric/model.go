package metric

import (
	"time"

	"github.com/google/uuid"
)

type Metric struct {
	Watcher Watcher `json:"watcher"`
}

type Watcher struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	StartedAt time.Time `json:"started_at"`
}
