package kafka

import (
	"github.com/segmentio/kafka-go"
	"ozonva/ova-competition-api/internal/models"
)

type CompetitionEventType uint64

const (
	CompetitionCreated CompetitionEventType = iota
	CompetitionUpdated
	CompetitionDeleted
)

type competitionEvent struct {
	Type        CompetitionEventType
	Competition *models.Competition
}

func NewCompetitionEvent(eventType CompetitionEventType, competition *models.Competition) Event {
	return &competitionEvent{
		Type:        eventType,
		Competition: competition,
	}
}

func (e *competitionEvent) toMessage() kafka.Message {
	return kafka.Message{
		Key:   []byte{byte(e.Type)},
		Value: []byte(e.Competition.String()),
	}
}
