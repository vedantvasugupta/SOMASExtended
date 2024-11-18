package BaseDiceAgent

import (
	common "SOMASExtended/BaseDiceGame/Common"
	rand "math/rand"

	baseAgent "github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
	// uuid "github.com/google/uuid"
)

type BaseDiceAgent struct {
	*baseAgent.BaseAgent[IBaseDiceAgent]
	team             common.Team
	score            int
	prevRole         int
	lastContribution int
	// memory map[uuid.UUID][]int
}

type IBaseDiceAgent interface {
	baseAgent.IAgent[IBaseDiceAgent]

	RollDice(IBaseDiceAgent)
	SetScore(int)
	GetScore() int
	SetTeam(common.Team)
	GetTeam() *common.Team
	SetPrevRoll(int)
	GetPrevRoll() int
	SetCOntribution(int)
	GetContribution() int

	// -------- The following functions are the ones that the specific agent should implement --
	DoIStick(int, int) bool
	MakeContribution() int // Returns the amount of resources that the agent wants to contribute to the common pool.
	TakeFromCommonPool() int
	// GetVoteForAudit and GetPreferredAoA will be checked each turn, and so the value should be updated each turn accordingly. 0/False if no preference.
	VoteForAudit() int     // Returns 1 if the agent votes for an audit, 0 for abstain, -1 for no audit.
	VotePreferredAoA() int // Returns the id of AoA that the agent prefers. 0 if no preference.
	// -----------------------------------------------------------------------------------------
}

// Sample for later
// func CreateDiceAgent (serv baseAgent.IExposedServerFunctions[IBaseDiceAgent]) IBaseDiceAgent{
// 	return &SpecificDiceAgent{ // REPLACE WITH YOUR SPECIFIC AGENT
// 		BaseAgent: baseAgent.CreateBaseAgent(serv),
// 		team:  common.CreateTeam(),
// 		score: 0,
// 	}
// }

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
	agent.SetPrevRoll(total)
	agent.SetScore(agent.score + total)
}

// / Returns the pointer to the Team object that this agent is assigned to
func (agent *BaseDiceAgent) GetTeam() *common.Team {
	return &agent.team
}

func (agent *BaseDiceAgent) SetScore(score int) {
	agent.score = score
}

func (agent *BaseDiceAgent) GetScore() int {
	return agent.score
}

func (agent *BaseDiceAgent) SetPrevRoll(prevRole int) {
	agent.prevRole = prevRole
}

func (agent *BaseDiceAgent) GetPrevRoll() int {
	return agent.prevRole
}

func (agent *BaseDiceAgent) SetLastCOntribution(contribution int) {
	agent.lastContribution = contribution
}

func (agent *BaseDiceAgent) GetLastContribution() int {
	return agent.lastContribution
}
