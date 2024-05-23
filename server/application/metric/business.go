package metric

import (
	"context"
	"fmt"

	"github.com/isaqueveras/outis-dashboard/server/database"
)

func Save(ctx context.Context) {
	tx, err := database.NewTx(ctx)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer tx.Rollback()

	var id *string
	if err := tx.QueryRow(`SELECT "name" FROM t_watcher WHERE id = '86f72420-3d80-4237-8cee-9d3006eb4be7'`).Scan(&id); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("id: %v\n", *id)

	if err = tx.Commit(); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}
