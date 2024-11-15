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
	MakeContribution() int  // Implemented by the specific agent
	DoIStick(int, int) bool  // Implemented by the specific agent
	// GetVoteForAudit and GetPreferredAoA can be checked multiple times per turn, and should just return False/0  
	GetVoteForAudit() bool  // Implemented by the specific agent. Returns true if the agent votes for an audit.
	GetPreferredAoA() int  // Implemented by the specific agent. Returns the id of AoA that the agent prefers. 0 if no preference.
	RollDice(IBaseDiceAgent)
	GetTeam() *common.Team
	SetScore(int)
	GetScore() int
}

func (agent *BaseDiceAgent) RollDice(specificAgent IBaseDiceAgent) {
	/*
	RollDice is a function that simulates the rolling of three dice.

	The loop runs until the agent decides to stick or goes bust.
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

/// Returns the pointer to the Team object that this agent is assigned to 
func (agent *BaseDiceAgent) GetTeam() *common.Team {
	return &agent.team
}

func (agent *BaseDiceAgent) SetScore(score int) {
	agent.score = score
}

func (agent *BaseDiceAgent) GetScore() int {
	return agent.score
}


