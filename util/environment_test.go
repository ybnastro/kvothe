package util

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckEnvironment(t *testing.T) {
	type tcase struct {
		sourceEnv     string
		expectedError error
	}
	testCases := make(map[string]tcase)

	testCases["invalid-env"] = tcase{
		sourceEnv:     "testing",
		expectedError: errors.New("Invalid environment, should be dev/stg/prod"),
	}

	testCases["valid-dev-env"] = tcase{
		sourceEnv:     "dev",
		expectedError: nil,
	}

	for ktc, vtc := range testCases {
		fmt.Println("doing test on TestCheckEnvironment with test case:", ktc)
		actualErr := CheckEnvironment(vtc.sourceEnv)
		assert.Equal(t, actualErr, vtc.expectedError)
	}
}
