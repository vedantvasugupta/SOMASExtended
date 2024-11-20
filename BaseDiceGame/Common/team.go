package Common

import "github.com/google/uuid"

type Team struct {
	teamID      uuid.UUID
	commonPool  int
	teamMembers []uuid.UUID
	articlesOfAssociation *ArticlesOfAssociation
	auditResult map[uuid.UUID]int // map of agentID -> int (number of rounds the agent defers from donation strategy)
}

// TODO: Do we want to make all previous rounds auditable or only the current round?
func CreateTeam() Team {
	aoa := CreateArticlesOfAssociation(None, NoCost, NoPenalty)
	return Team{
		teamID:      uuid.New(),    // Generate a unique TeamID
		commonPool:  0,             // Initialize commonPool to 0
		teamMembers: []uuid.UUID{}, // Initialize an empty slice of agent UUIDs
		articlesOfAssociation: aoa,
		auditResult: map[uuid.UUID]int{}, // Initialize an empty map of agentID -> int
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
	GetArticlesOfAssociation() *ArticlesOfAssociation
	SetArticlesOfAssociation(contributionRule ContributionRule, auditCost AuditCost, auditFailureStrategy AuditFailureStrategy)
	// Hidden from all the other agents until an audit is requested
	// TODO: Combine contribution and withdrawal into one composite score once withdrawal implemented
	SetContributionResult(agentID uuid.UUID, agentScore int, agentContribution int) // Can return a bool if we want to track whether the agent deferred that particular round
	// GetAuditResult() map[uuid.UUID]bool
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

func (t *Team) GetArticlesOfAssociation() *ArticlesOfAssociation {
	return t.articlesOfAssociation
}

func (t *Team) SetArticlesOfAssociation(contributionRule ContributionRule, auditCost AuditCost, auditFailureStrategy AuditFailureStrategy) {
	t.articlesOfAssociation = CreateArticlesOfAssociation(contributionRule, auditCost, auditFailureStrategy)
}

func (t *Team) SetContributionResult(agentID uuid.UUID, agentScore int, agentContribution int) {
	contributionPercentage := t.articlesOfAssociation.GetContributionRule()
	if float64(agentScore) * contributionPercentage > float64(agentContribution) {
		t.auditResult[agentID]++
	}
}

func CreateArticlesOfAssociation(contributionRule ContributionRule, auditCost AuditCost, auditFailureStrategy AuditFailureStrategy) *ArticlesOfAssociation {
    return &ArticlesOfAssociation{
        contributionRule: contributionRule,
        auditCost: auditCost,
        auditFailureStrategy: auditFailureStrategy,
    }
}
