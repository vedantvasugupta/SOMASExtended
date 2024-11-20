package Common

type IArticlesOfAssociation interface {
    GetContributionRule() float64
    SetContributionRule(contributionRule ContributionRule)

    GetAuditCost() float64
    SetAuditCost(auditCost AuditCost)

    GetAuditFailureStrategy() AuditFailureStrategy
    SetAuditFailureStrategy(auditFailureStrategy AuditFailureStrategy)
}

// Enum for various contribution rules
type ContributionRule int
const (
    None ContributionRule = iota
    ThirtyPercentCurrentScore
    SixtyPercentCurrentScore
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

// TODO: Add withdrawal rules

type ArticlesOfAssociation struct {
    contributionRule ContributionRule
    auditCost AuditCost
    auditFailureStrategy AuditFailureStrategy
}

func (aoa *ArticlesOfAssociation) GetContributionRule() float64 {
    switch aoa.contributionRule {
    case ThirtyPercentCurrentScore:
        return 0.3
    case SixtyPercentCurrentScore:
        return 0.6
    case All:
        return 2.0
    default:
        return 1.0
    }
}

func (aoa *ArticlesOfAssociation) SetContributionRule(contributionRule ContributionRule) {
    aoa.contributionRule = contributionRule
}

func (aoa *ArticlesOfAssociation) GetAuditCost() float64 {
    switch aoa.auditCost {
    case TenPercentCurrentScore:
        return 0.1
    case TwentyPercentCurrentScore:
        return 0.2
    default:
        return 0.0
    }
}

func (aoa *ArticlesOfAssociation) SetAuditCost(auditCost AuditCost) {
    aoa.auditCost = auditCost
}

func (aoa *ArticlesOfAssociation) GetAuditFailureStrategy() AuditFailureStrategy {
    return aoa.auditFailureStrategy
}

func (aoa *ArticlesOfAssociation) SetAuditFailureStrategy(auditFailureStrategy AuditFailureStrategy) {
    aoa.auditFailureStrategy = auditFailureStrategy
}
