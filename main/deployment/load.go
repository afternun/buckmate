package deployment

import (
	"buckmate/main/common/util"
	"buckmate/structs"
	"log"
)

func Load(env string) structs.Deployment {
	commonFile, err := util.LoadYaml("buckmate/Deployment.yaml")
	if err != nil {
		log.Fatalln("Loading common config failed")
	}

	envFile, err := util.LoadYaml("buckmate/" + env + "/Deployment.yaml")
	if err != nil {
		log.Fatalln("Loading environment config failed")
	}

	commonConfig := structs.Deployment{}
	envConfig := structs.Deployment{}

	err2 := util.YamlToStruct(commonFile, &commonConfig)
	if err2 != nil {
		log.Fatalln("Unmarshaling common config failed.")
	}

	err3 := util.YamlToStruct(envFile, &envConfig)
	if err3 != nil {
		log.Fatalln("Unmarshaling environment config failed.")
	}

	err4 := util.MergeStruct(&commonConfig, envConfig)
	if err4 != nil {
		log.Fatalln("Merging configs failed")
	}

	return commonConfig
}
