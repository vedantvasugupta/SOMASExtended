package common

import (
	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
	"github.com/google/uuid"
)

type IMI_256 interface {
	agent.IAgent[IMI_256]

	// Getters
	GetTeamID() uuid.UUID
	GetTrueScore() int

	// Setters
	SetTeamID(teamID uuid.UUID)
	StartRollingDice()
	StickOrAgain() bool
	DecideStick()
	DecideRollAgain()

	// team forming
	StartTeamForming(agentInfoList []ExposedAgentInfo)

	// Messaging functions
	// BroadcastMessageInTeam(T any)
	SendPrivateMessage(receiver uuid.UUID, msg IExtendedMessage)
	SendTeamMessage(msg IExtendedMessage)
	SendMessageBroadcast(msg IExtendedMessage)
	ReceiveMessage(msg any)

	// Info
	GetExposedInfo() ExposedAgentInfo
	LogSelfInfo()
}
