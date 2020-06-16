package util

import "errors"

//Environment defines the environment in which the app runs
type Environment string

const (
	//Development is used to represent Development Environment
	Development Environment = "dev"
	//Staging is used to represent Staging Environment
	Staging Environment = "stg"
	//Production is used to represent Production Environment
	Production Environment = "prod"
)

//CheckEnvironment checks if the given environment is valid
func CheckEnvironment(env string) error {
	tempEnv := Environment(env)
	switch tempEnv {
	case Development, Staging, Production:
		return nil
	default:
		return errors.New("Invalid environment, should be dev/stg/prod")
	}
}
