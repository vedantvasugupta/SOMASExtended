package baseDiceServer

import (
	baseDiceAgent "SOMASExtended/BaseDiceGame/BaseDiceAgent"
	common "SOMASExtended/BaseDiceGame/common"

	baseServer "github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
    uuid "github.com/google/uuid"
)
//based on methods defined in cw_structure_plan

type IBaseDiceServer interface{
	baseServer.IServer[baseDiceAgent.IBaseDiceAgent]
	createServer(int, int, int, int, int) *IBaseDiceServer 
	formTeams()
	voteforArticlesofAssociation()
	runTurn()
	manageResources()
	generateReport()
	audit()
	modifyRules()
	verifyThreshold()
}


type BaseDiceServer struct{
	*baseServer.BaseServer[baseDiceAgent.IBaseDiceAgent]
	teams map[uuid.UUID]common.Team //map of team IDs to their corresponding Team struct.
	turns int
	teamSize int
	numAgents int
	rounds int
	threshold int
}


// TEAM 2 METHODS BELOW

//TODO:
func (bds *BaseDiceServer) createServer(threshold, rounds, turns, teamSize, numAgents int) *IBaseDiceServer {

}

func (bds *BaseDiceServer) formTeams() {
		agents := bds.GetAgentMap()
		teamSize := bds.teamSize
		numOfAgents := bds.numAgents
		numTeams := numOfAgents/teamSize // calculate num of teams needed
		teamIDList := []uuid.UUID{}

		// STEP 1: Create Teams

		// create [numTeams] Team structs, initialised each with a different TeamID, empty agent slice and empty strategy / commonpool. 
		for i := 0; i < numTeams; i++ {
			//Create a new Team struct
			team := common.NewTeam()

			// fill out the mapping between teamID's and the team struct.
			bds.teams[team.TeamID] = team

			// keep a list of the team ids
			teamIDList = append(teamIDList, team.TeamID)

		}

		// Step 2: Assign each agent a team


		teamIndex := 0 // what teamID we are currently looking at
		agentCount := 0 // counts number of agents on a team

		// iterate over all agents, first adding the agent to their team struct, then populating the agent with their team struct.
		for _, ag := range agents {
			
			// find the teamID of the team we are currently working with
			currentTeamID := teamIDList[teamIndex] 

			// append this agents uuid to the list of agents in their team struct.
			teamAgentList := teams[currentTeamID].Agents
			teamAgentList = append(teamAgentList, ag.GetID())

			//assign agent the team represented by the current team id
			ag.team = bds.teams[currentTeamID]

			//increment num of agents on the team
			agentCount++ 


			// if we have reached the team size, move on to the next team index and reset the counter.
			if agentCount == teamSize {
				teamIndex++
				agentCount = 0
			}
		} 
		

}

//TODO:
func (bds *BaseDiceServer) voteforArticlesofAssociation() {
	
}

func (bds *BaseDiceServer) runTurn() {

	// Step 1: Get each agent to enter the Dice Roll loop and attain a score.
	for _, ag := range bds.GetAgentMap() {
		ag.RollDice(ag)
	}

	// Step 2: Agents make contribution to their team pool, and server redistributes based on team rules.
	bds.manageResources()

	// Step 3: Report Generation (and broadcast?)

	bds.generateReport()

	// Step 4: Run the Audit Process

	bds.audit()

	// Step 5: Run the Rule Mod Process

	bds.modifyRules()

}

func (bds *BaseDiceServer) manageResources() {
	
	// Stage 1: Contribution

	// iterate through agents and call on them to make their contribution to their teams common pool
	for _, ag := range bds.GetAgentMap() {
		agentTeam := bds.teams[ag.team.TeamID]
		agentTeam.CommonPool += ag.MakeContribution()
	}

	// Stage 2: Redistribution

	// iterate through the agents and give them part of their teams common pool, based on their teams strategy.
	for _, ag := range bds.GetAgentMap() {

		// determine this agents share of their teams common pool, given their teams strategy.
		shareOfPool := determineShare(ag.team.CommonPool, ag.team.Strategy,)
		
		// increase their score by what they are given from the pool
		ag.score += shareOfPool
	}

}





// TEAM 5 METHODS

