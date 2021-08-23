package models

import "fmt"

type Participant struct {
	Id       uint64
	FullName string
	Age      uint8
}

func NewParticipant(id uint64, fullName string, age uint8) Participant {
	return Participant{
		Id:       id,
		FullName: fullName,
		Age:      age,
	}
}

func (p *Participant) String() string {
	return fmt.Sprintf("Participant (Id: %d, Full name: \"%s\", Age: %d)", p.Id, p.FullName, p.Age)
}
