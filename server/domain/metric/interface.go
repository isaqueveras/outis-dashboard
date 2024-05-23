package metric

import (
	"github.com/google/uuid"
)

type IMetric interface {
	SaveWatcher(id uuid.UUID, name string) error
	ObtainWatcher(id uuid.UUID) (*Watcher, error)
}
