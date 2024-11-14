package baseDiceServer

import (
	baseDiceAgent "SOMASExtended/BaseDiceGame/BaseDiceAgent"
	common "SOMASExtended/BaseDiceGame/common"

	baseServer "github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
	"math/rand"
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

func (bds *BaseDiceServer) createServer(threshold, rounds, turns, teamSize, numAgents int) *IBaseDiceServer {
	
}

func (bds *BaseDiceServer) formTeams() {
	agents := bds.GetAgentMap()
	teamSize := bds.teamSize
	numOfAgents := bds.numAgents
	numTeams := numOfAgents/teamSize // calculate num of teams needed

	// create [numTeams] Team structs, initialised each with a different TeamID, empty agent slice and empty strategy / commonpool.
	for i := 0; i < numTeams; i++ {
		//Create a new Team struct
		team := common.NewTeam()

		// fill out the mapping between teamID's and the team struct.
		bds.teams[team.TeamID] = team
	}


	// TODO: Iterate through all the agents and populate their team field with the correct Team struct.
	for _, ag := range agents {

	}

}

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
		agentTeam := bds.teams[ag.team.TeamID]
		// TODO: Modify the agents score according to their teams strategy and common pool.
	}

}





// TEAM 5 METHODS

