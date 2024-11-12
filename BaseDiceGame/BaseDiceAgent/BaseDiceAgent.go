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
	RollDice(IBaseDiceAgent)
	MakeContribution() int
	BroadcastReport(int)
	VoteForAudit() uuid.UUID
	ProposeAoAChange() bool
	VoteForNewAoA() int
	DoIStick(int, int) bool  // Should be implemented by the specific agent
}


func (agent *BaseDiceAgent) RollDice (specificAgent IBaseDiceAgent) {
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
	//agent.-------------------------------------add 
	agent.scores = append(agent.scores, total)
}

///////////////////////////////////////////////////////////////////////////////////
// EXAMPLE IMPLEMENTATION OF A SPECIFIC AGENT WHICH IMPLEMENTS THE BaseDiceAgent //
///////////////////////////////////////////////////////////////////////////////////

type SpecificDiceAgent struct {
	*BaseDiceAgent
}

// All of the functions defined in the Interface IBaseDiceAgent must be implemented in the SpecificDiceAgent

func (agent *SpecificDiceAgent) MakeContribution() int {
	return 0
}

func (agent *SpecificDiceAgent) DoIStick(total int, prev int) bool {
	return total >= 15
}

func (agent *SpecificDiceAgent) BroadcastReport(score int) {
}

func (agent *SpecificDiceAgent) VoteForAudit() uuid.UUID {
	return uuid.New()
}

func (agent *SpecificDiceAgent) ProposeAoAChange() bool {
	return rand.Intn(2) == 1
}

func (agent *SpecificDiceAgent) VoteForNewAoA() int {
	return rand.Intn(2)
}

func main() {

	baseAgent := &BaseDiceAgent{}
	specificAgent := &SpecificDiceAgent{BaseDiceAgent: baseAgent}

	// Call RollDice, passing specificAgent, which implements DoIStick
	specificAgent.RollDice(specificAgent)
}