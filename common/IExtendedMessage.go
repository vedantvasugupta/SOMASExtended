package common

import (
	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/message"
	"github.com/google/uuid"
)

// IExtendedMessage defines the interface for messages in the system
type IExtendedMessage interface {
	message.IMessage[IMI_256]
	GetTeamID() uuid.UUID
}
