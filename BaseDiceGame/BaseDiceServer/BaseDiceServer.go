package BaseDiceServer

import (
	baseDiceAgent "SOMASExtended/BaseDiceGame/BaseDiceAgent"
	common "SOMASExtended/BaseDiceGame/Common"

	baseServer "github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
	uuid "github.com/google/uuid"

	"math/rand"
	"time"
	"fmt"
)

type IBaseDiceServer interface {
	baseServer.IServer[baseDiceAgent.IBaseDiceAgent]
	CreateServer(int, int, int, int, int) *IBaseDiceServer
	FormTeams()
	VoteforArticlesofAssociation()
	RunTurn()
	ManageResources()
	GenerateReport() // Report[]
	Audit(common.Team)
	rollAndContribute()
	RedistributeCommonPool()
}

type BaseDiceServer struct {
	*baseServer.BaseServer[baseDiceAgent.IBaseDiceAgent]
	teams     map[uuid.UUID]common.Team //map of team IDs to their corresponding Team struct.
	turns     int
	teamSize  int
	numAgents int
	rounds    int
	threshold int
}

func (bds *BaseDiceServer) FormTeams() {
	agents := bds.GetAgentMap()
	teamSize := bds.teamSize
	numOfAgents := bds.numAgents
	numTeams := numOfAgents / teamSize
	teamIDList := []uuid.UUID{}

	// Step 1: Create Teams

	// create [numTeams] Team structs, initialised each with a different TeamID, empty agent slice and empty strategy / commonpool.
	for i := 0; i < numTeams; i++ {
		//Create a new Team struct
		team := common.CreateTeam()
		teamId := team.GetTeamID()
		// fill out the mapping between teamID's and the team struct.
		bds.teams[teamId] = team
		// keep a list of the team ids
		teamIDList = append(teamIDList, teamId)
	}

	// Step 2: Assign each agent a team

	teamIndex := 0  // what teamID we are currently looking at
	agentCount := 0 // counts number of agents on a team

	// iterate through all the teams in the server
	for _, ag := range agents {

		for teamIndex < len(teamIDList) {

			// find the teamID of the team we are currently working with
			currentTeamID := teamIDList[teamIndex]

			// append this agents uuid to the list of agents in their team struct.
			teamAgentList := bds.teams[currentTeamID]
			teamAgentList.AddMember(ag.GetID())

			//increment num of agents on the team
			agentCount++

			// if we have reached the team size, move on to the next team index and reset the counter.
			if agentCount == teamSize {
				teamIndex++
				agentCount = 0
			}
		}
	}

}

func (bds *BaseDiceServer) RunTurn() {

	// Step 1: Vote for Articles of Association
	bds.VoteforArticlesofAssociation()

	bds.rollAndContribute()

	// TODO: Confirm if audit before redistribution is the correct order?

	// Step 4: Run the Audit Process
	for _, team := range bds.teams {
		agentMap := bds.GetAgentMap()
		var vote int = 0
		for _, agentId := range team.GetTeamMembers() {
			ag := agentMap[agentId]
			vote = ag.VoteForAudit()
		}
		if vote > 0 {
			bds.audit(team)
		} else {
			// If no agents voted for audit, set the audit result to an empty map
			team.SetAuditResult(make(map[uuid.UUID]bool))
		}

	}

	// Step 5: Redistribute Common Pool
	for _, team := range bds.teams {
		agentMap := bds.GetAgentMap()
		for _, agentId := range team.GetTeamMembers() {
			ag := agentMap[agentId]
			ag.TakeFromCommonPool()
		}
	}

}

/// Iterates though each team in the server and asks each agent to vote for Articles of Association.
/// The team's strategy is then set to the most common AoA number.
/// If there is a tie, the team's strategy is set to a random AoA number from the most common AoAs.
func (bds *BaseDiceServer) VoteforArticlesofAssociation() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, team := range bds.teams {
		agentMap := bds.GetAgentMap()
		voteCounts := make(map[int]int) // Map to count votes for
		maxVotes := 0
		var mostCommonAoAs []int

		for _, agentId := range team.GetTeamMembers() {
			ag := agentMap[agentId]
			AoA := ag.VotePreferredAoA()
			voteCounts[AoA]++

			// Update maxVotes and mostCommonAoAs in the same loop
			count := voteCounts[AoA]
			if count > maxVotes {
				maxVotes = count
				// reset to only have the current AoA in the mostCommonAoAs
				mostCommonAoAs = []int{AoA}
			} else if count == maxVotes {
				mostCommonAoAs = append(mostCommonAoAs, AoA)
			}
		}

		// Randomly select one AoA from the most common AoAs
		selectedAoA := mostCommonAoAs[rng.Intn(len(mostCommonAoAs))]
		team.SetStrategy(selectedAoA)
	}
}

