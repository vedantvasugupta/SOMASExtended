package messages

import (
	"SOMAS_Extended/common"
)

type introductionMessage struct {
	ExtendedMessage
	Message string
}

func (tm *introductionMessage) GetMessage() string {
	return tm.Message
}

func (tm *introductionMessage) InvokeMessageHandler(agent common.IExtendedAgent) {
	agent.ReceiveMessage(tm)
}

// override of InvokeHandler (visitor pattern)
// func (cm *CounterMessage) InvokeMessageHandler(agent ICounterAgent) {
// 	agent.HandleCounterMessage(cm)
// }
