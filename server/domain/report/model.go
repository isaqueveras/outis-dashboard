package report

import "github.com/lib/pq"

type ChartHistogram struct {
	Labels   pq.StringArray
	Datasets []Dataset
}

type Dataset struct {
	Label string
	Data  pq.Float64Array
	Color string
}
