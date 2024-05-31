package report

import (
	"github.com/google/uuid"
	"github.com/isaqueveras/outis-dashboard/server/database"
	"github.com/isaqueveras/outis-dashboard/server/domain/report"
)

type repository struct {
	db *database.Tx
}

func New(tx *database.Tx) report.IReport {
	return &repository{db: tx}
}

const getHistogramsSQL = `
SELECT
	T."key",
	ARRAY_AGG(T.ts)::TEXT[], ARRAY_AGG(COALESCE(T.total, 0))::FLOAT8[]
FROM (
	WITH serial AS (
		SELECT DATE_TRUNC('minute', ts) ts
		FROM GENERATE_SERIES((NOW() - INTERVAL '60 minute')::TIMESTAMP, NOW()::TIMESTAMP, '1 minute'::INTERVAL) ts
	) SELECT
		TH."key",
		TO_CHAR(serial.ts, 'HH24:MI') ts,
		SUM(COALESCE(TH.value, 0)) FILTER (WHERE DATE_TRUNC('minute', TH.created_at) = serial.ts) total
	FROM serial, public.t_histogram th
	JOIN public.t_execution TE ON TE.id = TH.id_routine_execution
	WHERE TH.created_at >= (NOW() - INTERVAL '60 minute')::TIMESTAMP AND TE.id_routine = $1
	GROUP BY 1, 2
) AS T
GROUP BY 1`

func (r *repository) GetHistograms(id uuid.UUID) (chart *report.ChartHistogram, err error) {
	chart = new(report.ChartHistogram)

	rows, err := r.db.Query(getHistogramsSQL, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var dataset report.Dataset
		if err = rows.Scan(&dataset.Label, &chart.Labels, &dataset.Data); err != nil {
			return nil, err
		}
		chart.Datasets = append(chart.Datasets, dataset)
	}

	return chart, nil
}
