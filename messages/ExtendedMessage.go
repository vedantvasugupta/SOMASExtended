package messages

import (
	"MI_256/common"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/message"
	"github.com/google/uuid"
)

type ExtendedMessage struct {
	message.BaseMessage
	TeamID uuid.UUID
}

func (m ExtendedMessage) GetTeamID() uuid.UUID {
	return m.TeamID
}

func (m *ExtendedMessage) InvokeMessageHandler(mi common.IMI_256) {
	mi.ReceiveMessage(m)
}
