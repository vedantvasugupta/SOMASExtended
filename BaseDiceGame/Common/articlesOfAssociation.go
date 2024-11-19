package Common

import "SOMASExtended/BaseDiceGame/BaseDiceAgent"

type IArticlesOfAssociation interface {
    // Method to calculate the expected contribution for a given agent
    GetExpectedContribution(agent BaseDiceAgent.IBaseDiceAgent) int

    // Method to check if an agent's action complies with the AoA
    IsAgentCompliant(agent BaseDiceAgent.IBaseDiceAgent) bool

    // Method to impose penalty on agent that fails an audit - T: Agent alive, F: Agent dead
    ImposePenalty(agent BaseDiceAgent.IBaseDiceAgent)

    // Method to check how much an agent should take from the common pool
    // GetExpectedWithdrawal(agent BaseDiceAgent.IBaseDiceAgent) int
}

// Enum for various contribution rules
type ContributionRule int
const (
    None ContributionRule = iota
    ThirtyPercentAboveThreshold
    SixtyPercentAboveThreshold
    All
)

// Enum for costs related to an audit
type AuditCost int
const (
    NoCost AuditCost = iota
    TenPercentCurrentScore
    TwentyPercentCurrentScore
)

// Enum for how to handle an agent that fails an audit
type AuditFailureStrategy int
const (
    NoPenalty AuditFailureStrategy = iota
    TwoWarnings
    ImmediateExpulsion
)

type ArticlesOfAssociation struct {
    contributionRule ContributionRule
    auditCost AuditCost
    auditFailureStrategy AuditFailureStrategy
}

// TODO: Store all previous scores of all agents on the server, and use that to see if the agents are behaving
func (aoa *ArticlesOfAssociation) IsAgentCompliant(agent BaseDiceAgent.IBaseDiceAgent) bool {
    // Basic implementation, but the ideas is the same
    agentScore := agent.GetScore()
    agentContribution := agent.MakeContribution()
    switch aoa.contributionRule {
    case ThirtyPercentAboveThreshold:
        return agentContribution >= int(float64(agentScore) * 0.3)
    case SixtyPercentAboveThreshold:
        return agentContribution >= int(float64(agentScore) * 0.6)
    case All:
        return agentContribution == agentScore
    default:
        return true
    }
}

func (aoa *ArticlesOfAssociation) ImposePenalty (agent BaseDiceAgent.IBaseDiceAgent) bool {
    switch aoa.auditFailureStrategy {
    case TwoWarnings:
        if agent.GetAuditViolations() == 1 {
            return false
        } else {
            agent.AddAuditViolation()
            return true
        }
    case ImmediateExpulsion:
        return false
    default:
        return true
    }
}