// For each team, collects contributions from all agents and populates the auditing map to see how many deferrals there are
// for this round
func (bds *BaseDiceServer) rollAndContribute() {
	for _, team := range bds.teams {
		var commonPoolIncrease int = 0
		for _, agentId := range team.GetTeamMembers() {
			agent := bds.GetAgentMap()[agentId]
			// Roll
			agent.RollDice(agent)
			// Contribute
			agentContribution := agent.MakeContribution()
			commonPoolIncrease += agentContribution
			// For the auditing process
			team.SetContributionResult(agentId, agent.GetScore(), agentContribution)
		}
		team.IncreaseCommonPool(commonPoolIncrease)
	}
}

// generateReport generates and returns a report for each agent, including team common pool and agent-specific history.
func (bds *BaseDiceServer) GenerateReport() []Report {
	var reports []Report // Slice to store reports for each agent

	for _, team := range bds.teams { // Iterate over each team
		teamCommonPool := team.GetCommonPool() // Retrieve the current team common pool

		for _, agentID := range team.GetTeamMembers() { // Iterate over each agent in the team
			agent := bds.GetAgentMap()[agentID] // Retrieve the agent by ID

			// Call the agent's BroadcastReport method, which returns a report slice
			agentReports := agent.BroadcastReport(teamCommonPool)

			// Append the agent's report(s) to the main reports slice
			reports = append(reports, agentReports...)
		}
	}

	return reports // Return the slice of all reports generated
}

func (bds *BaseDiceServer) eliminateAgentsBelowThreshold() {
	// Create a slice to store the IDs of agents who will be eliminated
	var agentsToEliminate []uuid.UUID

	// Iterate over all agents in the game
	for id, agent := range bds.GetAgentMap() {
		// Retrieve the agent's current score
		agentScore := agent.GetScore()

		// Check if the agent's score is below the elimination threshold
		if agentScore < bds.threshold {
			// Add the agent's ID to the list of agents to eliminate
			agentsToEliminate = append(agentsToEliminate, id)

			// Notify that the agent will be eliminated
			fmt.Printf("Agent %s did not meet the threshold of %d and will be eliminated.\n", id, bds.threshold)
		}
	}

	// Remove each agent who did not meet the threshold from the game
	for _, id := range agentsToEliminate {
		// Remove the agent from their team
		bds.removeAgent(id)
	}
}

func (bds *BaseDiceServer) removeAgent(id uuid.UUID) {
	agent := bds.GetAgentMap()[id]
	teamId := agent.GetTeamId()
	team := bds.teams[teamId]
	team.RemoveMember(id)
	// Implement logic to remove the agent from the server's agent map
	delete(bds.GetAgentMap(), id)
	// Additional cleanup might be necessary
}

func (bds *BaseDiceServer) audit(agent *baseDiceAgent.BaseDiceAgent, team *common.Team) bool {
	// Retrieve the contribution rule function and audit failure strategy
	contributionRuleFunction := team.ArticlesOfAssociation.ContributionRule // Function to calculate required contribution per round
	auditFailureStrategy := team.ArticlesOfAssociation.AuditFailureStrategy // Function to handle actions when audit fails
	auditCost := team.ArticlesOfAssociation.AuditCost                       // Retrieve audit cost from AoA

	// Iterate over each turn record in the agent’s memory
	for _, turnRecord := range agent.memory[agent.GetID()] {
		totalScore := turnRecord.TotalScore           // Total score rolled in this turn
		actualContribution := turnRecord.Contribution // Actual contribution made by the agent in this turn

		// Calculate expected contribution using the team strategy function
		expectedContribution := contributionRuleFunction(totalScore)

		// Check if actual contribution meets expected contribution
		if actualContribution < expectedContribution {
			// Audit fails: contribution does not meet the required amount
			team.DecreaseCommonPool(auditCost) // Deduct audit cost from team’s common pool

			// Call the team’s audit failure strategy
			auditFailureStrategy(agent, team) // Execute team-defined actions on audit failure

			// Return true indicating the agent cheated
			return true
		}
	}

	// Return false if no failures found (agent did not cheat)
	return false
}
