package metric

import (
	"context"
	"fmt"

	"github.com/isaqueveras/outis-dashboard/server/database"
	domain "github.com/isaqueveras/outis-dashboard/server/domain/metric"
	"github.com/isaqueveras/outis-dashboard/server/infrastructure/metric"
)

func Event(ctx context.Context, in *Metric) {
	tx, err := database.NewTx(ctx)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer tx.Rollback()

	repo := metric.New(tx)
	if err = repo.InsertMetadata(domain.Metric{
		Id:         in.Id,
		Latency:    in.Latency,
		Metadata:   in.Metadata,
		StartedAt:  in.StartedAt,
		FinishedAt: in.FinishedAt,
		RoutineID:  in.Routine.Id,
	}); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if err = tx.Commit(); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func Setup(ctx context.Context, in *SetupIn) {
	tx, err := database.NewTx(ctx)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer tx.Rollback()

	repo := metric.New(tx)
	watcherExists, err := repo.WatcherExists(in.Watcher.Id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if watcherExists {
		if err := repo.SetWatcher(domain.Watcher(in.Watcher)); err != nil {
			fmt.Printf("%v\n", err)
			return
		}
	} else {
		if err = repo.SetupWatcher(domain.Watcher(in.Watcher)); err != nil {
			fmt.Printf("%v\n", err)
			return
		}
	}

	routineExists, err := repo.RoutineExists(in.Routine.Id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if routineExists {
		if err = repo.SetRoutine(domain.Routine(in.Routine)); err != nil {
			fmt.Printf("%v\n", err)
			return
		}
	} else {
		in.Routine.WatcherID = in.Watcher.Id
		if err = repo.SetupRoutine(domain.Routine(in.Routine)); err != nil {
			fmt.Printf("%v\n", err)
			return
		}
	}

	if err = tx.Commit(); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
