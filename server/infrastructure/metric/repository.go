package metric

import (
	"database/sql"

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

const saveWatcherSQL = "INSERT INTO public.t_watcher (id, name) VALUES ($1, $2)"

func (r *repository) SaveWatcher(id uuid.UUID, name string) error {
	_, err := r.db.Execute(saveWatcherSQL, id, name)
	return err
}

const obtainWatcherSQL = "SELECT id, name FROM public.t_watcher WHERE id = $1"

func (r *repository) ObtainWatcher(id uuid.UUID) (*metric.Watcher, error) {
	watcher := new(metric.Watcher)
	if err := r.db.QueryRow(obtainWatcherSQL, id).Scan(&watcher.Id, &watcher.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return watcher, nil
}
