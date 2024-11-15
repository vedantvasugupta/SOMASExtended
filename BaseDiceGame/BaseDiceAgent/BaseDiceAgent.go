package BaseDiceAgent

import (
	common "SOMASExtended/BaseDiceGame/Common"
	rand "math/rand"

	baseAgent "github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
	uuid "github.com/google/uuid"
)

type BaseDiceAgent struct {
	*baseAgent.BaseAgent[IBaseDiceAgent]
	team   common.Team
	score  int
	memory map[uuid.UUID][]int
}

type IBaseDiceAgent interface {
	baseAgent.IAgent[IBaseDiceAgent]
	RollDice(IBaseDiceAgent)
	MakeContribution() int
	BroadcastReport(int)
	ProposeAudit() bool
	VoteForAudit() uuid.UUID
	ProposeAoAChange() bool
	VoteForNewAoA() int
	DoIStick(int, int) bool
	GetTeam() *common.Team
}

func (agent *BaseDiceAgent) RollDice(specificAgent IBaseDiceAgent) {
	/*
	RollDice is a function that simulates the rolling of three dice.

	The loop runs until the agent decides to stick or goes bust.
	This function MUST BE implemented by the specificAgent object.
	*/
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

			stick = specificAgent.DoIStick(total, prev)
		} else {
			bust = true
			score = 0
		}
	}

}

// If not taking specificAgent as a function parameter (ideal method, as done in RollDice)
// then you need to provide a basic implementation of the function in the BaseDiceAgent struct which should then be overrided by the specific agent
func (agent *BaseDiceAgent) MakeContribution(specificAgent IBaseDiceAgent) int {
	//agent.scores[1] just a check
	agent.team.GetStrategy()
	return 0

}

/// Returns the pointer to the Team object that this agent is assigned to 
func (agent *BaseDiceAgent) GetTeam() *common.Team {
	return &agent.team
}

func (agent *BaseDiceAgent) ProposeAudit() bool {
	return true
}

func (agent *BaseDiceAgent) VoteForAudit() uuid.UUID {
	return agent.GetID()
}

func ProposeAoAChange() bool {
	return true
}

func VoteForNewAoA() int {
	return 0
}
