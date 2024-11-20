package messages

import "SOMAS_Extended/common"

type TeamFormingInvitationMessage struct {
	ExtendedMessage
	AgentInfo common.ExposedAgentInfo
	Message   string
}

func (tm *TeamFormingInvitationMessage) GetMessage() string {
	return tm.Message
}

func (tm *TeamFormingInvitationMessage) GetAgentInfo() common.ExposedAgentInfo {
	return tm.AgentInfo
}

func (tm *TeamFormingInvitationMessage) InvokeMessageHandler(agent common.IExtendedAgent) {
	agent.ReceiveMessage(tm)
}
