package BaseDiceServer

import (
	baseDiceAgent "SOMASExtended/BaseDiceGame/BaseDiceAgent"
	common "SOMASExtended/BaseDiceGame/Common"
	baseAoA "SOMASExtended/BaseDiceGame/BaseArticlesOfAssociation"

	baseServer "github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
	uuid "github.com/google/uuid"

	"math/rand"
	"time"
)

// NOTES:
// Need the BaseDiceAgent to have a getter / setter functions for their team, and their score
// once this is implemented, change any instances of ag.team and ag.score, etc etc to the appropriate getter / setter func.

type IBaseDiceServer interface {
	baseServer.IServer[baseDiceAgent.IBaseDiceAgent]
	CreateServer(int, int, int, int, int) *IBaseDiceServer
	FormTeams()
	VoteforArticlesofAssociation()
	RunTurn()
	ManageResources()
	GenerateReport()
	Audit(common.Team)
	VerifyThreshold()
	CollectContributions()
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

func (bds *BaseDiceServer) runTurn() {

	// Step 1: Vote for Articles of Association
	bds.VoteforArticlesofAssociation()

	for _, team := range bds.teams {
		agentMap := bds.GetAgentMap()
		var commonPool int = 0
		for _, agentId := range team.GetTeamMembers() {
			ag := agentMap[agentId]
			// Step 2: Roll Dice
			ag.RollDice(ag)
			// Step 3: Make Contribution
			commonPool += ag.MakeContribution()
		}
		team.IncreaseCommonPool(commonPool)
	}

	// Step 4: Run the Audit Process
	for _, team := range bds.teams {
		agentMap := bds.GetAgentMap()
		var vote int = 0
		for _, agentId := range team.GetTeamMembers() {
			ag := agentMap[agentId]
			vote = ag.VoteForAudit()
		}
		if vote > 0 {
			bds.Audit(team)
		} else{
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

/// Audit verifies that each agent has contributed the expected amount to the common pool.
/// An audit is based on the team's strategy and the information the server has about the agents.
/// Provides the expected contributions
func (bds *BaseDiceServer) Audit(team common.Team) {
	agentsInTeam := team.GetTeamMembers()
	auditMap := make(map[uuid.UUID]bool)
	strategyNum := team.GetStrategy()

	for _, agentId := range agentsInTeam {
		agent := bds.GetAgentMap()[agentId]
		isAgentCompliant := baseAoA.IsAgentCompliant(agent, strategyNum)
		auditMap[agentId] = isAgentCompliant
	}

	team.SetAuditResult(auditMap)
}

/// Iterates though each team in the server and asks each agent to vote for Articles of Association.
/// The team's strategy is then set to the most common AoA number.
/// If there is a tie, the team's strategy is set to a random AoA number from the most common AoAs.
func (bds *BaseDiceServer) VoteforArticlesofAssociation() {
	// Seed the random number generator as rand is pseudo-random
	// If not seeded the random number generator will produce the same output given the same input
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, team := range bds.teams {
		agentMap := bds.GetAgentMap()
		voteCounts := make(map[int]int) // Map to count votes for each AoA number
		for _, agentId := range team.GetTeamMembers() {
			ag := agentMap[agentId]
			vote := ag.VotePreferredAoA()
			voteCounts[vote]++
		}

		// Find the maximum number of votes any AoA received
		maxVotes := 0
		for _, count := range voteCounts {
			if count > maxVotes {
				maxVotes = count
			}
		}

		// Find the AoA number that received the most votes
		mostCommonAoAs := []int{}
		for aoa, count := range voteCounts {
			if count == maxVotes {
				mostCommonAoAs = append(mostCommonAoAs, aoa)
			}
		}

		// randomly select index of the most common AoAs list
		selectedAoA := mostCommonAoAs[0]
		if len(mostCommonAoAs) > 1 {
			selectedAoA = mostCommonAoAs[rng.Intn(len(mostCommonAoAs))]
		}

		// Set the team's strategy to the selected AoA
		team.SetStrategy(selectedAoA)

	}
}

// TODO these function have been made redundant, keep for reference
// /// CollectContributions iterates through all the agents in the server and calls on them to make their contribution to their team's common pool.
// func (bds *BaseDiceServer) CollectContributions() {
// 	// iterate through agents and call on them to make their contribution to their teams common pool
// 	for _, ag := range bds.GetAgentMap() {
// 		agentTeam := bds.teams[ag.GetTeam().GetTeamID()]
// 		agentTeam.IncreaseCommonPool(ag.MakeContribution())
// 	}
// }

// /// RedistributeCommonPool calls the TakeFromCommonPool function for each agent in the server.
// /// Each agent will take from the common pool based on their team's strategy.
// func (bds *BaseDiceServer) RedistributeCommonPool() {
// 	for _, ag := range bds.GetAgentMap() {
// 		// Agents take from the common pool based on their team's strategy
// 		ag.TakeFromCommonPool()
// 	}
// }
