# Team4 SOMAS base v1

This is a workable base version Team 4 build for a MVP. 
Note: The project structure, function names, have slight difference to the base we had now. However, we focused on getting everything working while being expandable & maintainable.
Feel free to use any of our code here to speed up with your development. 

### What feature is implemented in this base:
- the game runs with pre-defined iterations & turns
- agents send team invitations at start of turn to each other (randomly)
- server keeps a list of all teams
- agent roll dices and randomly decides to stick or roll again
- agent has a score
- server checks if agent passes the score threshold at the end of each turn
- server kills any agent below threshold
- when the game finishes, log all team information and server information.

### Example Output
if everything works, you should see similar output:
```shell
main function started.
--------Start of iteration 2---------
[server] New round score threshold: 10

Turn 0, 0, current agent count: 3
------------- [server] Starting team formation -------------

159798ac-3dfb-420b-8d5f-b61f796413f4 is starting team formation
Sending invitation: sender=159798ac-3dfb-420b-8d5f-b61f796413f4, teamID=00000000-0000-0000-0000-000000000000, receiver=faf3309d-dc38-4436-bbb3-f4424dce6e4f
faf3309d-dc38-4436-bbb3-f4424dce6e4f is starting team formation
Sending invitation: sender=faf3309d-dc38-4436-bbb3-f4424dce6e4f, teamID=00000000-0000-0000-0000-000000000000, receiver=159798ac-3dfb-420b-8d5f-b61f796413f4

...
```
here, agents are sending invitations to each other. Threshold is defined randomly.


```shell
Agent faf3309d-dc38-4436-bbb3-f4424dce6e4f received team forming invitation from 159798ac-3dfb-420b-8d5f-b61f796413f4
debug_checkID: true
Agent faf3309d-dc38-4436-bbb3-f4424dce6e4f is creating a new team
Agent faf3309d-dc38-4436-bbb3-f4424dce6e4f received team forming invitation from 1b4976b4-79c0-4fa5-a380-a30b8760f2d4
debug_checkID: false
Agent faf3309d-dc38-4436-bbb3-f4424dce6e4f rejected invitation from 1b4976b4-79c0-4fa5-a380-a30b8760f2d4 - already in team 

...
```
here, agents are deciding to join a team or not, and rejecting invalid invitations.


```shell
Turn 0, 1, current agent count: 3
---------------------
faf3309d-dc38-4436-bbb3-f4424dce6e4f decides to ROLL AGAIN, last score: 7
faf3309d-dc38-4436-bbb3-f4424dce6e4f decides to [STICK], last score: 8
[server] Created team 37cce1bc-93af-45fa-aa89-de14b03dab05 with agents [faf3309d-dc38-4436-bbb3-f4424dce6e4f 159798ac-3dfb-420b-8d5f-b61f796413f4]
faf3309d-dc38-4436-bbb3-f4424dce6e4f's turn score: 15, total score: 15
---------------------
newTeamID: 37cce1bc-93af-45fa-aa89-de14b03dab05
Agent faf3309d-dc38-4436-bbb3-f4424dce6e4f created a new team with ID 37cce1bc-93af-45fa-aa89-de14b03dab05
1b4976b4-79c0-4fa5-a380-a30b8760f2d4 decides to [STICK], last score: 6
1b4976b4-79c0-4fa5-a380-a30b8760f2d4's turn score: 6, total score: 6
---------------------
159798ac-3dfb-420b-8d5f-b61f796413f4 decides to ROLL AGAIN, last score: 9
159798ac-3dfb-420b-8d5f-b61f796413f4 decides to ROLL AGAIN, last score: 14
159798ac-3dfb-420b-8d5f-b61f796413f4 **BURSTED!** round: 3, current score: 14
159798ac-3dfb-420b-8d5f-b61f796413f4's turn score: 0, total score: 0

...
```
dice rolling starts at turn 1 (turn 0 in each iteration is used for team forming).
repeat until game ends.


```shell
Agent count: 2
[Agent faf3309d-dc38-4436-bbb3-f4424dce6e4f] score: 26
[Agent 1b4976b4-79c0-4fa5-a380-a30b8760f2d4] score: 30
Agent 159798ac-3dfb-420b-8d5f-b61f796413f4 is dead
Team 00000000-0000-0000-0000-000000000000: [1b4976b4-79c0-4fa5-a380-a30b8760f2d4 faf3309d-dc38-4436-bbb3-f4424dce6e4f]
```
log agents thaat are alive and dead, their score, and teams that are formed.


## Project Structure
```
📦 SOMASExtended
├── 📂 agents
│   ├── ExtendedAgent.go
│   └── SOMAS_Extended_v1.go
├── 📂 common
│   ├── ExposedAgentInfo.go
│   ├── IExtendedMessage.go
│   ├── ISOMAS_Extended.go
│   ├── IServer.go
│   └── team.go
├── 📂 messages
│   ├── ExtendedMessage.go
│   ├── IntroductionMessage.go
│   └── TeamFormingInvitationMessage.go
├── 📂 server
│   └── EnvironmentServer.go
└── SOMAS_Extended.go
```

## Final Note
This platform is far from perfect - it has bugs (you will see if you run) and I fell like there is still so many work to do. 

However, I hope the work here will be useful for you! Let us know (Team 4, or Stanly Chen) if you want to chat about anything.


## Q&A

Why there are so many interfaces?
- because golang hates cyclic dependency. For example, for server to call agent functions and agent to call server functions at the same time, it will cause circular dependency. But - if we put interfaces in a common package, we can avoid this problem.

why is the main called SOMAS_Extended.go?
- because that's the name of agent in our original project (military intelligence agent 256). Rename it as you like!

This project architecture is very different from the current one!
- yes - at first I was planning to build this system for our team, so that we could work on the strategy part first, hence the architecture was customly designed. I did not post our project for replacement of our current architecture, but I just wanted to share our progress, and help speed up your development if you found it useful.
