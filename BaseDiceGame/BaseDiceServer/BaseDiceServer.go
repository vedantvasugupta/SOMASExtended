package baseDiceServer

import (
	common "SOMASExtended/BaseDiceGame/common"

	baseServer "github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
	uuid "github.com/google/uuid"
)

type IBaseDiceServer interface{
	baseServer.IServer[IBaseDiceServer]
	
}


type BaseDiceServer struct{
	*baseServer.BaseServer[IBaseDiceServer]
	team common.Team
	

}




