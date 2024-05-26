package metric

import (
	"github.com/google/uuid"
)

type IMetric interface {
	SetWatcher(Watcher) error
	SetRoutine(Routine) error

	InsertMetadata(Metric) error

	SetupWatcher(Watcher) error
	SetupRoutine(Routine) error

	UpdateWatcher(Watcher) error
	UpdateRoutine(Routine) error

	WatcherExists(id uuid.UUID) (bool, error)
	RoutineExists(id uuid.UUID) (bool, error)

	SetIndicator(uuid.UUID, Indicator) error
}
