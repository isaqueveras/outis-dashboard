package report

import "github.com/google/uuid"

type IReport interface {
	GetHistograms(uuid.UUID) (*ChartHistogram, error)
}
