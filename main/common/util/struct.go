package util

import "dario.cat/mergo"

func MergeStruct(dest interface{}, src interface{}) error {
	return mergo.Merge(dest, src, mergo.WithOverride)
}
