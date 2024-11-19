package common

import "SOMASExtended/BaseDiceGame/BaseDiceAgent"

type IArticlesOfAssociation interface {
    // Method to calculate the expected contribution for a given agent
    GetExpectedContribution(agent BaseDiceAgent.IBaseDiceAgent) int

    // Method to check if an agent's action complies with the AoA
    IsAgentCompliant(agent BaseDiceAgent.IBaseDiceAgent) bool

    // Method to check how much an agent should take from the common pool
    // GetExpectedWithdrawal(agent BaseDiceAgent.IBaseDiceAgent) int
}

// Enum for various contribution rules
type ContributionRule int
const (
    None ContributionRule = iota
    TenPercentAboveThreshold
    ThirtyPercentAboveThreshold
    SixtyPercentAboveThreshold
    NinetyPercentAboveThreshold
    All
)

type AuditCost int
const (
    NoCost AuditCost = iota
    TenPercentCurrentScore
    TwentyPercentCurrentScore
)

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

func (aoa *ArticlesOfAssociation) IsAgentCompliant(agent BaseDiceAgent.IBaseDiceAgent) bool {
    agentScore := agent.GetScore()
    agentContribution := agent.MakeContribution()
    switch aoa.contributionRule {
    case TenPercentAboveThreshold:
        return agentContribution >= int(float64(agentScore) * 0.1)
    case ThirtyPercentAboveThreshold:
        return agentContribution >= int(float64(agentScore) * 0.3)
    case SixtyPercentAboveThreshold:
        return agentContribution >= int(float64(agentScore) * 0.6)
    case NinetyPercentAboveThreshold:
        return agentContribution >= int(float64(agentScore) * 0.9)
    case All:
        return agentContribution == agentScore
    default:
        return true
    }
}
