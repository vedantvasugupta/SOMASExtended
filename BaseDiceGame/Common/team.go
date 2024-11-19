package common

import "github.com/google/uuid"

type Team struct {
	teamID      uuid.UUID
	commonPool  int
	teamMembers []uuid.UUID
	articlesOfAssociation ArticlesOfAssociation
	auditResult map[uuid.UUID]bool // map of agentID -> bool (true if agent is compliant)
}

func CreateTeam() Team {
	return Team{
		teamID:      uuid.New(),    // Generate a unique TeamID
		commonPool:  0,             // Initialize commonPool to 0
		teamMembers: []uuid.UUID{}, // Initialize an empty slice of agent UUIDs
	}
}

type ITeam interface {
	GetTeamID() uuid.UUID
	GetCommonPool() int
	GetTeamMembers() []uuid.UUID
	IncreaseCommonPool(int)
	DecreaseCommonPool(amount int)
	ResetCommonPool()
	AddMember(memberID uuid.UUID)
	RemoveMember(memberID uuid.UUID)
	SetAuditResult(agentID uuid.UUID, result bool)
	GetAuditResult() map[uuid.UUID]bool
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

func (t *Team) SetAuditResult(audit map[uuid.UUID]bool) {
	t.auditResult = audit
}

func (t *Team) GetAuditResult() map[uuid.UUID]bool {
	return t.auditResult
}
