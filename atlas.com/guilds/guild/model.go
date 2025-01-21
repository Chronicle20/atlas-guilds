package guild

import (
	"atlas-guilds/guild/member"
	"atlas-guilds/guild/title"
	"github.com/google/uuid"
)

type Model struct {
	tenantId            uuid.UUID
	id                  uint32
	worldId             byte
	name                string
	notice              string
	points              uint32
	capacity            uint32
	logo                uint16
	logoColor           byte
	logoBackground      uint16
	logoBackgroundColor byte
	leaderId            uint32
	members             []member.Model
	titles              []title.Model
}

func (m Model) Id() uint32 {
	return m.id
}
