package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	ProcessDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "image_process_duration_seconds",
			Help:    "Время обработки задачи ProcessTask в секундах",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"filter"},
	)

	ProcessedImages = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "image_processed_total",
			Help: "Количество обработанных изображений",
		},
		[]string{"filter"},
	)
)

func init() {
	prometheus.MustRegister(ProcessDuration, ProcessedImages)
}
