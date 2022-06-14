package utils

import (
	"testing"

	"github.com/SurgicalSteel/kvothe/resources"
)

func TestUtils_GeneratePostgreURL(t *testing.T) {
	testCases := []struct {
		name           string
		account        resources.DBAccount
		expectedResult string
	}{
		{
			name: "success case",
			account: resources.DBAccount{
				Username: "admin",
				Password: "admin",
				URL:      "localhost",
				Port:     "5432",
				DBName:   "inventory",
				Timeout:  "10",
			},
			expectedResult: "user=admin password=admin dbname=inventory host=localhost port=5432  sslmode=disable extra_float_digits=-1 connect_timeout=10",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := GeneratePostgreURL(tc.account)
			if actualResult != tc.expectedResult {
				t.Errorf("Expected result and actual result are not equal!\n expected : %s\n actual : %s\n", tc.expectedResult, actualResult)
			}
		})
	}
}
