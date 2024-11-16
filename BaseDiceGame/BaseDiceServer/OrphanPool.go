package BaseDiceServer

import (
	"fmt"
	"github.com/google/uuid"
)

/* Declare the orphan pool for keeping track of agents that are not currently
* part of a team. This maps agentID -> slice of teamIDs that agent wants to
* join. Note that the slice of teams is processed in order, so the agent should
* put the team it most wants to join at the start of the slice. */
type orphanPool map[uuid.UUID][]uuid.UUID

/* 
* Print the contents of the pool. Careful as this will not necessarily print
* the elements in the order that you added them. 
*/
func (pool orphanPool) PrintPool() {
    for i, v := range pool {
        // truncate the UUIDs to make it easier to read
        shortAgentId := i.String()[:8]
        shortTeamIds := make([]string, len(v))

        // go over all the teams in the wishlist and add to shortened IDs
        for _, teamID := range v {
            shortTeamIds = append(shortTeamIds, teamID.String()[:8])
        }

        fmt.Println(shortAgentId, " Wants to join : ", shortTeamIds)
    }
}

// test code ---------------------------------------
var TestPool = orphanPool{
    uuid.New(): {uuid.New(), uuid.New(), uuid.New()},
    uuid.New(): {uuid.New(), uuid.New(), uuid.New(), uuid.New(), uuid.New()},
    uuid.New(): {}, // outcast, does not want to join any team. 
    uuid.New(): {uuid.New()}, 
}

func RunTests() {
    TestPool.PrintPool()
}
