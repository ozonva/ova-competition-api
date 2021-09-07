package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type CompetitionMetrics interface {
	MultiCreated()
	ListCompetitions()
	CreateCompetition()
	DescribeCompetition()
	RemoveCompetition()
	UpdateCompetition()
}

type competitionMetrics struct {
	successMultiCreateCompetitionsCounter prometheus.Counter
	successListCompetitionsCounter        prometheus.Counter
	successCreateCompetitionCounter       prometheus.Counter
	successDescribeCompetitionCounter     prometheus.Counter
	successRemoveCompetitionCounter       prometheus.Counter
	successUpdateCompetitionCounter       prometheus.Counter
}

func NewMetrics(namespace string, subsystem string) CompetitionMetrics {
	return &competitionMetrics{
		successMultiCreateCompetitionsCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "success_multi_create_competitions_request_count",
		}),
		successListCompetitionsCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "success_list_competitions_request_count",
		}),
		successCreateCompetitionCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "success_create_competition_request_count",
		}),
		successDescribeCompetitionCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "success_describe_competition_request_count",
		}),
		successRemoveCompetitionCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "success_remove_competition_request_count",
		}),
		successUpdateCompetitionCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "success_update_competition_request_count",
		}),
	}
}

func (c *competitionMetrics) MultiCreated() {
	c.successMultiCreateCompetitionsCounter.Inc()
}

func (c *competitionMetrics) ListCompetitions() {
	c.successListCompetitionsCounter.Inc()
}

func (c *competitionMetrics) CreateCompetition() {
	c.successCreateCompetitionCounter.Inc()
}

func (c *competitionMetrics) DescribeCompetition() {
	c.successDescribeCompetitionCounter.Inc()
}

func (c *competitionMetrics) RemoveCompetition() {
	c.successRemoveCompetitionCounter.Inc()
}

func (c *competitionMetrics) UpdateCompetition() {
	c.successUpdateCompetitionCounter.Inc()
}
