package metric

import (
	"context"
	"fmt"

	"github.com/isaqueveras/outis-dashboard/server/database"
	"github.com/isaqueveras/outis-dashboard/server/infrastructure/metric"
)

func Save(ctx context.Context, in *Metric) {
	tx, err := database.NewTx(ctx)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer tx.Rollback()

	repo := metric.New(tx)

	watcher, err := repo.ObtainWatcher(in.Watcher.Id)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	if watcher == nil {
		if err := repo.SaveWatcher(in.Watcher.Id, in.Watcher.Name); err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}
