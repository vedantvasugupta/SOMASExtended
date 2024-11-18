package main

import (
	"fmt"
	"time"

	"MI_256/agents"
	envServer "MI_256/server"
)

func main() {
	fmt.Println("main function started.")

	// agent configurations
	agentConfig := agents.AgentConfig{
		InitScore:    0,
		VerboseLevel: 10,
	}

	// parameters: agent num, iterations, turns, max duration, max thread
	// note: the zero turn is used for team forming
	serv := envServer.MakeEnvServer(3, 2, 3, 1000*time.Millisecond, 10, agentConfig)

	//serv.ReportMessagingDiagnostics()
	serv.Start()

	// custom function to see agent result
	serv.LogAgentStatus()
	serv.LogTeamStatus()
}
