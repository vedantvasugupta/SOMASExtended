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
