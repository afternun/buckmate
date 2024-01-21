package deployment

import (
	"buckmate/main/common/exception"
	"buckmate/main/common/util"
	"buckmate/structs"
)

func Load() structs.Deployment {
	deploymentRaw, deploymentRawErr := util.LoadYaml("buckmate/Deployment.yaml")
	exception.Handle(structs.Exception{Err: deploymentRawErr, Message: "Loading deployment failed."})

	deployment := structs.Deployment{}

	errDeploymentYaml := util.YamlToStruct(deploymentRaw, &deployment)
	exception.Handle(structs.Exception{Err: errDeploymentYaml, Message: "Unmarshaling deployment failed."})

	return deployment
}
