package utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtil_GetUniqueElementsSliceOfString(t *testing.T) {
	testCases := []struct {
		name           string
		inputSlice     []string
		expectedResult []string
	}{
		{
			name:           "normal slice with duplication",
			inputSlice:     []string{"this", "one", "and", "that", "one", "are", "yours", "and", "shut", "up"},
			expectedResult: []string{"this", "one", "and", "that", "are", "yours", "shut", "up"},
		},
		{
			name:           "normal slice without duplication",
			inputSlice:     []string{"Armin", "van", "Buuren", "is", "the", "greatest", "DJ", "ever"},
			expectedResult: []string{"Armin", "van", "Buuren", "is", "the", "greatest", "DJ", "ever"},
		},
		{
			name:           "empty slice",
			inputSlice:     []string{},
			expectedResult: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := GetUniqueElementsSliceOfString(tc.inputSlice)
			if !assert.ElementsMatch(t, tc.expectedResult, actualResult) {
				t.Error("expected result slice and actual result slice are not the same")
			}
		})
	}
}

func TestUtils_ConvertSliceOfStringIntoSliceOfInt64(t *testing.T) {
	testCases := []struct {
		name           string
		rawSlice       []string
		expectedResult []int64
	}{
		{
			name:           "normal slice",
			rawSlice:       []string{"1", "2", "3", "666"},
			expectedResult: []int64{1, 2, 3, 666},
		},
		{
			name:           "empty slice",
			rawSlice:       []string{},
			expectedResult: []int64{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := ConvertSliceOfStringIntoSliceOfInt64(tc.rawSlice)
			if len(actualResult) != len(tc.expectedResult) {
				t.Error("slice length are not the same")
			}
			if !reflect.DeepEqual(actualResult, tc.expectedResult) {
				t.Error("actual and expected slice are not the same")
			}
		})
	}
}

func TestUtils_ConstructSliceOfInt64IntoString(t *testing.T) {
	testCases := []struct {
		name           string
		rawSlice       []int64
		separator      string
		expectedResult string
	}{
		{
			name:           "normal slice",
			rawSlice:       []int64{1, 2, 3, 4, 5, 6},
			separator:      ",",
			expectedResult: "1,2,3,4,5,6",
		},
		{
			name:           "empty slice",
			rawSlice:       []int64{},
			separator:      "-",
			expectedResult: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := ConstructSliceOfInt64IntoString(tc.rawSlice, tc.separator)
			if actualResult != tc.expectedResult {
				t.Error("actual result and expected result are not the same")
			}
			if !reflect.DeepEqual(actualResult, tc.expectedResult) {
				t.Error("actual and expected slice are not the same")
			}
		})
	}
}
