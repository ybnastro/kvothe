package utils

import "testing"

func TestUtils_GetInt(t *testing.T) {
	testCases := []struct {
		name           string
		envString      string
		expectedResult int
	}{
		{
			name:           "success positive with sign",
			envString:      "+2",
			expectedResult: 2,
		},
		{
			name:           "success negative with sign",
			envString:      "-2",
			expectedResult: -2,
		},
		{
			name:           "failed invalid string",
			envString:      "&&&",
			expectedResult: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := GetInt(tc.envString)
			if actualResult != tc.expectedResult {
				t.Errorf("Expected result and actual result are not equal!\n expected : %d\n actual : %d\n", tc.expectedResult, actualResult)
			}
		})
	}
}

func TestUtils_GetInt64(t *testing.T) {
	testCases := []struct {
		name           string
		envString      string
		expectedResult int64
	}{
		{
			name:           "success positive with sign",
			envString:      "+2",
			expectedResult: 2,
		},
		{
			name:           "success negative with sign",
			envString:      "-2",
			expectedResult: -2,
		},
		{
			name:           "failed invalid string",
			envString:      "&&&",
			expectedResult: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := GetInt64(tc.envString)
			if actualResult != tc.expectedResult {
				t.Errorf("Expected result and actual result are not equal!\n expected : %d\n actual : %d\n", tc.expectedResult, actualResult)
			}
		})
	}
}

func TestUtils_GetBool(t *testing.T) {
	testCases := []struct {
		name           string
		envString      string
		expectedResult bool
	}{
		{
			name:           "success true mixed case",
			envString:      "trUE",
			expectedResult: true,
		},
		{
			name:           "success false mixed case",
			envString:      "FALse",
			expectedResult: false,
		},
		{
			name:           "failed invalid string",
			envString:      "&&&",
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := GetBool(tc.envString)
			if actualResult != tc.expectedResult {
				t.Errorf("Expected result and actual result are not equal!\n expected : %t\n actual : %t\n", tc.expectedResult, actualResult)
			}
		})
	}
}
