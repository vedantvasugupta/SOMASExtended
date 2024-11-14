package baseDiceServer

import (
	baseDiceAgent "SOMASExtended/BaseDiceGame/BaseDiceAgent"
	common "SOMASExtended/BaseDiceGame/common"

	baseServer "github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
//	uuid "github.com/google/uuid"
)
//based on methods defined in cw_structure_plan
// unsure if necessary based on counter agent example
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
	team common.Team
	turns int
	teamSize int
	numAgents int
	rounds int
	threshold int
}


// TEAM 2 METHODS BELOW



// TEAM 5 METHODS

