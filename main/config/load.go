package config

import (
	"buckmate/main/common/exception"
	"buckmate/main/common/util"
	"buckmate/structs"
)

func Load(env string) structs.Config {
	commonRawConfig, errRawConfig := util.LoadYaml("buckmate/Config.yaml")
	exception.Handle(structs.Exception{Err: errRawConfig, Message: "Loading common config failed."})

	envRawConfig, errEnvConfig := util.LoadYaml("buckmate/" + env + "/Config.yaml")
	exception.Handle(structs.Exception{Err: errEnvConfig, Message: "Loading environment config failed."})

	commonConfig := structs.Config{}
	envConfig := structs.Config{}

	errCommonYaml := util.YamlToStruct(commonRawConfig, &commonConfig)
	exception.Handle(structs.Exception{Err: errCommonYaml, Message: "Unmarshaling common config failed."})

	errEnvYaml := util.YamlToStruct(envRawConfig, &envConfig)
	exception.Handle(structs.Exception{Err: errEnvYaml, Message: "Unmarshaling environment config failed."})

	errMerge := util.MergeStruct(&commonConfig, envConfig)
	exception.Handle(structs.Exception{Err: errMerge, Message: "Merging configs failed."})

	return commonConfig
}
