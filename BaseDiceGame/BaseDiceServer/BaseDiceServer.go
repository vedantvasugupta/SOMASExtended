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

