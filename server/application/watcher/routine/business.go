package routine

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func Obtain(ctx context.Context) (RoutineRes, error) {
	res := RoutineRes{
		Identifier: uuid.NewString(),
		Name:       "Technology image collector",
		Desc:       "Technology image collector",
		State:      StateRunning,
		CreatedAt:  "02/05/2024",
		LastRun:    "Today 13:29",
		Interval:   "2m (20m)",
		Hours:      "0h - 6h",
		Path:       "/etc/system/usr/example/memory/routines.go:10",
	}

	return res, nil
}

func Indicators(ctx context.Context) (IndicatorsRes, error) {
	res := IndicatorsRes{
		Timestamp: time.Now(),
		TotalRequests: Indicator{
			Name:  "Total requests",
			Desc:  "Overall sum of count",
			Value: 897,
		},
		ErrorRequests: Indicator{
			Name:  "Error requests",
			Desc:  "Overall sum of count",
			Value: 14,
		},
		AverageLatency: Indicator{
			Name:  "Average latency",
			Desc:  "General average of count",
			Value: "0.21ms",
		},
		ErrorPercentage: Indicator{
			Name:  "Error percentage",
			Desc:  "Overall error percentage",
			Value: "12.2%",
		},
	}

	return res, nil
}

func GetLogs(ctx context.Context) ([]LogsRes, error) {
	res := []LogsRes{
		{
			Identifier:  uuid.New(),
			Latency:     "1.23s",
			Initialized: "05/05/2014 14:47:16",
			Terminated:  "05/05/2014 14:47:16",
			Status:      StatusLogSuccess,
			CreatedAt:   "05/05/2014 14:47:16",
		},
	}

	return res, nil
}
