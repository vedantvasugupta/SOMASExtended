package baseDiceAgent

import (
	common "SOMASExtended/BaseDiceGame/common"
	rand "math/rand"

	baseAgent "github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
	uuid "github.com/google/uuid"
)

type Report struct {
	AgentID     uuid.UUID
	RollHistory []int
	CommonPool  int
}

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
	BroadcastReport(int) []Report
	ProposeAudit() bool
	VoteForAudit() uuid.UUID
	ProposeAoAChange() bool
	VoteForNewAoA() int
	DoIStick(int, int) bool
}

func (agent *BaseDiceAgent) RollDice(specificAgent IBaseDiceAgent) {
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
//func (agent *BaseDiceAgent) MakeContribution() int{
//agent.scores[1] just a check
//agent.team.strategy ----------bug that needs to be solved
//return 0

//}

func (agent *BaseDiceAgent) MakeContribution() int {

	// Get the agent's current score
	currentScore := agent.score

	// Get the proposed contribution from the team's strategy
	proposedContribution := agent.team.GetStrategy()

	// Validate and adjust the contribution if needed
	validContribution := proposedContribution
	if proposedContribution > currentScore {
		validContribution = currentScore // Cap the contribution at the current score
	} else if proposedContribution < 0 {
		validContribution = 0 // Prevent negative contributions
	}

	// Add the validated contribution to the team's pool
	agent.team.AddToPool(validContribution) //AddToPool is not defined in the team interface

	return validContribution
}

func (agent *BaseDiceAgent) BroadcastReport(commonPool int) []Report {

	if agent.memory == nil {
		agent.memory = make(map[uuid.UUID][]int)
	}

	// Create a single report containing all of this agent's rolls
	report := Report{
		AgentID:     agent.GetID(),
		RollHistory: agent.memory[agent.GetID()], // All rolls from memory
		CommonPool:  commonPool,
	}

	return []Report{report}
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
