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

//func createServer


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

//not complete and pro, just showing basic idea
func (bds *BaseDiceServer) runTurn() {
	for _, ag := range bds.GetAgentMap() {
		// get each agent to roll dice
		// something like...
		ag.RollDice(ag)
	}
}

func (bds *BaseDiceServer) manageResources() {

}





// TEAM 5 METHODS

