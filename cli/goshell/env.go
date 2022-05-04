package goshell

import "os"

func GetEnv(envVarName string, defaultValue string) (retval string) {
	retval = os.ExpandEnv(envVarName)
	if retval == "" {
		return defaultValue
	}
	return
}
