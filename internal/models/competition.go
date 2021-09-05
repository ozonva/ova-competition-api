package models

import (
	"errors"
	"fmt"
	"time"
)

type CompetitionStatus uint

const (
	Planned CompetitionStatus = iota
	Pending
	Finished
)

type Competition struct {
	Id   uint64 `db:"id"`
	Name string `db:"name"`

	StartTime    time.Time         `db:"start_time"`
	status       CompetitionStatus `db:"status"`
	participants []Participant
}

func (c *Competition) Status() CompetitionStatus {
	return c.status
}

func NewCompetition(id uint64, name string, startTime time.Time) Competition {
	participants := make([]Participant, 0)
	return Competition{
		Id:           id,
		Name:         name,
		StartTime:    startTime,
		status:       Planned,
		participants: participants,
	}
}

func (c *Competition) AddParticipant(participant Participant) error {
	existingParticipantIdx := c.findParticipant(participant.Id)
	if existingParticipantIdx == -1 {
		c.participants = append(c.participants, participant)
		return nil
	} else {
		return errors.New(fmt.Sprintf("participant %d already exists in competition", participant.Id))
	}
}

func (c *Competition) RemoveParticipant(participantId uint64) error {
	participantIdxToDelete := c.findParticipant(participantId)

	if participantIdxToDelete == -1 {
		return errors.New(fmt.Sprintf("could not find participant with id: %d", participantId))
	} else {
		c.participants = removeParticipant(c.participants, participantIdxToDelete)
		return nil
	}
}

func (c *Competition) ChangeStatus(newStatus CompetitionStatus) error {
	if newStatus == Pending && len(c.participants) == 0 {
		return errors.New("could not start competition without participants")
	}

	c.status = newStatus
	return nil
}

func (c *Competition) String() string {
	return fmt.Sprintf("Competition (Id: %d, Name: \"%s\", Start time: %v, Status: %d, participants: %v)",
		c.Id, c.Name, c.StartTime, c.Status(), c.participants)
}

func (c *Competition) findParticipant(participantId uint64) int {
	foundIndex := -1
	for participantIdx, participant := range c.participants {
		if participant.Id == participantId {
			foundIndex = participantIdx
			break
		}
	}
	return foundIndex
}

func removeParticipant(participants []Participant, index int) []Participant {
	capacity := len(participants) - 1
	if capacity < 0 {
		capacity = 0
	}
	res := make([]Participant, 0, capacity)
	res = append(res, participants[:index]...)
	return append(res, participants[index+1:]...)
}
