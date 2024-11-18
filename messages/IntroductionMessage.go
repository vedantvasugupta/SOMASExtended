package messages

import (
	"MI_256/common"
)

type introductionMessage struct {
	ExtendedMessage
	Message string
}

func (tm *introductionMessage) GetMessage() string {
	return tm.Message
}

func (tm *introductionMessage) InvokeMessageHandler(agent common.IMI_256) {
	agent.ReceiveMessage(tm)
}

// override of InvokeHandler (visitor pattern)
// func (cm *CounterMessage) InvokeMessageHandler(agent ICounterAgent) {
// 	agent.HandleCounterMessage(cm)
// }
