package deployment

import (
	"buckmate/main/common/util"
	"buckmate/structs"
	"log"
)

func Load(env string, rootDir string) structs.Deployment {
	commonPath := rootDir + "/Deployment.yaml"
	commonFile, err := util.LoadYaml(commonPath)
	if err != nil {
		log.Fatalln("Failed to load " + commonPath)
	}

	envConfig := structs.Deployment{}

	if len(env) > 0 {
		envPath := rootDir + "/" + env + "/Deployment.yaml"
		envFile, err := util.LoadYaml(envPath)
		if err != nil {
			log.Fatalln("Failed to load " + commonPath)
		}
		err3 := util.YamlToStruct(envFile, &envConfig)
		if err3 != nil {
			log.Fatalln(envPath + "\n" + err3.Error())
		}
	}

	commonConfig := structs.Deployment{}
	commonConfig.ConfigBoundary = "%%%"

	err2 := util.YamlToStruct(commonFile, &commonConfig)
	if err2 != nil {
		log.Fatalln(commonPath + "\n" + err2.Error())
	}

	err4 := util.MergeStruct(&commonConfig, envConfig)
	if err4 != nil {
		log.Fatalln("Merging configs failed")
	}

	return commonConfig
}
