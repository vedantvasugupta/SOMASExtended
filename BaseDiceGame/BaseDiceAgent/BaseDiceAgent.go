package baseDiceAgent

import (
	common "SOMASExtended/BaseDiceGame/common"
	rand "math/rand"

	baseAgent "github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
	uuid "github.com/google/uuid"
)


type BaseDiceAgent struct{

	*baseAgent.BaseAgent[IBaseDiceAgent]
	team common.Team
	scores []int  											//check

}

type IBaseDiceAgent interface {
	baseAgent.IAgent[IBaseDiceAgent]
	RollDice()
	MakeContribution() int
	BroadcastReport(int)
	VoteForAudit() uuid.UUID
	ProposeAoAChange() bool
	VoteForNewAoA() int
	DoIStick(int, int) bool

}
func (agent *BaseDiceAgent) DoIStick (int, int) bool{
	return true
}

func (agent *BaseDiceAgent) RollDice() {
	prev := 0
	total := 0
	stick := false
	bust := false

	for !stick && !bust {
		// Roll three dice
		r1, r2, r3 := (rand.Intn(6) + 1), (rand.Intn(6) + 1), (rand.Intn(6) + 1)
		score := r1 + r2 + r3

		if score > prev {
			total += score
			prev = score

			stick = agent.DoIStick(total, prev)
		} else {
			bust = true
			score = 0
		}
	}
	//agent.-------------------------------------add 

}


// func (agent *BaseDiceAgent) MakeContribution() int{
// 	//agent.scores[1] just a check
// 	agent.team.strategy
// 	return 0

// }


