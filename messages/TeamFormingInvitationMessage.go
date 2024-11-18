package messages

import "MI_256/common"

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

func (tm *TeamFormingInvitationMessage) InvokeMessageHandler(agent common.IMI_256) {
	agent.ReceiveMessage(tm)
}
