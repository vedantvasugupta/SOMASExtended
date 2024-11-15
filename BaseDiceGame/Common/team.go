package Common

import "github.com/google/uuid"

type Team struct {
	teamID      uuid.UUID
	commonPool  int
	teamMembers []uuid.UUID
	strategy    int
}

func CreateTeam() Team {
	return Team{
		teamID:      uuid.New(),    // Generate a unique TeamID
		commonPool:  0,             // Initialize commonPool to 0
		teamMembers: []uuid.UUID{}, // Initialize an empty slice of agent UUIDs
		strategy:    0,             // Initialize strategy as 0
	}
}

type ITeam interface {
	GetTeamID() uuid.UUID
	GetCommonPool() int
	GetTeamMembers() []uuid.UUID
	GetStrategy() int
	IncreaseCommonPool(int)
	DecreaseCommonPool(amount int)
	ResetCommonPool()
	AddMember(memberID uuid.UUID)
	RemoveMember(memberID uuid.UUID)
	SetStrategy(strategy int)
}

func (t *Team) GetTeamID() uuid.UUID {
	return t.teamID
}

func (t *Team) GetCommonPool() int {
	return t.commonPool
}

func (t *Team) GetTeamMembers() []uuid.UUID {
	return t.teamMembers
}

func (t *Team) GetStrategy() int {
	return t.strategy
}

func (t *Team) IncreaseCommonPool(amount int) {
	t.commonPool += amount
}

func (t *Team) DecreaseCommonPool(amount int) {
	t.commonPool -= amount
}

func (t *Team) ResetCommonPool() {
	t.commonPool = 0
}

func (t *Team) RemoveMember(memberID uuid.UUID) {
	for i, member := range t.teamMembers {
		if member == memberID {
			t.teamMembers = append(t.teamMembers[:i], t.teamMembers[i+1:]...)
			return 
		}
	}
} 


func (t *Team) AddMember(memberID uuid.UUID) {
	// Check if the member is already in the team
	for _, member := range t.teamMembers {
		if member == memberID {
			return 
		}
	}
	t.teamMembers = append(t.teamMembers, memberID)
}

func (t *Team) SetStrategy(strategy int) {
	t.strategy = strategy
}
