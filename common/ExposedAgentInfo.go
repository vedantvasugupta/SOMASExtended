package common

import (
	"github.com/google/uuid"
)

type ExposedAgentInfo struct {
	AgentUUID   uuid.UUID
	AgentTeamID uuid.UUID
}
