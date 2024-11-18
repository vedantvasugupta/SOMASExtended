package BaseArticlesOfAssociation

import "SOMASExtended/BaseDiceGame/BaseDiceAgent"


type IArticlesOfAssociation interface {
    // Method to calculate the expected contribution for a given agent
    GetExpectedContribution(agent BaseDiceAgent.IBaseDiceAgent, strategy int) int

    // Method to check if an agent's action complies with the AoA
    IsAgentCompliant(agent BaseDiceAgent.IBaseDiceAgent, strategy int) bool

    // Method to check how much an agent should take from the common pool
    GetExpectedWithdrawal(agent BaseDiceAgent.IBaseDiceAgent, strategy int) int
}

func IsAgentCompliant(agent BaseDiceAgent.IBaseDiceAgent, strategy int) bool {
    //TODO Placeholder
    return true
}
