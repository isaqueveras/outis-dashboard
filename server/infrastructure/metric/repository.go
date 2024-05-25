package metric

import (
	"time"

	"github.com/google/uuid"
	"github.com/isaqueveras/outis-dashboard/server/database"
	"github.com/isaqueveras/outis-dashboard/server/domain/metric"
)

type repository struct {
	db *database.Tx
}

func New(tx *database.Tx) metric.IMetric {
	return &repository{db: tx}
}

const setWatcherSQL = "UPDATE t_watcher SET run = $1 WHERE id = $2"

func (r *repository) SetWatcher(id uuid.UUID, startedAt time.Time) error {
	_, err := r.db.Execute(setWatcherSQL, startedAt, id)
	return err
}

const setRoutineSQL = "UPDATE t_routine SET run = $1, path = $2 WHERE id = $3"

func (r *repository) SetRoutine(in metric.Routine) error {
	_, err := r.db.Execute(setRoutineSQL, in.RunAt, in.Path, in.Id)
	return err
}

const insertExecutionMetadataSQL = `INSERT INTO t_execution (id, id_routine, latency, metadata, "start", "end") VALUES ($1, $2, $3, $4, $5, $6)`

func (r *repository) InsertMetadata(in metric.Metric) error {
	_, err := r.db.Execute(insertExecutionMetadataSQL, in.Id, in.RoutineID, in.Latency, in.Metadata, in.StartedAt, in.FinishedAt)
	return err
}

const setupWatcherSQL = "INSERT INTO public.t_watcher (id, name, run, created_at) VALUES ($1, $2, $3, $4)"

func (r *repository) SetupWatcher(in metric.Watcher) error {
	_, err := r.db.Execute(setupWatcherSQL, in.Id, in.Name, in.RunAt, in.RunAt)
	return err
}

const setupRoutineSQL = `INSERT INTO public.t_routine (id, id_watcher, "name", "desc", path, "interval", load_interval, "start", "end", run, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $10)`

func (r *repository) SetupRoutine(in metric.Routine) error {
	_, err := r.db.Execute(setupRoutineSQL, in.Id, in.WatcherID, in.Name, in.Desc, in.Path, in.Interval, in.LoadInterval, in.Start, in.End, in.RunAt)
	return err
}

const watcherExistsSQL = "SELECT COUNT(1) > 0 FROM public.t_watcher WHERE id = $1::uuid"

func (r *repository) WatcherExists(id uuid.UUID) (exists bool, err error) {
	if err = r.db.QueryRow(watcherExistsSQL, id).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

const routineExistsSQL = "SELECT COUNT(1) > 0 FROM public.t_routine WHERE id = $1::uuid"

func (r *repository) RoutineExists(id uuid.UUID) (exists bool, err error) {
	if err = r.db.QueryRow(routineExistsSQL, id).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

const updateWatcherSQL = "UPDATE public.t_watcher SET name = $2, run = $3 WHERE id = $1::uuid"

func (r *repository) UpdateWatcher(in metric.Watcher) error {
	_, err := r.db.Execute(updateWatcherSQL, in.Id, in.Name, in.RunAt)
	return err
}

const updateRoutineSQL = `UPDATE public.t_routine SET id_watcher = $2::uuid, "name" = $3, "desc" = $4, path = $5,
"interval" = $6, load_interval = $7, "start" = $8, "end" = $9, run = $10 WHERE id = $1::uuid`

func (r *repository) UpdateRoutine(in metric.Routine) error {
	_, err := r.db.Execute(updateRoutineSQL, in.Id, in.WatcherID, in.Name, in.Desc, in.Path, in.Interval, in.LoadInterval, in.Start, in.End, in.RunAt)
	return err
}
